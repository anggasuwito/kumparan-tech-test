package utils

import "time"

const (
	DefaultLoc = "Asia/Jakarta"
)

var (
	loc, _ = time.LoadLocation(DefaultLoc)
)

func TimeNow() time.Time {
	return time.Now().In(loc)
}

func ParseTime(t time.Time) string {
	if t.IsZero() {
		return ""
	}
	return t.Format(time.RFC3339)
}
