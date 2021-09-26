package tools

import "time"

func GetCurrentTimeStr() string {
	return time.Now().String()
}
