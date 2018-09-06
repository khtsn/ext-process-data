package process

import (
	"bufio"
	"fmt"
	"github.com/khtsn/ext-process-data/models"
	"github.com/khtsn/ext-process-data/utils"
	"strings"
)

// GetPhoneDetail lookup file for the phone number provided
func (p *Process) GetPhoneDetail(phone string) (models.PhoneDetail, bool, error) {
	//retro cursor to beginning of file
	p.File.Seek(0, 0)
	//scan line by line
	scanner := bufio.NewReader(p.File)
	ind := 0
	for {
		line, err := scanner.ReadString('\n')
		if err != nil {
			break
		}
		if strings.Contains(line, phone) {
			phoneDetail, err := utils.ParseLine(line)
			if err != nil {
				return phoneDetail, true, err
			}
			phoneDetail.CursorPosition = int64(ind)
			return phoneDetail, true, nil
		}
		ind += len(line)
	}

	return models.PhoneDetail{}, false, nil
}

// UpdatePhoneDetail update tmp file for the phone number provided with cursor number
func (p *Process) UpdatePhoneDetail(phoneDetail *models.PhoneDetail) error {
	line := fmt.Sprintf("%s,%s,%s\n", phoneDetail.Phone, phoneDetail.ActivationDate, phoneDetail.DeactivationDate)
	return utils.WriteToFile(p.File, line, phoneDetail.CursorPosition)
}

// NewPhoneDetail append new phone number detail
func (p *Process) NewPhoneDetail(phoneDetail *models.PhoneDetail) error {
	line := fmt.Sprintf("%s,%s,%s\n", phoneDetail.Phone, phoneDetail.ActivationDate, phoneDetail.DeactivationDate)
	return utils.AppendToFile(p.File, line)
}
