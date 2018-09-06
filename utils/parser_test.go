package utils

import (
	"testing"
)

func TestParseLineSuccess(t *testing.T) {
	line := `0987000002,2016-02-01,2016-03-01`
	_, err := ParseLine(line)
	if err != nil {
		t.Fail()
	}
}

func TestParseLineError(t *testing.T) {
	line := "0987000002,"
	_, err := ParseLine(line)
	if err == nil {
		t.Fail()
	}
}
