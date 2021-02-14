// package config 設定ファイル処理用パッケージ

package config

import (
	"reflect"
	"testing"
)

func Test_configer_Read(t *testing.T) {
	type args struct {
		configFilePath string
	}
	// functions
	// fileCreator := func(filepath string, writedata []byte, perm os.FileMode) {
	// 	err := ioutil.WriteFile(filepath, writedata, perm)
	// 	if err != nil {
	// 		fmt.Println("[TEST WARNING] failed to create data file.", filepath)
	// 	}
	// }
	// fileDeletor := func(filepath string) {
	// 	if err := os.Remove(filepath); err != nil {
	// 		fmt.Println("[TEST WARNING] failed to delete data file.", filepath)
	// 	}
	// }
	tests := []struct {
		name         string
		c            *configer
		args         args
		wantConfig   []byte
		wantErr      error
		testSetup    func()
		testTeardown func()
	}{
		// {
		// 	name:         "configer.Read 正常な設定ファイルが与えられたとき Nilが返ってくること",
		// 	c:            &configer{
		// 		config:
		// 	},
		// 	args:         args{configFilePath: "config.json"},
		// 	wantConfig:   []byte("configdata"),
		// 	wantErr:      nil,
		// 	testSetup:    func() { fileCreator("config.json", []byte("configdata"), 0666) },
		// 	testTeardown: func() { fileDeletor("config.json") },
		// },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.c.Read(tt.args.configFilePath); err != tt.wantErr {
				t.Errorf("configer.Read() error = %v, wantErr %v", err, tt.wantErr)
			}
			// if tt.c.config != tt.wantConfig {
			// 	t.Errorf("configer.config got = %v, wantConfig %v", tt.c.config, tt.wantConfig)
			// }
		})
	}
}

func Test_configer_Get(t *testing.T) {
	tests := []struct {
		name string
		c    *configer
		want config
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.Get(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("configer.Get() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewConfiger(t *testing.T) {
	tests := []struct {
		name string
		want *configer
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewConfiger(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewConfiger() = %v, want %v", got, tt.want)
			}
		})
	}
}
