package configs

import "github.com/spf13/viper"

// The package config with viper prevents global variables
var conf *config

type config struct {
	API APIConfig
	DB  DBConfig
	SW  SWConfig
}

type APIConfig struct {
	Port string
}

type SWConfig struct {
	Port string
}

type DBConfig struct {
	Host     string
	Port     string
	User     string
	Pass     string
	Database string
}

func init() {
	viper.SetDefault("api.port", "9000")
	viper.SetDefault("database.host", "localhost")
	viper.SetDefault("database.port", "5432")
}

func Load() error {
	viper.SetConfigName("config")
	viper.SetConfigType("toml")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			return err
		}
	}

	conf = new(config)

	conf.API = APIConfig{
		Port: viper.GetString("api.port"),
	}

	conf.SW = SWConfig{
		Port: viper.GetString("swagger.port"),
	}

	conf.DB = DBConfig{
		Host:     viper.GetString("database.host"),
		Port:     viper.GetString("database.port"),
		User:     viper.GetString("database.user"),
		Pass:     viper.GetString("database.pass"),
		Database: viper.GetString("database.name"),
	}

	return nil
}

func GetDB() DBConfig {
	return conf.DB
}

func GetServerPort() string {
	return conf.API.Port
}

func GetSwaggerPort() string {
	return conf.SW.Port
}
