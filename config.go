package main

import (
	"encoding/json"
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
		User:     "postgres",
		Password: "secret",
		Name:     "products",
	}
}

// EnvPostgresConfig loads the postgres configuration from the env vars
// TODO: handle env vars that were not provided.
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

// LoadConfig loads a configuration file
// TODO: Decide on how to pass secrets
func LoadConfig(configReq bool) Config {
	f, err := os.Open(".config")
	if err != nil {
		if configReq {
			panic(err)
		}
		// Config from env vars
		fmt.Println("Using the env vars to config...")
		return EnvVarsConfig()
	}
	var c Config
	dec := json.NewDecoder(f)
	err = dec.Decode(&c)
	if err != nil {
		panic(err)
	}
	fmt.Println("Successfully loaded .config")
	fmt.Println(c.Database.Host)
	return c
}

// LoadEnvVarsConfig creates the config based on the env vars.
// Since configuration is only based on env vars, if one is not set
// it will report it.
func LoadEnvVarsConfig() Config {
	_, isSet := os.LookupEnv("POSTGRES_PORT")
	if !isSet {
		panic("POSTGRES_PORT not set")
	}
	_, isSet = os.LookupEnv("POSTGRES_USER")
	if !isSet {
		panic("POSTGRES_USER not set")
	}
	_, isSet = os.LookupEnv("POSTGRES_PASSWORD")
	if !isSet {
		panic("POSTGRES_PASSWORD not set")
	}
	_, isSet = os.LookupEnv("POSTGRES_DB")
	if !isSet {
		panic("POSTGRES_DB not set")
	}
	_, isSet = os.LookupEnv("POSTGRES_HOST")
	if !isSet {
		panic("POSTGRES_HOST not set")
	}

	// TODO:map with which are set and which are not
	fmt.Println("Using env vars to config connection... \n POSTGRES_USER, POSTGRES_PASSWORD, POSTGRES_DB, POSTGRES_HOST, POSTGRES_PORT")
	return EnvVarsConfig()
}
