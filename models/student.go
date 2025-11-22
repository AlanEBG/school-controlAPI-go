package models

import "time"

type Student struct {
	StudentID int       `gorm:"primaryKey;autoIncrement" json:"student_id"`
	Name      string    `gorm:"not null" json:"name" binding:"required"`
	Group     string    `gorm:"not null" json:"group" binding:"required"`
	Email     string    `gorm:"unique;not null" json:"email" binding:"required,email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Grades    []Grade   `gorm:"foreignKey:StudentID;constraint:OnDelete:CASCADE" json:"grades,omitempty"`
}

func (Student) TableName() string {
	return "students"
}
