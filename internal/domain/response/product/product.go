package product

type ProductResponse struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	CategoryID   int    `json:"category_id"`
	Description  string `json:"description"`
	Price        int    `json:"price"`
	CountInStock int    `json:"count_in_stock"`
	Brand        string `json:"brand"`
	Weight       int    `json:"weight"`
	Rating       int    `json:"rating"`
	Slug         string `json:"slug"`
	ImagePath    string `json:"image_path"`
	CreatedAt    string `json:"created_at"`
	UpdatedAt    string `json:"updated_at"`
}
