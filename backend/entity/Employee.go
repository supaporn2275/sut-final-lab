package entity

import "gorm.io/gorm"
type Employee struct {
	gorm.Model
	Name string 		//`validator"required~Name not to be blank"`
	Email string
	EmployeeID string // รหัสพนักงานขึ7นต้นด้วย J หรือ M หรือ S แล้วตามด้วยตัวเลขจํานวน 8 ตัว
}