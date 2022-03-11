package models

import (
	"time"
)

// ExampleReq struct
type ExampleReq struct {
	Name        *string `json:"name"`
	Description *string `json:"description"`
}

// Example struct
type Example struct {
	ID *int64 `json:"id"`
	ExampleReq
	Updatedat *time.Time `json:"updatedat"`
	Createdat *time.Time `json:"createdat"`
}
