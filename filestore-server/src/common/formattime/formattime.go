package formattime

import (
	"database/sql/driver"
	"strings"
	"time"
)

var timeFormat string

// SetTimeFormat set time format
func SetTimeFormat(format string) {
	timeFormat = format
}

// Time a common Time struct
type Time struct {
	time.Time
}

// MarshalJSON marshal the Time
func (t Time) MarshalJSON() ([]byte, error) {
	if timeFormat != "" {
		return []byte(`"` + t.Time.Format(timeFormat) + `"`), nil
	}
	return []byte(`"` + t.Time.String() + `"`), nil
}

// UnmarshalJSON Unmarshal the JSON
func (t Time) UnmarshalJSON(data []byte) error {
	layout := timeFormat
	if layout == "" {
		layout = time.RFC3339
	}
	tt, err := time.Parse(layout, strings.Trim(string(data), `"`))
	if err != nil {
		return err
	}
	t.Time = tt

	return nil
}

// Scan data base scan value
func (t *Time) Scan(value interface{}) error {
	t.Time = value.(time.Time)
	return nil
}

// Value data base value
func (t *Time) Value() (driver.Value, error) {
	return t.Time, nil
}
