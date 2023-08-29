package monday

type API struct {
	ClientID     string
	ClientSecret string
	Domain       string
	Auth         string
	Debug        bool
}

type RequestParams struct {
	ID string `json:"id"`

	ClientID     string `json:"client_id"`
	RedirectURI  string `json:"redirect_uri"`   // optional, when more than 1 redirect_uri
	Scope        string `json:"scope"`          // optional
	State        string `json:"state"`          // optional
	AppVersionID string `json:"app_version_id"` // optional
}
