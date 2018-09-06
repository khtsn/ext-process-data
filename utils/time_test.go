package utils

import (
	"github.com/khtsn/ext-process-data/models"
	"testing"
	"time"
)

func TestGetActivationTime(t *testing.T) {
	activationDate := "2016-01-12"
	deactivationDate := "2016-02-11"
	phoneDetail := models.PhoneDetail{Phone: "0982323232", ActivationDate: activationDate, DeactivationDate: deactivationDate}

	expectedActivationTime, _ := time.Parse(TimeFormat, activationDate)
	actualActivationTime := GetActivationTime(phoneDetail)

	if expectedActivationTime != actualActivationTime {
		t.Fail()
	}
}

func TestGetDeactivationTime(t *testing.T) {
	activationDate := "2016-01-12"
	deactivationDate := "2016-02-11"
	phoneDetail := models.PhoneDetail{Phone: "0982323232", ActivationDate: activationDate, DeactivationDate: deactivationDate}

	expectedDeactivationTime, _ := time.Parse(TimeFormat, deactivationDate)
	actualDeactivationTime := GetDeactivationTime(phoneDetail)

	if expectedDeactivationTime != actualDeactivationTime {
		t.Fail()
	}
}
