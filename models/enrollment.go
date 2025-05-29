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

type EnrollmentInfo struct {
	ID             uint      `json:"id"`
	CourseName     string    `json:"courseName"`
	FirstName      string    `json:"firstName"`
	LastName       string    `json:"lastName"`
	StudentLevel   string    `json:"studentLevel"`
	CourseMode     string    `json:"courseMode"`
	EnrollmentDate time.Time `json:"enrollmentDate"`
}

type CreateEnrollmentReq struct {
	StudentID      uint      `json:"studentId"`
	CourseID       uint      `json:"courseId"`
	EnrollmentDate time.Time `json:"enrollmentDate"`
}

func (EnrollmentInfo) TableName() string {
	return "enrollments_info"
}
