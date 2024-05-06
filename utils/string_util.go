package utils

func ParseString(value string) *string {
	if value == "" {
		return nil
	}
	return &value
}
