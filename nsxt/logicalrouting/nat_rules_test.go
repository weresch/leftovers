package logicalrouting_test

import (
	"context"
	"errors"

	"github.com/genevieve/leftovers/nsxt/logicalrouting"
	"github.com/genevieve/leftovers/nsxt/logicalrouting/fakes"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/vmware/go-vmware-nsxt/manager"
)

var _ = Describe("NAT Rules", func() {
	var (
		client   *fakes.LogicalRoutingAndServicesAPI
		logger   *fakes.Logger
		ctx      context.Context
		natRules logicalrouting.NatRules
	)
	BeforeEach(func() {
		client = &fakes.LogicalRoutingAndServicesAPI{}
		logger = &fakes.Logger{}
		logger.PromptWithDetailsCall.Returns.Proceed = true

		ctx = context.WithValue(context.Background(), "fruit", "soursop")

		natRules = logicalrouting.NewNatRules(client, ctx, logger)
	})

	Describe("List", func() {
		BeforeEach(func() {
			client.ListNatRulesCall.Returns.ListResult = manager.NatRuleListResult{
				Results: []manager.NatRule{
					{
						DisplayName:     "banana",
						Id:              "fruit",
						LogicalRouterId: "router-id",
					},
					{
						DisplayName:     "mango",
						Id:              "fruit",
						LogicalRouterId: "router-id",
					},
				},
			}
		})

		It("lists, prompts and returns nat rules to delete", func() {
			list, err := natRules.List("ban")
			Expect(err).NotTo(HaveOccurred())

			Expect(client.ListNatRulesCall.CallCount).To(Equal(1))

			Expect(logger.PromptWithDetailsCall.CallCount).To(Equal(1))
			Expect(logger.PromptWithDetailsCall.Receives.Type).To(Equal("NAT Rule"))
			Expect(logger.PromptWithDetailsCall.Receives.Name).To(Equal("banana"))

			Expect(list).To(HaveLen(1))
		})

		Context("when the user does not want to proceed with deleting", func() {
			BeforeEach(func() {
				logger.PromptWithDetailsCall.Returns.Proceed = false
			})

			It("returns the list of deletable resources", func() {
				list, err := natRules.List("")
				Expect(err).NotTo(HaveOccurred())

				Expect(list).To(HaveLen(0))
			})
		})

		Context("when the client fails to list nat rules", func() {
			BeforeEach(func() {
				client.ListNatRulesCall.Returns.Error = errors.New("the-error-msg")
			})

			It("returns a helpful error", func() {
				_, err := natRules.List("")
				Expect(err).To(MatchError("List NAT Rules: the-error-msg"))
			})
		})
	})
})
