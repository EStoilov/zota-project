package config

import (
	"log"
	"os"
)

type Config struct {
	Port        string
	MerchantID  string
	SecretKey   string
	Currency    string
	EndpointID  string
	BaseURL     string
}

func LoadConfig() *Config {
	config := &Config{
		Port:       os.Getenv("PORT"),
		MerchantID: os.Getenv("MERCHANT_ID"),
		SecretKey:  os.Getenv("SECRET_KEY"),
		Currency:   os.Getenv("CURRENCY"),
		EndpointID: os.Getenv("ENDPOINT_ID"),
		BaseURL:    os.Getenv("BASE_URL"),
	}

	if config.Port == "" {
		config.Port = "8080"
	}

	if config.MerchantID == "" || config.SecretKey == "" || config.Currency == "" || config.EndpointID == "" || config.BaseURL == "" {
		log.Fatal("One or more required environment variables are missing")
	}

	return config
}