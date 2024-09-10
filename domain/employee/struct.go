package employee

import "time"

type EmployeeInfoStruct struct {
	FirstName      string `json:"firstName" binding:"required" db:"first_name"`
	LastName       string `json:"lastName" binding:"required" db:"last_name"`
	Email          string `json:"email" binding:"required" db:"email"`
	Floor          int    `json:"floor" binding:"required" db:"floor"`
	PhoneExtension string `json:"phoneExtension" binding:"required" db:"phone_extension"`
	Role           string `json:"role" binding:"required" db:"role"`
	BranchId       string `json:"branchId" binding:"required" db:"branch_id"`
}

type MetaData struct {
	CreatedAt time.Time `json:"createdAt" binding:"required" db:"created_at"`
	UpdatedAt time.Time `json:"updatedAt" binding:"required" db:"updated_at"`
}

type EmployeeStruct struct {
	Id string `json:"id" binding:"required"`
	EmployeeInfoStruct
	MetaData
}
