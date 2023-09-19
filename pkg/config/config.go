package config

import (
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
	"time"
)

const (
	ENV_LOCAL = "local"
	ENV_DEV   = "dev"
	ENV_PROD  = "prod"
)

type Config struct {
	Env         string     `yaml:"env" env:"ENV" env-default:"local" env-required:"true"`
	StoragePath string     `yaml:"storage_path" env-required:"true"`
	HTTPServer  HTTPServer `yaml:"http_server"`
	DBConfig    DBConfig   `yaml:"db_config"`
}

type HTTPServer struct {
	Address      string        `yaml:"address" env-default:"localhost:8080"`
	ReadTimeout  time.Duration `yaml:"read_timeout" env-default:"10s"`
	WriteTimeout time.Duration `yaml:"idle_timeout" env-default:"10s"`
}

type DBConfig struct {
	Host     string `yaml:"host" env-required:"true"`
	Port     string `yaml:"port" env-required:"true"`
	Username string `yaml:"username" env-required:"true"`
	//Password string `yaml:"password" env-required:"false"`
	DBName  string `yaml:"db_name" env-required:"true"`
	SSLMode string `yaml:"ssl_mode" env-required:"true"`
}

func MustLoad() *Config {
	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("could not load .env file: %s", err)
	}

	configPath := os.Getenv("CONFIG_PATH")
	if configPath == "" {
		logrus.Fatal("config path is empty")
	}

	// check if file exists
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		logrus.Fatalf("config file does not exist: %s", configPath)
	}

	var cfg Config

	viper.SetConfigFile(configPath)

	if err := viper.ReadInConfig(); err != nil {
		logrus.Fatalf("cannot read config: %s", err.Error())
	}

	logrus.Warning(viper.GetString("env"))

	if err := viper.Unmarshal(&cfg); err != nil {
		logrus.Fatalf("cannot read config: %s", err)
	}

	// cfg.DBConfig.Password = os.Getenv("DB_PASSWORD")
	return &cfg
}
