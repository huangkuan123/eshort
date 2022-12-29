package mtime

import (
	"database/sql/driver"
	"eshort/pkg/easylogger"
	"fmt"
	"time"
)

type Mtime time.Time

var format_str string = "2006-01-02 15:04:05"

// 返回当前时间戳
func (t *Mtime) GetTime() int64 {
	unix := time.Now().Unix()
	return unix
}

// 返回当前格式化后的时间
func (t *Mtime) GetFormatTime() string {
	return time.Unix(t.GetTime(), 0).Format(format_str)
}

// 将格式化的时间转为时间戳
func (t Mtime) StrtoTime(str string) int64 {
	parse, err := time.Parse(format_str, str)
	if err != nil {
		easylogger.LogError(err, "时间转换有误")
	}
	return parse.Unix()
}

// 将时间戳转为格式化时间
func (t Mtime) Date(timestamp int64) string {
	return time.Unix(timestamp, 0).Format(format_str)
}

func (t *Mtime) MarshalJSON() ([]byte, error) {
	tTime := time.Time(*t)
	return []byte(fmt.Sprintf("\"%v\"", tTime.Format(format_str))), nil
}

// 存储时会自动调用，解决时间为空
func (t Mtime) Value() (driver.Value, error) {
	var zeroTime time.Time
	tlt := time.Time(t)
	//判断给定时间是否和默认零时间的时间戳相同
	if tlt.UnixNano() == zeroTime.UnixNano() {
		return nil, nil
	}
	return tlt, nil
}

// 查询出来时会根据这个进行转换
func (t *Mtime) Scan(v interface{}) error {
	if value, ok := v.(time.Time); ok {
		*t = Mtime(value)
		return nil
	}
	return fmt.Errorf("can not convert %v to timestamp", v)
}
