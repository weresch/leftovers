package logicalrouting_test

import (
	"context"
	"errors"

	"github.com/genevieve/leftovers/nsxt/logicalrouting"
	"github.com/genevieve/leftovers/nsxt/logicalrouting/fakes"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Tier 1 Router", func() {
	var (
		client       *fakes.LogicalRoutingAndServicesAPI
		staticRoutes *fakes.StaticRoutes
		ctx          context.Context
		name         string
		id           string

		tier1Router logicalrouting.Tier1Router
	)

	BeforeEach(func() {
		client = &fakes.LogicalRoutingAndServicesAPI{}
		staticRoutes = &fakes.StaticRoutes{}
		name = "ackee"
		id = "ackee-123"

		ctx = context.WithValue(context.Background(), "fruit", "ackee")

		tier1Router = logicalrouting.NewTier1Router(client, ctx, name, id, staticRoutes)
	})

	Describe("Delete", func() {
		It("deletes the tier1 router", func() {
			err := tier1Router.Delete()
			Expect(err).NotTo(HaveOccurred())

			Expect(client.DeleteLogicalRouterCall.CallCount).To(Equal(1))
			Expect(client.DeleteLogicalRouterCall.Receives.ID).To(Equal(id))
			Expect(client.DeleteLogicalRouterCall.Receives.Context).To(Equal(ctx))
			Expect(client.DeleteLogicalRouterCall.Receives.LocalVarOptionals).To(HaveKeyWithValue("force", true))

			Expect(staticRoutes.DeleteCall.CallCount).To(Equal(1))
			Expect(staticRoutes.DeleteCall.Receives.RouterID).To(Equal("ackee-123"))
		})

		Context("when static routes fail to delete", func() {
			BeforeEach(func() {
				staticRoutes.DeleteCall.Returns.Error = errors.New("lp0 on fire")
			})

			It("returns the error", func() {
				err := tier1Router.Delete()
				Expect(err).To(MatchError("Delete static routes: lp0 on fire"))
			})
		})

		Context("when the client fails to delete the router", func() {
			BeforeEach(func() {
				client.DeleteLogicalRouterCall.Returns.Error = errors.New("insufficient funds")
			})

			It("returns the error", func() {
				err := tier1Router.Delete()
				Expect(err).To(MatchError("Delete: insufficient funds"))
			})
		})
	})

	Describe("Name", func() {
		It("returns the name", func() {
			Expect(tier1Router.Name()).To(Equal(name))
		})
	})

	Describe("Type", func() {
		It("returns the type", func() {
			Expect(tier1Router.Type()).To(Equal("Tier 1 Router"))
		})
	})
})
