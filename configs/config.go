package configs

import (
	"os"
	"payment-s1/pkg/db/mysql_db"

	yaml "github.com/go-yaml/yaml"
)

type Config struct {
	Mysqldb mysql_db.DatabaseConfig `yaml:"mysql"`
}

// NewConfig returns a new decoded Config struct
func NewConfig() (*Config, error) {
	// Create config structure
	config := &Config{}

	// Open config file
	file, err := os.Open("./config.yml")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	d := yaml.NewDecoder(file)

	// Start YAML decoding from file
	if err := d.Decode(&config); err != nil {
		return nil, err
	}

	return config, nil
}
