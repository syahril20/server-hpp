package questionnaire

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Assesemnt
type CreateAssesmentRequest struct {
	Id            primitive.ObjectID           `json:"_id" bson:"_id,omitempty"`
	Name          string                       `json:"name" bson:"name"`
	DeletedAt     *time.Time                   `json:"deleted_at" bson:"deleted_at"`
	Questionnaire []CreateQuestionnaireRequest `json:"questionnaire" bson:"questionnaire"`
	CreatedAt     time.Time                    `json:"created_at" bson:"created_at"`
	CreatedBy     string                       `json:"created_by" bson:"created_by"`
	UpdatedAt     time.Time                    `json:"updated_at" bson:"updated_at"`
	UpdatedBy     string                       `json:"updated_by" bson:"updated_by"`
}

type UpdateAssesmentRequest struct {
	Name      string    `json:"name" bson:"name"`
	UpdatedAt time.Time `json:"updated_at" bson:"updated_at"`
	UpdatedBy string    `json:"updated_by" bson:"updated_by"`
}

type ActiveAssesmentRequest struct {
	DeletedAt *time.Time `json:"deleted_at" bson:"deleted_at"`
	UpdatedAt time.Time  `json:"updated_at" bson:"updated_at"`
	UpdatedBy string     `json:"updated_by" bson:"updated_by"`
}

// Questionnaire
type CreateQuestionnaireRequest struct {
	Id          primitive.ObjectID    `json:"_id" bson:"_id,omitempty"`
	IdAsessment primitive.ObjectID    `json:"_id_assesment" bson:"_id_assesment"`
	Question    string                `json:"question" bson:"question"`
	DeletedAt   *time.Time            `json:"deleted_at" bson:"deleted_at"`
	Answer      []CreateAnswerRequest `json:"answer" bson:"answer"`
	CreatedAt   time.Time             `json:"created_at" bson:"created_at"`
	CreatedBy   string                `json:"created_by" bson:"created_by"`
	UpdatedAt   time.Time             `json:"updated_at" bson:"updated_at"`
	UpdatedBy   string                `json:"updated_by" bson:"updated_by"`
}

type UpdateQuestionnaireRequest struct {
	Question  string    `json:"question" bson:"question"`
	UpdatedAt time.Time `json:"updated_at" bson:"updated_at"`
	UpdatedBy string    `json:"updated_by" bson:"updated_by"`
}

type ActiveQuestionnaireRequest struct {
	DeletedAt *time.Time `json:"deleted_at" bson:"deleted_at"`
	UpdatedAt time.Time  `json:"updated_at" bson:"updated_at"`
	UpdatedBy string     `json:"updated_by" bson:"updated_by"`
}

// Answer
type CreateAnswerRequest struct {
	Id              primitive.ObjectID `json:"_id" bson:"_id,omitempty"`
	IdQuestionnaire primitive.ObjectID `json:"_id_questionnaire" bson:"_id_questionnaire"`
	Value           bool               `json:"value" bson:"value"`
	Description     string             `json:"description" bson:"description"`
	DeletedAt       *time.Time         `json:"deleted_at" bson:"deleted_at"`
	CreatedAt       time.Time          `json:"created_at" bson:"created_at"`
	CreatedBy       string             `json:"created_by" bson:"created_by"`
	UpdatedAt       time.Time          `json:"updated_at" bson:"updated_at"`
	UpdatedBy       string             `json:"updated_by" bson:"updated_by"`
}

type UpdateAnswerRequest struct {
	Description string    `json:"description" bson:"description"`
	Value       bool      `json:"value" bson:"value"`
	UpdatedAt   time.Time `json:"updated_at" bson:"updated_at"`
	UpdatedBy   string    `json:"updated_by" bson:"updated_by"`
}

type ActiveAnswerRequest struct {
	DeletedAt *time.Time `json:"deleted_at" bson:"deleted_at"`
	UpdatedAt time.Time  `json:"updated_at" bson:"updated_at"`
	UpdatedBy string     `json:"updated_by" bson:"updated_by"`
}
