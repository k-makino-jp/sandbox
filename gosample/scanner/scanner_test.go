package scanner

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"reflect"
	"regexp"
	"runtime"
	"testing"
)

func TestScannerImpl_lineValidation(t *testing.T) {
	errExpectedErrorOccurred := errors.New("Expected Error Occurred")
	type args struct {
		line string
	}
	tests := []struct {
		name         string
		s            *ScannerImpl
		args         args
		want         error
		testSetup    func()
		testTeardown func()
	}{
		{
			name: "ScannerImpl.lineValidation InputCorrectFormat ReturnsEqualsNil",
			s:    &ScannerImpl{},
			args: args{line: "HOGE1234-AABBC01"},
			want: nil,
		},
		{
			name: "ScannerImpl.lineValidation CompileFailed ReturnsEqualsError",
			s:    &ScannerImpl{},
			args: args{line: "HOGE1234-AABBC01"},
			want: errExpectedErrorOccurred,
			testSetup: func() {
				regexpCompile = func(string) (*regexp.Regexp, error) { return nil, errExpectedErrorOccurred }
			},
			testTeardown: func() { regexpCompile = regexp.Compile },
		},
		{
			name: "ScannerImpl.lineValidation InputCorrectFormatButIncludeSpace ReturnsEqualsError",
			s:    &ScannerImpl{},
			args: args{line: " HOGE1234-AABBC01"},
			want: errInvalidLine,
		},
		{
			name: "ScannerImpl.lineValidation InputCorrectFormatButIncludeSpaceAtEnd ReturnsEqualsError",
			s:    &ScannerImpl{},
			args: args{line: "HOGE1234-AABBC0 "},
			want: errInvalidLine,
		},
		{
			name: "ScannerImpl.lineValidation InputEmptyString ReturnsEqualsError",
			s:    &ScannerImpl{},
			args: args{line: ""},
			want: errInvalidLine,
		},
		{
			name: "ScannerImpl.lineValidation InputJapanese ReturnsEqualsError",
			s:    &ScannerImpl{},
			args: args{line: "あいうえお"},
			want: errInvalidLine,
		},
		{
			name: "ScannerImpl.lineValidation InputNoExistUnicode文字 ReturnsEqualsError",
			s:    &ScannerImpl{},
			args: args{line: string([]byte{0xff})},
			want: errInvalidLine,
		},
		{
			name: "ScannerImpl.lineValidation PrefixIsLowercase ReturnsEqualsError",
			s:    &ScannerImpl{},
			args: args{line: "hOGE1234-AABBC01"},
			want: errInvalidLine,
		},
		{
			name: "ScannerImpl.lineValidation PrefixIsPeriod ReturnsEqualsError",
			s:    &ScannerImpl{},
			args: args{line: ".OGE1234-AABBC01"},
			want: errInvalidLine,
		},
		{
			name: "ScannerImpl.lineValidation PrefixIsSpace ReturnsEqualsError",
			s:    &ScannerImpl{},
			args: args{line: " OGE1234-AABBC01"},
			want: errInvalidLine,
		},
		{
			name: "ScannerImpl.lineValidation AAisNotUpperCase ReturnsEqualsError",
			s:    &ScannerImpl{},
			args: args{line: "HOGE1234-aaBBC01"},
			want: errInvalidLine,
		},
		{
			name: "ScannerImpl.lineValidation BBisNotUpperCase ReturnsEqualsError",
			s:    &ScannerImpl{},
			args: args{line: "HOGE1234-AAbbC01"},
			want: errInvalidLine,
		},
		{
			name: "ScannerImpl.lineValidation CisNotC ReturnsEqualsError",
			s:    &ScannerImpl{},
			args: args{line: "HOGE1234-AABBX01"},
			want: errInvalidLine,
		},
		{
			name: "ScannerImpl.lineValidation 01isNotNumber ReturnsEqualsError",
			s:    &ScannerImpl{},
			args: args{line: "HOGE1234-AABBCXX"},
			want: errInvalidLine,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.testSetup != nil {
				tt.testSetup()
			}
			if got := tt.s.lineValidation(tt.args.line); got != tt.want {
				t.Errorf("ScannerImpl.lineValidation() error = %v, want %v", got, tt.want)
			}
			if tt.testTeardown != nil {
				tt.testTeardown()
			}
		})
	}
}

func TestScannerImpl_detectDuplicate(t *testing.T) {
	type args struct {
		lines []string
	}
	tests := []struct {
		name string
		s    *ScannerImpl
		args args
		want error
	}{
		{
			name: "ScannerImpl.removeDuplicate UnduplicatedLines ReturnsErrorEqualsNil",
			s:    &ScannerImpl{},
			args: args{lines: []string{"HOGEHOGE-AABBC01", "HOGEHOGE-AABBC02"}},
			want: nil,
		},
		{
			name: "ScannerImpl.removeDuplicate LinesIncludeDuplicate ReturnsErrorEqualsDuplicateDetected",
			s:    &ScannerImpl{},
			args: args{
				lines: []string{
					"HOGEHOGE-AABBC01",
					"HOGEHOGE-AABBC02",
					"HOGEHOGE-AABBC01",
					"HOGEHOGE-AABBC02",
					"HOGEHOGE-AABBC03",
				},
			},
			want: errDuplicatedLines,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.detectDuplicate(tt.args.lines); got != tt.want {
				t.Errorf("ScannerImpl.detectDuplicate() error = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestScannerImpl_Get(t *testing.T) {
	tests := []struct {
		name string
		s    *ScannerImpl
		want []string
	}{
		{
			name: "ScannerImpl.Get GetTwoLines ReturnsEqualsTwoLines",
			s: &ScannerImpl{
				validLines: []string{"HOGEHOGE-AABBC01", "FUGAFUGA-DDEEF02"},
			},
			want: []string{"HOGEHOGE-AABBC01", "FUGAFUGA-DDEEF02"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.Get(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ScannerImpl.Get() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestScannerImpl_Scan(t *testing.T) {
	testFilePath := "scannerImpl_scan.txt"
	fileCreator := func(filepath string, writeData []byte, perm os.FileMode) {
		err := ioutil.WriteFile(filepath, writeData, perm)
		if err != nil {
			fmt.Println("[TEST WARNING] failed to create data file.", filepath)
		}
	}
	fileDeletor := func(filepath string) {
		if err := os.Remove(filepath); err != nil {
			fmt.Println("[TEST WARNING] failed to delete data file.", filepath)
		}
	}
	type args struct {
		scanFilePath string
	}
	tests := []struct {
		name         string
		s            *ScannerImpl
		args         args
		want         error
		testSetup    func()
		testTeardown func()
	}{
		{
			name:         "ScannerImpl.Scan InputValidFile ReturnsEqualsNil",
			s:            &ScannerImpl{},
			args:         args{scanFilePath: testFilePath},
			want:         nil,
			testSetup:    func() { fileCreator(testFilePath, []byte("HOGEHOGE-AABBC01\nHOGEHOGE-AABBC02"), 0444) },
			testTeardown: func() { fileDeletor(testFilePath) },
		},
		{
			name:         "ScannerImpl.Scan NotExistFile ReturnsEqualsError",
			s:            &ScannerImpl{},
			args:         args{scanFilePath: testFilePath},
			want:         errFileNotExist,
			testSetup:    func() {},
			testTeardown: func() {},
		},
		{
			name: "ScannerImpl.Scan CannotOpenFile ReturnsEqualsError",
			s:    &ScannerImpl{},
			args: args{scanFilePath: testFilePath},
			want: errCannotOpenFile,
			testSetup: func() {
				if runtime.GOOS == "windows" {
					osOpen = func(string) (*os.File, error) { return nil, errCannotOpenFile }
				}
				fileCreator(testFilePath, []byte("HOGEHOGE-AABBC01\nHOGEHOGE-AABBC02"), 0200)
			},
			testTeardown: func() {
				fileDeletor(testFilePath)
				osOpen = os.Open
			},
		},
		{
			name:         "ScannerImpl.Scan InputInvalidLineFile ReturnsEqualsError",
			s:            &ScannerImpl{},
			args:         args{scanFilePath: testFilePath},
			want:         errInvalidLine,
			testSetup:    func() { fileCreator(testFilePath, []byte("hogeHOGE-AABBC01"), 0444) },
			testTeardown: func() { fileDeletor(testFilePath) },
		},
		{
			name: "ScannerImpl.Scan InputLargeFile ReturnsEqualNil",
			s:    &ScannerImpl{},
			args: args{scanFilePath: testFilePath},
			want: errInvalidLine,
			testSetup: func() {
				buf64byte := []byte("xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx")
				bufHasMaxScanTokenSize := bytes.Repeat(buf64byte, bufio.MaxScanTokenSize/len(buf64byte)-1) // MaxScanTokenSize = 64 * 1024 [byte]
				fileCreator(testFilePath, bufHasMaxScanTokenSize, 0444)
			},
			testTeardown: func() { fileDeletor(testFilePath) },
		},
		{
			name: "ScannerImpl.Scan InputLargeFile ReturnsEqualsError",
			s:    &ScannerImpl{},
			args: args{scanFilePath: testFilePath},
			want: errScanFailed,
			testSetup: func() {
				buf64byte := []byte("xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx")
				bufHasMaxScanTokenSize := bytes.Repeat(buf64byte, bufio.MaxScanTokenSize/len(buf64byte)) // MaxScanTokenSize = 64 * 1024 [byte]
				fileCreator(testFilePath, bufHasMaxScanTokenSize, 0444)
			},
			testTeardown: func() { fileDeletor(testFilePath) },
		},
		{
			name:         "ScannerImpl.Scan InputDuplicatedFile ReturnsEqualsError",
			s:            &ScannerImpl{},
			args:         args{scanFilePath: testFilePath},
			want:         errDuplicatedLines,
			testSetup:    func() { fileCreator(testFilePath, []byte("HOGEHOGE-AABBC01\nHOGEHOGE-AABBC01"), 0444) },
			testTeardown: func() { fileDeletor(testFilePath) },
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.testSetup()
			if got := tt.s.Scan(tt.args.scanFilePath); got != tt.want {
				t.Errorf("ScannerImpl.Scan() error = %v, want %v", got, tt.want)
			}
			tt.testTeardown()
		})
	}
}

func TestNewScannerImpl(t *testing.T) {
	tests := []struct {
		name string
		want *ScannerImpl
	}{
		{
			name: "NewScannerImpl CreateInstance ReturnEqualsInstance",
			want: &ScannerImpl{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewScannerImpl(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewScannerImpl() = %v, want %v", got, tt.want)
			}
		})
	}
}
