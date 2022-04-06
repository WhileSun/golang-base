package dbtypes

import (
	"database/sql/driver"
	"fmt"
	"time"
)

const TimeFormat = "2006-01-02 15:04:05"
type GTime time.Time

func (t GTime) Value() (driver.Value, error) {
	// 0001-01-01 00:00:00 属于空值，遇到空值解析成 null 即可
	if t.String() == "0001-01-01 00:00:00" {
		return nil, nil
	}
	return []byte(time.Time(t).Format(TimeFormat)), nil
}

func (t *GTime) Scan(v interface{}) error {
	// mysql 内部日期的格式可能是 2006-01-02 15:04:05 +0800 CST 格式，所以检出的时候还需要进行一次格式化
	tTime, _ := time.Parse("2006-01-02 15:04:05 +0800 CST", v.(time.Time).String())
	*t = GTime(tTime)
	return nil
}

// 用于 fmt.Println 和后续验证场景
func (t GTime) String() string {
	return time.Time(t).Format(TimeFormat)
}


func (t GTime) MarshalJSON() ([]byte, error) {
	fmt.Println("MarshalJSON")
	b := make([]byte, 0, len(TimeFormat)+2)
	b = append(b, '"')
	b = time.Time(t).AppendFormat(b, TimeFormat)
	b = append(b, '"')
	return b, nil
}


func (t *GTime) UnmarshalJSON(data []byte) (err error) {
	fmt.Println("UnmarshalJSON")
	// 空值不进行解析
	if len(data) == 2 {
		*t = GTime(time.Time{})
		return
	}

	// 指定解析的格式
	now, err := time.ParseInLocation(`"`+TimeFormat+`"`, string(data), time.Local)
	*t = GTime(now)
	return
}


