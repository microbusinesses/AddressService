// +build integration

package service_test

import (
	"fmt"
	"testing"

	"github.com/gocql/gocql"
	"github.com/golang/mock/gomock"
	"github.com/micro-business/AddressService/data/contract"
	"github.com/micro-business/AddressService/data/service"
	"github.com/micro-business/Micro-Business-Core/system"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Delete method behaviour", func() {
	var (
		mockCtrl                 *gomock.Controller
		addressDataService       *service.AddressDataService
		mockUUIDGeneratorService *MockUUIDGeneratorService
		tenantID                 system.UUID
		applicationID            system.UUID
		addressID                system.UUID
		clusterConfig            *gocql.ClusterConfig
	)

	BeforeEach(func() {
		clusterConfig = getClusterConfig()
		clusterConfig.Keyspace = keyspace

		mockCtrl = gomock.NewController(GinkgoT())
		mockUUIDGeneratorService = NewMockUUIDGeneratorService(mockCtrl)

		addressDataService = &service.AddressDataService{UUIDGeneratorService: mockUUIDGeneratorService, ClusterConfig: clusterConfig}

		tenantID, _ = system.RandomUUID()
		applicationID, _ = system.RandomUUID()
		addressID, _ = system.RandomUUID()
	})

	AfterEach(func() {
		mockCtrl.Finish()
	})

	Context("when deleting existing address", func() {
		It("should return error if address does not exist", func() {
			err := addressDataService.Delete(tenantID, applicationID, addressID)

			Expect(err).To(Equal(fmt.Errorf("Address not found. Address ID: %s", addressID.String())))
		})

		It("should remove the records from address table", func() {
			mockUUIDGeneratorService.
				EXPECT().
				GenerateRandomUUID().
				Return(addressID, nil)

			expectedAddressDetails := createRandomAddressDetails()

			returnedAddressID, err := addressDataService.Create(
				tenantID,
				applicationID,
				contract.Address{AddressDetails: expectedAddressDetails})

			Expect(err).To(BeNil())

			err = addressDataService.Delete(
				tenantID,
				applicationID,
				returnedAddressID)

			Expect(err).To(BeNil())

			config := getClusterConfig()
			config.Keyspace = keyspace

			session, err := config.CreateSession()

			Expect(err).To(BeNil())

			defer session.Close()

			var key string
			var value string

			iter := session.Query(
				"SELECT address_key, address_value"+
					" FROM address"+
					" WHERE"+
					" tenant_id = ?"+
					" AND application_id = ?"+
					" AND address_id = ?",
				tenantID.String(),
				applicationID.String(),
				returnedAddressID.String()).Iter()

			defer iter.Close()

			Expect(iter.Scan(&key, &value)).To(BeFalse())
		})

		It("should remove all the index records from address_indexed_by_address_key table", func() {
			mockUUIDGeneratorService.
				EXPECT().
				GenerateRandomUUID().
				Return(addressID, nil)

			expectedAddressDetails := createRandomAddressDetails()

			returnedAddressID, err := addressDataService.Create(
				tenantID,
				applicationID,
				contract.Address{AddressDetails: expectedAddressDetails})

			Expect(err).To(BeNil())

			err = addressDataService.Delete(
				tenantID,
				applicationID,
				returnedAddressID)

			Expect(err).To(BeNil())

			config := getClusterConfig()
			config.Keyspace = keyspace

			session, err := config.CreateSession()

			Expect(err).To(BeNil())

			defer session.Close()

			for key := range expectedAddressDetails {
				var id gocql.UUID
				var addressValue string

				iter := session.Query(
					"SELECT address_id, address_value"+
						" FROM address_indexed_by_address_key"+
						" WHERE"+
						" tenant_id = ?"+
						" AND application_id = ?"+
						" AND address_key = ?",
					tenantID.String(),
					applicationID.String(),
					key).Iter()

				defer iter.Close()

				Expect(iter.Scan(&id, &addressValue)).To(BeFalse())

			}
		})
	})
})

func TestDeleteBehaviour(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Delete method behaviour")
}
