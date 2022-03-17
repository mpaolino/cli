package config

import (
	"os"
	"testing"

	"github.com/spf13/viper"

	assert "github.com/stretchr/testify/assert"
)

func Test__WriteConfigValue_Write_and_Verify(t *testing.T) {
	tempFile, err := os.CreateTemp("", "sem-test.*.yaml")
	defer os.Remove(tempFile.Name())
	if err != nil {
		t.Error("Failed to create temporary config file")
	}
	viper.SetConfigFile(tempFile.Name())
	assert.NoError(t, Set("test_key", "test_value"))
	assert.Equal(t, Get("test_key"), "test_value")
}

func Test__WriteConfigValue_Fail_To_Write(t *testing.T) {
	// A simple way to trigger a write failure is configuring viper to find the default
	// config filename "config" but never specify any valid paths to search for it.
	// An actual real scenario where this might get triggered is if the config file changes owners or their FS
	// gets mounted in read-only mode.
	viper.SetConfigFile("")
	assert.Error(t, Set("test_key", "test_value"))
}
