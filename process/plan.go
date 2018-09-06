package process

import (
	"github.com/khtsn/ext-process-data/models"
	"github.com/khtsn/ext-process-data/utils"
)

// ProcessPhonePlan check phone detail business logic
func (p *Process) ProcessPhonePlan(new *models.PhoneDetail) error {
	//check if exist in file
	current, found, err := p.GetPhoneDetail(new.Phone)
	if err != nil {
		return err
	}
	if !found {
		err = p.NewPhoneDetail(new)
		return nil
	}
	//recent owner
	if isRecentOwner(current, *new) {
		return p.UpdatePhoneDetail(&models.PhoneDetail{
			Phone:            current.Phone,
			ActivationDate:   new.ActivationDate,
			DeactivationDate: new.DeactivationDate,
			CursorPosition:   current.CursorPosition,
		})
	}
	//renew plan
	if isRenew(current, *new) {
		return p.UpdatePhoneDetail(&models.PhoneDetail{
			Phone:            current.Phone,
			ActivationDate:   current.ActivationDate,
			DeactivationDate: new.DeactivationDate,
			CursorPosition:   current.CursorPosition,
		})
	}
	//just historical data
	if isHistory(current, *new) {
		return p.UpdatePhoneDetail(&models.PhoneDetail{
			Phone:            current.Phone,
			ActivationDate:   new.ActivationDate,
			DeactivationDate: current.DeactivationDate,
			CursorPosition:   current.CursorPosition,
		})
	}

	return nil
}

//isRecentOwner is most recent owner
func isRecentOwner(current models.PhoneDetail, new models.PhoneDetail) bool {
	newActivationTime := utils.GetActivationTime(new)
	currentActivationTime := utils.GetActivationTime(current)
	newDeactivationTime := utils.GetDeactivationTime(new)
	currentDeactivationTime := utils.GetDeactivationTime(current)

	return newActivationTime.After(currentActivationTime) &&
		newDeactivationTime.After(currentDeactivationTime) &&
		new.DeactivationDate != current.ActivationDate &&
		new.ActivationDate != current.DeactivationDate
}

//isRenew is extended plan
func isRenew(current models.PhoneDetail, new models.PhoneDetail) bool {
	newActivationTime := utils.GetActivationTime(new)
	currentActivationTime := utils.GetActivationTime(current)

	return newActivationTime.After(currentActivationTime) &&
		(new.DeactivationDate == current.ActivationDate || new.ActivationDate == current.DeactivationDate)
}

//isHistory is history plan log
func isHistory(current models.PhoneDetail, new models.PhoneDetail) bool {
	newActivationTime := utils.GetActivationTime(new)
	currentActivationTime := utils.GetActivationTime(current)

	return newActivationTime.Before(currentActivationTime) &&
		(new.ActivationDate == current.ActivationDate || new.DeactivationDate == current.ActivationDate)
}
