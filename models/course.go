package models

type mode string

const (
	IN_PERSON mode = "IN_PERSON"
	ONLINE    mode = "ONLINE"
)

type Course struct {
	ID          uint   `json:"id" gorm:"primaryKey"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Mode        mode   `json:"mode"`
}
