package logicalrouting

import (
	"context"
	"fmt"
)

type NatRule struct {
	client   logicalroutingClient
	ctx      context.Context
	name     string
	routerId string
	id       string
}

func NewNatRule(client logicalroutingClient, ctx context.Context, name string, routerId string, id string) NatRule {
	return NatRule{
		client:   client,
		ctx:      ctx,
		name:     name,
		routerId: routerId,
		id:       id,
	}
}

func (n NatRule) Delete() error {
	_, err := n.client.DeleteNatRule(n.ctx, n.routerId, n.id)
	if err != nil {
		return fmt.Errorf("Delete: %s", err)
	}

	return nil
}

func (n NatRule) Type() string {
	return "NAT Rule"
}

func (n NatRule) Name() string {
	return n.name
}
