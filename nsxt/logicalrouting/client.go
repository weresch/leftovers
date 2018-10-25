package logicalrouting

import (
	"context"
	"net/http"

	"github.com/vmware/go-vmware-nsxt/manager"
)

type logicalroutingClient interface {
	DeleteLogicalRouter(context.Context, string, map[string]interface{}) (*http.Response, error)
	ListLogicalRouters(context.Context, map[string]interface{}) (manager.LogicalRouterListResult, *http.Response, error)
	ListNatRules(context.Context, map[string]interface{}) (manager.NatRuleListResult, *http.Response, error)
	DeleteNatRule(ctx context.Context, routerID string, ruleID string) (*http.Response, error)
}
