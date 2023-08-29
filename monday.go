package monday

import (
	"fmt"
)

func NewAPI(clientID, clientSecret string) *API {
	return &API{
		ClientID:     clientID,
		ClientSecret: clientSecret,
	}
}

func (api *API) SetOptions(domain, auth string, debug bool) error {
	api.Domain = domain

	if auth != "" {
		api.Auth = auth
	} else {
		return fmt.Errorf("accessToken is not set")
	}

	api.Debug = debug

	return nil
}
