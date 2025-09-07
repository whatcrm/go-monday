package monday

import logging "github.com/whatcrm/go-monday/logger"

func NewAPI(clientID, clientSecret string) *API {
	return &API{
		ClientID:     clientID,
		ClientSecret: clientSecret,
	}
}

func (api *API) SetOptions(domain, auth string, debug bool) {
	api.Domain = domain
	api.Auth = auth
	api.Debug = debug
	api.Version = April2025
	api.l = logging.GetLogger()
}

func (api *API) SetVersion(version string) {
	if version == "" {
		return
	}

	api.Version = version
}
