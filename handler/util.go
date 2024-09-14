package handler

func UnixToMinute(timestamp int64) int64 {
	return (timestamp % 3600) / 60
}

func ExtractUnixDate(timestamp int64) int64 {
	return (timestamp  / 86400) * 86400
}