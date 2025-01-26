package dto

type FitnessDTO struct {
	Title        string     `json:"title"`
	Image        string     `json:"image"`
	Description  string     `json:"description"`
	Category     string     `json:"category"`
	BodyCategory string     `json:"body_category"`
	Trait        string     `json:"trait"`
	Video        []VideoDTO `json:"video"`
	Workout      string     `json:"workout"`
	Deleted      bool       `json:"deleted"`
}

type VideoDTO struct {
	Title     string `json:"title"`
	URL       string `json:"url"`
	Duration  string `json:"duration"`
	Thumbnail string `json:"thumbnail"`
}
