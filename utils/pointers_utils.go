package utils

import "time"

func IntPtr(i int) *int {
	return &i
}

func UintPtr(i uint) *uint {
	return &i
}

func TimePtr(i time.Time) *time.Time {
	return &i
}

func StringPtr(str string) *string {
	return &str
}
