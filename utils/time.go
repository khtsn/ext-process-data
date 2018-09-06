package utils

import (
	"github.com/khtsn/ext-process-data/models"
	"time"
)

// GetActivationTime get activation date in time
func GetActivationTime(p models.PhoneDetail) time.Time {
	activationTime, err := time.Parse(TimeFormat, p.ActivationDate)
	if err != nil {
		panic(err)
	}

	return activationTime
}

// GetDeactivationTime get activation date in time
func GetDeactivationTime(p models.PhoneDetail) time.Time {
	deactivationTime, err := time.Parse(TimeFormat, p.DeactivationDate)
	if err != nil {
		panic(err)
	}

	return deactivationTime
}
