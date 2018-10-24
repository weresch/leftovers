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

var _ = Describe("Static Routes", func() {
	var (
		client       *fakes.LogicalRoutingAndServicesAPI
		ctx          context.Context
		staticRoutes logicalrouting.StaticRoutes
	)

	BeforeEach(func() {
		client = &fakes.LogicalRoutingAndServicesAPI{}

		ctx = context.WithValue(context.Background(), "fruit", "yuzu")

		staticRoutes = logicalrouting.NewStaticRoutes(client, ctx)
	})

	Describe("Delete", func() {
		BeforeEach(func() {
			client.ListStaticRoutesCall.Returns.ListResult = manager.StaticRouteListResult{
				Results: []manager.StaticRoute{
					manager.StaticRoute{
						Id:              "tamarind-123",
						LogicalRouterId: "pitaya-789",
						DisplayName:     "tamarind",
					},
					manager.StaticRoute{
						Id:              "starfruit-456",
						LogicalRouterId: "pitaya-789",
						DisplayName:     "starfruit",
					},
				},
			}
		})

		It("lists and deletes static routes for the provided tier 1 router ID", func() {
			err := staticRoutes.Delete("pitaya-789")
			Expect(err).NotTo(HaveOccurred())

			Expect(client.ListStaticRoutesCall.CallCount).To(Equal(1))
			Expect(client.ListStaticRoutesCall.Receives.Context).To(Equal(ctx))
			Expect(client.ListStaticRoutesCall.Receives.LogicalRouterID).To(Equal("pitaya-789"))

			Expect(client.DeleteStaticRouteCall.CallCount).To(Equal(2))
			Expect(client.DeleteStaticRouteCall.Receives[0].ID).To(Equal("tamarind-123"))
			Expect(client.DeleteStaticRouteCall.Receives[0].RouterID).To(Equal("pitaya-789"))
			Expect(client.DeleteStaticRouteCall.Receives[0].Context).To(Equal(ctx))

			Expect(client.DeleteStaticRouteCall.Receives[1].ID).To(Equal("starfruit-456"))
			Expect(client.DeleteStaticRouteCall.Receives[1].RouterID).To(Equal("pitaya-789"))
			Expect(client.DeleteStaticRouteCall.Receives[1].Context).To(Equal(ctx))
		})

		Context("when the client fails to list static routes", func() {
			BeforeEach(func() {
				client.ListStaticRoutesCall.Returns.Error = errors.New("Not a typewriter")
			})

			It("returns the error", func() {
				err := staticRoutes.Delete("pitaya-123")
				Expect(err).To(MatchError("List Static Routes: Not a typewriter"))
			})
		})

		Context("when the client fails to delete static routes", func() {
			BeforeEach(func() {
				client.DeleteStaticRouteCall.Returns = []fakes.DeleteStaticRouteCallReturns{
					{
						Response: nil,
						Error:    errors.New("Not a typewriter"),
					},
				}
			})

			It("returns the error", func() {
				err := staticRoutes.Delete("pitaya-123")
				Expect(err).To(MatchError("Delete Static Route: Not a typewriter"))
			})
		})
	})
})
