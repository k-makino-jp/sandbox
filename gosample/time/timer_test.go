// time is sample package

package time

import (
	"reflect"
	"testing"
	"time"
)

var (
	fixedTime = time.Date(2021, 12, 31, 01, 23, 45, 500, time.UTC)
)

func Test_timerImpl_Now(t *testing.T) {
	nowImpl = func() time.Time {
		return fixedTime
	}
	tests := []struct {
		name string
		t    *timerImpl
		want time.Time
	}{
		{
			name: "timerImpl.Now GetCurrentTime ReturnsTimeEquals" + fixedTime.Format(layoutRFC3389),
			t:    &timerImpl{},
			want: fixedTime,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.t.Now(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("timerImpl.Now() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_timerImpl_NowRFC3389(t *testing.T) {
	tests := []struct {
		name string
		t    *timerImpl
		want string
	}{
		{
			name: "timerImpl.NowRFC3389 GetCurrentTimeString ReturnsTimeStringEquals" + fixedTime.Format(layoutRFC3389),
			t:    &timerImpl{},
			want: fixedTime.Format(layoutRFC3389),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.t.NowRFC3389(); got != tt.want {
				t.Errorf("timerImpl.NowRFC3389() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewTimerImpl(t *testing.T) {
	tests := []struct {
		name string
		want *timerImpl
	}{
		{
			name: "NewTimerImpl GetTimerImplPointer ReturnsTimeEqualsTimerImplPointer",
			want: &timerImpl{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewTimerImpl(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewTimerImpl() = %v, want %v", got, tt.want)
			}
		})
	}
}
