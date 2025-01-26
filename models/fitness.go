package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Video struct {
	Title     string `json:"title" bson:"title"`
	URL       string `json:"url" bson:"url"`
	Duration  string `json:"duration" bson:"duration"`
	Thumbnail string `json:"thumbnail" bson:"thumbnail"`
}

type Fitness struct {
	ID           primitive.ObjectID `json:"id" bson:"_id"`
	Title        string             `json:"title" bson:"title"`
	Image        string             `json:"image" bson:"image"`
	Description  string             `json:"description" bson:"description"`
	Category     string             `json:"category" bson:"category"`
	BodyCategory string             `json:"body_category" bson:"body_category"`
	Trait        string             `json:"trait" bson:"trait"`
	Video        []Video            `json:"video" bson:"video"`
	Workout      string             `json:"workout" bson:"workout"`
	Deleted      bool               `json:"deleted" bson:"deleted"`
	CreatedAt    time.Time          `json:"created_at" bson:"created_at"`
	CreatedBy    string             `json:"created_by" bson:"created_by"`
	UpdatedAt    time.Time          `json:"updated_at" bson:"updated_at"`
	UpdatedBy    string             `json:"updated_by" bson:"updated_by"`
}
