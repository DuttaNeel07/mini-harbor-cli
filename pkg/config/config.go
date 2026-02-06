package config

import (
	"os"
	"path/filepath"

	"github.com/spf13/viper"
)

func InitConfig() error {
	home, err := os.UserHomeDir()
	if err != nil {
		return err
	}

	configDir := filepath.Join(home, ".mini")
	configFile := filepath.Join(configDir, "config.yaml")

	// create ~/.mini if it doesn't exist
	if _, err := os.Stat(configDir); os.IsNotExist(err) {
		err = os.Mkdir(configDir, 0755)
		if err != nil {
			return err
		}
	}

	viper.SetConfigFile(configFile)
	viper.SetConfigType("yaml")

	// create config file if it doesn't exist
	if _, err := os.Stat(configFile); os.IsNotExist(err) {
		file, err := os.Create(configFile)
		if err != nil {
			return err
		}
		file.Close()
	}

	return viper.ReadInConfig()
}
