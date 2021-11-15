package main

import (
	"fmt"
	"gojson/json"
	"gojson/os"
	"reflect"
)

const (
	userConfigFilePath = "data/user_config.json"
)

type UserConfig struct {
	HTTPSProxy string `json:"https.proxy"`
}

type configCmd struct {
	UserConfig
	os   os.OsInterface
	json json.JsonInterface
}

func NewConfigCmd() *configCmd {
	return &configCmd{
		UserConfig: UserConfig{
			HTTPSProxy: "http://proxy.example.com",
		},
		os:   os.NewOs(),
		json: json.NewJson(),
	}
}

func (c configCmd) createConfigDat() error {
	const prefix = ""
	const indent = "  "
	jsonBytes, err := c.json.MarshalIndent(c.UserConfig, prefix, indent)
	if err != nil {
		return err
	}
	return c.os.WriteFile(userConfigFilePath, jsonBytes, 0644)
}

func (c configCmd) listConfigDat() error {
	UserConfigType := reflect.TypeOf(c.UserConfig)
	UserConfigValue := reflect.ValueOf(c.UserConfig)
	for i := 0; i < UserConfigType.NumField(); i++ {
		key := UserConfigType.Field(i).Tag.Get("json")
		value := UserConfigValue.Field(i).Interface()
		output := fmt.Sprintf("%s=%s", key, value)
		fmt.Println(output)
	}
	return nil
}
