package util

import (
	"testing"
	"time"
)

func TestBeforeEpoc_No(t *testing.T) {
	epoch, _ := time.Parse(time.RFC822, "01 Jan 70 00:01 UTC")

	isBeforeEpoch := BeforeEpoc(&epoch)

	if isBeforeEpoch {
		t.Error("Shouldn't have been before epoch")
	}
}

func TestBeforeEpoc_Yes(t *testing.T) {
	epoch, _ := time.Parse(time.RFC822, "01 Jan 69 00:01 UTC")

	isBeforeEpoch := BeforeEpoc(&epoch)

	if !isBeforeEpoch {
		t.Error("Shouldn't have been before epoch")
	}
}

func TestNewTimeNow(t *testing.T) {
	now := time.Now()
	newTimeNow := NewTimeNow()

	expectedYear, expectedMonth, expectedDay := now.Date()
	receivedYear, receivedMonth, receivedDay := newTimeNow.Date()

	//Hopefully this NEVER fails... every now and then it could I suppose.
	expectedHour, expectedMinute, expectedSecond := now.Hour(), now.Minute(), now.Second()
	receivedHour, receivedMinute, receivedSecond := newTimeNow.Hour(), newTimeNow.Minute(), newTimeNow.Second()

	if expectedYear != receivedYear {
		t.Errorf("Received year: (%d) but expected: (%d)", receivedYear, expectedYear)
	}

	if expectedMonth != receivedMonth {
		t.Errorf("Received month: (%d) but expected: (%d)", receivedMonth, expectedMonth)
	}

	if expectedDay != receivedDay {
		t.Errorf("Received day: (%d) but expected: (%d)", receivedDay, expectedDay)
	}

	if expectedHour != receivedHour {
		t.Errorf("Received hour: (%d) but expected: (%d)", receivedHour, expectedHour)
	}

	if expectedMinute != receivedMinute {
		t.Errorf("Received minute: (%d) but expected: (%d)", receivedMinute, expectedMinute)
	}

	if expectedSecond != receivedSecond {
		t.Errorf("Received second: (%d) but expected: (%d)", receivedSecond, expectedSecond)
	}
}
