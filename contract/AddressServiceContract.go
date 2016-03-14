// Defines the Address Service contract.
package contract

import (
	. "github.com/microbusinesses/AddressService/domain"
	. "github.com/microbusinesses/Micro-Businesses-Core/system"
)

// The address service contract, it can add new address and update/retrieve/remove an existing address.
type AddressService interface {
	// Creates a new address.
	// tenantId: Mandatory. The unique identifier of the tenant owning the address.
	// address: Mandatory. The reference to the new address information.
	// Returns either the unique identifier of the new address or error if something goes wrong.
	Create(tenantId UUID, address Address) (UUID, error)

	// Updates an existing address.
	// tenantId: Mandatory. The unique identifier of the tenant owning the address.
	// addressId: Mandatory. The unique identifier of the existing address.
	// address: Mandatory. The reeference to the updated address information.
	// Returns error if something goes wrong.
	Update(tenantId UUID, addressId UUID, address Address) error

	// Retrieves an existing address information and returns the detail of it.
	// tenantId: Mandatory. The unique identifier of the tenant owning the address.
	// addressId: Mandatory. The unique identifier of the existing address.
	// Returns either the address information or error if something goes wrong.
	Read(tenantId UUID, addressId UUID) (Address, error)

	// Deletes an existing address information.
	// tenantId: Mandatory. The unique identifier of the tenant owning the address.
	// addressId: Mandatory. The unique identifier of the existing address to remove.
	// Returns error if something goes wrong.
	Delete(tenantId UUID, addressId UUID) error
}