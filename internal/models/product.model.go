package models

type Product struct {
	Id              int     `json:"id"`
	CategoryId      int     `json:"category_id"`
	FavoriteProduct bool    `json:"favorite_product"`
	Name            string  `json:"name"`
	Description     string  `json:"description"`
	Price           float64 `json:"price"`
	CampaignId      int     `json:"campaign_id"`
	Stock           string  `json:"stock"`
}

type CreateProductRequest struct {
	CategoryId      int     `json:"category_id"`
	FavoriteProduct bool    `json:"favorite_product"`
	Name            string  `json:"name"`
	Description     string  `json:"description"`
	Price           float64 `json:"price"`
	CampaignId      int     `json:"campaign_id"`
	Stock           string  `json:"stock"`
}

type GetProductResponse struct {
	Id              int     `json:"id"`
	CategoryId      int     `json:"category_id"`
	FavoriteProduct bool    `json:"favorite_product"`
	Name            string  `json:"name"`
	Description     string  `json:"description"`
	Price           float64 `json:"price"`
	CampaignId      int     `json:"campaign_id"`
	Stock           string  `json:"stock"`
}

type UpdateProductRequest struct {
	Id              int     `json:"id"`
	CategoryId      int     `json:"category_id"`
	FavoriteProduct bool    `json:"favorite_product"`
	Name            string  `json:"name"`
	Description     string  `json:"description"`
	Price           float64 `json:"price"`
	CampaignId      int     `json:"campaign_id"`
	Stock           string  `json:"stock"`
}

type DeleteProductRequest struct {
	Id int `json:"id"`
}
