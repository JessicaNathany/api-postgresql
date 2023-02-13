package configs

import "github.com/spf13/viper"

// variável privada de configuração
var cfg *config

type config struct {
	API ApiConfig
	DB  DbConfig
}

type ApiConfig struct {
	Port string
}

// configuraçãos do banco de dados para fazer as transações
type DbConfig struct {
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
	viper.SetConfigType("")
	viper.AddConfigPath(".")
	error := viper.ReadInConfig()
	if error != nil {
		if _, ok := error.(viper.ConfigFileNotFoundError); !ok {
			return error
		}
	}

	cfg = new(config)
	cfg.API = ApiConfig{
		Port: viper.GetString("api.port"),
	}

	cfg.DB = DbConfig{
		Host:     viper.GetString("database.host"),
		Port:     viper.GetString("database.host"),
		User:     viper.GetString("database.host"),
		Pass:     viper.GetString("database.host"),
		Database: viper.GetString("database.name"),
	}

	return nil
}

func GetDB() DbConfig {
	return cfg.DB
}

func SetServerPort() string {
	return cfg.API.Port
}
