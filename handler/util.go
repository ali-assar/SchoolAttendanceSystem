package handler

func unixToMinute(timestamp int64) int64 {
	return (timestamp % 3600) / 60
}
