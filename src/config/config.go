package config

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
