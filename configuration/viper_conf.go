package configuration

import (
	"github.com/spf13/viper"
)

// ReadConfig default
func ReadConfig() {

	// set viper default
	// viper.SetDefault("USER_NAME", "root")
	// viper.SetDefault("PASSWORD", "P@ssw0rd")
	// viper.SetDefault("HOST_NAME", "localhost:3306")
	// viper.SetDefault("NAME", "ms_account_dev")

	// viper.AutomaticEnv()

	viper.SetConfigName("properties-staging")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./resource")

	// if you wanna directly read file properties
	// viper.SetConfigFile("./resource/properties-staging.yaml")

	viper.ReadInConfig()
}
