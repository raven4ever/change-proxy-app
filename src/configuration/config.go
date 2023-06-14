package configuration

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type Configuration struct {
	// the credentials
	Credentials struct {
		Username string `yaml:"username"`
		Password string `yaml:"password"`
	} `yaml:"credentials"`

	// the proxy URLs
	HttpProxy  string `yaml:"http_proxy_url"`
	HttpsProxy string `yaml:"https_proxy_url"`

	// the files to edit
	Files []struct {
		Path      string   `yaml:"path"`
		Variables []string `yaml:"variables"`
	} `yaml:"files"`
}

// function to load the YAML config file from the given path
func LoadConfig(configFilePath string) (*Configuration, error) {
	buf, err := os.ReadFile(configFilePath)
	if err != nil {
		return nil, err
	}

	c := &Configuration{}
	err = yaml.Unmarshal(buf, c)
	if err != nil {
		return nil, fmt.Errorf("in file %q: %w", configFilePath, err)
	}

	return c, err
}
