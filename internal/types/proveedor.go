package types

import "github.com/google/uuid"

type Proveedor struct {
	ID           uuid.UUID `json:"id"`
	SupplierID   string    `json:"supplier_id"`
	Name         string    `json:"name"`
	ContactName  *string   `json:"contact_name"`
	ContactEmail *string   `json:"contact_email"`
	ContactPhone *string   `json:"contact_phone"`
}
