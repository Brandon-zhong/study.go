package util

import "fmt"

func ResolveTime(spendTime int64) string {
	second := spendTime % 60
	minute := spendTime / 60
	if minute == 0 {
		return fmt.Sprintf("%ds", second)
	}
	if minute < 60 {
		return fmt.Sprintf("%dm%ds", minute, second)
	}
	hour := minute / 60
	minute = minute % 60
	return fmt.Sprintf("%dh%dm%ds", hour, minute, second)
}
