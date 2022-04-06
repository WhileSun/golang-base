package dbtypes

import (
	"time"
)

const (
	format="2006-01-02 15:04:05"
)

type MyTime time.Time


func (t MyTime) MarshalJSON() ([]byte, error) {
	b := make([]byte, 0, len(format)+2)
	b = append(b, '"')
	b = (time.Time(t)).AppendFormat(b, format)
	b = append(b, '"')
	return b, nil
}

func (t *MyTime) UnmarshalJSON(data []byte) (err error) {
	now, err := time.ParseInLocation(`"`+format+`"`, string(data), time.Local)
	*t = MyTime(now)
	return
}
