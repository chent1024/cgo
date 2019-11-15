package util

import (
	"database/sql/driver"
	"fmt"
	"time"
)

const GoTimeStart = "2006-01-02 15:04:05"

type LocalDate struct {
	time.Time
}

// MarshalJSON on JSONTime format Time field with %Y-%m-%d %H:%M:%S
func (l LocalDate) MarshalJSON() ([]byte, error) {
	formatted := fmt.Sprintf("\"%s\"", l.Format(GoTimeStart))
	return []byte(formatted), nil
}

// Value insert timestamp into mysql need this function.
func (l LocalDate) Value() (driver.Value, error) {
	var zeroTime time.Time
	if l.Time.UnixNano() == zeroTime.UnixNano() {
		return nil, nil
	}
	return l.Time, nil
}

// Scan valueof time.Time
func (l *LocalDate) Scan(v interface{}) error {
	value, ok := v.(time.Time)
	if ok {
		*l = LocalDate{Time: value}
		return nil
	}
	return fmt.Errorf("can not convert %v to timestamp", v)
}

func TodayUnix() int64 {
	t := time.Now()
	tm1 := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
	return tm1.Unix()
}

func DateNow() string {
	return time.Now().Format(GoTimeStart)
}

func DateAdd(day int) string {
	t := time.Now()
	tm1 := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
	return tm1.AddDate(0, 0, day).Format(GoTimeStart)
}

func DateTime() int64 {
	return time.Now().Unix()
}

func Str2Time(s string) (int64, error) {
	t, err := time.Parse(GoTimeStart, s)
	if err != nil {
		return 0, err
	}
	return t.Unix(), nil
}
