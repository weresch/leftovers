package logicalrouting

import (
	"context"
	"fmt"
	"strings"

	"github.com/genevieve/leftovers/common"
)

type NatRules struct {
	client logicalroutingClient
	ctx    context.Context
	logger logger
}

func NewNatRules(client logicalroutingClient, ctx context.Context, logger logger) NatRules {
	return NatRules{
		client: client,
		ctx:    ctx,
		logger: logger,
	}
}

func (n NatRules) List(filter string) ([]common.Deletable, error) {
	result, _, err := n.client.ListNatRules(n.ctx, map[string]interface{}{})
	if err != nil {
		return []common.Deletable{}, fmt.Errorf("List NAT Rules: %s", err)
	}

	resources := []common.Deletable{}
	for _, rule := range result.Results {
		if !strings.Contains(rule.DisplayName, filter) {
			continue
		}

		proceed := n.logger.PromptWithDetails("NAT Rule", rule.DisplayName)
		if !proceed {
			continue
		}

		resources = append(resources, NewNatRule(n.client, n.ctx, rule.DisplayName, rule.LogicalRouterId, rule.Id))
	}

	return resources, nil
}
