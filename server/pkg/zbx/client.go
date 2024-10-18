package zbx

import (
	"encoding/json"
	"fmt"
	"strings"
	"sync"
	"sync/atomic"
	"time"

	"github.com/imroc/req/v3"
	"github.com/wangxin688/narvis/server/core"

	zs "github.com/wangxin688/narvis/server/pkg/zbx/zschema"
)

var requestID uint64

var once sync.Once
var zbxInstance *Zbx

type Zbx struct {
	*req.Client
	id uint64
}

type ErrorMessage struct {
	Message string `json:"message"`
}

func (e *ErrorMessage) Error() string {
	return fmt.Sprintf("request metrics system error: %s", e.Message)
}

func NewZbxClient() *Zbx {
	url := core.Settings.Zbx.Url
	token := core.Settings.Zbx.Token
	url = strings.TrimSuffix(url, "/")
	once.Do(func() {
		zbxInstance = &Zbx{
			Client: req.C().SetBaseURL(url).SetCommonContentType("application/json").
				SetCommonBearerAuthToken(token).SetCommonRetryCount(2).SetTimeout(time.Duration(5) * time.Second).
				OnAfterResponse(func(_ *req.Client, resp *req.Response) error {
					if resp.Err != nil {
						return resp.Err
					}
					if errMsg, ok := resp.ErrorResult().(*ErrorMessage); ok {
						resp.Err = errMsg
						return nil
					}
					if !resp.IsSuccessState() {
						resp.Err = fmt.Errorf("request metrics system error: %s, status code: %s", resp.Dump(), resp.Status)
					}
					return nil
				}),
			id: requestID,
		}
	})
	return zbxInstance
}

func (zbx *Zbx) Rpc(request *zs.ZbxRequest) (rsp *zs.ZbxResponse, err error) {
	if zbx == nil {
		return nil, fmt.Errorf("zbx client is nil")
	}
	request.ID = atomic.AddUint64(&requestID, 1)
	request.JsonRpc = "2.0"

	_, err = zbx.R().SetSuccessResult(&rsp).SetBody(request).Post("/api_jsonrpc.php")
	if err != nil {
		return nil, err
	}
	err = rsp.HasError()
	if err != nil {
		return nil, err
	}
	return rsp, nil
}

func (z *Zbx) HostGroupCreate(params *zs.HostGroupCreate) (string, error) {
	req := &zs.ZbxRequest{
		Params: params,
		Method: "hostgroup.create",
	}

	rsp, err := z.Rpc(req)
	if err != nil {
		return "", err
	}
	hgr := zs.HostGroupCreateResult{}
	rsp.GetResult(&hgr)
	return hgr.GroupIDs[0], nil
}

func (z *Zbx) HostGroupUpdate(params *zs.HostGroupUpdate) (string, error) {
	req := &zs.ZbxRequest{
		Params: params,
		Method: "hostgroup.update",
	}

	rsp, err := z.Rpc(req)
	if err != nil {
		return "", err
	}
	hgr := zs.HostGroupUpdateResult{}
	rsp.GetResult(&hgr)
	return hgr.GroupIDs[0], nil
}

func (z *Zbx) HostGroupDelete(hostGroupIDs []string) ([]string, error) {
	req := &zs.ZbxRequest{
		Params: hostGroupIDs,
		Method: "hostgroup.delete",
	}

	rsp, err := z.Rpc(req)
	if err != nil {
		return []string{}, err
	}
	hgr := zs.HostGroupDeleteResult{}
	rsp.GetResult(&hgr)
	return hgr.GroupIDs, nil
}

func (z *Zbx) HostGroupGet(params *zs.HostGroupGet) (res []*zs.HostGroup, err error) {
	req := &zs.ZbxRequest{
		Params: params,
		Method: "hostgroup.get",
	}

	rsp, err := z.Rpc(req)
	if err != nil {
		return nil, err
	}

	json.Unmarshal([]byte(rsp.Result), &res)
	return
}
func (z *Zbx) HostCreate(params *zs.HostCreate) (res string, err error) {
	req := &zs.ZbxRequest{
		Params: params,
		Method: "host.create",
	}

	rsp, err := z.Rpc(req)
	if err != nil {
		return "", err
	}
	host := zs.HostCreateResult{}
	rsp.GetResult(&host)
	return host.HostIDs[0], nil
}

func (z *Zbx) HostUpdate(params *zs.HostUpdate) (res string, err error) {
	req := &zs.ZbxRequest{
		Params: params,
		Method: "host.update",
	}

	rsp, err := z.Rpc(req)
	if err != nil {
		return "", err
	}
	host := zs.HostUpdateResult{}
	rsp.GetResult(&host)
	return host.HostIDs[0], nil
}

func (z *Zbx) HostMassUpdate(params *zs.HostMassUpdate) (res string, err error) {
	req := &zs.ZbxRequest{
		Params: params,
		Method: "host.massupdate",
	}

	rsp, err := z.Rpc(req)
	if err != nil {
		return "", err
	}
	host := zs.HostUpdateResult{}
	rsp.GetResult(&host)
	return host.HostIDs[0], nil
}

func (z *Zbx) HostDelete(hostIDs []string) (res []string, err error) {
	req := &zs.ZbxRequest{
		Params: hostIDs,
		Method: "host.delete",
	}

	rsp, err := z.Rpc(req)
	if err != nil {
		return []string{}, err
	}
	host := zs.HostDeleteResult{}
	rsp.GetResult(&host)
	return host.HostIDs, nil
}

func (z *Zbx) HostInterfaceGet(params *zs.HostInterfaceGet) (res []*zs.HostInterface, err error) {
	req := &zs.ZbxRequest{
		Params: params,
		Method: "hostinterface.get",
	}

	rsp, err := z.Rpc(req)
	if err != nil {
		return nil, err
	}
	rsp.GetResult(&res)
	json.Unmarshal([]byte(rsp.Result), &res)
	return
}

func (z *Zbx) HostInterfaceReplaceHostInterfaces(params *zs.HostInterfaceReplace) (res []string, err error) {
	req := &zs.ZbxRequest{
		Params: params,
		Method: "hostinterface.replacehostinterfaces",
	}

	rsp, err := z.Rpc(req)
	if err != nil {
		return []string{}, err
	}
	host := zs.HostInterfaceUpdateResult{}
	rsp.GetResult(&host)
	return host.InterfaceIds, nil
}

func (z *Zbx) HostGet(params *zs.HostGet) (res []*zs.Host, err error) {
	req := &zs.ZbxRequest{
		Params: params,
		Method: "host.get",
	}

	rsp, err := z.Rpc(req)
	if err != nil {
		return nil, err
	}
	rsp.GetResult(&res)
	json.Unmarshal([]byte(rsp.Result), &res)
	return
}

func (z *Zbx) TemplateCreate(params *zs.TemplateCreate) (res string, err error) {
	req := &zs.ZbxRequest{
		Params: params,
		Method: "template.create",
	}

	rsp, err := z.Rpc(req)
	if err != nil {
		return "", err
	}
	host := zs.TemplateCreateResult{}
	rsp.GetResult(&host)
	return host.TemplateIDs[0], nil
}

func (z *Zbx) TemplateUpdate(params *zs.TemplateUpdate) (res string, err error) {
	req := &zs.ZbxRequest{
		Params: params,
		Method: "template.update",
	}

	rsp, err := z.Rpc(req)
	if err != nil {
		return "", err
	}
	host := zs.TemplateUpdateResult{}
	rsp.GetResult(&host)
	return host.TemplateIDs[0], nil
}

func (z *Zbx) TemplateDelete(templateIDs []string) (res []string, err error) {
	req := &zs.ZbxRequest{
		Params: templateIDs,
		Method: "template.delete",
	}

	rsp, err := z.Rpc(req)
	if err != nil {
		return []string{}, err
	}
	host := zs.TemplateDeleteResult{}
	rsp.GetResult(&host)
	return host.TemplateIDs, nil
}

func (z *Zbx) TemplateGet(params *zs.TemplateGet) (res []*zs.TemplateGetResult, err error) {
	req := &zs.ZbxRequest{
		Params: params,
		Method: "template.get",
	}

	rsp, err := z.Rpc(req)
	if err != nil {
		return nil, err
	}
	rsp.GetResult(&res)
	json.Unmarshal([]byte(rsp.Result), &res)
	return
}

func (z *Zbx) ConfigurationImport(config string) (res bool, err error) {
	params := make(map[string]any)
	params["format"] = "yaml"
	params["source"] = config
	params["rules"] = map[string]map[string]bool{
		"templates":       {"createMissing": true, "updateExisting": true},
		"items":           {"createMissing": true, "updateExisting": true, "deleteMissing": true},
		"triggers":        {"createMissing": true, "updateExisting": true, "deleteMissing": true},
		"valueMaps":       {"createMissing": true, "updateExisting": true},
		"discoveryRules":  {"createMissing": true, "updateExisting": true, "deleteMissing": true},
		"graphs":          {"createMissing": true, "updateExisting": true, "deleteMissing": true},
		"template_groups": {"createMissing": true, "updateExisting": true},
		"templateLinkage": {"createMissing": true, "deleteMissing": true},
	}

	req := &zs.ZbxRequest{
		Params: params,
		Method: "configuration.import",
	}

	rsp, err := z.Rpc(req)
	if err != nil {
		return false, err
	}
	rsp.GetResult(&res)
	return res, nil
}

func (z *Zbx) EventGet(params *zs.EventGet) (res []*zs.Event, err error) {
	req := &zs.ZbxRequest{
		Params: params,
		Method: "event.get",
	}

	rsp, err := z.Rpc(req)
	if err != nil {
		return nil, err
	}
	rsp.GetResult(&res)
	json.Unmarshal([]byte(rsp.Result), &res)
	return
}

func (z *Zbx) EventAcknowledge(eventIDs []string) (res []string, err error) {
	req := &zs.ZbxRequest{
		Params: eventIDs,
		Method: "event.acknowledge",
	}

	rsp, err := z.Rpc(req)
	if err != nil {
		return []string{}, err
	}
	rsp.GetResult(&res)
	return res, nil
}

func (z *Zbx) UserMacroCreateGlobal(params *zs.Macro) (res string, err error) {
	req := &zs.ZbxRequest{
		Params: params,
		Method: "usermacro.createglobal",
	}

	rsp, err := z.Rpc(req)
	if err != nil {
		return "", err
	}
	result := zs.GlobalMacroCreateResult{}
	rsp.GetResult(&result)
	return result.GlobalMacroIDs[0], nil
}

func (z *Zbx) UserMacroUpdateGlobal(params *zs.GlobalMacroUpdate) (res string, err error) {
	req := &zs.ZbxRequest{
		Params: params,
		Method: "usermacro.updateglobal",
	}
	rsp, err := z.Rpc(req)
	if err != nil {
		return "", err
	}
	result := zs.GlobalMacroCreateResult{}
	rsp.GetResult(&result)
	return result.GlobalMacroIDs[0], nil
}

func (z *Zbx) ProxyCreate(params *zs.ProxyCreate) (res string, err error) {
	req := &zs.ZbxRequest{
		Params: params,
		Method: "proxy.create",
	}
	resp, err := z.Rpc(req)
	if err != nil {
		return "", err
	}
	result := zs.ProxyCreateResult{}
	resp.GetResult(&result)
	return result.ProxyIDs[0], nil
}
