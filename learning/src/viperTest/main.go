package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/spf13/viper"
)

type DbConfig struct {
	Host     string
	Port     int
	Database string
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
}

func main() {

	var environment = flag.String("env", "dev", "which env here?(prod/dev/test)")
	flag.Parse()

	viper.SetConfigName("default")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("src/viperTest/config/db")
	err := viper.ReadInConfig()
	if err != nil { // Handle errors reading the config file
		log.Fatalf("read default db config error: %v", err)
	}

	fmt.Println(*environment)
	switch *environment {
	case "prod":
		viper.SetConfigName("production")
	case "test":
		viper.SetConfigName("development")
	default: // dev
		viper.SetConfigName("development")
	}

	err = viper.MergeInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		log.Fatalf("db merge error: %v", err)
	}

	viper.SetEnvPrefix("mysql")
	viper.AutomaticEnv() // merge environment

	dbConfig := DbConfig{}
	err = viper.Unmarshal(&dbConfig)
	if err != nil {
		log.Fatalf("read db config error: %v", err)
	}

	fmt.Printf("%+v", dbConfig)
}
