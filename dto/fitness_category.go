package dto

type FitnessCategoryDTO struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Image       string `json:"image"`
	Deleted     bool   `json:"deleted"`
}
