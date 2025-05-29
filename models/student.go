package models

type level string

const (
	BACHELOR  level = "BACHELOR"
	MASTER    level = "MASTER"
	DOCTORATE level = "DOCTORATE"
)

type gender string

const (
	MALE        gender = "MALE"
	FEMALE      gender = "FEMALE"
	OTHER       gender = "OTHER"
	UNSPECIFIED gender = "UNSPECIFIED"
)

type Student struct {
	ID        uint   `json:"id" gorm:"primaryKey"`
	FirstName string `json:"firstName" gorm:"column:first_name"`
	LastName  string `json:"lastName" gorm:"column:last_name"`
	Email     string `json:"email"`
	Level     level  `json:"level"`
	Gender    gender `json:"gender"`
}
