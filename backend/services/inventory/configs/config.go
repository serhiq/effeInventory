package config

import (
	"fmt"
	envconfig "github.com/kelseyhightower/envconfig"
	"gopkg.in/yaml.v3"
	"os"
)

type Config struct {
	Project  Project  `yaml:"project"`
	DBConfig DBConfig `yaml:"database"`
}

type Project struct {
	Name        string `yaml:"name"`
	ServiceName string `yaml:"serviceName"`
	Version     string
	Timezone    string `yaml:"timezone"`
}

type DBConfig struct {
	Host         string `yaml:"host" envconfig:"DB_HOST"`
	Port         int    `yaml:"port" envconfig:"DB_PORT"`
	DatabaseName string `yaml:"database_name" envconfig:"DB_DATABASE_NAME"`
	Username     string `yaml:"username" envconfig:"DB_USERNAME"`
	Password     string `yaml:"password" envconfig:"DB_PASSWORD"`
}

func New() (*Config, error) {

	path := os.Getenv("CONFIG_PATH")
	if path == "" {
		path = "./configs/config.yaml"
	}

	config := &Config{}

	config.DBConfig.Username = "user_app"
	config.DBConfig.DatabaseName = "my_db"
	config.DBConfig.Password = "password"
	config.DBConfig.Host = "mysql"
	config.DBConfig.Port = 3306

	//if err := fromYaml(path, config); err != nil {
	//	fmt.Printf("couldn'n load config from %s: %s\r\n", path, err.Error())
	//}
	//
	//if err := fromEnv(config); err != nil {
	//	fmt.Printf("couldn'n load config from env: %s\r\n", err.Error())
	//}

	if err := validate(config); err != nil {
		return nil, err
	}

	return config, nil
}

func fromYaml(path string, config *Config) error {
	if path == "" {
		return nil
	}

	data, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	return yaml.Unmarshal(data, config)
}

func fromEnv(config *Config) error {
	return envconfig.Process("", config)
}

func validate(cfg *Config) error {

	if cfg.DBConfig.DatabaseName == "" {
		return fmt.Errorf("config: %s is not set", "DB_DATABASE_NAME")
	}

	if cfg.DBConfig.Username == "" {
		return fmt.Errorf("config: %s is not set", "DB_USERNAME")
	}

	return nil
}
