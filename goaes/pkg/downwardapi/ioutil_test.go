// Package downwardapi Pod情報取得処理用パッケージ

// ---

// mock作成コマンド

// mockgen -self_package=downwardapi -source=ioutil.go -destination=ioutil_mock.go -package=downwardapi

package downwardapi

import (
	"io/ioutil"
	"log"
	"os"
	"reflect"
	"testing"
)

func Test_ioUtilImpl_ReadFile(t *testing.T) {
	fileCreator := func(filepath string, writedata []byte, perm os.FileMode) {
		if err := ioutil.WriteFile(filepath, writedata, perm); err != nil {
			log.Fatal(err)
		}
	}
	fileDeletor := func(filepath string) {
		if err := os.RemoveAll(filepath); err != nil {
			log.Fatal(err)
		}
	}
	type args struct {
		filename string
	}
	tests := []struct {
		name         string
		i            ioUtilImpl
		args         args
		want         []byte
		wantErr      bool
		testSetup    func()
		testTeardown func()
	}{
		{
			name:         "ioUtilImpl.ReadFile ファイル読み込みに成功するとき ファイル内容が返ること",
			i:            ioUtilImpl{},
			args:         args{filename: "readfile.txt"},
			want:         []byte("readfile"),
			wantErr:      false,
			testSetup:    func() { fileCreator("readfile.txt", []byte("readfile"), 0666) },
			testTeardown: func() { fileDeletor("readfile.txt") },
		},
		{
			name:         "ioUtilImpl.ReadFile ファイル読み込みに失敗するとき Errorが返ること",
			i:            ioUtilImpl{},
			args:         args{filename: "readfile.txt"},
			want:         nil,
			wantErr:      true,
			testSetup:    func() {},
			testTeardown: func() {},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.testSetup()
			got, err := tt.i.ReadFile(tt.args.filename)
			if (err != nil) != tt.wantErr {
				t.Errorf("ioUtilImpl.ReadFile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ioUtilImpl.ReadFile() = %v, want %v", got, tt.want)
			}
			tt.testTeardown()
		})
	}
}
