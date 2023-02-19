package model

import (
	"time"

	"gorm.io/gorm"
)

type Medicine struct {
	gorm.Model
	Animal              string    `json:"animal"`
	MedicineName        string    `json:"medicineName"`
	ApplicationDate     time.Time `json:"applicationDate"`
	NextApplicationDate time.Time `json:"nextApplicationDate"`
}
