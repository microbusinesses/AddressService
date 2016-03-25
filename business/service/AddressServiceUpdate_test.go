package service_test

import (
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/microbusinesses/AddressService/business/domain"
	"github.com/microbusinesses/AddressService/business/service"
	dataServiceMocks "github.com/microbusinesses/AddressService/data/contract/mocks"
	dataShared "github.com/microbusinesses/AddressService/data/shared"
	"github.com/microbusinesses/Micro-Businesses-Core/system"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Update method input parameters", func() {
	var (
		mockCtrl               *gomock.Controller
		addressService         *service.AddressService
		mockAddressDataService *dataServiceMocks.MockAddressDataService
		tenantId               system.UUID
		applicationId          system.UUID
		addressId              system.UUID
		validAddress           domain.Address
		emptyAddress           domain.Address
	)

	BeforeEach(func() {
		mockCtrl = gomock.NewController(GinkgoT())
		mockAddressDataService = dataServiceMocks.NewMockAddressDataService(mockCtrl)

		addressService = &service.AddressService{AddressDataService: mockAddressDataService}

		tenantId, _ = system.RandomUUID()
		applicationId, _ = system.RandomUUID()
		addressId, _ = system.RandomUUID()
		validAddress = domain.Address{AddressParts: map[string]string{"FirstName": "Morteza"}}
		emptyAddress = domain.Address{}
	})

	AfterEach(func() {
		mockCtrl.Finish()
	})

	Context("when address data service not provided", func() {
		It("should panic", func() {
			addressService.AddressDataService = nil

			Ω(func() { addressService.Update(tenantId, applicationId, addressId, validAddress) }).Should(Panic())
		})
	})

	Context("when empty tenant unique identifier provided", func() {
		It("should panic", func() {
			Ω(func() { addressService.Update(system.EmptyUUID, applicationId, addressId, validAddress) }).Should(Panic())
		})
	})

	Context("when empty application unique identifier provided", func() {
		It("should panic", func() {
			Ω(func() { addressService.Update(tenantId, system.EmptyUUID, addressId, validAddress) }).Should(Panic())
		})
	})

	Context("when empty address unique identifier provided", func() {
		It("should panic", func() {
			Ω(func() { addressService.Update(tenantId, applicationId, system.EmptyUUID, validAddress) }).Should(Panic())
		})
	})

	Context("when address without address parts provided", func() {
		It("should panic", func() {
			Ω(func() { addressService.Update(tenantId, applicationId, addressId, emptyAddress) }).Should(Panic())
		})
	})
})

var _ = Describe("Update method behaviour", func() {
	var (
		mockCtrl               *gomock.Controller
		addressService         *service.AddressService
		mockAddressDataService *dataServiceMocks.MockAddressDataService
		tenantId               system.UUID
		applicationId          system.UUID
		addressId              system.UUID
		validAddress           domain.Address
	)

	BeforeEach(func() {
		mockCtrl = gomock.NewController(GinkgoT())
		mockAddressDataService = dataServiceMocks.NewMockAddressDataService(mockCtrl)

		addressService = &service.AddressService{AddressDataService: mockAddressDataService}

		tenantId, _ = system.RandomUUID()
		applicationId, _ = system.RandomUUID()
		addressId, _ = system.RandomUUID()
		validAddress = domain.Address{AddressParts: map[string]string{"FirstName": "Morteza"}}
	})

	AfterEach(func() {
		mockCtrl.Finish()
	})

	It("should call address data service Update function", func() {
		mappedAddress := dataShared.Address{AddressParts: validAddress.AddressParts}

		mockAddressDataService.EXPECT().Update(tenantId, applicationId, addressId, mappedAddress)

		addressService.Update(tenantId, applicationId, addressId, validAddress)
	})

	Context("when address data service succeeds to update the requested address", func() {
		It("should return no error", func() {
			mappedAddress := dataShared.Address{AddressParts: validAddress.AddressParts}

			mockAddressDataService.
				EXPECT().
				Update(tenantId, applicationId, addressId, mappedAddress).
				Return(nil)

			err := addressService.Update(tenantId, applicationId, addressId, validAddress)

			Expect(err).To(BeNil())
		})
	})

	Context("when address data service fails to update the requested address", func() {
		It("should return the error returned by address data service", func() {
			mappedAddress := dataShared.Address{AddressParts: validAddress.AddressParts}

			expectedErrorId, _ := system.RandomUUID()
			expectedError := errors.New(expectedErrorId.String())
			mockAddressDataService.
				EXPECT().
				Update(tenantId, applicationId, addressId, mappedAddress).
				Return(expectedError)

			err := addressService.Update(tenantId, applicationId, addressId, validAddress)

			Expect(err).To(Equal(expectedError))
		})
	})
})

func TestUpdate(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Update method input parameters")
	RunSpecs(t, "Update method behaviour")
}