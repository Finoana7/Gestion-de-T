package helper

import "time"

func GetDate() string {
	currentTime := time.Now()
	formattedDate := currentTime.Format("02/01/06")
	return formattedDate
}
