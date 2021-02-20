package impl

import "time"

type TimeFormat time.Time

//MarshalJSON jsonTime序列化调用的方法
func (jsonTime TimeFormat) MarshalJSON() ([]byte, error) {
	//当返回时间为空时，需特殊处理
	if time.Time(jsonTime).IsZero() {
		return []byte(`""`), nil
	}
	return []byte(`"` + time.Time(jsonTime).Format("2006-01-02 15:04:05") + `"`), nil
}
