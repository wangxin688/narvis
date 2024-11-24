package infra_biz

import (
	"errors"

	"github.com/samber/lo"
	"github.com/wangxin688/narvis/server/dal/gen"
	"github.com/wangxin688/narvis/server/features/infra/schemas"
	"github.com/wangxin688/narvis/server/models"
	"github.com/wangxin688/narvis/server/pkg/contextvar"
	te "github.com/wangxin688/narvis/server/tools/errors"
	"gorm.io/gorm"
)

type CliCredentialService struct{}

func NewCliCredentialService() *CliCredentialService {
	return &CliCredentialService{}
}

func (s *CliCredentialService) CreateCredential(deviceId string, credential *schemas.CliCredentialCreate) (string, error) {
	cred := &models.CliCredential{
		OrganizationId: contextvar.OrganizationId.Get(),
		DeviceId:       &deviceId,
		Username:       credential.Username,
		Password:       credential.Password,
	}
	if err := gen.CliCredential.Create(cred); err != nil {
		return "", err
	}
	return cred.Id, nil
}

func (s *CliCredentialService) CreateServerCredential(serverId string, credential *schemas.CliCredentialCreate) (string, error) {
	cred := &models.ServerCredential{
		OrganizationId: contextvar.OrganizationId.Get(),
		ServerId:       &serverId,
		Username:       credential.Username,
		Password:       credential.Password,
	}
	if err := gen.ServerCredential.Create(cred); err != nil {
		return "", err
	}
	return cred.Id, nil
}

// func (s *CliCredentialService) validateCreateCredential(credential *schemas.CliCredentialCreate) error {
// 	if credential.DeviceId == nil {
// 		orgCred, err := s.GetCredentialByOrgId(contextvar.OrganizationId.Get())
// 		if err != nil {
// 			return err
// 		}
// 		if orgCred != nil {
// 			return te.NewError(te.CodeCredentialDeviceIdMissing, te.MsgCredentialDeviceIdMissing)
// 		}
// 	}
// 	return nil
// }

func (s *CliCredentialService) UpdateCredential(deviceId string, credential *schemas.CliCredentialUpdate) (diff map[string]map[string]*contextvar.Diff, err error) {
	dbCred, err := gen.CliCredential.Where(gen.CliCredential.DeviceId.Eq(deviceId), gen.CliCredential.OrganizationId.Eq(contextvar.OrganizationId.Get())).First()
	if err != nil {
		return nil, err
	}
	updateFields := make(map[string]*contextvar.Diff)
	if credential.Username != nil && *credential.Username != dbCred.Username {
		updateFields["username"] = &contextvar.Diff{Before: dbCred.Username, After: *credential.Username}
		dbCred.Username = *credential.Username
	}
	if credential.Password != nil && *credential.Password != dbCred.Password {
		updateFields["password"] = &contextvar.Diff{Before: dbCred.Password, After: *credential.Password}
		dbCred.Password = *credential.Password
	}
	if len(updateFields) == 0 {
		return nil, nil
	}

	diff = make(map[string]map[string]*contextvar.Diff)
	diff[dbCred.Id] = updateFields
	contextvar.OrmDiff.Set(diff)
	err = gen.CliCredential.UnderlyingDB().Save(dbCred).Error
	if err != nil {
		return nil, err
	}
	return diff, nil
}

func (s *CliCredentialService) UpdateServerCredential(serverId string, credential *schemas.CliCredentialUpdate) (diff map[string]map[string]*contextvar.Diff, err error) {
	dbCred, err := gen.ServerCredential.Where(gen.ServerCredential.ServerId.Eq(serverId), gen.ServerCredential.OrganizationId.Eq(contextvar.OrganizationId.Get())).First()
	if err != nil {
		return nil, err
	}
	updateFields := make(map[string]*contextvar.Diff)
	if credential.Username != nil && *credential.Username != dbCred.Username {
		updateFields["username"] = &contextvar.Diff{Before: dbCred.Username, After: *credential.Username}
		dbCred.Username = *credential.Username
	}
	if credential.Password != nil && *credential.Password != dbCred.Password {
		updateFields["password"] = &contextvar.Diff{Before: dbCred.Password, After: *credential.Password}
		dbCred.Password = *credential.Password
	}
	if len(updateFields) == 0 {
		return nil, nil
	}

	diff = make(map[string]map[string]*contextvar.Diff)
	diff[dbCred.Id] = updateFields
	contextvar.OrmDiff.Set(diff)
	err = gen.ServerCredential.UnderlyingDB().Save(dbCred).Error
	if err != nil {
		return nil, err
	}
	return diff, nil
}

func (s *CliCredentialService) GetCredentialByOrgId(orgId string) (*models.CliCredential, error) {
	cred, err := gen.CliCredential.Where(gen.CliCredential.OrganizationId.Eq(orgId)).First()
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return cred, err

}

func (s *CliCredentialService) GetServerCredentialByOrgId(orgId string) (*models.ServerCredential, error) {
	cred, err := gen.ServerCredential.Where(gen.ServerCredential.OrganizationId.Eq(orgId)).First()
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return cred, err

}

func (s *CliCredentialService) DeleteCredential(deviceId string) error {
	dbCred, err := gen.CliCredential.Where(
		gen.CliCredential.Id.Eq(deviceId),
		gen.CliCredential.OrganizationId.Eq(contextvar.OrganizationId.Get()),
	).First()
	if err != nil {
		return err
	}
	_, err = gen.CliCredential.Delete(dbCred)
	if err != nil {
		return err
	}
	return nil
}

func (s *CliCredentialService) DeleteServerCredential(serverId string) error {
	dbCred, err := gen.ServerCredential.Where(
		gen.ServerCredential.Id.Eq(serverId),
		gen.ServerCredential.OrganizationId.Eq(contextvar.OrganizationId.Get()),
	).First()
	if err != nil {
		return err
	}
	_, err = gen.ServerCredential.Delete(dbCred)
	if err != nil {
		return err
	}
	return nil
}

func (s *CliCredentialService) GetCredentialByDeviceId(deviceId string) (*schemas.CliCredential, error) {
	orgId := contextvar.OrganizationId.Get()

	cred, err := gen.CliCredential.Where(
		gen.CliCredential.DeviceId.Eq(deviceId),
		gen.CliCredential.OrganizationId.Eq(orgId),
	).First()
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			globalCred, err := s.GetCredentialByOrgId(orgId)
			if err != nil {
				return nil, err
			}
			if globalCred == nil {
				return nil, te.NewError(te.CodeNotFound, te.MsgNotFound, gen.CliCredential.TableName(), "deviceId", deviceId)
			}
			return &schemas.CliCredential{
				Username:       globalCred.Username,
				Password:       globalCred.Password,
				Port:           globalCred.Port,
				InheritFromOrg: true,
			}, nil
		}
		return nil, err
	}

	return &schemas.CliCredential{
		Username:       cred.Username,
		Password:       cred.Password,
		Port:           cred.Port,
		InheritFromOrg: false,
	}, nil
}

func (s *CliCredentialService) GetServerCredentialByDeviceId(serverId string) (*schemas.CliCredential, error) {
	orgId := contextvar.OrganizationId.Get()

	cred, err := gen.ServerCredential.Where(
		gen.ServerCredential.ServerId.Eq(serverId),
		gen.ServerCredential.OrganizationId.Eq(orgId),
	).First()
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			globalCred, err := s.GetServerCredentialByOrgId(orgId)
			if err != nil {
				return nil, err
			}
			if globalCred == nil {
				return nil, te.NewError(te.CodeNotFound, te.MsgNotFound, gen.ServerCredential.TableName(), "deviceId", serverId)
			}
			return &schemas.CliCredential{
				Username:       globalCred.Username,
				Password:       globalCred.Password,
				Port:           globalCred.Port,
				InheritFromOrg: true,
			}, nil
		}
		return nil, err
	}

	return &schemas.CliCredential{
		Username:       cred.Username,
		Password:       cred.Password,
		Port:           cred.Port,
		InheritFromOrg: false,
	}, nil
}

// GetCredentialByDeviceIds returns a map of device id to cli credential.
// If a device does not have a cli credential, it will use the global cli credential for the organization.
func (s *CliCredentialService) GetCredentialByDeviceIds(deviceIds []string) (map[string]*schemas.CliCredential, error) {
	deviceIds = lo.Uniq(deviceIds) // remove duplicates deviceIds
	orgId := contextvar.OrganizationId.Get()
	creds, err := gen.CliCredential.Where(
		gen.CliCredential.DeviceId.In(deviceIds...),
		gen.CliCredential.OrganizationId.Eq(orgId)).Find()
	if err != nil {
		return nil, err
	}

	// if the size of creds is equal to the size of deviceIds, it means that all devices have credentials
	if len(creds) == len(deviceIds) {
		results := make(map[string]*schemas.CliCredential)
		for _, cred := range creds {
			results[*cred.DeviceId] = &schemas.CliCredential{
				Username:       cred.Username,
				Password:       cred.Password,
				Port:           cred.Port,
				InheritFromOrg: false,
			}
		}
		return results, nil
	}

	// get the global credential
	orgCred, err := s.GetCredentialByOrgId(orgId)
	if err != nil {
		return nil, err
	}

	// fill the results map
	results := make(map[string]*schemas.CliCredential)
	for _, cred := range creds {
		results[*cred.DeviceId] = &schemas.CliCredential{
			Username:       cred.Username,
			Password:       cred.Password,
			Port:           cred.Port,
			InheritFromOrg: false,
		}
	}
	var missing []string
	for _, deviceId := range deviceIds {
		if _, ok := results[deviceId]; !ok {
			missing = append(missing, deviceId)
		}
	}

	// if there are missing devices, fill the results with the global credential
	if len(missing) > 0 {
		for _, deviceId := range missing {
			results[deviceId] = &schemas.CliCredential{
				Username:       orgCred.Username,
				Password:       orgCred.Password,
				Port:           orgCred.Port,
				InheritFromOrg: true,
			}
		}
	}
	return results, nil
}

func (s *CliCredentialService) GetCredentialByServerIds(serverIds []string) (map[string]*schemas.CliCredential, error) {
	serverIds = lo.Uniq(serverIds) // remove duplicates serverIds
	orgId := contextvar.OrganizationId.Get()
	creds, err := gen.ServerCredential.Where(
		gen.ServerCredential.ServerId.In(serverIds...),
		gen.ServerCredential.OrganizationId.Eq(orgId)).Find()
	if err != nil {
		return nil, err
	}
	if len(creds) == len(serverIds) {
		results := make(map[string]*schemas.CliCredential)
		for _, cred := range creds {
			results[*cred.ServerId] = &schemas.CliCredential{
				Username:       cred.Username,
				Password:       cred.Password,
				Port:           cred.Port,
				InheritFromOrg: false,
			}
		}
		return results, nil
	}

	orgCred, err := s.GetCredentialByOrgId(orgId)
	if err != nil {
		return nil, err
	}

	results := make(map[string]*schemas.CliCredential)
	for _, cred := range creds {
		results[*cred.ServerId] = &schemas.CliCredential{
			Username:       cred.Username,
			Password:       cred.Password,
			Port:           cred.Port,
			InheritFromOrg: false,
		}
	}
	var missing []string
	for _, serverId := range serverIds {
		if _, ok := results[serverId]; !ok {
			missing = append(missing, serverId)
		}
	}
	if len(missing) > 0 {
		for _, serverId := range missing {
			results[serverId] = &schemas.CliCredential{
				Username:       orgCred.Username,
				Password:       orgCred.Password,
				Port:           orgCred.Port,
				InheritFromOrg: true,
			}
		}
	}
	return results, nil
}
