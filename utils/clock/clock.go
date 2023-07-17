package clock

import "time"

type Clocker interface {
	Now() time.Time
}

type RealClocker struct{}

// 現在の時刻を返却する
//
// @return 現在時刻
func (r RealClocker) Now() time.Time {
	return time.Now()
}
