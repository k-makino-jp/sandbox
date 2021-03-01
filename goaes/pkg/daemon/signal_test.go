package daemon

import (
	"fmt"
	"os"
	"testing"

	"gosample/pkg/daemon/mock_daemon"

	"github.com/golang/mock/gomock"
)

func Test_daemonImpl_WaitForInputSignal(t *testing.T) {
	// configure mock
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockSignal := mock_daemon.NewMockSignal(ctrl)
	tests := []struct {
		name      string
		d         daemonImpl
		testSetup func()
	}{
		{
			name: "daemonImpl SIGINTシグナルを受信するとき ",
			d:    daemonImpl{signal: mockSignal},
			testSetup: func() {
				// sig := make(chan os.Signal, 1)
				mockSignal.EXPECT().Notify(gomock.Any(), os.Interrupt, os.Kill).DoAndReturn(
					func(c chan<- os.Signal, sig ...os.Signal) {
						fmt.Println("test func")
						c <- os.Interrupt
					},
				).Times(1)
			},
		},
	}
	for _, tt := range tests {
		tt.testSetup()
		t.Run(tt.name, func(t *testing.T) {
			tt.d.WaitForInputSignal()
		})
	}
}
