// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package gen

import (
	"context"
	"database/sql"

	"gorm.io/gorm"

	"gorm.io/gen"

	"gorm.io/plugin/dbresolver"
)

var (
	Q                    = new(Query)
	AP                   *aP
	Alert                *alert
	AlertActionLog       *alertActionLog
	AlertGroup           *alertGroup
	ApLLDPNeighbor       *apLLDPNeighbor
	AuditLog             *auditLog
	Circuit              *circuit
	CliCredential        *cliCredential
	Device               *device
	DeviceConfig         *deviceConfig
	DeviceInterface      *deviceInterface
	DeviceStack          *deviceStack
	IpAddress            *ipAddress
	LLDPNeighbor         *lLDPNeighbor
	MacAddress           *macAddress
	Maintenance          *maintenance
	Menu                 *menu
	Organization         *organization
	Permission           *permission
	Prefix               *prefix
	Proxy                *proxy
	Rack                 *rack
	RestconfCredential   *restconfCredential
	Role                 *role
	RootCause            *rootCause
	ScanDevice           *scanDevice
	Server               *server
	ServerCredential     *serverCredential
	ServerSnmpCredential *serverSnmpCredential
	Site                 *site
	SnmpV2Credential     *snmpV2Credential
	Subscription         *subscription
	SubscriptionRecord   *subscriptionRecord
	TaskResult           *taskResult
	Template             *template
	User                 *user
)

func SetDefault(db *gorm.DB, opts ...gen.DOOption) {
	*Q = *Use(db, opts...)
	AP = &Q.AP
	Alert = &Q.Alert
	AlertActionLog = &Q.AlertActionLog
	AlertGroup = &Q.AlertGroup
	ApLLDPNeighbor = &Q.ApLLDPNeighbor
	AuditLog = &Q.AuditLog
	Circuit = &Q.Circuit
	CliCredential = &Q.CliCredential
	Device = &Q.Device
	DeviceConfig = &Q.DeviceConfig
	DeviceInterface = &Q.DeviceInterface
	DeviceStack = &Q.DeviceStack
	IpAddress = &Q.IpAddress
	LLDPNeighbor = &Q.LLDPNeighbor
	MacAddress = &Q.MacAddress
	Maintenance = &Q.Maintenance
	Menu = &Q.Menu
	Organization = &Q.Organization
	Permission = &Q.Permission
	Prefix = &Q.Prefix
	Proxy = &Q.Proxy
	Rack = &Q.Rack
	RestconfCredential = &Q.RestconfCredential
	Role = &Q.Role
	RootCause = &Q.RootCause
	ScanDevice = &Q.ScanDevice
	Server = &Q.Server
	ServerCredential = &Q.ServerCredential
	ServerSnmpCredential = &Q.ServerSnmpCredential
	Site = &Q.Site
	SnmpV2Credential = &Q.SnmpV2Credential
	Subscription = &Q.Subscription
	SubscriptionRecord = &Q.SubscriptionRecord
	TaskResult = &Q.TaskResult
	Template = &Q.Template
	User = &Q.User
}

func Use(db *gorm.DB, opts ...gen.DOOption) *Query {
	return &Query{
		db:                   db,
		AP:                   newAP(db, opts...),
		Alert:                newAlert(db, opts...),
		AlertActionLog:       newAlertActionLog(db, opts...),
		AlertGroup:           newAlertGroup(db, opts...),
		ApLLDPNeighbor:       newApLLDPNeighbor(db, opts...),
		AuditLog:             newAuditLog(db, opts...),
		Circuit:              newCircuit(db, opts...),
		CliCredential:        newCliCredential(db, opts...),
		Device:               newDevice(db, opts...),
		DeviceConfig:         newDeviceConfig(db, opts...),
		DeviceInterface:      newDeviceInterface(db, opts...),
		DeviceStack:          newDeviceStack(db, opts...),
		IpAddress:            newIpAddress(db, opts...),
		LLDPNeighbor:         newLLDPNeighbor(db, opts...),
		MacAddress:           newMacAddress(db, opts...),
		Maintenance:          newMaintenance(db, opts...),
		Menu:                 newMenu(db, opts...),
		Organization:         newOrganization(db, opts...),
		Permission:           newPermission(db, opts...),
		Prefix:               newPrefix(db, opts...),
		Proxy:                newProxy(db, opts...),
		Rack:                 newRack(db, opts...),
		RestconfCredential:   newRestconfCredential(db, opts...),
		Role:                 newRole(db, opts...),
		RootCause:            newRootCause(db, opts...),
		ScanDevice:           newScanDevice(db, opts...),
		Server:               newServer(db, opts...),
		ServerCredential:     newServerCredential(db, opts...),
		ServerSnmpCredential: newServerSnmpCredential(db, opts...),
		Site:                 newSite(db, opts...),
		SnmpV2Credential:     newSnmpV2Credential(db, opts...),
		Subscription:         newSubscription(db, opts...),
		SubscriptionRecord:   newSubscriptionRecord(db, opts...),
		TaskResult:           newTaskResult(db, opts...),
		Template:             newTemplate(db, opts...),
		User:                 newUser(db, opts...),
	}
}

type Query struct {
	db *gorm.DB

	AP                   aP
	Alert                alert
	AlertActionLog       alertActionLog
	AlertGroup           alertGroup
	ApLLDPNeighbor       apLLDPNeighbor
	AuditLog             auditLog
	Circuit              circuit
	CliCredential        cliCredential
	Device               device
	DeviceConfig         deviceConfig
	DeviceInterface      deviceInterface
	DeviceStack          deviceStack
	IpAddress            ipAddress
	LLDPNeighbor         lLDPNeighbor
	MacAddress           macAddress
	Maintenance          maintenance
	Menu                 menu
	Organization         organization
	Permission           permission
	Prefix               prefix
	Proxy                proxy
	Rack                 rack
	RestconfCredential   restconfCredential
	Role                 role
	RootCause            rootCause
	ScanDevice           scanDevice
	Server               server
	ServerCredential     serverCredential
	ServerSnmpCredential serverSnmpCredential
	Site                 site
	SnmpV2Credential     snmpV2Credential
	Subscription         subscription
	SubscriptionRecord   subscriptionRecord
	TaskResult           taskResult
	Template             template
	User                 user
}

func (q *Query) Available() bool { return q.db != nil }

func (q *Query) clone(db *gorm.DB) *Query {
	return &Query{
		db:                   db,
		AP:                   q.AP.clone(db),
		Alert:                q.Alert.clone(db),
		AlertActionLog:       q.AlertActionLog.clone(db),
		AlertGroup:           q.AlertGroup.clone(db),
		ApLLDPNeighbor:       q.ApLLDPNeighbor.clone(db),
		AuditLog:             q.AuditLog.clone(db),
		Circuit:              q.Circuit.clone(db),
		CliCredential:        q.CliCredential.clone(db),
		Device:               q.Device.clone(db),
		DeviceConfig:         q.DeviceConfig.clone(db),
		DeviceInterface:      q.DeviceInterface.clone(db),
		DeviceStack:          q.DeviceStack.clone(db),
		IpAddress:            q.IpAddress.clone(db),
		LLDPNeighbor:         q.LLDPNeighbor.clone(db),
		MacAddress:           q.MacAddress.clone(db),
		Maintenance:          q.Maintenance.clone(db),
		Menu:                 q.Menu.clone(db),
		Organization:         q.Organization.clone(db),
		Permission:           q.Permission.clone(db),
		Prefix:               q.Prefix.clone(db),
		Proxy:                q.Proxy.clone(db),
		Rack:                 q.Rack.clone(db),
		RestconfCredential:   q.RestconfCredential.clone(db),
		Role:                 q.Role.clone(db),
		RootCause:            q.RootCause.clone(db),
		ScanDevice:           q.ScanDevice.clone(db),
		Server:               q.Server.clone(db),
		ServerCredential:     q.ServerCredential.clone(db),
		ServerSnmpCredential: q.ServerSnmpCredential.clone(db),
		Site:                 q.Site.clone(db),
		SnmpV2Credential:     q.SnmpV2Credential.clone(db),
		Subscription:         q.Subscription.clone(db),
		SubscriptionRecord:   q.SubscriptionRecord.clone(db),
		TaskResult:           q.TaskResult.clone(db),
		Template:             q.Template.clone(db),
		User:                 q.User.clone(db),
	}
}

func (q *Query) ReadDB() *Query {
	return q.ReplaceDB(q.db.Clauses(dbresolver.Read))
}

func (q *Query) WriteDB() *Query {
	return q.ReplaceDB(q.db.Clauses(dbresolver.Write))
}

func (q *Query) ReplaceDB(db *gorm.DB) *Query {
	return &Query{
		db:                   db,
		AP:                   q.AP.replaceDB(db),
		Alert:                q.Alert.replaceDB(db),
		AlertActionLog:       q.AlertActionLog.replaceDB(db),
		AlertGroup:           q.AlertGroup.replaceDB(db),
		ApLLDPNeighbor:       q.ApLLDPNeighbor.replaceDB(db),
		AuditLog:             q.AuditLog.replaceDB(db),
		Circuit:              q.Circuit.replaceDB(db),
		CliCredential:        q.CliCredential.replaceDB(db),
		Device:               q.Device.replaceDB(db),
		DeviceConfig:         q.DeviceConfig.replaceDB(db),
		DeviceInterface:      q.DeviceInterface.replaceDB(db),
		DeviceStack:          q.DeviceStack.replaceDB(db),
		IpAddress:            q.IpAddress.replaceDB(db),
		LLDPNeighbor:         q.LLDPNeighbor.replaceDB(db),
		MacAddress:           q.MacAddress.replaceDB(db),
		Maintenance:          q.Maintenance.replaceDB(db),
		Menu:                 q.Menu.replaceDB(db),
		Organization:         q.Organization.replaceDB(db),
		Permission:           q.Permission.replaceDB(db),
		Prefix:               q.Prefix.replaceDB(db),
		Proxy:                q.Proxy.replaceDB(db),
		Rack:                 q.Rack.replaceDB(db),
		RestconfCredential:   q.RestconfCredential.replaceDB(db),
		Role:                 q.Role.replaceDB(db),
		RootCause:            q.RootCause.replaceDB(db),
		ScanDevice:           q.ScanDevice.replaceDB(db),
		Server:               q.Server.replaceDB(db),
		ServerCredential:     q.ServerCredential.replaceDB(db),
		ServerSnmpCredential: q.ServerSnmpCredential.replaceDB(db),
		Site:                 q.Site.replaceDB(db),
		SnmpV2Credential:     q.SnmpV2Credential.replaceDB(db),
		Subscription:         q.Subscription.replaceDB(db),
		SubscriptionRecord:   q.SubscriptionRecord.replaceDB(db),
		TaskResult:           q.TaskResult.replaceDB(db),
		Template:             q.Template.replaceDB(db),
		User:                 q.User.replaceDB(db),
	}
}

type queryCtx struct {
	AP                   IAPDo
	Alert                IAlertDo
	AlertActionLog       IAlertActionLogDo
	AlertGroup           IAlertGroupDo
	ApLLDPNeighbor       IApLLDPNeighborDo
	AuditLog             IAuditLogDo
	Circuit              ICircuitDo
	CliCredential        ICliCredentialDo
	Device               IDeviceDo
	DeviceConfig         IDeviceConfigDo
	DeviceInterface      IDeviceInterfaceDo
	DeviceStack          IDeviceStackDo
	IpAddress            IIpAddressDo
	LLDPNeighbor         ILLDPNeighborDo
	MacAddress           IMacAddressDo
	Maintenance          IMaintenanceDo
	Menu                 IMenuDo
	Organization         IOrganizationDo
	Permission           IPermissionDo
	Prefix               IPrefixDo
	Proxy                IProxyDo
	Rack                 IRackDo
	RestconfCredential   IRestconfCredentialDo
	Role                 IRoleDo
	RootCause            IRootCauseDo
	ScanDevice           IScanDeviceDo
	Server               IServerDo
	ServerCredential     IServerCredentialDo
	ServerSnmpCredential IServerSnmpCredentialDo
	Site                 ISiteDo
	SnmpV2Credential     ISnmpV2CredentialDo
	Subscription         ISubscriptionDo
	SubscriptionRecord   ISubscriptionRecordDo
	TaskResult           ITaskResultDo
	Template             ITemplateDo
	User                 IUserDo
}

func (q *Query) WithContext(ctx context.Context) *queryCtx {
	return &queryCtx{
		AP:                   q.AP.WithContext(ctx),
		Alert:                q.Alert.WithContext(ctx),
		AlertActionLog:       q.AlertActionLog.WithContext(ctx),
		AlertGroup:           q.AlertGroup.WithContext(ctx),
		ApLLDPNeighbor:       q.ApLLDPNeighbor.WithContext(ctx),
		AuditLog:             q.AuditLog.WithContext(ctx),
		Circuit:              q.Circuit.WithContext(ctx),
		CliCredential:        q.CliCredential.WithContext(ctx),
		Device:               q.Device.WithContext(ctx),
		DeviceConfig:         q.DeviceConfig.WithContext(ctx),
		DeviceInterface:      q.DeviceInterface.WithContext(ctx),
		DeviceStack:          q.DeviceStack.WithContext(ctx),
		IpAddress:            q.IpAddress.WithContext(ctx),
		LLDPNeighbor:         q.LLDPNeighbor.WithContext(ctx),
		MacAddress:           q.MacAddress.WithContext(ctx),
		Maintenance:          q.Maintenance.WithContext(ctx),
		Menu:                 q.Menu.WithContext(ctx),
		Organization:         q.Organization.WithContext(ctx),
		Permission:           q.Permission.WithContext(ctx),
		Prefix:               q.Prefix.WithContext(ctx),
		Proxy:                q.Proxy.WithContext(ctx),
		Rack:                 q.Rack.WithContext(ctx),
		RestconfCredential:   q.RestconfCredential.WithContext(ctx),
		Role:                 q.Role.WithContext(ctx),
		RootCause:            q.RootCause.WithContext(ctx),
		ScanDevice:           q.ScanDevice.WithContext(ctx),
		Server:               q.Server.WithContext(ctx),
		ServerCredential:     q.ServerCredential.WithContext(ctx),
		ServerSnmpCredential: q.ServerSnmpCredential.WithContext(ctx),
		Site:                 q.Site.WithContext(ctx),
		SnmpV2Credential:     q.SnmpV2Credential.WithContext(ctx),
		Subscription:         q.Subscription.WithContext(ctx),
		SubscriptionRecord:   q.SubscriptionRecord.WithContext(ctx),
		TaskResult:           q.TaskResult.WithContext(ctx),
		Template:             q.Template.WithContext(ctx),
		User:                 q.User.WithContext(ctx),
	}
}

func (q *Query) Transaction(fc func(tx *Query) error, opts ...*sql.TxOptions) error {
	return q.db.Transaction(func(tx *gorm.DB) error { return fc(q.clone(tx)) }, opts...)
}

func (q *Query) Begin(opts ...*sql.TxOptions) *QueryTx {
	tx := q.db.Begin(opts...)
	return &QueryTx{Query: q.clone(tx), Error: tx.Error}
}

type QueryTx struct {
	*Query
	Error error
}

func (q *QueryTx) Commit() error {
	return q.db.Commit().Error
}

func (q *QueryTx) Rollback() error {
	return q.db.Rollback().Error
}

func (q *QueryTx) SavePoint(name string) error {
	return q.db.SavePoint(name).Error
}

func (q *QueryTx) RollbackTo(name string) error {
	return q.db.RollbackTo(name).Error
}
