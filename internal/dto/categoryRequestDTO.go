package dto

type CreateCategoryRequest struct {
	Name         string `json:"name"`
	ParentID     int   `json:"parent_id"`
	ImageUrl     string `json:"image_url"`
	DisplayOrder int    `json:"display_order"`
}
