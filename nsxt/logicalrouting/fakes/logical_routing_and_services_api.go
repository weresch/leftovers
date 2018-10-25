package fakes

import (
	"context"
	"net/http"

	"github.com/vmware/go-vmware-nsxt/manager"
)

type LogicalRoutingAndServicesAPI struct {
	ListNatRulesCall struct {
		CallCount int
		Receives  struct {
			Context           context.Context
			LocalVarOptionals map[string]interface{}
		}
		Returns struct {
			ListResult manager.NatRuleListResult
			Response   *http.Response
			Error      error
		}
	}

	DeleteNatRuleCall struct {
		CallCount int
		Receives  struct {
			Context         context.Context
			LogicalRouterID string
			ID              string
		}
		Returns struct {
			Response *http.Response
			Error    error
		}
	}

	ListLogicalRoutersCall struct {
		CallCount int
		Receives  struct {
			Context           context.Context
			LocalVarOptionals map[string]interface{}
		}
		Returns struct {
			ListResult manager.LogicalRouterListResult
			Response   *http.Response
			Error      error
		}
	}

	DeleteLogicalRouterCall struct {
		CallCount int
		Receives  struct {
			Context           context.Context
			ID                string
			LocalVarOptionals map[string]interface{}
		}
		Returns struct {
			Response *http.Response
			Error    error
		}
	}
}

func (l *LogicalRoutingAndServicesAPI) ListNatRules(ctx context.Context, localVarOptionals map[string]interface{}) (manager.NatRuleListResult, *http.Response, error) {
	l.ListNatRulesCall.CallCount++
	l.ListNatRulesCall.Receives.Context = ctx
	l.ListNatRulesCall.Receives.LocalVarOptionals = localVarOptionals

	return l.ListNatRulesCall.Returns.ListResult, l.ListNatRulesCall.Returns.Response, l.ListNatRulesCall.Returns.Error
}

func (l *LogicalRoutingAndServicesAPI) DeleteNatRule(ctx context.Context, routerID string, id string) (*http.Response, error) {
	l.DeleteNatRuleCall.CallCount++
	l.DeleteNatRuleCall.Receives.Context = ctx
	l.DeleteNatRuleCall.Receives.LogicalRouterID = routerID
	l.DeleteNatRuleCall.Receives.ID = id

	return l.DeleteNatRuleCall.Returns.Response, l.DeleteNatRuleCall.Returns.Error
}

func (l *LogicalRoutingAndServicesAPI) ListLogicalRouters(ctx context.Context, localVarOptionals map[string]interface{}) (manager.LogicalRouterListResult, *http.Response, error) {
	l.ListLogicalRoutersCall.CallCount++
	l.ListLogicalRoutersCall.Receives.Context = ctx
	l.ListLogicalRoutersCall.Receives.LocalVarOptionals = localVarOptionals

	return l.ListLogicalRoutersCall.Returns.ListResult, l.ListLogicalRoutersCall.Returns.Response, l.ListLogicalRoutersCall.Returns.Error
}

func (l *LogicalRoutingAndServicesAPI) DeleteLogicalRouter(ctx context.Context, id string, localVarOptionals map[string]interface{}) (*http.Response, error) {
	l.DeleteLogicalRouterCall.CallCount++
	l.DeleteLogicalRouterCall.Receives.Context = ctx
	l.DeleteLogicalRouterCall.Receives.ID = id
	l.DeleteLogicalRouterCall.Receives.LocalVarOptionals = localVarOptionals

	return l.DeleteLogicalRouterCall.Returns.Response, l.DeleteLogicalRouterCall.Returns.Error
}
