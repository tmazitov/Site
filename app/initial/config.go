package initial

import (

	"github.com/spf13/viper"
)

func Config() error {
	viper.AddConfigPath("./configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
