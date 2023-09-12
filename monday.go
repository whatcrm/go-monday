package monday

func NewAPI(clientID, clientSecret string) *API {
	return &API{
		ClientID:     clientID,
		ClientSecret: clientSecret,
	}
}

func (api *API) SetOptions(domain, auth string, debug bool) error {
	api.Domain = domain
	api.Auth = auth
	api.Debug = debug
	return nil
}
