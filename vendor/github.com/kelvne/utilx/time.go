package utilx

import "time"

const (
	FormatYmd = "2006-01-02"
	FormatHms = "15:04:05"
)

// =========================
// Time enhancer
// =========================

type completeTime struct {
	time.Time
}

// CompleteTime is a more complete time.Time
func CompleteTime(t time.Time) completeTime {
	return completeTime{
		Time: t,
	}
}

// T returns completeTime as time.Time
func (c completeTime) T() time.Time {
	return c.Time
}

// Hms returns hours, minutes, seconds since 00:00
func (c completeTime) Hms() (int, int, int) {
	return c.Hour(), c.Minute(), c.Second()
}

// DailySeconds returns how many seconds passed since 00:00
func (c completeTime) DailySeconds() int {
	return (c.Hour() * 3600) + (c.Minute() * 60) + c.Second()
}

// FormatYmd returns time.Time as Year-Month-Day TimeString
func (c completeTime) FormatYmd() TimeString {
	return TimeString(c.Format(FormatYmd))
}

// FormatYmd returns time.Time as Hours:Minutes:Seconds TimeString
func (c completeTime) FormatHms() TimeString {
	return TimeString(c.Format(FormatHms))
}

// =========================
// TimeString
// =========================

type TimeString string

// S returns TimeString as string
func (c TimeString) S() string {
	return string(c)
}

// Ymd returns completeTime from string using Year-Month-Day format
func (t TimeString) Ymd() (completeTime, error) {
	raw, err := time.Parse(FormatYmd, string(t))
	if err != nil {
		return completeTime{}, err
	}
	return CompleteTime(raw), nil
}

// Hms returns completeTime from string using Hours:Minutes:Seconds
func (t TimeString) Hms() (completeTime, error) {
	raw, err := time.Parse(FormatHms, string(t))
	if err != nil {
		return completeTime{}, err
	}
	return CompleteTime(raw), nil
}

// =========================
// Comparasions
// =========================

// CompareYmd compares two time.Time as 2006-01-02
func CompareYmd(a, b time.Time) bool {
	return CompleteTime(a).FormatYmd() == CompleteTime(b).FormatYmd()
}
