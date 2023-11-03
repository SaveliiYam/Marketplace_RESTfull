package marketplace

type ProductList struct {
	Id          int    `json:"id" db:"id"`
	Title       string `json:"title" db:"title" binding:"required"`
	Description string `json:"description" db:"description"`
	Brand       string `json:"brand" db:"brand"`
	Price       string `json:"price" binding:"required"`
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
