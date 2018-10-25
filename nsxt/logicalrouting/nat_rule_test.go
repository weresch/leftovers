package logicalrouting_test

import (
	"context"
	"errors"

	"github.com/genevieve/leftovers/nsxt/logicalrouting"
	"github.com/genevieve/leftovers/nsxt/logicalrouting/fakes"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("NAT Rule", func() {
	var (
		client   *fakes.LogicalRoutingAndServicesAPI
		ctx      context.Context
		name     string
		routerID string
		id       string

		natRule logicalrouting.NatRule
	)

	BeforeEach(func() {
		client = &fakes.LogicalRoutingAndServicesAPI{}
		ctx = context.WithValue(context.Background(), "fruit", "soursop")
		name = "banana"
		routerID = "router-id"
		id = "fruit"

		natRule = logicalrouting.NewNatRule(client, ctx, name, routerID, id)
	})

	Describe("Delete", func() {
		It("deletes the nat rule", func() {
			err := natRule.Delete()
			Expect(err).NotTo(HaveOccurred())

			Expect(client.DeleteNatRuleCall.CallCount).To(Equal(1))
			Expect(client.DeleteNatRuleCall.Receives.LogicalRouterID).To(Equal(routerID))
			Expect(client.DeleteNatRuleCall.Receives.ID).To(Equal(id))
		})

		Context("when the client fails to delete the nat rule", func() {
			BeforeEach(func() {
				client.DeleteNatRuleCall.Returns.Error = errors.New("the-error-msg")
			})

			It("returns a helpful error", func() {
				err := natRule.Delete()
				Expect(err).To(MatchError("Delete: the-error-msg"))
			})
		})
	})

	Describe("Type", func() {
		It("returns the type of the resource", func() {
			Expect(natRule.Type()).To(Equal("NAT Rule"))
		})
	})

	Describe("Name", func() {
		It("returns the name of the resource", func() {
			Expect(natRule.Name()).To(Equal(name))
		})
	})
})
