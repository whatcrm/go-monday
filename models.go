package monday

import logging "github.com/whatcrm/go-monday/logger"

type API struct {
	ClientID     string
	ClientSecret string
	Domain       string
	Auth         string
	Version      string
	Debug        bool
	l            logging.Logger
}

type Tokens struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	Scope       string `json:"scope"`
}
