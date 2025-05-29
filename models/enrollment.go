package models

import "time"

type Enrollment struct {
	ID             uint      `json:"id" gorm:"primaryKey"`
	StudentID      uint      `json:"studentId" gorm:"column:student_id"`
	Student        Student   `json:"student"`
	CourseID       uint      `json:"courseId" gorm:"column:course_id"`
	Course         Course    `json:"course"`
	EnrollmentDate time.Time `json:"enrollmentDate" gorm:"column:enrollment_date"`
}
