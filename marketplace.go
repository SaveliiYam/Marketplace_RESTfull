package marketplace

type ProductList struct {
	Id          int    `json:"id" db:"id"`
	Title       string `json:"title" db:"title" binding:"required"`
	Description string `json:"description" db:"description"`
	Price       string `json:"price" binding:"required"`
	Brand       string `json:"brand" db:"brand_id"`
	Category    string `json:"category" db:"categories_id"`
}

type CategoriesList struct {
	Id    int    `json:"id" db:"id"`
	Title string `json:"title" binding:"required"`
}

type BrandsList struct {
	Id          int    `json:"id" db:"id"`
	Title       string `json:"title" binding:"required"`
	Description string `json:"description" db:"description"`
}
