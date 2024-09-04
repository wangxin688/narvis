package infra_biz

import (
	"errors"

	"github.com/samber/lo"
	"github.com/wangxin688/narvis/server/dal/gen"
	"github.com/wangxin688/narvis/server/features/infra/schemas"
	"github.com/wangxin688/narvis/server/global"
	"github.com/wangxin688/narvis/server/models"
	te "github.com/wangxin688/narvis/server/tools/errors"
	"gorm.io/gorm"
)

type SnmpCredentialService struct{}

func NewSnmpCredentialService() *SnmpCredentialService {
	return &SnmpCredentialService{}
}

func (s *SnmpCredentialService) CreateSnmpCredential(snmp *schemas.SnmpV2CredentialCreate) (string, error) {
	snmp.SetDefaultValue()

	if err := s.validateCreateSnmpCredential(snmp); err != nil {
		return "", err
	}
	cred := &models.SnmpV2Credential{
		OrganizationId: global.OrganizationId.Get(),
		DeviceId:       snmp.DeviceId,
		Community:      snmp.Community,
		Port:           *snmp.Port,
		Timeout:        *snmp.Timeout,
		MaxRepetitions: *snmp.MaxRepetitions,
	}
	err := gen.SnmpV2Credential.Create(cred)
	if err != nil {
		return "", err
	}
	return cred.Id, nil
}

func (s *SnmpCredentialService) validateCreateSnmpCredential(snmp *schemas.SnmpV2CredentialCreate) error {

	if snmp.DeviceId == nil {
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

func (s *SnmpCredentialService) UpdateSnmpCredential(credId string, snmp *schemas.SnmpV2CredentialUpdate) error {
	dbCred, err := gen.SnmpV2Credential.Where(
		gen.SnmpV2Credential.OrganizationId.Eq(global.OrganizationId.Get()),
		gen.SnmpV2Credential.Id.Eq(credId),
	).First()
	if err != nil {
		return err
	}
	if dbCred == nil {
		return te.NewError(te.CodeNotFound, te.MsgNotFound, gen.SnmpV2Credential.TableName(), "id", credId)
	}
	updateFields := make(map[string]any)

	if snmp.Community != nil {
		updateFields["community"] = *snmp.Community
	}
	if snmp.Port != nil {
		updateFields["port"] = *snmp.Port
	}
	if snmp.Timeout != nil {
		updateFields["timeout"] = *snmp.Timeout
	}
	if snmp.MaxRepetitions != nil {
		updateFields["max_repetitions"] = *snmp.MaxRepetitions
	}
	if len(updateFields) == 0 {
		return nil
	}
	_, err = gen.SnmpV2Credential.Where(gen.SnmpV2Credential.Id.Eq(credId)).Updates(updateFields)
	if err != nil {
		return err
	}
	return nil

}

func (s *SnmpCredentialService) GetCredentialByOrgId(orgId string) (*models.SnmpV2Credential, error) {
	cred, err := gen.SnmpV2Credential.Where(gen.SnmpV2Credential.OrganizationId.Eq(orgId)).First()
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return cred, err
}

func (s *SnmpCredentialService) DeleteCredential(credId string) error {
	_, err := gen.SnmpV2Credential.Where(
		gen.SnmpV2Credential.OrganizationId.Eq(global.OrganizationId.Get()),
		gen.SnmpV2Credential.Id.Eq(credId),
	).Delete()
	if err != nil {
		return err
	}
	return nil
}

func (s *SnmpCredentialService) GetCredentialByDeviceId(deviceId string) (*schemas.SnmpV2Credential, error) {
	orgId := global.OrganizationId.Get()
	cred, err := gen.SnmpV2Credential.Where(
		gen.SnmpV2Credential.OrganizationId.Eq(orgId),
		gen.SnmpV2Credential.DeviceId.Eq(deviceId),
	).First()
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			globalCred, err := s.GetCredentialByOrgId(orgId)
			if err != nil {
				return nil, err
			}
			if globalCred == nil {
				return nil, te.NewError(te.CodeNotFound, te.MsgNotFound, gen.SnmpV2Credential.TableName(), "deviceId", deviceId)
			}
			return &schemas.SnmpV2Credential{
				Community:      globalCred.Community,
				MaxRepetitions: globalCred.MaxRepetitions,
				Timeout:        globalCred.Timeout,
				Port:           globalCred.Port,
			}, nil
		}
		return nil, err
	}
	return &schemas.SnmpV2Credential{
		Community:      cred.Community,
		MaxRepetitions: cred.MaxRepetitions,
		Timeout:        cred.Timeout,
		Port:           cred.Port,
	}, err
}

// GetCredentialByDeviceIds returns a map of device id to snmp credential.
// If a device does not have a snmp credential, it will use the global snmp credential for the organization.
// make sure deviceIds is a set of unique device ids
func (s *SnmpCredentialService) GetCredentialByDeviceIds(deviceIds []string) (map[string]*schemas.SnmpV2Credential, error) {
	deviceIds = lo.Uniq(deviceIds) // remove duplicates deviceIds
	orgId := global.OrganizationId.Get()
	creds, err := gen.SnmpV2Credential.Where(
		gen.SnmpV2Credential.OrganizationId.Eq(orgId),
		gen.SnmpV2Credential.DeviceId.In(deviceIds...),
	).Find()
	if err != nil {
		return nil, err
	}
	// if the size of creds is equal to the size of deviceIds, it means that all devices have a snmp credential
	if len(creds) == len(deviceIds) {
		results := make(map[string]*schemas.SnmpV2Credential)
		for _, cred := range creds {
			results[*cred.DeviceId] = &schemas.SnmpV2Credential{
				Community:      cred.Community,
				MaxRepetitions: cred.MaxRepetitions,
				Timeout:        cred.Timeout,
				Port:           cred.Port,
			}
		}
		return results, nil
	}
	orgCred, err := s.GetCredentialByOrgId(orgId)
	if err != nil {
		return nil, err
	}
	results := make(map[string]*schemas.SnmpV2Credential)
	for _, deviceId := range deviceIds {
		results[deviceId] = &schemas.SnmpV2Credential{
			Community:      orgCred.Community,
			MaxRepetitions: orgCred.MaxRepetitions,
			Timeout:        orgCred.Timeout,
			Port:           orgCred.Port,
		}
	}
	var missing []string
	for _, deviceId := range deviceIds {
		if _, ok := results[deviceId]; !ok {
			missing = append(missing, deviceId)
		}
	}
	if len(missing) > 0 {
		for _, deviceId := range missing {
			results[deviceId] = &schemas.SnmpV2Credential{
				Community:      orgCred.Community,
				MaxRepetitions: orgCred.MaxRepetitions,
				Timeout:        orgCred.Timeout,
				Port:           orgCred.Port,
			}
		}
	}
	return results, nil
}
