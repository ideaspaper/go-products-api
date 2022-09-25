package response

type Product struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Quantity    int    `json:"quantity"`
}

type ProductWithPagination struct {
	Count     int        `json:"count"`
	Page      int        `json:"page"`
	TotalPage int        `json:"total_page"`
	Products  []*Product `json:"products"`
}
