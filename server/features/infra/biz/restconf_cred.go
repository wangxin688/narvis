package infra_biz

import (
	"errors"

	"github.com/samber/lo"
	"github.com/wangxin688/narvis/server/dal/gen"
	"github.com/wangxin688/narvis/server/features/infra/schemas"
	"github.com/wangxin688/narvis/server/global"
	"github.com/wangxin688/narvis/server/models"
	te "github.com/wangxin688/narvis/server/tools/errors"
	ts "github.com/wangxin688/narvis/server/tools/schemas"
	"gorm.io/gorm"
)

type RestConfCredentialService struct{}

func NewRestConfCredentialService() *RestConfCredentialService {
	return &RestConfCredentialService{}
}

func (r *RestConfCredentialService) CreateCredential(deviceId string, credential *schemas.RestconfCredentialCreate) (string, error) {
	cred := &models.RestconfCredential{
		OrganizationId: global.OrganizationId.Get(),
		DeviceId:       &deviceId,
		Username:       credential.Username,
		Password:       credential.Password,
	}
	if err := gen.RestconfCredential.Create(cred); err != nil {
		return "", err
	}
	return cred.Id, nil
}

// func (r *RestConfCredentialService) validateCreateCredential(deviceId string, credential *schemas.RestconfCredentialCreate) error {
// 	dbCred, err := gen.RestconfCredential.Where(
// 		gen.RestconfCredential.OrganizationId.Eq(global.OrganizationId.Get()),
// 		gen.RestconfCredential.DeviceId.Eq(deviceId),
// 	).First()
// 	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
// 		return err
// 	}
// 	if dbCred != nil {
// 		return te.NewError(te.CodeCredentialDeviceIdMissing, te.MsgCredentialDeviceIdMissing)
// 	}
// 	return nil
// }

func (r *RestConfCredentialService) UpdateCredential(deviceId string, credential *schemas.RestconfCredentialUpdate) (diff map[string]map[string]*ts.OrmDiff, err error) {
	dbCred, err := gen.RestconfCredential.Where(
		gen.RestconfCredential.OrganizationId.Eq(global.OrganizationId.Get()),
		gen.RestconfCredential.DeviceId.Eq(deviceId),
	).First()
	if err != nil {
		return nil, err
	}
	updateFields := make(map[string]*ts.OrmDiff)
	if credential.Url != nil && *credential.Url != dbCred.Url {
		updateFields["url"] = &ts.OrmDiff{Before: dbCred.Url, After: *credential.Url}
		dbCred.Url = *credential.Url
	}
	if credential.Username != nil && *credential.Username != dbCred.Username {
		updateFields["username"] = &ts.OrmDiff{Before: dbCred.Username, After: *credential.Username}
		dbCred.Username = *credential.Username
	}
	if credential.Password != nil && *credential.Password != dbCred.Password {
		updateFields["password"] = &ts.OrmDiff{Before: dbCred.Password, After: *credential.Password}
		dbCred.Password = *credential.Password
	}
	if len(updateFields) == 0 {
		return nil, nil
	}
	diff = make(map[string]map[string]*ts.OrmDiff)
	diff[dbCred.Id] = updateFields
	global.OrmDiff.Set(diff)
	err = gen.RestconfCredential.UnderlyingDB().Save(dbCred).Error
	if err != nil {
		return nil, err
	}
	return diff, nil
}

func (r *RestConfCredentialService) GetCredentialByOrgId(id string) (*models.RestconfCredential, error) {
	cred, err := gen.RestconfCredential.Where(gen.RestconfCredential.OrganizationId.Eq(id)).First()
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return cred, nil
}

func (r *RestConfCredentialService) DeleteCredential(deviceId string) error {

	dbCred, err := gen.RestconfCredential.Where(
		gen.RestconfCredential.OrganizationId.Eq(global.OrganizationId.Get()),
		gen.RestconfCredential.DeviceId.Eq(deviceId),
	).First()
	if err != nil {
		return err
	}
	_, err = gen.RestconfCredential.Delete(dbCred)
	if err != nil {
		return err
	}
	return nil
}

func (r *RestConfCredentialService) GetCredentialByDeviceId(deviceId string) (*schemas.RestconfCredential, error) {
	orgId := global.OrganizationId.Get()
	cred, err := gen.RestconfCredential.Where(
		gen.RestconfCredential.DeviceId.Eq(deviceId),
		gen.RestconfCredential.OrganizationId.Eq(orgId),
	).First()
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			globalCred, err := r.GetCredentialByOrgId(orgId)
			if err != nil {
				return nil, err
			}
			if globalCred != nil {
				return &schemas.RestconfCredential{
					Username: globalCred.Username,
					Password: globalCred.Password,
					Url:      globalCred.Url,
					InheritFromOrg: true,
				}, nil
			}
			return nil, te.NewError(te.CodeNotFound, te.MsgNotFound, gen.RestconfCredential.TableName(), "device_id", deviceId)
		}
		return nil, err
	}
	return &schemas.RestconfCredential{
		Username: cred.Username,
		Password: cred.Password,
		Url:      cred.Url,
		InheritFromOrg: false,
	}, nil
}

// GetCredentialByDeviceIds returns a map of restconf credential by device ids
// If a device does not have a restconf credential, it will use the global restconf credential for the organization.
// make sure deviceIds is a set of unique device ids
func (r *RestConfCredentialService) GetCredentialByDeviceIds(deviceIds []string) (*map[string]*schemas.RestconfCredential, error) {
	deviceIds = lo.Uniq(deviceIds) // remove duplicates deviceIds
	orgId := global.OrganizationId.Get()
	creds, err := gen.RestconfCredential.Where(
		gen.RestconfCredential.DeviceId.In(deviceIds...),
		gen.RestconfCredential.OrganizationId.Eq(orgId)).Find()
	if err != nil {
		return nil, err
	}
	if len(creds) == len(deviceIds) {
		results := make(map[string]*schemas.RestconfCredential)
		for _, cred := range creds {
			results[*cred.DeviceId] = &schemas.RestconfCredential{
				Url:      cred.Url,
				Username: cred.Username,
				Password: cred.Password,
				InheritFromOrg: false,
			}
		}
		return &results, nil
	}

	orgCred, err := r.GetCredentialByOrgId(orgId)
	if err != nil {
		return nil, err
	}

	results := make(map[string]*schemas.RestconfCredential)
	for _, cred := range creds {
		results[*cred.DeviceId] = &schemas.RestconfCredential{
			Url:      cred.Url,
			Username: cred.Username,
			Password: cred.Password,
			InheritFromOrg: false,
		}
	}
	var missing []string
	for _, id := range deviceIds {
		if _, ok := results[id]; !ok {
			missing = append(missing, id)
		}
	}
	if len(missing) > 0 {
		for _, deviceId := range missing {
			results[deviceId] = &schemas.RestconfCredential{
				Url:      orgCred.Url,
				Username: orgCred.Username,
				Password: orgCred.Password,
				InheritFromOrg: true,
			}
		}
	}
	return &results, nil
}

func (r *RestConfCredentialService) GetById(id string) (*schemas.RestconfCredential, error) {

	cred, err := gen.RestconfCredential.Where(
		gen.RestconfCredential.Id.Eq(id),
		gen.RestconfCredential.OrganizationId.Eq(global.OrganizationId.Get()),
	).First()

	if err != nil {
		return nil, err
	}
	return &schemas.RestconfCredential{
		Username: cred.Username,
		Password: cred.Password,
		Url:      cred.Url,
	}, nil

}
