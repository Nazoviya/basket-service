package util

import "github.com/spf13/viper"

// Config struct to store environmental values.
type Config struct {
	DBDriver                string  `mapstructure:"DB_DRIVER"`
	DBSource                string  `mapstructure:"DB_SOURCE"`
	ServerAddress           string  `mapstructure:"SERVER_ADDRESS"`
	MinBasketTotal          float32 `mapstructure:"MIN_BASKET_TOTAL"`
	MinTotalPurchaseInMonth float32 `mapstructure:"MIN_TOTAL_PURCHASE_IN_MONTH"`
}

// LoadConfig reads configuration from app.env file and loads data into required fields.
func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}
	err = viper.Unmarshal(&config)
	return
}
