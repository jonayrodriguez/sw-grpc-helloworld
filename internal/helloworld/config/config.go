package helloword

import (
	"io/ioutil"

	"github.com/kelseyhightower/envconfig"
	"gopkg.in/validator.v2"
	"gopkg.in/yaml.v2"
)

// Config represents an application configuration.
type Config struct {
	App    App    `yaml:"app"`
	Server Server `yaml:"server"`
}

// App represents the application information.
type App struct {
	Name       string `yaml:"name" envconfig:"APP_NAME" validate:"nonzero"`
	Repository string `yaml:"repository" envconfig:"APP_REPOSITORY" validate:"nonzero"`
}

// Server represents an server configuration.
type Server struct {
	Port int    `yaml:"port" envconfig:"SERVER_PORT" validate:"nonzero"`
	Host string `yaml:"host" envconfig:"SERVER_HOST" validate:"nonzero"`
}

// Validate validates the application configuration.
func (c Config) Validate() error {
	return validator.Validate(&c)

}

// LoadConfiguration returns an application configuration which is populated from the given configuration file and environment variables.
func LoadConfiguration(file string) (*Config, error) {
	// default empty config
	c := Config{}

	// load from YAML config file
	bytes, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}
	if err = yaml.Unmarshal(bytes, &c); err != nil {
		return nil, err
	}

	// load from environment variables prefixed with "APP_"
	if err = envconfig.Process("", &c); err != nil {
		return nil, err
	}

	// validation
	if err = c.Validate(); err != nil {
		return nil, err
	}

	return &c, err
}
