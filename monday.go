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
	api.l = logging.GetLogger()
}

func (api *API) SetLogPath(path string) {
	logging.SetPath(path)
}
