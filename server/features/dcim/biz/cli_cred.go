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

type CliCredentialService struct{}

func NewCliCredentialService() *CliCredentialService {
	return &CliCredentialService{}
}

func (s *CliCredentialService) CreateCredential(credential *schemas.CliCredentialCreate) (string, error) {

	if err := s.validateCreateCredential(credential); err != nil {
		return "", err
	}

	cred := &models.CliCredential{
		OrganizationId: global.OrganizationId.Get(),
		DeviceId:       credential.DeviceId,
		Username:       credential.Username,
		Password:       credential.Password,
	}
	if err := gen.CliCredential.Create(cred); err != nil {
		return "", err
	}
	return cred.Id, nil
}

func (s *CliCredentialService) validateCreateCredential(credential *schemas.CliCredentialCreate) error {
	if credential.DeviceId == nil {
		orgCred, err := s.GetCredentialByOrgId(global.OrganizationId.Get())
		if err != nil {
			return err
		}
		if orgCred != nil {
			return te.NewError(te.CodeCredentialDeviceIdMissing, te.MsgCredentialDeviceIdMissing)
		}
	}
	return nil
}

func (s *CliCredentialService) UpdateCredential(credId string, credential *schemas.CliCredentialUpdate) error {
	dbCred, err := gen.CliCredential.Where(gen.CliCredential.Id.Eq(credId), gen.CliCredential.OrganizationId.Eq(global.OrganizationId.Get())).First()
	if err != nil {
		return err
	}
	updateFields := make(map[string]any)
	if dbCred == nil {
		return te.NewError(te.CodeNotFound, te.MsgNotFound, gen.CliCredential.TableName(), "id", credId)
	}
	if credential.Username != nil {
		updateFields["username"] = *credential.Username
	}
	if credential.Password != nil {
		updateFields["password"] = *credential.Password
	}
	_, err = gen.CliCredential.Where(gen.CliCredential.Id.Eq(credId)).Updates(updateFields)
	if err != nil {
		return err
	}
	return nil
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

func (s *CliCredentialService) DeleteCredential(credId string) error {
	_, err := gen.CliCredential.Where(
		gen.CliCredential.Id.Eq(credId),
		gen.CliCredential.OrganizationId.Eq(global.OrganizationId.Get()),
	).Delete()
	if err != nil {
		return err
	}
	return nil
}

func (s *CliCredentialService) GetCredentialByDeviceId(deviceId string) (*schemas.CliCredential, error) {
	orgId := global.OrganizationId.Get()

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
				Username: globalCred.Username,
				Password: globalCred.Password,
				Port:     globalCred.Port,
			}, nil
		}
		return nil, err
	}

	return &schemas.CliCredential{
		Username: cred.Username,
		Password: cred.Password,
		Port:     cred.Port,
	}, nil
}

// GetCredentialByDeviceIds returns a map of device id to cli credential.
// If a device does not have a cli credential, it will use the global cli credential for the organization.
func (s *CliCredentialService) GetCredentialByDeviceIds(deviceIds []string) (map[string]*schemas.CliCredential, error) {
	deviceIds = lo.Uniq(deviceIds) // remove duplicates deviceIds
	orgId := global.OrganizationId.Get()
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
				Username: cred.Username,
				Password: cred.Password,
				Port:     cred.Port,
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
			Username: cred.Username,
			Password: cred.Password,
			Port:     cred.Port,
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
				Username: orgCred.Username,
				Password: orgCred.Password,
				Port:     orgCred.Port,
			}
		}
	}
	return results, nil
}

