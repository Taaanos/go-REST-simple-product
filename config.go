package main

import (
	"fmt"
	"os"
	"strconv"
)

// PostgresConfig is the struct for the neccessary parameters to connect to
// your postgres db.
type PostgresConfig struct {
	Host     string `json: "host"`
	Port     int    `json: "port"`
	User     string `json: "user"`
	Password string `json: "password"`
	Name     string `json: "name"`
}

// Dialect returns the dialect for the db driver
func (c PostgresConfig) Dialect() string {
	return "postgres"
}

// ConnectionInfo builds the string to be used in the sql driver to open the db
func (c PostgresConfig) ConnectionInfo() string {
	if c.Password == "" {
		return fmt.Sprintf("host=%s port=%d user=%s  dbname=%s sslmode=disable",
			c.Host, c.Port, c.User, c.Name)
	}
	return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		c.Host, c.Port, c.User, c.Password, c.Name)
}

// DefaultPostgresConfig returns a default hardcoded postgres configuration
func DefaultPostgresConfig() PostgresConfig {
	return PostgresConfig{
		Host:     "localhost",
		Port:     5432,
		User:     "root",
		Password: "secret",
		Name:     "products",
	}
}

// EnvPostgresConfig loads the postgres configuration from the env vars
func EnvPostgresConfig() PostgresConfig {
	port, _ := strconv.Atoi(os.Getenv("POSTGRES_PORT"))
	return PostgresConfig{
		Host:     os.Getenv("POSTGRES_HOST"),
		Port:     port,
		User:     os.Getenv("POSTGRES_USER"),
		Password: os.Getenv("POSTGRES_PASSWORD"),
		Name:     os.Getenv("POSTGRES_DB"),
	}
}

// Config is the struct to hold the configuration parameters for the app
type Config struct {
	Port     int            `json:"port"`
	Env      string         `json:"env"`
	Database PostgresConfig `json:"database"`
}

// DefaultConfig is the default configuration of the app.
func DefaultConfig() Config {
	return Config{
		Port:     3010,
		Env:      "dev",
		Database: DefaultPostgresConfig(),
	}
}

// EnvVarsConfig returns a configuration based on env vars
func EnvVarsConfig() Config {
	return Config{
		Port:     3010,
		Env:      "stage",
		Database: EnvPostgresConfig(),
	}
}

// LoadConfig creates the config based on the env vars if they exist
// or else gets the default.
// Configuration in production only based on env vars, if one is not set
// it will report it.
func LoadConfig() Config {
	notSet := 0
	_, isSet := os.LookupEnv("POSTGRES_PORT")
	if !isSet {
		fmt.Println("POSTGRES_PORT not set")
		notSet++
	}
	_, isSet = os.LookupEnv("POSTGRES_USER")
	if !isSet {
		fmt.Println("POSTGRES_USER not set")
		notSet++
	}
	_, isSet = os.LookupEnv("POSTGRES_PASSWORD")
	if !isSet {
		fmt.Println("POSTGRES_PASSWORD not set")
		notSet++
	}
	_, isSet = os.LookupEnv("POSTGRES_DB")
	if !isSet {
		fmt.Println("POSTGRES_DB not set")
		notSet++
	}
	_, isSet = os.LookupEnv("POSTGRES_HOST")
	if !isSet {
		fmt.Println("POSTGRES_HOST not set")
		notSet++
	}
	if notSet != 0 {
		fmt.Println("Using default config...")
		return DefaultConfig()
	}
	fmt.Println("Using env vars to config connection... \n POSTGRES_USER, POSTGRES_PASSWORD, POSTGRES_DB, POSTGRES_HOST, POSTGRES_PORT")
	return EnvVarsConfig()
}
