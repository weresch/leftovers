package fakes

import (
	"context"
	"net/http"

	"github.com/vmware/go-vmware-nsxt/manager"
)

type LogicalRoutingAndServicesAPI struct {
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
	ListStaticRoutesCall struct {
		CallCount int
		Receives  struct {
			Context           context.Context
			LogicalRouterID   string
			LocalVarOptionals map[string]interface{}
		}
		Returns struct {
			ListResult manager.StaticRouteListResult
			Response   *http.Response
			Error      error
		}
	}
	DeleteStaticRouteCall struct {
		CallCount int
		Receives  []DeleteStaticRouteCallReceives
		Returns   []DeleteStaticRouteCallReturns
	}
}

type DeleteStaticRouteCallReceives struct {
	Context  context.Context
	ID       string
	RouterID string
}

type DeleteStaticRouteCallReturns struct {
	Response *http.Response
	Error    error
}

func (l *LogicalRoutingAndServicesAPI) DeleteLogicalRouter(ctx context.Context, id string, localVarOptionals map[string]interface{}) (*http.Response, error) {
	l.DeleteLogicalRouterCall.CallCount++
	l.DeleteLogicalRouterCall.Receives.Context = ctx
	l.DeleteLogicalRouterCall.Receives.ID = id
	l.DeleteLogicalRouterCall.Receives.LocalVarOptionals = localVarOptionals

	return l.DeleteLogicalRouterCall.Returns.Response, l.DeleteLogicalRouterCall.Returns.Error
}

func (l *LogicalRoutingAndServicesAPI) ListLogicalRouters(ctx context.Context, localVarOptionals map[string]interface{}) (manager.LogicalRouterListResult, *http.Response, error) {
	l.ListLogicalRoutersCall.CallCount++
	l.ListLogicalRoutersCall.Receives.Context = ctx
	l.ListLogicalRoutersCall.Receives.LocalVarOptionals = localVarOptionals

	return l.ListLogicalRoutersCall.Returns.ListResult, l.ListLogicalRoutersCall.Returns.Response, l.ListLogicalRoutersCall.Returns.Error
}

func (l *LogicalRoutingAndServicesAPI) DeleteStaticRoute(ctx context.Context, routerID string, id string) (*http.Response, error) {
	l.DeleteStaticRouteCall.CallCount++

	l.DeleteStaticRouteCall.Receives = append(l.DeleteStaticRouteCall.Receives, DeleteStaticRouteCallReceives{
		Context:  ctx,
		ID:       id,
		RouterID: routerID,
	})

	if len(l.DeleteStaticRouteCall.Returns) < l.DeleteStaticRouteCall.CallCount {
		return nil, nil
	}

	return l.DeleteStaticRouteCall.Returns[l.DeleteStaticRouteCall.CallCount-1].Response, l.DeleteStaticRouteCall.Returns[l.DeleteStaticRouteCall.CallCount-1].Error
}

func (l *LogicalRoutingAndServicesAPI) ListStaticRoutes(ctx context.Context, logicalRouterID string, localVarOptionals map[string]interface{}) (manager.StaticRouteListResult, *http.Response, error) {
	l.ListStaticRoutesCall.CallCount++
	l.ListStaticRoutesCall.Receives.Context = ctx
	l.ListStaticRoutesCall.Receives.LogicalRouterID = logicalRouterID
	l.ListStaticRoutesCall.Receives.LocalVarOptionals = localVarOptionals

	return l.ListStaticRoutesCall.Returns.ListResult, l.ListStaticRoutesCall.Returns.Response, l.ListStaticRoutesCall.Returns.Error
}
