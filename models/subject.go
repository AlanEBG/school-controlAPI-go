package models

import "time"

type Subject struct {
	SubjectID int       `gorm:"primaryKey;autoIncrement" json:"subject_id"`
	Name      string    `gorm:"not null" json:"name" binding:"required"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Grades    []Grade   `gorm:"foreignKey:SubjectID;constraint:OnDelete:CASCADE" json:"grades,omitempty"`
}

func (Subject) TableName() string {
	return "subjects"
}
