package logicalrouting

import (
	"context"
	"fmt"
)

type StaticRoutes struct {
	client logicalRoutingAPI
	ctx    context.Context
}

type staticRoutes interface {
	Delete(routerID string) error
}

func NewStaticRoutes(client logicalRoutingAPI, ctx context.Context) StaticRoutes {
	return StaticRoutes{
		client: client,
		ctx:    ctx,
	}
}

func (s StaticRoutes) Delete(routerID string) error {
	result, _, err := s.client.ListStaticRoutes(s.ctx, routerID, map[string]interface{}{})
	if err != nil {
		return fmt.Errorf("List Static Routes: %s", err)
	}

	for _, route := range result.Results {
		fmt.Printf("I would be deleting route ID %s on router %s but I decided not to", route.Id, routerID)
		// _, err := s.client.DeleteStaticRoute(s.ctx, routerID, route.Id)
		// if err != nil {
		// 	return fmt.Errorf("Delete Static Route: %s", err)
		// }
	}
	return nil
}

func (s StaticRoutes) Type() string {
	return "Static Route"
}
