package zbx

import (
	"encoding/json"
	"fmt"
	"strings"
	"sync"
	"sync/atomic"
	"time"

	"github.com/imroc/req/v3"
)

var requestId uint64

var once sync.Once
var zbxInstance *Zbx

type Zbx struct {
	*req.Client
	id    uint64
	token string
	// HostGroup *HostGroupImpl
	// Host      *HostImpl
	// Template  *TemplateImpl
	// Interface *InterfaceImpl

}

type ErrorMessage struct {
	Message string `json:"message"`
}

func (e *ErrorMessage) Error() string {
	return fmt.Sprintf("request metrics system error: %s", e.Message)
}

func NewZbxClient(url, token string) *Zbx {
	url = strings.TrimSuffix(url, "/")
	once.Do(func() {
		zbxInstance = &Zbx{
			Client: req.C().SetBaseURL(url).SetCommonContentType("application/json").
				SetCommonBearerAuthToken(token).SetCommonRetryCount(2).SetTimeout(time.Duration(5) * time.Second).
				OnAfterResponse(func(client *req.Client, resp *req.Response) error {
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
			id: requestId,
		}
	})
	return zbxInstance
}

func (zbx *Zbx) Rpc(request *ZbxRequest) (rsp *ZbxResponse, err error) {
	if zbx == nil {
		return nil, fmt.Errorf("zbx client is nil")
	}
	if zbx.token == "" {
		return nil, fmt.Errorf("token is nil")
	}
	request.Id = atomic.AddUint64(&requestId, 1)
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

type HostGroupImpl struct {
	z *Zbx
}

func (hg *HostGroupImpl) Create(params HostGroupCreate) (groupId string, err error) {
	req := &ZbxRequest{
		Params: params,
		Method: "hostgroup.create",
	}

	rsp, err := hg.z.Rpc(req)
	if err != nil {
		return "", err
	}
	hgr := HostGroupCreateResult{}
	rsp.GetResult(&hgr)
	return hgr.GroupIds[0], nil
}

func (hg *HostGroupImpl) Update(params HostGroupUpdate) (groupId string, err error) {
	req := &ZbxRequest{
		Params: params,
		Method: "hostgroup.update",
	}

	rsp, err := hg.z.Rpc(req)
	if err != nil {
		return "", err
	}
	hgr := HostGroupUpdateResult{}
	rsp.GetResult(&hgr)
	return hgr.GroupIds[0], nil
}

func (hg *HostGroupImpl) Delete(hostGroupIds []string) (groupId []string, err error) {
	req := &ZbxRequest{
		Params: hostGroupIds,
		Method: "hostgroup.delete",
	}

	rsp, err := hg.z.Rpc(req)
	if err != nil {
		return []string{}, err
	}
	hgr := HostGroupDeleteResult{}
	rsp.GetResult(&hgr)
	return hgr.GroupIds, nil
}

func (hg *HostGroupImpl) Get(params HostGroupGet) (res []HostGroup, err error) {
	req := &ZbxRequest{
		Params: params,
		Method: "hostgroup.get",
	}

	rsp, err := hg.z.Rpc(req)
	if err != nil {
		return nil, err
	}

	json.Unmarshal([]byte(rsp.Result), &res)
	return
}

type HostImpl struct {
	z *Zbx
}

func (h *HostImpl) Create(params HostGet) (res string, err error) {
	req := &ZbxRequest{
		Params: params,
		Method: "host.create",
	}

	rsp, err := h.z.Rpc(req)
	if err != nil {
		return "", err
	}
	host := HostCreateResult{}
	rsp.GetResult(&host)
	return host.HostIds[0], nil
}

func (h *HostImpl) Update(params HostUpdate) (res string, err error) {
	req := &ZbxRequest{
		Params: params,
		Method: "host.update",
	}

	rsp, err := h.z.Rpc(req)
	if err != nil {
		return "", err
	}
	host := HostUpdateResult{}
	rsp.GetResult(&host)
	return host.HostIds[0], nil
}

func (h *HostImpl) MassUpdate(params HostMassUpdate) (res string, err error) {
	req := &ZbxRequest{
		Params: params,
		Method: "host.massupdate",
	}

	rsp, err := h.z.Rpc(req)
	if err != nil {
		return "", err
	}
	host := HostUpdateResult{}
	rsp.GetResult(&host)
	return host.HostIds[0], nil
}

func (h *HostImpl) Delete(hostIds []string) (res []string, err error) {
	req := &ZbxRequest{
		Params: hostIds,
		Method: "host.delete",
	}

	rsp, err := h.z.Rpc(req)
	if err != nil {
		return []string{}, err
	}
	host := HostDeleteResult{}
	rsp.GetResult(&host)
	return host.HostIds, nil
}

func (h *HostImpl) Get(params HostGet) (res []Host, err error) {
	req := &ZbxRequest{
		Params: params,
		Method: "host.get",
	}

	rsp, err := h.z.Rpc(req)
	if err != nil {
		return nil, err
	}
	rsp.GetResult(&res)
	json.Unmarshal([]byte(rsp.Result), &res)
	return
}

type TemplateImpl struct {
	z *Zbx
}

func (t *TemplateImpl) Create(params TemplateCreate) (res string, err error) {
	req := &ZbxRequest{
		Params: params,
		Method: "template.create",
	}

	rsp, err := t.z.Rpc(req)
	if err != nil {
		return "", err
	}
	host := TemplateCreateResult{}
	rsp.GetResult(&host)
	return host.TemplateIds[0], nil
}

func (t *TemplateImpl) Update(params TemplateUpdate) (res string, err error) {
	req := &ZbxRequest{
		Params: params,
		Method: "template.update",
	}

	rsp, err := t.z.Rpc(req)
	if err != nil {
		return "", err
	}
	host := TemplateUpdateResult{}
	rsp.GetResult(&host)
	return host.TemplateIds[0], nil
}

func (t *TemplateImpl) Delete(templateIds []string) (res []string, err error) {
	req := &ZbxRequest{
		Params: templateIds,
		Method: "template.delete",
	}

	rsp, err := t.z.Rpc(req)
	if err != nil {
		return []string{}, err
	}
	host := TemplateDeleteResult{}
	rsp.GetResult(&host)
	return host.TemplateIds, nil
}

type ConfigurationImpl struct {
	z *Zbx
}

func (t *ConfigurationImpl) Import(config string) (res bool, err error) {
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
	}

	req := &ZbxRequest{
		Params: params,
		Method: "configuration.import",
	}

	rsp, err := t.z.Rpc(req)
	if err != nil {
		return false, err
	}
	rsp.GetResult(&res)
	return res, nil
}

type EventImpl struct {
	z *Zbx
}

func (t *EventImpl) Get(params EventGet) (res []Event, err error) {
	req := &ZbxRequest{
		Params: params,
		Method: "event.get",
	}

	rsp, err := t.z.Rpc(req)
	if err != nil {
		return nil, err
	}
	rsp.GetResult(&res)
	json.Unmarshal([]byte(rsp.Result), &res)
	return
}

func (t *EventImpl) Acknowledge(eventIds []string) (res []string, err error) {
	req := &ZbxRequest{
		Params: eventIds,
		Method: "event.acknowledge",
	}

	rsp, err := t.z.Rpc(req)
	if err != nil {
		return []string{}, err
	}
	rsp.GetResult(&res)
	return res, nil
}
