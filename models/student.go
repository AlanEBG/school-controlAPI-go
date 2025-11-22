package models

import "gorm.io/gorm"

type Student struct {
	gorm.Model
	StudentID int     `gorm:"primaryKey;autoIncrement" json:"student_id"`
	Name      string  `gorm:"not null" json:"name" binding:"required"`
	Group     string  `gorm:"not null" json:"group" binding:"required"`
	Email     string  `gorm:"unique;not null" json:"email" binding:"required,email"`
	Grades    []Grade `gorm:"foreignKey:StudentID;constraint:OnDelete:CASCADE" json:"grades,omitempty"`
}

// TableName especifica el nombre de la tabla
func (Student) TableName() string {
	return "students"
}
