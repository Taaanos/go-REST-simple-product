package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type PostgresConfig struct {
	Host     string `json: "host"`
	Port     int    `json: "port"`
	User     string `json: "user"`
	Password string `json: "password"`
	Name     string `json: "name"`
}

func (c PostgresConfig) Dialect() string {
	return "postgres"
}

func (c PostgresConfig) ConnectionInfo() string {
	if c.Password == "" {
		return fmt.Sprintf("host=%s port=%d user=%s  dbname=%s sslmode=disable",
			c.Host, c.Port, c.User, c.Name)
	}
	return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		c.Host, c.Port, c.User, c.Password, c.Name)
}

func DefaultPostgresConfig() PostgresConfig {
	return PostgresConfig{
		Host:     "localhost",
		Port:     5432,
		User:     "root",
		Password: "secret",
		Name:     "products",
	}
}

type Config struct {
	Port     int            `json:"port"`
	Env      string         `json:"env"`
	Database PostgresConfig `json:"database"`
}

func (c Config) IsProd() bool {
	return c.Env == "prod"
}

func DefaultConfig() Config {
	return Config{
		Port:     3000,
		Env:      "dev",
		Database: DefaultPostgresConfig(),
	}
}

func LoadConfig(configReq bool) Config {
	f, err := os.Open(".config")
	if err != nil {
		if configReq {
			panic(err)
		}
		fmt.Println("Using the default config...")
		return DefaultConfig()
	}
	var c Config
	dec := json.NewDecoder(f)
	err = dec.Decode(&c)
	if err != nil {
		panic(err)
	}
	fmt.Println("Successfully loaded .config")
	return c
}

// func LoadEnvVars() Config {
// 	var c Config
// 	c.Database.Name = os.Getenv("APP_DB_NAME")
// 	c.Database.User = os.Getenv("APP_DB_USERNAME")
// 	c.Database.Password = os.Getenv("APP_DB_PASSWORD")
// 	return c
// }
