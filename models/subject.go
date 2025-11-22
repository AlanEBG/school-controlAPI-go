package models

import "gorm.io/gorm"

type Subject struct {
	gorm.Model
	SubjectID int     `gorm:"primaryKey;autoIncrement" json:"subject_id"`
	Name      string  `gorm:"not null" json:"name" binding:"required"`
	Grades    []Grade `gorm:"foreignKey:SubjectID;constraint:OnDelete:CASCADE" json:"grades,omitempty"`
}

// TableName especifica el nombre de la tabla
func (Subject) TableName() string {
	return "subjects"
}
