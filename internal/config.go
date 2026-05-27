package killgrave

import (
	"errors"
)

// Config representation of config file yaml
type Config struct {
	ImpostersPath string      `yaml:"imposters_path"`
	Port          int         `yaml:"port"`
	Host          string      `yaml:"host"`
	CORS          ConfigCORS  `yaml:"cors"`
	Proxy         ConfigProxy `yaml:"proxy"`
	Secure        bool        `yaml:"secure"`
	Watcher       bool        `yaml:"watcher"`
}

// ConfigCORS representation of section CORS of the yaml
type ConfigCORS struct {
	Methods          []string `yaml:"methods"`
	Headers          []string `yaml:"headers"`
	Origins          []string `yaml:"origins"`
	ExposedHeaders   []string `yaml:"exposed_headers"`
	AllowCredentials bool     `yaml:"allow_credentials"`
}

// ConfigProxy is a representation of section proxy of the yaml
type ConfigProxy struct {
	Url  string    `yaml:"url"`
	Mode ProxyMode `yaml:"mode"`
}

// ProxyMode is enumeration of proxy server modes
type ProxyMode uint8

const (
	// ProxyNone server is off
	ProxyNone ProxyMode = iota
	// ProxyMissing handle only missing requests are proxied
	ProxyMissing
	// ProxyAll all requests are proxied
	ProxyAll
)

var (
	errInvalidConfigPath  = errors.New("invalid config file")
	errEmptyImpostersPath = errors.New("imposters path can not be blank")
	errEmptyHost          = errors.New("host can not be blank")
	errInvalidPort        = errors.New("invalid port")
)

func (p ProxyMode) String() string { _ = "STUB: not implemented"; return "" }

// StringToProxyMode convert string into a ProxyMode if not exists return a none mode and an error
func StringToProxyMode(t string) (ProxyMode, error) {
	_ = "STUB: not implemented"
	return *new(ProxyMode), nil
}

// UnmarshalYAML implementation of yaml.Unmarshaler interface
func (p *ProxyMode) UnmarshalYAML(unmarshal func(interface{}) error) error {
	_ = "STUB: not implemented"
	return nil
}

// ConfigureProxy preparing the server with the proxy configuration that the user has indicated
func (cfg *Config) ConfigureProxy(proxyMode ProxyMode, proxyURL string) {
	_ = "STUB: not implemented"
	return
}

// ConfigOpt function to encapsulate optional parameters
type ConfigOpt func(cfg *Config) error

// NewConfig initialize the config
func NewConfig(impostersPath, host string, port int, secure bool) (Config, error) {
	_ = "STUB: not implemented"
	return *new(Config), nil
}

// NewConfigFromFile  unmarshal content of config file to initialize a Config struct
func NewConfigFromFile(cfgPath string) (Config, error) {
	_ = "STUB: not implemented"
	return *new(Config), nil
}
