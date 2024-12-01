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

type SnmpCredentialService struct{}

func NewSnmpCredentialService() *SnmpCredentialService {
	return &SnmpCredentialService{}
}

func (s *SnmpCredentialService) CreateSnmpCredential(deviceId string, snmp *schemas.SnmpV2CredentialCreate) (string, error) {
	snmp.SetDefaultValue()
	cred := &models.SnmpV2Credential{
		OrganizationId: contextvar.OrganizationId.Get(),
		DeviceId:       &deviceId,
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

func (s *SnmpCredentialService) CreateServerSnmpCredential(serverId string, snmp *schemas.SnmpV2CredentialCreate) (string, error) {
	snmp.SetDefaultValue()
	cred := &models.ServerSnmpCredential{
		OrganizationId: contextvar.OrganizationId.Get(),
		ServerId:       &serverId,
		Community:      snmp.Community,
		Port:           *snmp.Port,
		Timeout:        *snmp.Timeout,
		MaxRepetitions: *snmp.MaxRepetitions,
	}
	err := gen.ServerSnmpCredential.Create(cred)
	if err != nil {
		return "", err
	}
	return cred.Id, nil
}

// func (s *SnmpCredentialService) validateCreateSnmpCredential(snmp *schemas.SnmpV2CredentialCreate) error {

// 	if snmp.DeviceId == nil {
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

func (s *SnmpCredentialService) UpdateSnmpCredential(deviceId string, snmp *schemas.SnmpV2CredentialUpdate) (credId string, diff map[string]map[string]*contextvar.Diff, err error) {
	dbCred, err := gen.SnmpV2Credential.Where(
		gen.SnmpV2Credential.OrganizationId.Eq(contextvar.OrganizationId.Get()),
		gen.SnmpV2Credential.DeviceId.Eq(deviceId),
	).First()
	if err != nil {
		return "", nil, &te.GenericError{
			Code:    te.CodeCredentialUpdateNotFound,
			Message: te.MsgCredentialUpdateNotFound,
		}
	}
	updateFields := make(map[string]*contextvar.Diff)

	if snmp.Community != nil && dbCred.Community != *snmp.Community {
		updateFields["community"] = &contextvar.Diff{Before: dbCred.Community, After: *snmp.Community}
		dbCred.Community = *snmp.Community
	}
	if snmp.Port != nil && dbCred.Port != *snmp.Port {
		updateFields["port"] = &contextvar.Diff{Before: dbCred.Port, After: *snmp.Port}
		dbCred.Port = *snmp.Port
	}
	if snmp.Timeout != nil && dbCred.Timeout != *snmp.Timeout {
		updateFields["timeout"] = &contextvar.Diff{Before: dbCred.Timeout, After: *snmp.Timeout}
		dbCred.Timeout = *snmp.Timeout
	}
	if snmp.MaxRepetitions != nil && dbCred.MaxRepetitions != *snmp.MaxRepetitions {
		updateFields["max_repetitions"] = &contextvar.Diff{Before: dbCred.MaxRepetitions, After: *snmp.MaxRepetitions}
		dbCred.MaxRepetitions = *snmp.MaxRepetitions
	}
	if len(updateFields) == 0 {
		return "", nil, nil
	}
	diffValue := make(map[string]map[string]*contextvar.Diff)
	diffValue[dbCred.Id] = updateFields
	contextvar.OrmDiff.Set(diffValue)
	err = gen.SnmpV2Credential.UnderlyingDB().Save(dbCred).Error
	if err != nil {
		return "", nil, err
	}
	return dbCred.Id, diffValue, nil
}

func (s *SnmpCredentialService) UpdateServerSnmpCredential(serverId string, snmp *schemas.SnmpV2CredentialUpdate) (credId string, diff map[string]map[string]*contextvar.Diff, err error) {
	dbCred, err := gen.ServerSnmpCredential.Where(
		gen.ServerSnmpCredential.OrganizationId.Eq(contextvar.OrganizationId.Get()),
		gen.ServerSnmpCredential.ServerId.Eq(serverId),
	).First()
	if err != nil {
		return "", nil, &te.GenericError{
			Code:    te.CodeCredentialUpdateNotFound,
			Message: te.MsgCredentialUpdateNotFound,
		}
	}
	updateFields := make(map[string]*contextvar.Diff)

	if snmp.Community != nil && dbCred.Community != *snmp.Community {
		updateFields["community"] = &contextvar.Diff{Before: dbCred.Community, After: *snmp.Community}
		dbCred.Community = *snmp.Community
	}
	if snmp.Port != nil && dbCred.Port != *snmp.Port {
		updateFields["port"] = &contextvar.Diff{Before: dbCred.Port, After: *snmp.Port}
		dbCred.Port = *snmp.Port
	}
	if snmp.Timeout != nil && dbCred.Timeout != *snmp.Timeout {
		updateFields["timeout"] = &contextvar.Diff{Before: dbCred.Timeout, After: *snmp.Timeout}
		dbCred.Timeout = *snmp.Timeout
	}
	if snmp.MaxRepetitions != nil && dbCred.MaxRepetitions != *snmp.MaxRepetitions {
		updateFields["max_repetitions"] = &contextvar.Diff{Before: dbCred.MaxRepetitions, After: *snmp.MaxRepetitions}
		dbCred.MaxRepetitions = *snmp.MaxRepetitions
	}
	if len(updateFields) == 0 {
		return "", nil, nil
	}
	diffValue := make(map[string]map[string]*contextvar.Diff)
	diffValue[dbCred.Id] = updateFields
	contextvar.OrmDiff.Set(diffValue)
	err = gen.ServerSnmpCredential.UnderlyingDB().Save(dbCred).Error
	if err != nil {
		return "", nil, err
	}
	return dbCred.Id, diffValue, nil
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

func (s *SnmpCredentialService) GetServerCredentialByOrgId(orgId string) (*models.ServerSnmpCredential, error) {
	cred, err := gen.ServerSnmpCredential.Where(gen.ServerSnmpCredential.OrganizationId.Eq(orgId)).First()
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return cred, err
}

func (s *SnmpCredentialService) DeleteCredential(deviceId string) (cred *models.SnmpV2Credential, err error) {
	dbCred, err := gen.SnmpV2Credential.Where(
		gen.SnmpV2Credential.OrganizationId.Eq(contextvar.OrganizationId.Get()),
		gen.SnmpV2Credential.DeviceId.Eq(deviceId),
	).First()
	if err != nil {
		return nil, err
	}
	_, err = gen.SnmpV2Credential.Delete(dbCred)
	if err != nil {
		return nil, err
	}
	return dbCred, nil
}

func (s *SnmpCredentialService) DeleteServerCredential(serverId string) (cred *models.ServerSnmpCredential, err error) {
	dbCred, err := gen.ServerSnmpCredential.Where(
		gen.ServerSnmpCredential.OrganizationId.Eq(contextvar.OrganizationId.Get()),
		gen.ServerSnmpCredential.ServerId.Eq(serverId),
	).First()
	if err != nil {
		return nil, err
	}
	_, err = gen.ServerSnmpCredential.Delete(dbCred)
	if err != nil {
		return nil, err
	}
	return dbCred, nil
}

func (s *SnmpCredentialService) GetCredentialByDeviceId(deviceId string) (*schemas.SnmpV2Credential, error) {
	orgId := contextvar.OrganizationId.Get()
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
				InheritFromOrg: true,
			}, nil
		}
		return nil, err
	}
	return &schemas.SnmpV2Credential{
		Community:      cred.Community,
		MaxRepetitions: cred.MaxRepetitions,
		Timeout:        cred.Timeout,
		Port:           cred.Port,
		InheritFromOrg: false,
	}, err
}

func (s *SnmpCredentialService) GetServerCredentialByDeviceId(deviceId string) (*schemas.SnmpV2Credential, error) {
	orgId := contextvar.OrganizationId.Get()
	cred, err := gen.ServerSnmpCredential.Where(
		gen.ServerSnmpCredential.OrganizationId.Eq(orgId),
		gen.ServerSnmpCredential.ServerId.Eq(deviceId),
	).First()
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			globalCred, err := s.GetServerCredentialByOrgId(orgId)
			if err != nil {
				return nil, err
			}
			if globalCred == nil {
				return &schemas.SnmpV2Credential{}, nil
			}
			return &schemas.SnmpV2Credential{
				Community:      globalCred.Community,
				MaxRepetitions: globalCred.MaxRepetitions,
				Timeout:        globalCred.Timeout,
				Port:           globalCred.Port,
				InheritFromOrg: true,
			}, nil
		}
		return nil, err
	}
	return &schemas.SnmpV2Credential{
		Community:      cred.Community,
		MaxRepetitions: cred.MaxRepetitions,
		Timeout:        cred.Timeout,
		Port:           cred.Port,
		InheritFromOrg: false,
	}, err
}

// GetCredentialByDeviceIds returns a map of device id to snmp credential.
// If a device does not have a snmp credential, it will use the global snmp credential for the organization.
// make sure deviceIds is a set of unique device ids
func (s *SnmpCredentialService) GetCredentialByDeviceIds(deviceIds []string) (map[string]*schemas.SnmpV2Credential, error) {
	deviceIds = lo.Uniq(deviceIds) // remove duplicates deviceIds
	orgId := contextvar.OrganizationId.Get()
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
				InheritFromOrg: false,
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
			InheritFromOrg: false,
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
				InheritFromOrg: true,
			}
		}
	}
	return results, nil
}

func (s *SnmpCredentialService) GetCredentialByServerIds(serverIds []string) (map[string]*schemas.SnmpV2Credential, error) {
	serverIds = lo.Uniq(serverIds) // remove duplicates serverIds
	orgId := contextvar.OrganizationId.Get()
	creds, err := gen.ServerSnmpCredential.Where(
		gen.ServerSnmpCredential.OrganizationId.Eq(orgId),
		gen.ServerSnmpCredential.ServerId.In(serverIds...),
	).Find()
	if err != nil {
		return nil, err
	}
	if len(creds) == len(serverIds) {
		results := make(map[string]*schemas.SnmpV2Credential)
		for _, cred := range creds {
			results[*cred.ServerId] = &schemas.SnmpV2Credential{
				Community:      cred.Community,
				MaxRepetitions: cred.MaxRepetitions,
				Timeout:        cred.Timeout,
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
	results := make(map[string]*schemas.SnmpV2Credential)
	for _, cred := range creds {
		results[*cred.ServerId] = &schemas.SnmpV2Credential{
			Community:      cred.Community,
			MaxRepetitions: cred.MaxRepetitions,
			Timeout:        cred.Timeout,
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
			results[serverId] = &schemas.SnmpV2Credential{
				Community:      orgCred.Community,
				MaxRepetitions: orgCred.MaxRepetitions,
				Timeout:        orgCred.Timeout,
				Port:           orgCred.Port,
				InheritFromOrg: true,
			}
		}
	}
	return results, nil
}
