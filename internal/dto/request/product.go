package request

type ProductInsert struct {
	Name        string `json:"name" validate:"required"`
	Description string `json:"description" validate:"required"`
	Quantity    *int   `json:"quantity" validate:"required"`
}

type ProductUpdate struct {
	Name        string `json:"name" validate:"required"`
	Description string `json:"description" validate:"required"`
}

type ProductUpdateQuantity struct {
	Quantity *int `json:"quantity" validate:"required"`
}
