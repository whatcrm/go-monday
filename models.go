package monday

type API struct {
	ClientID     string
	ClientSecret string
	Domain       string
	Auth         string
	Debug        bool
}

type Tokens struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	Scope       string `json:"scope"`
}
