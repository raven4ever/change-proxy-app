package utils

import (
	"feynman/config"
	"fmt"
	"net/url"
	"os"

	"gopkg.in/yaml.v3"
)

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

func InsertCredentialsIntoProxyURLs(config *config.Configuration) (string, string, error) {
	// parse the proxy URLs
	http_proxy_url, err := url.Parse(config.HttpProxy)
	if err != nil {
		return "", "", err
	}

	https_proxy_url, err := url.Parse(config.HttpsProxy)
	if err != nil {
		return "", "", err
	}

	// insert the credentials into the proxy URLs
	http_proxy_url.User = url.UserPassword(config.Credentials.Username, config.Credentials.Password)
	https_proxy_url.User = url.UserPassword(config.Credentials.Username, config.Credentials.Password)

	return http_proxy_url.String(), https_proxy_url.String(), nil
}
