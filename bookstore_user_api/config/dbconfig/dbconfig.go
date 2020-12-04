package dbconfig

import (
	"fmt"

	"github.com/labstack/gommon/log"
	"github.com/leozhao0709/golang/bookstore_user_api/env"
	"github.com/spf13/viper"
)

// IDbConfig db config interface
type IDbConfig interface {
	GetDataSource() string
	GetDriver() string
	GetMaxIdleConns() int
	GetMaxOpenConns() int
	GetConnMaxLifetime() int
}

type dbConfig struct {
	Driver          string
	Host            string
	Port            int
	Database        string
	Username        string
	Password        string
	MaxIdleConns    int `mapstructure:"max_idle_conns"`
	MaxOpenConns    int `mapstructure:"max_open_conns"`
	ConnMaxLifetime int `mapstructure:"conn_max_lifetime"`
}

// GetConfig getdbConfig
func GetConfig() IDbConfig {
	dbViper := viper.New()

	environment := env.GetCurrentEnv()

	switch environment {
	case "prod":
		dbViper.SetConfigFile("config/dbconfig/db_prod.yml")
	case "test":
		dbViper.SetConfigFile("config/dbconfig/db_test.yml")
	default: // dev
		dbViper.SetConfigFile("config/dbconfig/db_dev.yml")
		dbViper.BindEnv("username", "MYSQL_USER")
		dbViper.BindEnv("password", "MYSQL_PASSWORD")
	}

	err := dbViper.ReadInConfig()
	if err != nil {
		log.Fatal("DB Config read err: ", err)
	}

	config := dbConfig{}
	err = dbViper.Unmarshal(&config)
	if err != nil {
		log.Fatal("DB Config unmarshal err: ", err)
	}

	return config
}

func (config dbConfig) GetDataSource() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true", config.Username, config.Password, config.Host, config.Port, config.Database)
}

func (config dbConfig) GetDriver() string {
	return config.Driver
}

func (config dbConfig) GetMaxIdleConns() int {
	return config.MaxIdleConns
}

func (config dbConfig) GetMaxOpenConns() int {
	return config.MaxOpenConns
}

func (config dbConfig) GetConnMaxLifetime() int {
	return config.ConnMaxLifetime
}
