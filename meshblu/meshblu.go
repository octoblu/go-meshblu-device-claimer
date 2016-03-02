package meshblu

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

// Meshblu interfaces with a remote meshblu server
type Meshblu interface {
	Register(deviceType string) (*Config, error)
}

// HTTPClient interfaces with a remote meshblu server
type HTTPClient struct {
	meshbluURI string
}

// New constructs a new Meshblu instance
func New(meshbluURI string) Meshblu {
	return &HTTPClient{meshbluURI}
}

// Register creates a new device of the specified type and returns it
func (meshblu *HTTPClient) Register(deviceType string) (*Config, error) {
	meshbluURL, err := ParseURL(meshblu.meshbluURI)
	if err != nil {
		return nil, err
	}

	meshbluURL.SetPath("/devices")

	responseBody, err := doRegister(meshbluURL.String(), url.Values{"type": {deviceType}})
	if err != nil {
		return nil, err
	}

	config, err := ParseConfig(responseBody)
	if err != nil {
		return nil, err
	}

	config.Server = meshbluURL.HostName()
	config.Port = meshbluURL.Port()

	return config, nil
}

func doRegister(uri string, values url.Values) ([]byte, error) {
	response, err := http.PostForm(uri, values)
	if err != nil {
		return nil, err
	}
	if response.StatusCode != 201 {
		return nil, fmt.Errorf("Meshblu register returned invalid response code: %v", response.StatusCode)
	}

	return ioutil.ReadAll(response.Body)
}
