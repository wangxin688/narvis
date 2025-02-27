package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/imroc/req/v3"
	rabbithole "github.com/michaelklishin/rabbit-hole/v2"
	"github.com/wangxin688/narvis/intend/logger"
	"github.com/wangxin688/narvis/server/cmd/bootstrap/asset"
	"github.com/wangxin688/narvis/server/cmd/bootstrap/mock"
	"github.com/wangxin688/narvis/server/config"
	"github.com/wangxin688/narvis/server/dal/gen"
	infra_hooks "github.com/wangxin688/narvis/server/features/infra/hooks"
	"github.com/wangxin688/narvis/server/features/organization/biz"
	"github.com/wangxin688/narvis/server/features/organization/hooks"
	"github.com/wangxin688/narvis/server/features/organization/schemas"
	"github.com/wangxin688/narvis/server/infra"
	"github.com/wangxin688/narvis/server/models"
	"github.com/wangxin688/narvis/server/pkg/contextvar"
	"github.com/wangxin688/narvis/server/pkg/zbx"
	"github.com/wangxin688/narvis/server/pkg/zbx/zschema"
	"gopkg.in/yaml.v2"

	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	var err error
	defer func() {
		if err != nil {
			panic(err)
		}
	}()
	config.InitConfig()
	config.InitLogger()
	db := connectDB()
	err = infra.AutoMigration(db)
	gen.SetDefault(db)
	initMacAddress()
	initZbx()
	initRabbitMQ()
	if err != nil {
		return
	}
	if config.Settings.Env == "on_prem" || config.Settings.Env == "dev" {
		orgId := initOrganization()
		if config.Settings.Env == "dev" {
			mock.GenerateMockData(orgId, db)
			logger.Logger.Info("[bootstrap]: mock data generate successfully.")
		}
		config.InitConfig()
		initProxy(orgId)
		initNarvisCliCredential(orgId)        //nolint: errcheck
		initNarvisSnmpCredential(orgId)       //nolint: errcheck
		initNarvisServerCredential(orgId)     //nolint: errcheck
		initNarvisServerSnmpCredential(orgId) //nolint: errcheck

	}
	config.InitConfig()
	err = initZbxTemplates()
	err = initNarvisTemplates()
	logger.Logger.Info("[bootstrap]: bootstrap completed")
}

func connectDB() *gorm.DB {
	dsn := config.Settings.Postgres.BuildPgDsn()
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		logger.Logger.Fatal("[bootstrap]: failed to connect database", zap.Error(err))
	}
	err = db.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\"").Error
	if err != nil {
		logger.Logger.Fatal("[infraConnectDb]: failed to create extension: %v", zap.Error(err))
	}
	return db
}

func initOrganization() string {
	service := biz.NewOrganizationService()
	if config.Settings.BootstrapConfig.EnterpriseCode == "" {
		logger.Logger.Error("[bootstrap]: enterprise code is empty")
		panic("enterprise code is empty")
	}

	org, err := gen.Organization.Where(gen.Organization.EnterpriseCode.Eq(config.Settings.BootstrapConfig.EnterpriseCode)).Find()
	if err != nil {
		logger.Logger.Error("[bootstrap]: failed to get organization", zap.Error(err))
		panic(err)
	}

	if len(org) > 0 {
		contextvar.OrganizationId.Set(org[0].Id)
		logger.Logger.Info("[bootstrap]: organization already exists", zap.String("id", org[0].Id))
		return org[0].Id
	}

	newOrg, err := service.CreateOrganization(&schemas.OrganizationCreate{
		Name:           config.Settings.BootstrapConfig.Organization,
		EnterpriseCode: config.Settings.BootstrapConfig.EnterpriseCode,
		DomainName:     config.Settings.BootstrapConfig.DomainName,
		Active:         true,
		LicenseCount:   100000,
		AuthType:       0,
		AdminPassword:  config.Settings.BootstrapConfig.AdminPassword,
	})
	if err != nil {
		logger.Logger.Error("[bootstrap]: failed to create organization", zap.Error(err))
		panic(err)
	}
	contextvar.OrganizationId.Set(newOrg.Id)
	logger.Logger.Info("[bootstrap]: organization created", zap.String("id", newOrg.Id))
	return newOrg.Id
}

func initProxy(orgId string) {
	dbProxy, err := gen.Proxy.Where(gen.Proxy.Name.Eq(config.Settings.BootstrapConfig.EnterpriseCode)).Find()
	if err != nil {
		logger.Logger.Error("[bootstrap]: failed to get proxy", zap.Error(err))
		panic(err)
	}
	if len(dbProxy) > 0 {
		logger.Logger.Info("[bootstrap]: proxy already exists", zap.String("id", dbProxy[0].Id))
		for _, proxy := range dbProxy {
			if proxy.Active && proxy.ProxyId == nil {
				hooks.CreateZbxProxy(proxy)
			} else {
				logger.Logger.Info("[bootstrap]: monitor proxy already exists", zap.String("proxyId", *proxy.ProxyId))
			}
		}
	} else {
		proxy, err := biz.NewProxyService().CreateProxy(&schemas.ProxyCreate{
			OrganizationId: orgId,
			Name:           config.Settings.BootstrapConfig.EnterpriseCode,
			Active:         true,
		})
		if err != nil {
			logger.Logger.Error("[bootstrap]: failed to create proxy", zap.Error(err))
			panic(err)
		}
		logger.Logger.Info("[bootstrap]: proxy created", zap.String("id", proxy.Id))
		hooks.CreateZbxProxy(proxy)
	}
}

func initMacAddress() {
	mac, err := gen.MacAddress.Count()
	if err != nil {
		logger.Logger.Error("[bootstrap]: failed to get mac address", zap.Error(err))
		panic(err)
	}
	if mac >= 1 {
		logger.Logger.Info("[bootstrap]: mac address already exists")
		return
	}
	var macAddresses []*models.MacAddress
	macAddressesBytes, err := asset.Asset("appdata/mac_address.json")
	if err != nil {
		logger.Logger.Error("[bootstrap]: failed to get mac address", zap.Error(err))
		panic(err)
	}
	err = json.Unmarshal(macAddressesBytes, &macAddresses)
	if err != nil {
		logger.Logger.Error("[bootstrap]: failed to decode mac address file", zap.Error(err))
		panic(err)
	}
	err = gen.MacAddress.CreateInBatches(macAddresses, 100)
	if err != nil {
		logger.Logger.Error("[bootstrap]: failed to create mac address", zap.Error(err))
		panic(err)
	}
}

func initRabbitMQ() {
	url := fmt.Sprintf("http://%s:%d", config.Settings.RabbitMQ.Host, config.Settings.RabbitMQ.Port)
	client, err := rabbithole.NewClient(
		url,
		config.Settings.RabbitMQ.Username,
		config.Settings.RabbitMQ.Password,
	)
	if err != nil {
		logger.Logger.Error("[bootstrap]: failed to connect rabbitmq", zap.Error(err))
		return
	}
	vhosts, err := client.ListVhosts()
	if err != nil {
		logger.Logger.Error("[bootstrap]: failed to get rabbitmq vhosts", zap.Error(err))
		return
	}
	if len(vhosts) >= 2 {
		logger.Logger.Info("[bootstrap]: rabbitmq vhost already exists")
		return
	}
	_, err = client.PutVhost("server", rabbithole.VhostSettings{Description: "narvis server vhost"})
	if err != nil {
		logger.Logger.Error("[bootstrap]: failed to create rabbitmq vhost", zap.Error(err))
	}
	logger.Logger.Info("[bootstrap]: rabbitmq vhost server created")
	_, err = client.PutVhost("proxy", rabbithole.VhostSettings{Description: "narvis monitor vhost"})
	if err != nil {
		logger.Logger.Error("[bootstrap]: failed to create rabbitmq vhost", zap.Error(err))
	}
	logger.Logger.Info("[bootstrap]: rabbitmq vhost proxy created")

	users, err := client.ListUsers()
	if err != nil {
		logger.Logger.Error("[bootstrap]: failed to list rabbitmq users", zap.Error(err))
		return
	}
	if len(users) >= 2 {
		logger.Logger.Info("[bootstrap]: rabbitmq user already exists")
		return
	}
	_, err = client.PutUser("narvis-proxy", rabbithole.UserSettings{
		Name:     "narvis-proxy",
		Password: "851b090b967a89f802e72a0baf1d230e",
	})
	if err != nil {
		logger.Logger.Error("[bootstrap]: failed to create rabbitmq user", zap.Error(err))
	}
	logger.Logger.Info("[bootstrap]: rabbitmq proxy user created")
	_, err = client.UpdatePermissionsIn("proxy", "narvis-server", rabbithole.Permissions{
		Read:      ".*",
		Write:     ".*",
		Configure: ".*",
	})
	if err != nil {
		logger.Logger.Error("[bootstrap]: failed to update rabbitmq user permissions", zap.Error(err))
	}
	logger.Logger.Info("[bootstrap]: rabbitmq server user permissions updated")
	_, err = client.UpdatePermissionsIn("proxy", "narvis-proxy", rabbithole.Permissions{
		Read:      ".*",
		Write:     ".*",
		Configure: ".*",
	})
	if err != nil {
		logger.Logger.Error("[bootstrap]: failed to update rabbitmq proxy user permissions", zap.Error(err))
	}
	logger.Logger.Info("[bootstrap]: rabbitmq proxy user permissions updated")

	logger.Logger.Info("[bootstrap]: rabbitmq initialized")
}

func initZbx() {
	client := getClient()
	if client == nil {
		return
	}
	mediaTypeId, err := initZbxMediaType(client)
	if err != nil {
		logger.Logger.Info("[bootstrap]: failed to init media type", zap.Error(err))
		return
	}
	initZbxGlobalMacro(client)
	superUserId, err := initZbxSuperUser(client, mediaTypeId)
	if err != nil {
		logger.Logger.Info("[bootstrap]: failed to init monitor super user", zap.Error(err))
		return
	}
	initZbxAction(client, mediaTypeId, superUserId)
	token, err := initZbxSuperToken(client, superUserId)
	if err != nil {
		logger.Logger.Info("[bootstrap]: failed to init monitor super token", zap.Error(err))
		return
	}
	initZbxDisableDefaultAdmin(client, token)
}

func getClient() *req.Client {
	url := config.Settings.Zbx.Url
	client := req.C().SetBaseURL(url).SetCommonHeader("Content-Type", "application/json-rpc")
	login := map[string]any{
		"jsonrpc": "2.0",
		"method":  "user.login",
		"params":  map[string]string{"username": "Admin", "password": "zabbix"},
		"id":      1,
	}
	var loginResponse zschema.LoginResponse
	resp, err := client.R().SetBody(login).SetSuccessResult(&loginResponse).Post("/api_jsonrpc.php")
	if err != nil || !resp.IsSuccessState() || loginResponse.Error != nil {
		logger.Logger.Info("[bootstrap]: monitor has been already init.")
		return nil
	}
	logger.Logger.Info("[bootstrap]: login to monitor success")
	token := loginResponse.Result
	client.SetCommonBearerAuthToken(token)
	return client
}

func initZbxMediaType(client *req.Client) (string, error) {
	mediaType := map[string]any{
		"type":   4,
		"status": "0",
		"name":   "Narvis Alerts",
		"script": `
try {
    var result = {},
        params = JSON.parse(value),
        req = new HttpRequest(),
        fields = {},
        resp;
    req.addHeader("Content-Type: application/json");
    req.addHeader("Authorization: Bearer " + params.token);
    fields.alertName = params.alertName;
    fields.hostId = params.hostId;
    fields.itemName = params.itemName;
    fields.opData = params.opData;
    fields.status = params.status;
    fields.triggerId = params.triggerId;
    fields.url = params.url;
    fields.eventId = params.eventId;
    for (var i in fields) {
        if (fields[i].startsWith('{')) {
            delete fields[i]
        }
    }
	fields.tags = JSON.parse(params.tags);
    if (JSON.stringify(fields) === '{}') {
        result = {}
    } else {
        resp = req.post(fields.url, JSON.stringify(fields));
        Zabbix.log('fields: ' + JSON.stringify(fields));
        if (req.getStatus() != 200) {
            throw 'Response Code: ' + req.getStatus();
        }
        resp = JSON.parse(resp);
        result.code = resp.code;
        result.data = resp.data;
        result.msg = resp.msg;
    }
} catch (error) {
    Zabbix.log(4,
        'send alert to narvis alert failed: eventId: ' + JSON.stringify(fields.eventId));
    result = {};
}
return JSON.stringify(result);
		`,
		"parameters": []map[string]string{
			{"name": "alertName", "value": "{TRIGGER.NAME}"},
			{"name": "hostId", "value": "{HOSTNAME}"},
			{"name": "itemName", "value": "{ITEM.NAME}"},
			{"name": "opData", "value": "{EVENT.OPDATA}"},
			{"name": "status", "value": "{TRIGGER.STATUS}"},
			{"name": "triggerId", "value": "{TRIGGER.ID}"},
			{"name": "eventId", "value": "{EVENT.ID}"},
			{"name": "token", "value": "{$XAUTHTOKEN}"},
			{"name": "tags", "value": "{EVENT.TAGSJSON}"},
			{"name": "url", "value": fmt.Sprintf("%s/api/v1/alert/alerts", config.Settings.System.BaseUrl)},
		},
	}
	mediaTypeBody := map[string]any{
		"jsonrpc": "2.0", "method": "mediatype.create", "params": mediaType, "id": 1,
	}
	type Resp struct {
		JsonRpc string `json:"jsonrpc"`
		Result  struct {
			Mediatypeids []string `json:"mediatypeids"`
		} `json:"result"`
		ID uint64 `json:"id"`
	}
	mediaTypeResponse := new(Resp)
	newAlertMedia, err := client.R().SetBody(mediaTypeBody).SetSuccessResult(&mediaTypeResponse).Post("/api_jsonrpc.php")
	fmt.Printf("%v", mediaTypeResponse)
	if err != nil || !newAlertMedia.IsSuccessState() {
		logger.Logger.Error("[bootstrap]: init monitor failed to create alert media type", zap.Error(err))
		return "", err
	}
	newMediaTypeId := mediaTypeResponse.Result.Mediatypeids[0]
	logger.Logger.Info(fmt.Sprintf("[bootstrap]: init monitor create new media type %s success", newMediaTypeId))
	return newMediaTypeId, nil
}

func initZbxGlobalMacro(client *req.Client) error {
	newMacroBody := map[string]any{
		"jsonrpc": "2.0",
		"method":  "usermacro.createglobal",
		"params": map[string]any{
			"macro": "{$XAUTHTOKEN}",
			"value": config.Settings.Jwt.PublicAuthKey,
			"type":  1,
		},
		"id": 1,
	}
	resp, err := client.R().SetBody(newMacroBody).Post("/api_jsonrpc.php")
	if err != nil {
		logger.Logger.Info("[bootstrap]: failed to create new global macro", zap.Error(err))
		return err
	}
	if resp.IsErrorState() {
		logger.Logger.Info("[bootstrap]: failed to create new global macro")
	}
	logger.Logger.Info("[bootstrap]: create new global macro success")
	return nil
}

func initZbxSuperUser(client *req.Client, mediaTypeId string) (string, error) {
	superUser := map[string]any{
		"username": "narvis",
		"passwd":   "50a8c8858b1ddca756db990053830303",
		"roleid":   "3",
		"usrgrps":  []map[string]string{{"usrgrpid": "7"}},
		"medias": []map[string]any{
			{
				"mediatypeid": mediaTypeId,
				"sendto":      []string{"narvis"},
				"active":      0,
				"severity":    63,
				"period":      "1-7,00:00-24:00",
			},
		},
	}
	superUserBody := map[string]any{
		"jsonrpc": "2.0", "method": "user.create", "params": superUser, "id": 1,
	}
	type Resp struct {
		JsonRpc string `json:"jsonrpc"`
		Result  struct {
			Userids []string `json:"userids"`
		} `json:"result"`
		ID uint64 `json:"id"`
	}
	superUserResponse := new(Resp)
	newSuperUser, err := client.R().SetBody(superUserBody).SetSuccessResult(&superUserResponse).Post("/api_jsonrpc.php")
	if err != nil || !newSuperUser.IsSuccessState() {
		logger.Logger.Error("[bootstrap]: init monitor failed to create super user", zap.Error(err))
		return "", err
	}
	newSuperUserId := superUserResponse.Result.Userids[0]
	logger.Logger.Info(fmt.Sprintf("[bootstrap]: init monitor create new super user %s success", newSuperUserId))
	return newSuperUserId, nil
}

func initZbxAction(client *req.Client, mediaTypeId string, superUserId string) (string, error) {
	newAction := map[string]any{
		"name":               "Narvis Alert",
		"eventsource":        "0",
		"status":             "0",
		"esc_period":         "1h",
		"pause_suppressed":   "1",
		"notify_if_canceled": "1",
		"pause_symptoms":     "1",
		"filter": map[string]any{
			"evaltype": "0",
			"conditions": []map[string]any{
				{"conditiontype": "16", "operator": "11"},
				{"conditiontype": "0", "operator": "1", "value": "4"},
			},
		},
		"operations": []map[string]any{
			{
				"operationtype": "0",
				"esc_period":    "20m",
				"esc_step_from": "1",
				"esc_step_to":   "0",
				"evaltype":      "0",
				"opconditions":  []map[string]any{{"conditiontype": "14", "operator": "0", "value": "0"}},
				"opmessage": map[string]any{
					"default_msg": "0",
					"subject":     "Problem: {EVENT.NAME}",
					"message":     "",
					"mediatypeid": mediaTypeId,
				},
				"opmessage_grp": []map[string]any{},
				"opmessage_usr": []map[string]any{{"userid": superUserId}},
			},
		},
		"recovery_operations": []map[string]any{
			{
				"operationtype": "0",
				"opmessage": map[string]any{
					"default_msg": "0",
					"subject":     "Resolved: {EVENT.NAME}",
					"message":     "",
					"mediatypeid": mediaTypeId,
				},
				"opmessage_grp": []map[string]any{},
				"opmessage_usr": []map[string]any{{"userid": superUserId}},
			},
		},
	}

	newActionBody := map[string]any{
		"jsonrpc": "2.0",
		"method":  "action.create",
		"params":  newAction,
		"id":      1,
	}
	type Resp struct {
		JsonRpc string `json:"jsonrpc"`
		Result  struct {
			ActionIds []string `json:"actionids"`
		} `json:"result"`
		ID uint64 `json:"id"`
	}
	newActionResponse := new(Resp)
	resp, err := client.R().SetBody(newActionBody).SetSuccessResult(&newActionResponse).Post("/api_jsonrpc.php")
	if err != nil || !resp.IsSuccessState() {
		logger.Logger.Error("[bootstrap]: init monitor failed to create new action", zap.Error(err))
		return "", err
	}
	newActionId := newActionResponse.Result.ActionIds[0]
	logger.Logger.Info(fmt.Sprintf("[bootstrap]: init monitor create new action %s success", newActionId))
	return newActionId, nil
}

func initZbxSuperToken(client *req.Client, superUserId string) (string, error) {
	newTokenBody := map[string]any{
		"jsonrpc": "2.0",
		"method":  "token.create",
		"params":  map[string]any{"name": "narvis token", "userid": superUserId},
		"id":      1,
	}
	type Resp struct {
		JsonRpc string `json:"jsonrpc"`
		Result  struct {
			TokenIds []string `json:"tokenids"`
		} `json:"result"`
		ID uint64 `json:"id"`
	}
	newTokenResponse := new(Resp)
	newToken, err := client.R().SetBody(newTokenBody).SetSuccessResult(&newTokenResponse).Post("/api_jsonrpc.php")
	if err != nil || !newToken.IsSuccessState() {
		logger.Logger.Error("[bootstrap]: init monitor failed to create super token", zap.Error(err))
		return "", err
	}
	newTokenId := newTokenResponse.Result.TokenIds[0]
	nowTokenGenerateBody := map[string]any{
		"jsonrpc": "2.0",
		"method":  "token.generate",
		"params":  []string{newTokenId},
		"id":      1,
	}
	type NowTokenGenerateResp struct {
		JsonRpc string `json:"jsonrpc"`
		Result  []struct {
			Token string `json:"token"`
		} `json:"result"`
		ID uint64 `json:"id"`
	}
	nowTokenGenerateResponse := new(NowTokenGenerateResp)
	nowTokenGenerate, err := client.R().SetBody(nowTokenGenerateBody).SetSuccessResult(&nowTokenGenerateResponse).Post("/api_jsonrpc.php")
	if err != nil || !nowTokenGenerate.IsSuccessState() {
		logger.Logger.Error("[bootstrap]: init_monitor failed to create super token", zap.Error(err))
		return "", err
	}
	nowToken := nowTokenGenerateResponse.Result[0].Token
	logger.Logger.Info(fmt.Sprintf("[bootstrap]: init monitor create new super token %s success", nowToken))
	writeTokenToYamlConfig(nowToken)
	return nowToken, nil
}

func writeTokenToYamlConfig(token string) {
	path, err := os.Getwd()
	if err != nil {
		logger.Logger.Error("[bootstrap]: failed to get current path", zap.Error(err))
		return
	}
	configPath := filepath.Join(path, "config.yaml")
	yamlConfig, err := os.ReadFile(configPath)
	if err != nil {
		logger.Logger.Error("[bootstrap]: failed to read config file", zap.Error(err))
		return
	}
	var config config.Config
	err = yaml.Unmarshal(yamlConfig, &config)
	if err != nil {
		logger.Logger.Error("[bootstrap]: failed to unmarshal config file", zap.Error(err))
		return
	}
	config.Zbx.Token = token
	yamlConfig, err = yaml.Marshal(&config)
	if err != nil {
		logger.Logger.Error("[bootstrap]: failed to marshal config file", zap.Error(err))
		return
	}
	err = os.WriteFile("config.yaml", yamlConfig, 0644)
	if err != nil {
		logger.Logger.Error("[bootstrap]: failed to write config file", zap.Error(err))
		return
	}
}

func initZbxDisableDefaultAdmin(client *req.Client, token string) error {
	client.SetCommonBearerAuthToken(token)
	updateBody := map[string]any{
		"jsonrpc": "2.0",
		"method":  "user.update",
		"params": map[string]any{
			"userid":  "1",
			"usrgrps": []map[string]string{{"usrgrpid": "9"}},
		},
		"id": 1,
	}
	resp, err := client.R().SetBody(updateBody).Post("/api_jsonrpc.php")
	if err != nil || !resp.IsSuccessState() {
		logger.Logger.Error("[bootstrap]: init monitor failed to disable default admin", zap.Error(err))
		return err
	}
	logger.Logger.Info("[bootstrap]: init monitor disable default admin success")
	return nil
}

func initNarvisSnmpCredential(orgId string) error {
	cred, err := gen.SnmpV2Credential.Where(gen.SnmpV2Credential.OrganizationId.Eq(orgId)).Find()
	if err != nil {
		logger.Logger.Error("[bootstrap]: failed to get snmp credential", zap.Error(err))
		return err
	}
	if len(cred) > 0 {
		logger.Logger.Info("[bootstrap]: snmp credential already exists", zap.String("id", cred[0].Id))
		for _, cr := range cred {
			if cr.DeviceId == nil && cr.GlobalMacroId == nil {
				snmpCred := &models.SnmpV2Credential{
					OrganizationId: orgId,
					Community:      config.Settings.BootstrapConfig.SnmpCommunity,
					MaxRepetitions: 50,
					Timeout:        config.Settings.BootstrapConfig.SnmpTimeout,
					Port:           config.Settings.BootstrapConfig.SnmpPort,
				}
				err = gen.SnmpV2Credential.Create(snmpCred)
				if err != nil {
					logger.Logger.Error("[bootstrap]: failed to create snmp credential", zap.Error(err))
					return err
				}
				logger.Logger.Info("[bootstrap]: snmp credential created", zap.String("id", snmpCred.Id))
				infra_hooks.SnmpCredCreateHooks(snmpCred.Id)
				break
			}
		}
	} else {
		snmpCred := &models.SnmpV2Credential{
			OrganizationId: orgId,
			Community:      config.Settings.BootstrapConfig.SnmpCommunity,
			MaxRepetitions: 50,
			Timeout:        config.Settings.BootstrapConfig.SnmpTimeout,
			Port:           config.Settings.BootstrapConfig.SnmpPort,
		}
		err = gen.SnmpV2Credential.Create(snmpCred)
		if err != nil {
			logger.Logger.Error("[bootstrap]: failed to create snmp credential", zap.Error(err))
			return err
		}
		logger.Logger.Info("[bootstrap]: snmp credential created", zap.String("id", snmpCred.Id))
		infra_hooks.SnmpCredCreateHooks(snmpCred.Id)
	}
	return nil
}

func initNarvisCliCredential(orgId string) error {
	cred, err := gen.CliCredential.Where(gen.CliCredential.OrganizationId.Eq(orgId)).Find()
	if err != nil {
		logger.Logger.Error("[bootstrap]: failed to get client credential", zap.Error(err))
		return err
	}
	if len(cred) > 0 {
		logger.Logger.Info("[bootstrap]: client credential already exists", zap.String("id", cred[0].Id))
		return nil
	}
	clientCred := &models.CliCredential{
		OrganizationId: orgId,
		Username:       config.Settings.BootstrapConfig.CliUser,
		Password:       config.Settings.BootstrapConfig.CliPassword,
	}
	err = gen.CliCredential.Create(clientCred)
	if err != nil {
		logger.Logger.Error("[bootstrap]: failed to create client credential", zap.Error(err))
		return err
	}
	logger.Logger.Info("[bootstrap]: client credential created", zap.String("id", clientCred.Id))
	return nil
}

func initNarvisServerCredential(orgId string) error {
	cred, err := gen.ServerCredential.Where(gen.ServerCredential.OrganizationId.Eq(orgId)).Find()
	if err != nil {
		logger.Logger.Error("[bootstrap]: failed to get server credential", zap.Error(err))
		return err
	}
	if len(cred) > 0 {
		logger.Logger.Info("[bootstrap]: server credential already exists", zap.String("id", cred[0].Id))
		return nil
	}
	serverCred := &models.ServerCredential{
		OrganizationId: orgId,
		Username:       config.Settings.BootstrapConfig.CliUser,
		Password:       config.Settings.BootstrapConfig.CliPassword,
	}
	err = gen.ServerCredential.Create(serverCred)
	if err != nil {
		logger.Logger.Error("[bootstrap]: failed to create server credential", zap.Error(err))
		return err
	}
	logger.Logger.Info("[bootstrap]: server credential created", zap.String("id", serverCred.Id))
	return nil
}

func initNarvisServerSnmpCredential(orgId string) error {
	cred, err := gen.ServerSnmpCredential.Where(gen.ServerSnmpCredential.OrganizationId.Eq(orgId)).Find()
	if err != nil {
		logger.Logger.Error("[bootstrap]: failed to get server snmp credential", zap.Error(err))
		return err
	}
	if len(cred) > 0 {
		logger.Logger.Info("[bootstrap]: server snmp credential already exists", zap.String("id", cred[0].Id))
		return nil
	}
	serverSnmpCred := &models.ServerSnmpCredential{
		OrganizationId: orgId,
		Community:      config.Settings.BootstrapConfig.SnmpCommunity,
		MaxRepetitions: 50,
		Timeout:        config.Settings.BootstrapConfig.SnmpTimeout,
		Port:           config.Settings.BootstrapConfig.SnmpPort,
	}
	err = gen.ServerSnmpCredential.Create(serverSnmpCred)
	if err != nil {
		logger.Logger.Error("[bootstrap]: failed to create server snmp credential", zap.Error(err))
		return err
	}
	logger.Logger.Info("[bootstrap]: server snmp credential created", zap.String("id", serverSnmpCred.Id))
	return nil
}

func initZbxTemplates() error {
	templates := asset.AssetNames()
	zbxTemplates := make([]string, 0)
	firstImportTemplates := make([]string, 0)
	for _, template := range templates {
		if strings.Contains(template, ".yaml") {
			if strings.Contains(template, "template_interface_basic") || strings.Contains(template, "template_interface_circuit") {
				firstImportTemplates = append(firstImportTemplates, template)
			} else {
				zbxTemplates = append(zbxTemplates, template)
			}
		}
	}
	zbxClient := zbx.NewZbxClient()
	for _, template := range firstImportTemplates {
		templateBytes, err := asset.Asset(template)
		if err != nil {
			logger.Logger.Error("[bootstrap]: failed to read template file", zap.Error(err))
			return err
		}
		_, err = zbxClient.ConfigurationImport(string(templateBytes))
		if err != nil {
			logger.Logger.Error("[bootstrap]: failed to import template", zap.Error(err))
			return err
		}
		logger.Logger.Info("[bootstrap]: init monitor template", zap.String("name", template))
	}
	for _, template := range zbxTemplates {
		templateBytes, err := asset.Asset(template)
		if err != nil {
			logger.Logger.Error("[bootstrap]: failed to read template file", zap.Error(err))
			return err
		}
		_, err = zbxClient.ConfigurationImport(string(templateBytes))
		if err != nil {
			logger.Logger.Error("[bootstrap]: failed to import template", zap.Error(err))
			return err
		}
		logger.Logger.Info("[bootstrap]: init monitor template", zap.String("name", template))
	}
	return nil

}

func initNarvisTemplates() error {
	templateMetaBytes, err := asset.Asset("appdata/template_meta.json")
	if err != nil {
		logger.Logger.Error("[bootstrap]: failed to get template meta file", zap.Error(err))
		return err
	}
	var templates []map[string]string
	err = json.Unmarshal(templateMetaBytes, &templates)
	if err != nil {
		logger.Logger.Error("[bootstrap]: failed to decode template meta file", zap.Error(err))
		return err
	}
	zbxClient := zbx.NewZbxClient()
	for _, template := range templates {
		templateName := template["platform"] + " " + template["deviceRole"]
		if strings.Contains(templateName, "*") {
			templateName = template["basicTemplate"]
		}
		dbTemplate, err := gen.Template.Where(
			gen.Template.TemplateName.Eq(templateName),
		).Find()
		if err != nil {
			logger.Logger.Error("[bootstrap]: failed to get template", zap.Error(err))
			return err
		}
		if len(dbTemplate) > 0 {
			logger.Logger.Info("[bootstrap]: template already exists", zap.String("id", dbTemplate[0].Id))
			continue
		}
		output := "templateid"
		templateId, err := zbxClient.TemplateGet(
			&zschema.TemplateGet{
				Output: &output,
				Filter: &map[string]string{
					"host": template["basicTemplate"],
				},
			},
		)
		if err != nil {
			logger.Logger.Error("[bootstrap]: failed to get template", zap.String("name", template["basicTemplate"]), zap.Error(err))
			return err
		}
		if len(templateId) == 0 {
			logger.Logger.Error("[bootstrap]: failed to get template", zap.String("name", template["basicTemplate"]), zap.Error(err))
			continue
		}
		newdbTemplate := &models.Template{
			TemplateName: templateName,
			DeviceRole:   template["deviceRole"],
			Platform:     template["platform"],
			TemplateId:   templateId[0].TemplateId,
		}
		err = gen.Template.Create(newdbTemplate)
		if err != nil {
			logger.Logger.Error("[bootstrap]: failed to create db template", zap.Error(err))
			return err
		}
		logger.Logger.Info("[bootstrap]: template created", zap.String("id", newdbTemplate.Id))
	}
	return nil
}
