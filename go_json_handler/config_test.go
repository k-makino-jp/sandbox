package main

import (
	"gojson/json"
	"gojson/os"
	"io/fs"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestNewConfigCmd_CreateInstance_ReturnsInstancePointer(t *testing.T) {
	expect := &configCmd{
		UserConfig: UserConfig{
			HTTPSProxy: "http://proxy.example.com",
		},
		os:   os.NewOs(),
		json: json.NewJson(),
	}
	actual := NewConfigCmd()
	assert.Equal(t, expect, actual)
}

func TestConfigCmdCreateConfigDat_UserConfigHasValues_CreatesConfigDat(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockJsonInterface := json.NewMockJsonInterface(ctrl)
	userConfig := UserConfig{HTTPSProxy: "http://proxy.example.com"}
	const prefix = ""
	const indent = "  "
	jsonBytes := []byte(`{
  "https.proxy": "http://proxy.example.com"
}`)
	mockJsonInterface.EXPECT().MarshalIndent(userConfig, prefix, indent).Return(jsonBytes, nil)

	mockOsInterface := os.NewMockOsInterface(ctrl)
	const perm fs.FileMode = 0644
	mockOsInterface.EXPECT().WriteFile(userConfigFilePath, jsonBytes, perm).Return(nil)

	configCmd := configCmd{
		UserConfig: UserConfig{
			HTTPSProxy: "http://proxy.example.com",
		},
		json: mockJsonInterface,
		os:   mockOsInterface,
	}
	actual := configCmd.createConfigDat()
	assert.NoError(t, actual)
}
