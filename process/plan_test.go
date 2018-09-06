package process

import (
	"github.com/khtsn/ext-process-data/models"
	"github.com/khtsn/ext-process-data/utils"
	"testing"
)

func TestIsRecentOwner(t *testing.T) {
	currentPhoneDetail := models.PhoneDetail{Phone: "09012345678", ActivationDate: "2016-01-01", DeactivationDate: "2016-03-01"}
	newPhoneDetail := models.PhoneDetail{Phone: "09012345678", ActivationDate: "2016-12-01", DeactivationDate: utils.NoDeactivateDate}

	actual := isRecentOwner(currentPhoneDetail, newPhoneDetail)
	if !actual {
		t.Fail()
	}
}

func TestIsRenew(t *testing.T) {
	currentPhoneDetail := models.PhoneDetail{Phone: "09012345678", ActivationDate: "2016-02-01", DeactivationDate: "2016-03-01"}
	newPhoneDetail := models.PhoneDetail{Phone: "09012345678", ActivationDate: "2016-03-01", DeactivationDate: "2016-05-01"}

	actual := isRenew(currentPhoneDetail, newPhoneDetail)
	if !actual {
		t.Fail()
	}
}

func TestIsHistory(t *testing.T) {
	currentPhoneDetail := models.PhoneDetail{Phone: "09012345678", ActivationDate: "2016-03-01", DeactivationDate: "2016-05-01"}
	newPhoneDetail := models.PhoneDetail{Phone: "09012345678", ActivationDate: "2016-02-01", DeactivationDate: "2016-03-01"}

	actual := isHistory(currentPhoneDetail, newPhoneDetail)
	if !actual {
		t.Fail()
	}
}
