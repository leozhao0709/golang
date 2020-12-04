package main

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

type prodConfig struct {
	Username string
	Password string
	Database string
}

type DbConfig struct {
	Host     string
	Port     int
	Database string

	Dev struct {
		Username string
		Password string
	}

	Prod prodConfig
}

func main() {
	viper.SetConfigFile("src/viperTest/config/config1/db.yml")
	viper.BindEnv("password", "MYSQL_PASSWORD")
	viper.BindEnv("username", "MYSQL_USER")
	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		log.Fatalf("read db config error: %v", err)
	}

	// merge db2.yml
	viper.SetConfigFile("src/viperTest/config/config2/db.yml")
	err = viper.ReadInConfig()
	if err != nil { // Handle errors reading the config file
		log.Fatalf("read db config error: %v", err)
	}

	dbConfig := DbConfig{}
	err = viper.Unmarshal(&dbConfig)
	if err != nil {
		log.Fatalf("read db config error: %v", err)
	}

	dbConfig.Prod.Username = viper.GetString("username")
	dbConfig.Prod.Password = viper.GetString("password")

	fmt.Printf("%+v", dbConfig)
}
