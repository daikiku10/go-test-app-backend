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

type FixedClocker struct{}

// 固定の時刻を返却する
//
// @return 固定時刻
func (f FixedClocker) Now() time.Time {
	return time.Date(2022, 5, 10, 12, 34, 56, 0, time.UTC)
}
