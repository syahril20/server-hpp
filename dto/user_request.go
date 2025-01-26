package dto

// CreateUserRequest adalah struct untuk permintaan pembuatan user baru
type CreateUserRequest struct {
	Email     Email    `json:"email" binding:"required"`
	Password  Password `json:"password" binding:"required"`
	BirthDate string   `json:"birthDate" binding:"required"`
	Level     int      `json:"level" binding:"required"`
	Name      string   `json:"name" binding:"required"`
	Gender    string   `json:"gender" binding:"required"`
	Ring      Ring     `json:"ring"`
	Personal  Personal `json:"personal"`
}

// Ring adalah struct untuk data cincin
type Ring struct {
	Size       int    `json:"size"`
	Color      string `json:"color"`
	Connection bool   `json:"connection"`
}

// Personal adalah struct untuk data pribadi pengguna
type Personal struct {
	Health   Health   `json:"health"`
	Physical Physical `json:"physical"`
	Habit    Habit    `json:"habit"`
}

// Health adalah struct untuk data kesehatan pengguna
type Health struct {
	Allergies     []string `json:"allergies"` // Harus string, sesuai dengan format JSON yang diberikan
	Diseases      []string `json:"diseases"`  // Harus string
	Goals         []string `json:"goals"`     // Harus string
	BloodType     string   `json:"bloodType"`
	Issues        []string `json:"issues"` // Harus string
	SpecificGoals []string `json:"specific_goals"`
}

// Physical adalah struct untuk data fisik pengguna
type Physical struct {
	Height    float64 `json:"height"`
	Weight    float64 `json:"weight"`
	Abdominal float64 `json:"abdominal"`
	Waist     float64 `json:"waist"`
	Unit      string  `json:"unit"`
}

// Habit adalah struct untuk kebiasaan pengguna
type Habit struct {
	Smoke   bool `json:"smoke"`
	Alcohol bool `json:"alcohol"`
}

// UpdateUserRequest adalah struct untuk permintaan pembaruan data pengguna
type UpdateUserRequest struct {
	Email    Email    `json:"email"`
	Password Password `json:"password"`
	Data     UserData `json:"data"`
	Personal Personal `json:"personal"`
}

// Password adalah struct untuk data password
type Password struct {
	Value         string  `json:"value"`
	RequestForgot bool    `json:"request_forgot"`
	RequestChange bool    `json:"request_change"`
	History       *string `json:"history"` // Bisa null, jadi tipe pointer string
}

// UserData adalah struct untuk data pengguna
type UserData struct {
	Name            string `json:"name"`
	BirthDate       string `json:"birth_date"`
	Gender          string `json:"gender"`
	AcknowledgedTOS bool   `json:"acknowledged_tos"`
	FirstLogin      bool   `json:"first_login"`
}

// Email adalah struct untuk data email
type Email struct {
	Value         string  `json:"value"`
	RequestForgot bool    `json:"request_forgot"`
	RequestChange bool    `json:"request_change"`
	History       *string `json:"history"` // Bisa null, jadi tipe pointer string
}
