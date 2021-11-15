package main

import (
	"encoding/json"
	"fmt"
	"os"
	"reflect"
)

type config struct {
	NoProxy    string `json:"noproxy"`
	HTTPProxy  string `json:"http.proxy"`
	HTTPSProxy string `json:"https.proxy"`
}

type rootCmd struct {
	config
}

func NewRootCmd() *rootCmd {
	return &rootCmd{
		config{
			NoProxy:    "proxy.example.com",
			HTTPProxy:  "http://proxy.example.com",
			HTTPSProxy: "http://proxy2.example.com",
		},
	}
}

var jsonMarshalIndent = json.MarshalIndent
var osWriteFile = os.WriteFile

func (r rootCmd) createConfigDat() error {
	prefix := ""
	indent := "  "
	jsonBytes, err := jsonMarshalIndent(r.config, prefix, indent)
	if err != nil {
		return err
	}
	return osWriteFile("data/config.json", jsonBytes, 0644)
}

func (r rootCmd) listConfigDat() error {
	configType := reflect.TypeOf(r.config)
	configValue := reflect.ValueOf(r.config)
	for i := 0; i < configType.NumField(); i++ {
		key := configType.Field(i).Tag.Get("json")
		value := configValue.Field(i).Interface()
		output := fmt.Sprintf("%s=%s", key, value)
		fmt.Println(output)
	}
	return nil
}
