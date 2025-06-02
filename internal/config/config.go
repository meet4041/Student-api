package config

import (
	"flag"
	"log"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

type HttpServer struct {
	Address string `yaml:"address" env-required:"true"`
}

// env-default:"production"

type Config struct {
	Env         string     `yaml:"env" env:"ENV" env-required:"true"`
	StoragePath string     `yaml:"storage_path" env-required:"true"`
	HttpServer  HttpServer `yaml:"http_server"`
}

func MustLoad() *Config {
	var configPath string
	configPath = os.Getenv("CONFIG_PATH")

	if configPath == "" {
		flags := flag.String("config", "", "path to the configration file")
		flag.Parse()

		configPath = *flags

		if configPath == "" {
			log.Fatal("Path not set")
		}
	}

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatalf("file doesn't exist: %s", configPath)
	}

	var cfg Config
	err := cleanenv.ReadConfig(configPath, &cfg)
	if err != nil {
		log.Fatalf("can't read file %s:", err.Error())
	}

	return &cfg
}
