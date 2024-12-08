package utils

import (
	"log"
	"time"
)

func FormatDate(date time.Time) string {
	dateFormated := date.Format("02/01/2006")
	return dateFormated
}

func ParseDate(date string) (time.Time, error) {
	dateFormated, err := time.Parse("02/01/2006", date)
	if err != nil {
		log.Println("Error parsing date:", err)
		return time.Time{}, err
	}

	return dateFormated, nil
}
