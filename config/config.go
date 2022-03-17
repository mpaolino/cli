package config

import (
	"flag"
	"fmt"
	"os"

	"github.com/spf13/viper"
)

func ContextList() ([]string, error) {
	res := []string{}

	m := make(map[string]interface{})

	err := viper.UnmarshalKey("contexts", &m)

	if err != nil {
		return []string{}, err
	}

	for k := range m {
		res = append(res, k)
	}

	return res, nil
}

func SetActiveContext(name string) {
	Set("active-context", name)
}

func GetActiveContext() string {
	if flag.Lookup("test.v") == nil {
		return viper.GetString("active-context")
	} else {
		return "org-semaphoretext-xyz"
	}
}

func GetAuth() string {
	if flag.Lookup("test.v") == nil {
		context := GetActiveContext()
		key_path := fmt.Sprintf("contexts.%s.auth.token", context)

		return Get(key_path)
	} else {
		return "123456789"
	}
}

func GetEditor() string {
	if flag.Lookup("test.v") == nil {
		editor := viper.GetString("editor")

		if editor != "" {
			return editor
		}

		editor = os.Getenv("EDITOR")

		if editor != "" {
			return editor
		}

		return "vim"
	} else {
		return "true" // Bash 'true' command, do nothing in tests
	}
}

func SetAuth(token string) error {
	context := GetActiveContext()
	key_path := fmt.Sprintf("contexts.%s.auth.token", context)

	return Set(key_path, token)
}

func GetHost() string {
	if flag.Lookup("test.v") == nil {
		context := GetActiveContext()
		key_path := fmt.Sprintf("contexts.%s.host", context)

		return Get(key_path)
	} else {
		return "org.semaphoretext.xyz"
	}
}

func SetHost(token string) error {
	context := GetActiveContext()
	key_path := fmt.Sprintf("contexts.%s.host", context)

	return Set(key_path, token)
}

func Set(key string, value string) error {
	viper.Set(key, value)
	err := viper.WriteConfig()
	if err != nil {
		return fmt.Errorf("failed to write value to config file: %s", err)
	}
	return nil
}

func Get(key string) string {
	return viper.GetString(key)
}

func IsSet(key string) bool {
	return viper.IsSet(key)
}
