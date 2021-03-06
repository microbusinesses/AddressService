package service_test

import (
	"errors"
	"math/rand"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/micro-business/AddressService/business/domain"
	"github.com/micro-business/AddressService/business/service"
	"github.com/micro-business/AddressService/data/contract"
	"github.com/micro-business/Micro-Business-Core/system"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("ReadAll method input parameters and dependency test", func() {
	var (
		mockCtrl               *gomock.Controller
		addressService         *service.AddressService
		mockAddressDataService *MockAddressDataService
		tenantID               system.UUID
		applicationID          system.UUID
		addressID              system.UUID
	)

	BeforeEach(func() {
		mockCtrl = gomock.NewController(GinkgoT())
		mockAddressDataService = NewMockAddressDataService(mockCtrl)

		addressService = &service.AddressService{AddressDataService: mockAddressDataService}

		tenantID, _ = system.RandomUUID()
		applicationID, _ = system.RandomUUID()
		addressID, _ = system.RandomUUID()
	})

	AfterEach(func() {
		mockCtrl.Finish()
	})

	Context("when address data service not provided", func() {
		It("should panic", func() {
			addressService.AddressDataService = nil

			Ω(func() { addressService.ReadAll(tenantID, applicationID, addressID) }).Should(Panic())
		})
	})

	Describe("Input Parameters", func() {
		It("should panic when empty tenant unique identifier provided", func() {
			Ω(func() { addressService.ReadAll(system.EmptyUUID, applicationID, addressID) }).Should(Panic())
		})

		It("should panic when empty application unique identifier provided", func() {
			Ω(func() { addressService.ReadAll(tenantID, system.EmptyUUID, addressID) }).Should(Panic())
		})

		It("should panic when empty address unique identifier provided", func() {
			Ω(func() { addressService.ReadAll(tenantID, applicationID, system.EmptyUUID) }).Should(Panic())
		})
	})
})

var _ = Describe("ReadAll method behaviour", func() {
	var (
		mockCtrl               *gomock.Controller
		addressService         *service.AddressService
		mockAddressDataService *MockAddressDataService
		tenantID               system.UUID
		applicationID          system.UUID
		addressID              system.UUID
	)

	BeforeEach(func() {
		mockCtrl = gomock.NewController(GinkgoT())
		mockAddressDataService = NewMockAddressDataService(mockCtrl)

		addressService = &service.AddressService{AddressDataService: mockAddressDataService}

		tenantID, _ = system.RandomUUID()
		applicationID, _ = system.RandomUUID()
		addressID, _ = system.RandomUUID()
	})

	AfterEach(func() {
		mockCtrl.Finish()
	})

	It("should call address data service ReadAll function", func() {
		mockAddressDataService.EXPECT().ReadAll(tenantID, applicationID, addressID)

		addressService.ReadAll(tenantID, applicationID, addressID)
	})

	Context("when address data service succeeds to read the requested address", func() {
		It("should return no error", func() {
			addressDetails := make(map[string]string)

			for idx := 0; idx < rand.Intn(10)+1; idx++ {
				key, _ := system.RandomUUID()
				value, _ := system.RandomUUID()

				addressDetails[key.String()] = value.String()
			}

			expectedAddress := domain.Address{AddressDetails: addressDetails}
			mockAddressDataService.
				EXPECT().
				ReadAll(tenantID, applicationID, addressID).
				Return(contract.Address{AddressDetails: expectedAddress.AddressDetails}, nil)

			address, err := addressService.ReadAll(tenantID, applicationID, addressID)

			Expect(address).To(Equal(expectedAddress))
			Expect(err).To(BeNil())
		})
	})

	Context("when address data service fails to read the requested address", func() {
		It("should return the error returned by address data service", func() {
			expectedErrorID, _ := system.RandomUUID()
			expectedError := errors.New(expectedErrorID.String())
			mockAddressDataService.
				EXPECT().
				ReadAll(tenantID, applicationID, addressID).
				Return(contract.Address{}, expectedError)

			expectedAddress, err := addressService.ReadAll(tenantID, applicationID, addressID)

			Expect(expectedAddress).To(Equal(domain.Address{}))
			Expect(err).To(Equal(expectedError))
		})
	})
})

func TestReadAll(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "ReadAll method input parameters and dependency test")
	RunSpecs(t, "ReadAll method behaviour")
}
