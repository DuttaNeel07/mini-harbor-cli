package auth

import "github.com/spf13/viper"

func GetToken() string {
	return viper.GetString("token")
}
