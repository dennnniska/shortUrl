package config

import (
	"log"
	"os"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
)

type Config struct {
	Env             string     `yaml:"env" env-default:"local"`
	InMemoryStorage bool       `yaml:"in_memory_storage" env-default:"false"`
	GRPC            GRPCConfig `yaml:"grpc"`
	Http            HttpConfig `yaml:"http"`
	Postgres        Postgres   `yaml:"postgres"`
}

type GRPCConfig struct {
	Port    string        `yaml:"port"`
	Timeout time.Duration `yaml:"timeout"`
}

type HttpConfig struct {
	Address     string        `yaml:"address"`
	Timeout     time.Duration `yaml:"timeout"`
	IdleTimeout time.Duration `yaml:"idle_timeout"`
}

type Postgres struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	Username string `yaml:"username"`
	Dbname   string `yaml:"dbname"`
	Sslmode  string `yaml:"sslmode"`
}

func MustLoad() *Config {

	//configPath := "./config/local.yaml"

	if err := godotenv.Load(); err != nil {
		log.Fatal("error loading env", err)
	}
	configPath := os.Getenv("CONFIG_PATH")
	if configPath == "" {
		log.Fatal("CONFIG_PATH is not set")
	}

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatalf("config file does not exist: %s", configPath)
	}

	var cfg Config

	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		log.Fatalf("cannot read config: %s", err)
	}
	host := os.Getenv("SPRING_DATASOURCE_URL")
	if host != "" {
		cfg.Postgres.Host = host
	}
	portGRPC := os.Getenv("PORT_GRPC")
	if portGRPC != "" {
		cfg.GRPC.Port = portGRPC
	}
	portHTTP := os.Getenv("PORT_HTTP")
	if portHTTP != "" {
		cfg.Http.Address = portHTTP
	}
	portPostgres := os.Getenv("PORT_POSTGRES")
	if portPostgres != "" {
		cfg.Postgres.Port = portPostgres
	}
	return &cfg
}
