package config

type Credentials struct {
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}

type File struct {
	Path      string   `yaml:"path"`
	Variables []string `yaml:"variables"`
}

type Configuration struct {
	// the credentials
	Credentials Credentials `yaml:"credentials"`

	// the proxy URLs
	ProxyURL string `yaml:"proxy_url"`

	// the files to edit
	Files []File `yaml:"files"`
}
