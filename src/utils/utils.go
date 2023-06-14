package utils

import (
	"feynman/config"
	"fmt"
	"net/url"
	"os"

	"gopkg.in/yaml.v3"
)

// function to transform the credentials to a URL encoded string
func CredentialsToURLString(config *config.Configuration) string {
	return fmt.Sprintf("%s:%s", url.QueryEscape(config.Credentials.Username), url.QueryEscape(config.Credentials.Password))
}

// function to load the YAML config file from the given path
func LoadConfig(configFilePath string) (*config.Configuration, error) {
	buf, err := os.ReadFile(configFilePath)
	if err != nil {
		return nil, err
	}

	c := &config.Configuration{}
	err = yaml.Unmarshal(buf, c)
	if err != nil {
		return nil, fmt.Errorf("in file %q: %w", configFilePath, err)
	}

	return c, err
}
