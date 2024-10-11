package config

import (
	"errors"
	"flag"
	"fmt"
	"log"

	"github.com/spf13/viper"
)

type GRPCConfig struct {
	Port           int `yaml:"port"`
	TimeoutConnect int `yaml:"timeout"`
	StoragePath    string `yaml:"storagepath"`
}

type Config struct {
	ModeLog    string          `yaml:"modelog" env-default:"debug"`
	GRPC       *GRPCConfig     `yaml:"grpc"`
}

func LoadConfig(fileName, fileType string) (*Config, error) {
	var cfg *Config
	v := viper.New()
	v.SetConfigType(fileType)
	v.SetConfigName(fileName)
	v.AddConfigPath("./configs")
	v.AutomaticEnv()

	if err := v.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
		return nil, err
	}
	err := v.Unmarshal(&cfg)
	if err != nil {
		log.Printf("Unable to decode into struct, %v", err)
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			return nil, errors.New("config file not found")
		}
		return nil, err
	}

	return cfg, nil
}

func fetchConfig() string {
	var res string
	flag.StringVar(&res, "config", "", "path to config file")
	fmt.Println(res)
	flag.Parse()

	return res
}