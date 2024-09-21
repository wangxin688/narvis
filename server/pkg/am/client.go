package am

import (
	"fmt"
	"strings"
	"sync"
	"time"

	"github.com/imroc/req/v3"
)

var once sync.Once
var amInstance *AlertManager

type ErrorMessage struct {
	Message string `json:"message"`
}

func (e *ErrorMessage) Error() string {
	return fmt.Sprintf("request AlertManager  error: %s", e.Message)
}

type AlertManager struct {
	*req.Client
}

func NewAlertManager(url, username, password string) *AlertManager {
	url = strings.TrimSuffix(url, "/")

	once.Do(func() {
		amInstance = &AlertManager{
			Client: req.C().SetBaseURL(url).SetCommonContentType("application/json").
				SetCommonBasicAuth(username, password).SetCommonRetryCount(2).SetTimeout(time.Duration(2) * time.Second).
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
		}
	})
	return amInstance

}

func (am *AlertManager) CreateAlerts(alert []*Alert) error {
	path := "/api/v2/alerts"
	resp, err := am.R().SetBody(alert).Post(path)
	if err != nil {
		return err
	}

	if !resp.IsSuccessState() {
		return fmt.Errorf("request metrics system error: %s, status code: %s", resp.Dump(), resp.Status)
	}
	return nil
}

func (am *AlertManager) GetAlerts(query *AlertRequest) ([]*AlertResponse, error) {
	results := make([]*AlertResponse, 0)
	path := "/api/v2/alerts"
	requestParams := map[string]any{}
	if query.Active != nil {
		requestParams["active"] = query.Active
	} else {
		requestParams["active"] = true
	}

	if query.Silenced != nil {
		requestParams["silenced"] = query.Silenced
	} else {
		requestParams["silenced"] = true
	}

	if query.Inhibited != nil {
		requestParams["inhibited"] = query.Inhibited
	} else {
		requestParams["inhibited"] = true
	}

	if query.Unprocessed != nil {
		requestParams["unprocessed"] = query.Unprocessed
	} else {
		requestParams["unprocessed"] = true
	}
	if query.Filter != nil {
		requestParams["filter"] = query.Filter
	}
	_, err := am.R().SetQueryParamsAnyType(requestParams).SetSuccessResult(&results).Get(path)
	if err != nil {
		return results, err
	}
	return results, nil
}

func (am *AlertManager) GetAlertGroups(query *AlertRequest) ([]*AlertGroupResponse, error) {
	results := make([]*AlertGroupResponse, 0)
	path := "/api/v2/alerts/groups"
	requestParams := map[string]any{}
	if query.Active != nil {
		requestParams["active"] = query.Active
	} else {
		requestParams["active"] = true
	}

	if query.Silenced != nil {
		requestParams["silenced"] = query.Silenced
	} else {
		requestParams["silenced"] = true
	}

	if query.Inhibited != nil {
		requestParams["inhibited"] = query.Inhibited
	} else {
		requestParams["inhibited"] = true
	}

	if query.Unprocessed != nil {
		requestParams["unprocessed"] = query.Unprocessed
	} else {
		requestParams["unprocessed"] = true
	}
	if query.Filter != nil {
		requestParams["filter"] = query.Filter
	}
	_, err := am.R().SetQueryParamsAnyType(requestParams).SetSuccessResult(&results).Get(path)
	if err != nil {
		return results, err
	}
	return results, nil
}
