package example

import (
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"local.packages/filereader"
	"local.packages/filereader/mock_filereader"
	"local.packages/printer"
	"local.packages/printer/mock_printer"
)

const (
	PrefixNormalCase = "[NormalCase]"
	PrefixErrorCase  = "[ErrorCase]"
)

func Test_execFileReader(t *testing.T) {
	// create controller
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	// create mock
	mock := mock_filereader.NewMockFileReader(ctrl)

	type args struct {
		f filereader.FileReader
	}
	tests := []struct {
		name         string
		args         args
		want         uint8
		testSetup    func()
		testTeardown func()
	}{
		{
			name: PrefixNormalCase + "execute",
			args: args{
				f: filereader.NewFileReaderInstance(),
			},
			want:         exitStatusSuccess,
			testSetup:    func() {},
			testTeardown: func() {},
		},
		{
			name: PrefixNormalCase + "execute with mock",
			args: args{
				f: mock,
			},
			want: exitStatusSuccess,
			testSetup: func() {
				// configure method
				mock.EXPECT().Read(filepath).Return(nil)
			},
			testTeardown: func() {},
		},
		{
			name: PrefixErrorCase + "execute with mock",
			args: args{
				f: mock,
			},
			want: exitStatusFailed,
			testSetup: func() {
				// configure method
				mock.EXPECT().Read(filepath).Return(errors.New("an error occured"))
			},
			testTeardown: func() {},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.testSetup()
			if got := execFileReader(tt.args.f); got != tt.want {
				t.Errorf("execFileReader() = %v, want %v", got, tt.want)
			}
			tt.testTeardown()
		})
	}
}

func Test_execPrinter(t *testing.T) {
	// create controller
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	// create mock
	mock := mock_printer.NewMockPrinter(ctrl)
	mock.EXPECT().Print().Return()

	type args struct {
		p printer.Printer
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: PrefixNormalCase + "execute",
			args: args{
				p: printer.NewPrinterImpl(),
			},
		},
		{
			name: PrefixNormalCase + "execute with mock",
			args: args{
				p: mock,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			execPrinter(tt.args.p)
		})
	}
}
