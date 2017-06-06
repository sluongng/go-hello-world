package main

import (
	"github.com/satori/uuid"
	"time"
)

// Course struct: Exported fields must capitalize first letter
type Course struct {
	ID        uuid.UUID	`json:"Id"`
	Name      string	`json:"Name"`
	StartTime time.Time	`json:"StartTime"`
	EndTime   time.Time	`json:"ExpireTime"`
}

// Courses struct
type Courses map[string]*Course