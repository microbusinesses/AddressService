package service_test

import (
	"testing"

	"github.com/gocql/gocql"
	"github.com/micro-business/AddressService/data/service"
	"github.com/micro-business/Micro-Business-Core/system"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("ReadAll method input parameters and dependency test", func() {
	var (
		addressDataService *service.AddressDataService
		tenantID           system.UUID
		applicationID      system.UUID
		addressID          system.UUID
	)

	BeforeEach(func() {
		addressDataService = &service.AddressDataService{ClusterConfig: &gocql.ClusterConfig{}}

		addressDataService = &service.AddressDataService{}
		tenantID, _ = system.RandomUUID()
		applicationID, _ = system.RandomUUID()
		addressID, _ = system.RandomUUID()
	})

	Context("when cluster configuration not provided", func() {
		It("should panic", func() {
			addressDataService.ClusterConfig = nil

			Ω(func() { addressDataService.ReadAll(tenantID, applicationID, addressID) }).Should(Panic())
		})
	})
})

func TestReadAll(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "ReadAll method input parameters and dependency test")
}
