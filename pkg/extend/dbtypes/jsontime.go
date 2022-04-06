package dbtypes

import (
	"database/sql/driver"
	"fmt"
	"time"
)

const (
	YYYYMMDDHHMISS = "2006-01-02 15:04:05" //常规类型
)

type JSONTime struct {
	time.Time
}

// Value insert timestamp into mysql need this function.
func (t JSONTime) Value() (driver.Value, error) {
	var zeroTime time.Time
	if t.Time.UnixNano() == zeroTime.UnixNano() {
		return nil, nil
	}
	return t.Time, nil
}

// Scan valueof time.Time
func (t *JSONTime) Scan(v interface{}) error {
	value, ok := v.(time.Time)
	if ok {
		*t = JSONTime{Time: value}
		return nil
	}
	return fmt.Errorf("can not convert %v to timestamp", v)
}


// MarshalJSON on JSONTime format Time field with %Y-%m-%d %H:%M:%S
func (t JSONTime) MarshalJSON() ([]byte, error) {
	formatted := "null"
	if t.IsZero(){
		formatted = "null"
	}else{
		formatted = fmt.Sprintf("\"%s\"", t.Format(YYYYMMDDHHMISS))
	}
	return []byte(formatted), nil
}

// UnmarshalJSON implements the json.Unmarshaler interface.
// The time is expected to be a quoted string in RFC 3339 format.
func (t *JSONTime) UnmarshalJSON(data []byte) error {
	// Ignore null, like in the main JSON package.
	if string(data) == "null" {
		return nil
	}
	// Fractional seconds are handled implicitly by Parse.
	var err error
	(*t).Time, err = time.ParseInLocation(`"`+YYYYMMDDHHMISS+`"`, string(data), time.Local)
	return err
}