package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	Database struct {
		Host     string
		Port     int
		User     string
		Password string
		Name     string
		SSLMode  string
	}
}

var Cfg *Config

// LoadConfig lee config.yaml o variables de entorno
func LoadConfig() *Config {
	v := viper.New()
	v.SetConfigName("config")
	v.SetConfigType("yaml")
	v.AddConfigPath("./config") // busca en ./config/config.yaml

	// Variables de entorno tienen prioridad
	v.AutomaticEnv()

	if err := v.ReadInConfig(); err != nil {
		log.Fatalf("Error leyendo archivo de configuración: %v", err)
	}

	cfg := &Config{}
	if err := v.Unmarshal(&cfg); err != nil {
		log.Fatalf("Error parseando configuración: %v", err)
	}

	Cfg = cfg
	return cfg
}
