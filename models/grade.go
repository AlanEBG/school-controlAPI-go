package models

import "time"

type Grade struct {
	GradeID   int       `gorm:"primaryKey;autoIncrement" json:"grade_id"`
	StudentID int       `gorm:"not null" json:"student_id" binding:"required"`
	SubjectID int       `gorm:"not null" json:"subject_id" binding:"required"`
	Grade     float64   `gorm:"not null" json:"grade" binding:"required,min=0,max=100"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Student   *Student  `gorm:"foreignKey:StudentID;constraint:OnDelete:CASCADE" json:"student,omitempty"`
	Subject   *Subject  `gorm:"foreignKey:SubjectID;constraint:OnDelete:CASCADE" json:"subject,omitempty"`
}

func (Grade) TableName() string {
	return "grades"
}
