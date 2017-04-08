package util

import "time"

func BeforeEpoc(timeToCheck *time.Time) bool {
	epoch, _ := time.Parse(time.RFC822, "01 Jan 70 00:01 UTC")

	if timeToCheck == nil {
		return true
	}

	return timeToCheck.Before(epoch)
}

func NewTimeNow() *time.Time {
	now := time.Now()
	now = time.Date(now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second(), now.Nanosecond(), now.Location())
	return &now
}
