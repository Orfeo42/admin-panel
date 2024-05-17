package utils

import "time"

func IntPtr(i int) *int {
	return &i
}

func TimePtr(i time.Time) *time.Time {
	return &i
}
