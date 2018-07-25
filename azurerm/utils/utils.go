package utils

func Bool(input bool) *bool {
	return &input
}

func Int32(input int32) *int32 {
	return &input
}

func Int64(input int64) *int64 {
	return &input
}

func String(input string) *string {
	return &input
}

func SafeStringPtr(input string) *string {
	if input == "" {
		return nil
	} else {
		return &input
	}
}
