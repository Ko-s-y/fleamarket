package dto

type CreateItemInput struct {
	Name        string `json:"name" binding:"required, min=3"`
	Price       uint   `json:"price" binding:"required, min=1,max=1000000"`
	Description string `json:"description"`
}
