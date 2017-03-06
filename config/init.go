package config

import (
	"fmt"
	"os"
	"strings"

	"path/filepath"

	"github.com/BurntSushi/toml"
)

// CheckFlags verifies that supplied flags are no
// file system paths and point to existing files.
func CheckFlags(envName string, configName string) (string, string, error) {

	// Check if envName contains a separator
	// and fail if that is the case.
	if strings.ContainsRune(envName, filepath.Separator) {
		return "", "", fmt.Errorf("do not specify path but only name of environment file")
	}

	// Check if configName contains a separator
	// and fail if that is the case.
	if strings.ContainsRune(configName, filepath.Separator) {
		return "", "", fmt.Errorf("do not specify path but only name of configuration file")
	}

	// Check if environment file exists.
	_, err := os.Stat(envName)
	if os.IsNotExist(err) {
		return "", "", fmt.Errorf("specified environment file does not exist")
	}

	// Check if configuration file exists.
	_, err = os.Stat(configName)
	if os.IsNotExist(err) {
		return "", "", fmt.Errorf("specified configuration file does not exist")
	}

	return envName, configName, nil
}

// LoadConfig expects a name for protokollamt's
// configuration file. It attempts to parse its
// contents into above configuration structs.
func LoadConfig(configName string, dbPassword string, mailPassword string) (*Config, error) {

	c := new(Config)

	// Parse values from TOML file into struct.
	_, err := toml.DecodeFile(configName, c)
	if err != nil {
		return nil, fmt.Errorf("error decoding TOML config file: %v", err)
	}

	// Enrich config with sensitive password values
	// retrieved prior to this function via .env files.
	c.Database.Password = dbPassword
	c.Mail.Password = mailPassword

	return c, nil
}