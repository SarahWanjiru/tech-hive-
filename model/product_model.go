package model

type ProductModel struct {
	Id          string  `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Stock       int32   `json:"stock"`
	ImageUrl    string  `json:"image_url"`
}

type ProductCreateOrUpdateModel struct {
 	Name        string  `json:"name" validate:"required"`
 	Description string  `json:"description"`
 	Price       float64 `json:"price" validate:"required,min=0"`
 	Stock       int32   `json:"stock" validate:"required,min=0"`
 	ImageUrl    string  `json:"image_url"`
 }

type ProductSearchModel struct {
 	Name        string  `json:"name,omitempty"`
 	MinPrice    float64 `json:"min_price,omitempty"`
 	MaxPrice    float64 `json:"max_price,omitempty"`
 	InStock     *bool   `json:"in_stock,omitempty"`
 	Page        int     `json:"page,omitempty"`
 	Limit       int     `json:"limit,omitempty"`
 	SortBy      string  `json:"sort_by,omitempty"`      // name, price, created_at
 	SortOrder   string  `json:"sort_order,omitempty"`   // asc, desc
 }
