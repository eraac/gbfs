package gbfsspec

import (
	"strconv"
	"strings"
	"time"
)

// Enum from the spec
type (
	AlertType    string
	UserType     string
	Day          string
	RentalMethod string
)

type (
	// Service day in the YYYY-MM-DD format. Example: 2019-09-13 for September 13th, 2019.
	Date string

	// Service hour in the HH:MM:SS format. Example 22:26:02 or 02:16:00
	Time string

	Timestamp int64
)

// Custom type for handle many format used by different provider (or the spec)
type (
	Price string
	Boolean bool
)

const (
	DateFormat = "2006-01-02"
	TimeFormat = "15:04:05"
)

// ToTime convert string to a time.Time
func (d Date) ToTime(tz string) (time.Time, error) {
	loc := time.Local

	if tz != "" {
		var err error
		loc, err = time.LoadLocation(tz)

		if err != nil {
			return time.Time{}, err
		}
	}

	return time.ParseInLocation(DateFormat, string(d), loc)
}

// Clock parse the hour, minute, and second and split it
func (t Time) Clock() (h int, m int, s int) {
	ss := strings.Split(string(t), ":")
	if len(ss) != 3 {
		return
	}

	h, _ = strconv.Atoi(ss[0])
	m, _ = strconv.Atoi(ss[1])
	s, _ = strconv.Atoi(ss[2])

	return
}

// ToTime convert timestamp to a time.Time
func (t Timestamp) ToTime() time.Time {
	return time.Unix(int64(t), 0)
}

// UnmarshalJSON try to parse the price in JSON,
// who can be either a int, a float or a string
func (p *Price) UnmarshalJSON(bs []byte) error {
	*p = Price(strings.Trim(string(bs), "\""))

	return nil
}

// UnmarshalJSON try to parse boolean in JSON,
// who can be either a true boolean, a int (1/0) or a string ("false"/"true")
func (b *Boolean) UnmarshalJSON(bs []byte) error {
	v := strings.ToLower(strings.Trim(string(bs), "\""))

	switch v {
	case "1", "true":
		*b = true
	default:
		*b = false
	}

	return nil
}
