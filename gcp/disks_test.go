package gcp_test

import (
	"errors"

	"github.com/genevievelesperance/leftovers/gcp"
	"github.com/genevievelesperance/leftovers/gcp/fakes"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	compute "google.golang.org/api/compute/v1"
)

var _ = Describe("Disks", func() {
	var (
		client *fakes.DisksClient
		logger *fakes.Logger
		zones  []string

		disks gcp.Disks
	)

	BeforeEach(func() {
		client = &fakes.DisksClient{}
		logger = &fakes.Logger{}
		zones = []string{"zone-1", "zone-2"}

		disks = gcp.NewDisks(client, logger, zones)
	})

	Describe("Delete", func() {
		BeforeEach(func() {
			logger.PromptCall.Returns.Proceed = true
			client.ListDisksCall.Returns.Output = &compute.DiskList{
				Items: []*compute.Disk{{
					Name: "banana",
					Zone: "the-zone",
				}},
			}
		})

		It("deletes disks", func() {
			err := disks.Delete()
			Expect(err).NotTo(HaveOccurred())

			Expect(client.ListDisksCall.CallCount).To(Equal(1))

			Expect(logger.PromptCall.Receives.Message).To(Equal("Are you sure you want to delete disk banana?"))

			Expect(client.DeleteDiskCall.CallCount).To(Equal(1))
			Expect(client.DeleteDiskCall.Receives.Zone).To(Equal("the-zone"))
			Expect(client.DeleteDiskCall.Receives.Disk).To(Equal("banana"))

			Expect(logger.PrintfCall.Messages).To(Equal([]string{"SUCCESS deleting disk banana\n"}))
		})

		Context("when the client fails to list disks", func() {
			BeforeEach(func() {
				client.ListDisksCall.Returns.Error = errors.New("some error")
			})

			It("returns the error", func() {
				err := disks.Delete()
				Expect(err).To(MatchError("Listing disks: some error"))
			})
		})

		Context("when the client fails to delete the disk", func() {
			BeforeEach(func() {
				client.DeleteDiskCall.Returns.Error = errors.New("some error")
			})

			It("logs the error", func() {
				err := disks.Delete()
				Expect(err).NotTo(HaveOccurred())

				Expect(logger.PrintfCall.Messages).To(Equal([]string{"ERROR deleting disk banana: some error\n"}))
			})
		})

		Context("when the user says no to the prompt", func() {
			BeforeEach(func() {
				logger.PromptCall.Returns.Proceed = false
			})

			It("does not delete the disk", func() {
				err := disks.Delete()
				Expect(err).NotTo(HaveOccurred())

				Expect(client.DeleteDiskCall.CallCount).To(Equal(0))
			})
		})
	})
})