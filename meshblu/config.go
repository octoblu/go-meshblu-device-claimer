package meshblu

import "encoding/json"

// Config interfaces with a remote meshblu server
type Config struct {
	UUID   string `json:"uuid"`
	Token  string `json:"token"`
	Server string `json:"server"`
	Port   int    `json:"port"`
}

// NewConfig constructs a new Meshblu instance
func NewConfig(UUID, Token, Server string, Port int) *Config {
	return &Config{UUID, Token, Server, Port}
}

// ParseConfig creates a config with the UUID and Token from a JSON byte array
func ParseConfig(data []byte) (*Config, error) {
	config := &Config{}
	err := json.Unmarshal(data, config)
	return config, err
}

// ToJSON serializes the object to the meshblu.json format
func (config *Config) ToJSON() ([]byte, error) {
	return json.Marshal(config)
}
