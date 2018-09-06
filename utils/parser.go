package utils

import (
	"encoding/csv"
	"errors"
	"fmt"
	"github.com/khtsn/ext-process-data/models"
	"strings"
)

// ParseLine get line data and parse to model
func ParseLine(line string) (phoneDetail models.PhoneDetail, err error) {
	reader := csv.NewReader(strings.NewReader(line))

	record, err := reader.Read()
	if err != nil {
		return phoneDetail, err
	}

	//verify records
	if len(record) != 3 {
		return phoneDetail, errors.New(fmt.Sprintf("invalid input"))
	}

	//new phoneDetail object
	phoneDetail = models.PhoneDetail{Phone: record[0], ActivationDate: record[1], DeactivationDate: record[2]}
	if strings.TrimSpace(record[2]) == "" { //in use
		phoneDetail.DeactivationDate = NoDeactivateDate
	}

	return phoneDetail, nil
}
