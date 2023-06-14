package utils

import (
	"feynman/config"
	"fmt"
	"net/url"
	"os"
	"strings"

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

func InsertCredentialsIntoProxyURLs(config *config.Configuration) string {
	proxyUrl := config.ProxyURL

	// escape the username and password
	username := url.QueryEscape(config.Username)
	password := url.QueryEscape(config.Password)

	// create the auth string
	auth_string := username + ":" + password + "@"

	// insert credentials into proxy url
	if strings.HasPrefix(proxyUrl, "http://") || strings.HasPrefix(proxyUrl, "https://") {
		proxyUrl = strings.Replace(proxyUrl, "://", "://"+auth_string, 1)
	} else {
		proxyUrl = "https://" + auth_string + proxyUrl
	}

	return proxyUrl
}
