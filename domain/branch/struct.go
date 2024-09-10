package branch

import "time"

type BranchInfoStruct struct {
	Name        string    `json:"name" binding:"required"`
	Address     string    `json:"address" binding:"required"`
	PhoneNumber string    `json:"phoneNumber" binding:"required"`
}

type BranchStruct struct {
	BranchInfoStruct
	Id          string    `json:"id" binding:"required"`
	CreatedAt   time.Time `json:"createdAt" binding:"required"`
	UpdatedAt   time.Time `json:"updatedAt" binding:"required"`
}
