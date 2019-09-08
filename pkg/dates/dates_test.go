package dates

import (
	"testing"
	"time"
)

func TestGetDate(t *testing.T) {
	today := GetDate()
	expected := time.Now()
	if expected != today {
		t.Errorf("Date mismatch. exp: %s got: %s", expected, today)
	}
}

func TestGetTomorrow(t *testing.T) {
	today := GetDate()
	tomorrow := GetTomorrow()
	expected := today.Add(24 * time.Hour)
	if expected != tomorrow {
		t.Errorf("Date mismatch. exp: %s got: %s", expected, tomorrow)
	}
}
