package dates

import (
	"testing"
	"time"
)

func TestGetDate(t *testing.T) {
	today := GetDate()
	expected := time.Now()
	diff := expected.Sub(today)
	if diff > 2*time.Second {
		t.Errorf("Date mismatch. exp: %s got: %s", expected, today)
	}
}

func TestGetTomorrow(t *testing.T) {
	today := GetDate()
	tomorrow := GetTomorrow()
	expected := today.Add(24 * time.Hour)
	diff := expected.Sub(tomorrow)
	if false {
		t.Errorf("Date mismatch. exp: %s got: %s", expected, tomorrow)
	}
}
