package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/imroc/req/v3"
	rabbithole "github.com/michaelklishin/rabbit-hole/v2"
	"github.com/wangxin688/narvis/server/core"
	"github.com/wangxin688/narvis/server/core/config"
	"github.com/wangxin688/narvis/server/dal/gen"
	"github.com/wangxin688/narvis/server/features/organization/biz"
	"github.com/wangxin688/narvis/server/features/organization/schemas"
	"github.com/wangxin688/narvis/server/global"
	"github.com/wangxin688/narvis/server/models"
	"github.com/wangxin688/narvis/server/pkg/zbx/zschema"
	"gopkg.in/yaml.v2"

	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	initOrganization()
	initMacAddress()
	initZbx()
	initRabbitMQ()
}

func connectDB() *gorm.DB {
	core.SetUpConfig()
	dsn := core.Settings.Postgres.BuildPgDsn()
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		core.Logger.Fatal("[bootstrap]: failed to connect database", zap.Error(err))
	}
	return db
}

func initOrganization() string {
	gen.SetDefault(connectDB())
	core.SetUpLogger()
	service := biz.NewOrganizationService()

	org, err := gen.Organization.Where(gen.Organization.Name.Eq("NarvisDemo")).Find()
	if err != nil {
		core.Logger.Error("[bootstrap]: failed to get organization", zap.Error(err))
		panic(err)
	}

	if len(org) > 0 {
		global.OrganizationId.Set(org[0].Id)
		core.Logger.Info("[bootstrap]: organization already exists", zap.String("id", org[0].Id))
		return org[0].Id
	}

	newOrg, err := service.CreateOrganization(&schemas.OrganizationCreate{
		Name:           "NarvisDemo",
		EnterpriseCode: "narvis-demo",
		DomainName:     "navis-demo@narvis.com",
		Active:         true,
		LicenseCount:   100000,
		AuthType:       0,
		AdminPassword:  "admin123456",
	})
	if err != nil {
		core.Logger.Error("[bootstrap]: failed to create organization", zap.Error(err))
		panic(err)
	}
	global.OrganizationId.Set(newOrg.Id)
	core.Logger.Info("[bootstrap]: organization created", zap.String("id", newOrg.Id))
	return newOrg.Id
}

func initMacAddress() {

	gen.SetDefault(connectDB())
	core.SetUpLogger()
	core.SetUpConfig()
	mac, err := gen.MacAddress.Count()
	if err != nil {
		core.Logger.Error("[bootstrap]: failed to get mac address", zap.Error(err))
		panic(err)
	}
	if mac >= 1 {
		core.Logger.Info("[bootstrap]: mac address already exists")
		return
	}
	macAddressFilePath := core.ProjectPath + "/cmd/bootstrap/appdata/mac_address.json"
	file, err := os.Open(macAddressFilePath)
	if err != nil {
		core.Logger.Error("[bootstrap]: failed to open mac address file", zap.Error(err))
		panic(err)
	}
	defer file.Close()
	var macAddresses []*models.MacAddress
	if err := json.NewDecoder(file).Decode(&macAddresses); err != nil {
		core.Logger.Error("[bootstrap]: failed to decode mac address file", zap.Error(err))
		panic(err)
	}
	err = gen.MacAddress.CreateInBatches(macAddresses, 100)
	if err != nil {
		core.Logger.Error("[bootstrap]: failed to create mac address", zap.Error(err))
		panic(err)
	}
	core.Logger.Info("[bootstrap]: mac address created")
}

func initRabbitMQ() {
	client, err := rabbithole.NewClient(
		"http://localhost:15672",
		"narvis-server",
		"26cc7abbea97a17b9f7860ee0dabb051",
	)
	if err != nil {
		core.Logger.Error("[bootstrap]: failed to connect rabbitmq", zap.Error(err))
		return
	}
	vhosts, err := client.ListVhosts()
	if err != nil {
		core.Logger.Error("[bootstrap]: failed to get rabbitmq vhosts", zap.Error(err))
		return
	}
	if len(vhosts) >= 2 {
		core.Logger.Info("[bootstrap]: rabbitmq vhost already exists")
		return
	} else {
		_, err = client.PutVhost("server", rabbithole.VhostSettings{Description: "narvis server vhost"})
		if err != nil {
			core.Logger.Error("[bootstrap]: failed to create rabbitmq vhost", zap.Error(err))
		}
		core.Logger.Info("[bootstrap]: rabbitmq vhost server created")
		_, err = client.PutVhost("proxy", rabbithole.VhostSettings{Description: "narvis monitor vhost"})
		if err != nil {
			core.Logger.Error("[bootstrap]: failed to create rabbitmq vhost", zap.Error(err))
		}
		core.Logger.Info("[bootstrap]: rabbitmq vhost proxy created")
	}
	users, err := client.ListUsers()
	if err != nil {
		core.Logger.Error("[bootstrap]: failed to list rabbitmq users", zap.Error(err))
		return
	}
	if len(users) >= 2 {
		core.Logger.Info("[bootstrap]: rabbitmq user already exists")
		return
	} else {
		_, err = client.PutUser("narvis-proxy", rabbithole.UserSettings{
			Name:     "narvis-proxy",
			Password: "851b090b967a89f802e72a0baf1d230e",
		})
		if err != nil {
			core.Logger.Error("[bootstrap]: failed to create rabbitmq user", zap.Error(err))
		}
		core.Logger.Info("[bootstrap]: rabbitmq proxy user created")
		_, err = client.UpdatePermissionsIn("proxy", "narvis-server", rabbithole.Permissions{
			Read:      ".*",
			Write:     ".*",
			Configure: ".*",
		})
		if err != nil {
			core.Logger.Error("[bootstrap]: failed to update rabbitmq user permissions", zap.Error(err))
		}
		core.Logger.Info("[bootstrap]: rabbitmq server user permissions updated")
		_, err = client.UpdatePermissionsIn("proxy", "narvis-proxy", rabbithole.Permissions{
			Read:      ".*",
			Write:     ".*",
			Configure: ".*",
		})
		if err != nil {
			core.Logger.Error("[bootstrap]: failed to update rabbitmq proxy user permissions", zap.Error(err))
		}
		core.Logger.Info("[bootstrap]: rabbitmq proxy user permissions updated")
	}
	core.Logger.Info("[bootstrap]: rabbitmq initialized")
}

func initZbx() {
	client := getClient()
	if client == nil {
		return
	}
	medieTypeId, err := initZbxMediaType(client)
	if err != nil {
		core.Logger.Info("[bootstrap]: failed to init zbx media type", zap.Error(err))
		return
	}
	initZbxGlobalMacro(client)
	superUserId, err := initZbxSuperUser(client, medieTypeId)
	if err != nil {
		core.Logger.Info("[bootstrap]: failed to init zbx super user", zap.Error(err))
		return
	}
	initZbxAction(client, medieTypeId, superUserId)
	token, err := initZbxSuperToken(client, superUserId)
	if err != nil {
		core.Logger.Info("[bootstrap]: failed to init zbx super token", zap.Error(err))
		return
	}
	initZbxDisableDefaultAdmin(client, token)
}

func getClient() *req.Client {
	url := core.Settings.Zbx.Url
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
		core.Logger.Info("[bootstrap]: zbx has been already init.")
		return nil
	}
	core.Logger.Info("[bootstrap]: login to zabbix success")
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
    fields.HostId = params.HostId;
    fields.itemName = params.itemName;
    fields.opData = params.opData;
    fields.status = params.status;
    fields.triggerId = params.triggerId;
    fields.url = params.url;
    fields.eventId =
        params.eventId;
    fields.tags = params.tags;
    for (var i in fields) {
        if (fields[i].startsWith('{')) {
            delete fields[i]
        }
    }
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
			{"name": "HostId", "value": "{HOSTNAME}"},
			{"name": "itemName", "value": "{ITEM.NAME}"},
			{"name": "opData", "value": "{EVENT.OPDATA}"},
			{"name": "status", "value": "{TRIGGER.STATUS}"},
			{"name": "triggerId", "value": "{TRIGGER.ID}"},
			{"name": "eventId", "value": "{EVENT.ID}"},
			{"name": "token", "value": "{$XAUTHTOKEN}"},
			{"name": "tags", "value": "{EVENT.TAGSJSON}"},
			{"name": "url", "value": fmt.Sprintf("%s/api/v1/alert/alerts", core.Settings.System.BaseUrl)},
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
		core.Logger.Error("[bootstrap]: init_zbx failed to create alert media type", zap.Error(err))
		return "", err
	}
	newMediaTypeId := mediaTypeResponse.Result.Mediatypeids[0]
	core.Logger.Info(fmt.Sprintf("[bootstrap]: init_zbx create new media type %s success", newMediaTypeId))
	return newMediaTypeId, nil
}

func initZbxGlobalMacro(client *req.Client) error {
	newMacroBody := map[string]any{
		"jsonrpc": "2.0",
		"method":  "usermacro.createglobal",
		"params": map[string]any{
			"macro": "{$XAUTHTOKEN}",
			"value": core.Settings.Jwt.PublicAuthKey,
			"type":  1,
		},
		"id": 1,
	}
	resp, err := client.R().SetBody(newMacroBody).Post("/api_jsonrpc.php")
	if err != nil {
		core.Logger.Info("[bootstrap]: failed to create new global macro", zap.Error(err))
		return err
	}
	if resp.IsErrorState() {
		core.Logger.Info("[bootstrap]: failed to create new global macro")
	}
	core.Logger.Info("[bootstrap]: create new global macro success")
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
		core.Logger.Error("[bootstrap]: init_zbx failed to create super user", zap.Error(err))
		return "", err
	}
	newSuperUserId := superUserResponse.Result.Userids[0]
	core.Logger.Info(fmt.Sprintf("[bootstrap]: init_zbx create new super user %s success", newSuperUserId))
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
		core.Logger.Error("[bootstrap]: init_zbx failed to create new action", zap.Error(err))
		return "", err
	}
	newActionId := newActionResponse.Result.ActionIds[0]
	core.Logger.Info(fmt.Sprintf("[bootstrap]: init_zbx create new action %s success", newActionId))
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
		core.Logger.Error("[bootstrap]: init_zbx failed to create super token", zap.Error(err))
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
		core.Logger.Error("[bootstrap]: init_zbx failed to create super token", zap.Error(err))
		return "", err
	}
	nowToken := nowTokenGenerateResponse.Result[0].Token
	core.Logger.Info(fmt.Sprintf("[bootstrap]: init_zbx create new super token %s success", nowToken))
	writeTokenToYamlConfig(nowToken)
	return nowToken, nil
}

func writeTokenToYamlConfig(token string) {
	yamlConfig, err := os.ReadFile(core.ProjectPath + "/config.yaml")
	if err != nil {
		core.Logger.Error("[bootstrap]: failed to read config file", zap.Error(err))
		return
	}
	var config config.Settings
	err = yaml.Unmarshal(yamlConfig, &config)
	if err != nil {
		core.Logger.Error("[bootstrap]: failed to unmarshal config file", zap.Error(err))
		return
	}
	config.Zbx.Token = token
	yamlConfig, err = yaml.Marshal(&config)
	if err != nil {
		core.Logger.Error("[bootstrap]: failed to marshal config file", zap.Error(err))
		return
	}
	err = os.WriteFile(core.ProjectPath+"config.yaml", yamlConfig, 0644)
	if err != nil {
		core.Logger.Error("[bootstrap]: failed to write config file", zap.Error(err))
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
		core.Logger.Error("[bootstrap]: init_zbx failed to disable default admin", zap.Error(err))
		return err
	}
	core.Logger.Info("[bootstrap]: init_zbx disable default admin success")
	return nil
}
