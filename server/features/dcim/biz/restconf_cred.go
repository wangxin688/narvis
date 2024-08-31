package biz

import (
	"errors"

	"github.com/samber/lo"
	"github.com/wangxin688/narvis/server/dal/gen"
	"github.com/wangxin688/narvis/server/features/dcim/schemas"
	"github.com/wangxin688/narvis/server/global"
	"github.com/wangxin688/narvis/server/models"
	te "github.com/wangxin688/narvis/server/tools/errors"
	"gorm.io/gorm"
)

type RestConfCredentialService struct{}

func NewRestConfCredentialService() *RestConfCredentialService {
	return &RestConfCredentialService{}
}

func (r *RestConfCredentialService) CreateCredential(credential *schemas.RestconfCredentialCreate) (string, error) {

	if err := r.validateCreateCredential(credential); err != nil {
		return "", err
	}

	cred := &models.RestconfCredential{
		OrganizationId: global.OrganizationId.Get(),
		DeviceId:       credential.DeviceId,
		Username:       credential.Username,
		Password:       credential.Password,
	}
	if err := gen.RestconfCredential.Create(cred); err != nil {
		return "", err
	}
	return cred.Id, nil
}

func (r *RestConfCredentialService) validateCreateCredential(credential *schemas.RestconfCredentialCreate) error {
	if credential.DeviceId == nil {
		orgCred, err := r.GetCredentialByOrgId(global.OrganizationId.Get())
		if err != nil {
			return err
		}
		if orgCred != nil {
			return te.NewError(te.CodeCredentialDeviceIdMissing, te.MsgCredentialDeviceIdMissing)
		}
		return nil
	}
	return nil
}

func (r *RestConfCredentialService) UpdateCredential(credId string, credential *schemas.RestconfCredentialUpdate) error {
	dbCred, err := gen.RestconfCredential.Where(
		gen.RestconfCredential.OrganizationId.Eq(global.OrganizationId.Get()),
		gen.RestconfCredential.Id.Eq(credId),
	).First()
	if err != nil {
		return err
	}
	if dbCred == nil {
		return te.NewError(te.CodeNotFound, te.MsgNotFound, gen.RestconfCredential.TableName(), "id", credId)
	}

	updateFields := make(map[string]any)
	if credential.Url != nil {
		updateFields["url"] = *credential.Url
	}
	if credential.Username != nil {
		updateFields["username"] = *credential.Username
	}
	if credential.Password != nil {
		updateFields["password"] = *credential.Password
	}
	_, err = gen.RestconfCredential.Where(gen.RestconfCredential.Id.Eq(credId)).Updates(updateFields)
	if err != nil {
		return err
	}
	return nil
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

func (r *RestConfCredentialService) DeleteCredential(id string) error {

	_, err := gen.RestconfCredential.Where(
		gen.RestconfCredential.OrganizationId.Eq(global.OrganizationId.Get()),
		gen.RestconfCredential.Id.Eq(id),
	).Delete()
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
			}
		}
	}
	return &results, nil
}
