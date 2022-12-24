/*
Package student test contain everything for student data,
such as model, route handler, utility, etc.
*/
package student

import (
	"time"
)

// Student student data model
type Student struct {
	ID            int       `json:"id" gorm:"primarykey"`
	StudentNumber int       `json:"student_number" gorm:"student_number;not null;unique"`
	FullName      string    `json:"full_name" gorm:"full_name;not null"`
	FullAddress   string    `json:"full_address" gorm:"full_address;not null"`
	YearOfEnroll  int       `json:"year_of_enroll" gorm:"year_of_enroll;not null"`
	CreatedAt     time.Time `json:"created_at" gorm:"created_at;not null"`
	UpdatedAt     time.Time `json:"updated_at" gorm:"updated_at;not null"`
}
