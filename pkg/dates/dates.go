package dates

import "time"

func GetDate() time.Time {
	return time.Now()
}

func GetTomorrow() time.Time {
	return time.Now().Add(1 * time.Hour * 24)
}
