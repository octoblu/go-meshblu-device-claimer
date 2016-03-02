package gateblu

import (
	"fmt"
	"net/url"

	"github.com/skratchdot/open-golang/open"
)

// Gateblu interfaces with Octoblu UI
type Gateblu struct {
	claimURI string
}

// New constucts a new Gateblu instance
func New(claimURI string) *Gateblu {
	return &Gateblu{claimURI}
}

// Claim opens the device in the claim uri
func (octoblu *Gateblu) Claim(uuid, token string) error {
	uri, err := generateURI(octoblu.claimURI, uuid, token)
	if err != nil {
		return err
	}
	open.Run(uri.String())
	return nil
}

func generateURI(claimURI, uuid, token string) (*url.URL, error) {
	uri, err := url.Parse(claimURI)
	if err != nil {
		return nil, err
	}
	uri.Path = fmt.Sprintf("/gateblu/%v/claim", uuid)
	query := uri.Query()
	query.Set("token", token)
	uri.RawQuery = query.Encode()
	return uri, nil
}
