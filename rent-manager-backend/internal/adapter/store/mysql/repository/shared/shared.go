package shared

func NullInt64(value int64) any {
	if value == 0 {
		return nil
	}

	return value
}
