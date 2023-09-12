package monday

import "github.com/whatcrm/go-monday/models"

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

type RequestParams struct {
	ID string `json:"id"`

	ClientID     string `json:"client_id"`
	Subdomain    string `json:"subdomain"`
	RedirectURI  string `json:"redirect_uri"`   // optional, when more than 1 redirect_uri
	Scope        string `json:"scope"`          // optional
	State        string `json:"state"`          // optional
	AppVersionID string `json:"app_version_id"` // optional

	ClientSecret string `json:"client_secret"`
	AuthCode     string `json:"code"`
}

type Data struct {
	Users []models.User `json:"users,omitempty"`
}

type ResponseData struct {
	AccountID float64 `json:"account_id"`
	Data      *Data   `json:"data"`
}

type ErrorResponse struct {
	Errors []struct {
		Message   string `json:"message"`
		Locations []struct {
			Line   int `json:"line"`
			Column int `json:"column"`
		} `json:"locations"`
		Path       []string `json:"path"`
		Extensions struct {
			Code     string `json:"code"`
			NodeName string `json:"nodeName"`
			TypeName string `json:"typeName"`
		} `json:"extensions"`
	} `json:"errors"`
	AccountID        int    `json:"account_id"`
	Error            string `json:"error"`
	ErrorDescription string `json:"error_description"`
}

type ErrorResp struct {
	Error            string `json:"error"`
	ErrorDescription string `json:"error_description"`
}
