package domain

import (
	"time"
)

type Medicine struct {
	Animal              string
	MedicineName        string
	ApplicationDate     time.Time
	NextApplicationDate time.Time
}
