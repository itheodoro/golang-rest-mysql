package models

// Dog model
type Dog struct {
	ID     int    `gorm:"AUTO_INCREMENT" json:"id"`
	Name   string `gorm:"not null" json:"name"`
	Gender string `gorm:"not null" json:"gender"`
	Age    uint   `gorm:"not null" json:"age"`
}
