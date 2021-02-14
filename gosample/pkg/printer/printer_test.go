// printer is sample package

package printer

import (
	"gosample/hoge/enver"
	"gosample/hoge/enver/mock_enver"
	"testing"

	"github.com/golang/mock/gomock"
)

func Test_printImpl_Print(t *testing.T) {
	// create controller
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	// create mock
	mockEnver := mock_enver.NewMockEnver(ctrl)

	tests := []struct {
		name      string
		p         *printImpl
		testSetup func()
	}{
		{
			name:      "sample test without mock",
			p:         NewPrintImpl(),
			testSetup: func() {},
		},
		{
			name: "sample test with mock",
			p: &printImpl{
				mockEnver,
			},
			testSetup: func() {
				mockEnver.EXPECT().Initialize().Times(1)
				e := enver.ModelEnv{}
				e.Endpoint = "hoge"
				mockEnver.EXPECT().GetEnv().Return(e).Times(1)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.testSetup()
			tt.p.Print()
		})
	}
}
