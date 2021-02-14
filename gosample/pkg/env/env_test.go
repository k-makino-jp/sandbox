package env

import (
	"errors"
	"os"
	"reflect"
	"testing"
)

func Test_env_hasValidLevel(t *testing.T) {
	tests := []struct {
		name string
		e    *env
		want *env
	}{
		{
			name: "env.hasValidLevel レベルINFOが与えられたとき レベルにINFOが設定されること",
			e:    &env{Level: "INFO"},
			want: &env{Level: "INFO"},
		},
		{
			name: "env.hasValidLevel レベルHOGEが与えられたとき レベルにINFOが設定されること",
			e:    &env{Level: "HOGE"},
			want: &env{Level: "INFO"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.e.hasValidLevel()
			if !reflect.DeepEqual(tt.e, tt.want) {
				t.Errorf("hasValidLevel() got = %v, want %v", tt.e, tt.want)
			}
		})
	}
}

func TestInit(t *testing.T) {
	tests := []struct {
		name         string
		want         error
		testSetup    func()
		testTeardown func()
	}{
		{
			name:         "Init 環境変数を取得するとき ErrorとしてNilが返ること",
			want:         nil,
			testSetup:    func() { os.Setenv("ENDPOINT", "endpoint") },
			testTeardown: func() { os.Unsetenv("ENDPOINT") },
		},
		{
			name:         "Init 環境変数が未設定の時 Errorが返ること",
			want:         errors.New("required key ENDPOINT missing value"),
			testSetup:    func() {},
			testTeardown: func() {},
		},
	}
	isSameError := func(err, want error) bool {
		var errString, wantString string
		if err != nil {
			errString = err.Error()
		}
		if want != nil {
			wantString = want.Error()
		}
		if errString == wantString {
			return true
		}
		return false
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.testSetup()
			if err := Init(); !isSameError(err, tt.want) {
				t.Errorf("Init() error = %v, want %v", err, tt.want)
			}
			tt.testTeardown()
		})
	}
}
