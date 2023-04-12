package configuration

import "os"

type Config struct {
	LoRaWANHost string
}

var appConfig *Config

func (entity *Config) Load() {
	entity.LoRaWANHost = os.Getenv("LORAWAN_HOST")

	appConfig = entity
}

func GetConfig() *Config {
	return appConfig
	// if appConfig != nil {
	// 	return appConfig
	// }
	// var conf Configuration
	// conf.Load()
	// appConfig = &conf
	// return appConfig
}
