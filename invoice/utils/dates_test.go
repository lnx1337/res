package utils

import (
	"testing"
	"time"
)

func TestGetDateInDays(t *testing.T) {
	var time = time.Now()
	var result = GetDateInDays(time)
	if result == 0 {
		t.Error("Expected > 0", "result: ", result)
	}
}

func TestDateToHours(t *testing.T) {
	var time = time.Now()
	var result = DateToHours(time)

	if result == 0 {
		t.Error("Expected > 0", "result: ", result)
	}
}

func TestGetMidDate(t *testing.T) {
	var midDate = "2017-06-25"
	var result = GetMidDate("2017-01-01", "2017-12-30")
	if result != midDate {
		t.Error("expected", midDate, "result", result)
	}
}
