package config

type File struct {
	Path      string   `yaml:"path"`
	Variables []string `yaml:"variables"`
	Export    bool     `yaml:"export"`
}

type Configuration struct {
	// the credentials
	Username string `yaml:"username"`
	Password string `yaml:"password"`

	// the proxy URLs
	ProxyURL string `yaml:"proxy_url"`

	// the files to edit
	Files []File `yaml:"files"`
}
