package models

import "time"

type User struct {
	ID        string       `bson:"_id,omitempty"`
	Email     UserEmail    `bson:"email"`
	Level     int          `bson:"level"`
	Password  UserPassword `bson:"password"`
	Data      UserData     `bson:"data"`
	Ring      UserRing     `bson:"ring"`
	Personal  UserPersonal `bson:"personal"`
	DeletedAt *time.Time   `bson:"deleted_at,omitempty"`
	CreatedAt time.Time    `json:"created_at" bson:"created_at"`
	CreatedBy string       `json:"created_by" bson:"created_by"`
	UpdatedAt time.Time    `json:"updated_at" bson:"updated_at"`
	UpdatedBy string       `json:"updated_by" bson:"updated_by"`
}

type UserEmail struct {
	Value         string     `bson:"value"`
	RequestValue  *string    `bson:"request_value,omitempty"`
	RequestForgot bool       `bson:"request_forgot"`
	RequestChange bool       `bson:"request_change"`
	RequestExpire *time.Time `bson:"request_expire,omitempty"`
	Token         *string    `bson:"token,omitempty"`
	History       []string   `bson:"history"`
}

type UserPassword struct {
	Value         string     `bson:"value"`
	RequestValue  *string    `bson:"request_value,omitempty"`
	RequestForgot bool       `bson:"request_forgot"`
	RequestChange bool       `bson:"request_change"`
	RequestExpire *time.Time `bson:"request_expire,omitempty"`
	Token         *string    `bson:"token,omitempty"`
	History       []string   `bson:"history"`
}

type UserData struct {
	Name       string    `bson:"name"`
	BirthDate  time.Time `bson:"birth_date"`
	Gender     string    `bson:"gender"`
	AckTOS     bool      `bson:"acknowledged_tos"`
	FirstLogin bool      `bson:"first_login"`
}

type UserRing struct {
	MAC          *string    `bson:"mac,omitempty"`
	PurchaseDate *time.Time `bson:"purchase_date,omitempty"`
	Size         int        `bson:"size"`
	Color        string     `bson:"color"`
	Connection   bool       `bson:"connection"`
}

type UserPersonal struct {
	Health    UserHealth   `bson:"health"`
	Physical  UserPhysical `bson:"physical"`
	Habit     UserHabit    `bson:"habit"`
	Goals     UserGoals    `bson:"goals"`
	UpdatedAt time.Time    `bson:"updated_at"`
}

type UserHealth struct {
	Allergies     []string `bson:"allergies"`
	Diseases      []string `bson:"diseases"`
	Goals         []string `bson:"goals"`
	BloodType     string   `bson:"blood_type"`
	Issues        []string `bson:"issues"`
	SpecificGoals []string `bson:"specific_goals"`
}

type UserPhysical struct {
	Height    float64 `bson:"height"`
	Weight    float64 `bson:"weight"`
	Abdominal float64 `bson:"abdominal"`
	Waist     float64 `bson:"waist"`
	Unit      string  `bson:"unit"`
}

type UserHabit struct {
	Smoke   bool `bson:"smoke"`
	Alcohol bool `bson:"alcohol"`
}

type UserGoals struct {
	Health  []string `bson:"health"`
	Sports  []string `bson:"sports"`
	Medical []string `bson:"medical"`
}
