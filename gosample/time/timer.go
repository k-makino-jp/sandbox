// time is sample package
package time

import (
	"time"
)

const (
	layoutRFC3389 = "2006-01-02T15:04:05Z.000"
)

var (
	nowImpl = time.Now
)

type timer interface {
	Now() time.Time
	NowRFC3389() string
}

type timerImpl struct {
}

func (t *timerImpl) Now() time.Time {
	now := nowImpl().UTC()
	return now
}

func (t *timerImpl) NowRFC3389() string {
	nowRFC3389 := t.Now().Format(layoutRFC3389)
	return nowRFC3389
}

func NewTimerImpl() *timerImpl {
	return &timerImpl{}
}
