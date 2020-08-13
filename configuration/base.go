package configuration

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

type Configuration struct {
	Sql struct {
		Sqlite3     bool `json:"sqlite"`
		Addr     string `json:"addr"`
	} `json:"sql"`
	Redis struct {
		Addr     string `json:"addr"`
		Password     string `json:"password"`
		Db     int `json:"db"`
	} `json:"redis"`
	Admin     struct {
		Name     string `json:"name"`
		Mail     string `json:"mail"`
		Password string `json:"password"`
	} `json:"admin"`
	Address   string `json:"address"`
	Secretkey string `json:"secretkey"`
}

var C Configuration

func init() {
	var err error

	viper.SetConfigName("configuration")
	viper.AddConfigPath("./")
	viper.SetConfigType("yaml")

	if err = viper.ReadInConfig(); err != nil {
		fmt.Printf("config file error: %s\n", err)
		os.Exit(1)
	}
	if err = viper.Unmarshal(&C); err != nil {
		fmt.Println("config file error:", err)
		os.Exit(1)
	}

	fmt.Println("Configuration.C", C)
}