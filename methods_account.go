package monday

import (
	"github.com/gofiber/fiber/v2"
	"github.com/whatcrm/go-monday/models"
)

func (c *Get) Token(redirectURI string) (out *Tokens, err error) {
	options := callMethodOptions{
		Method:  fiber.MethodPost,
		BaseURL: tokenRequest,
		In:      nil,
		Out:     &out,
		Params: &RequestParams{
			ClientID:     c.api.ClientID,
			ClientSecret: c.api.ClientSecret,
			RedirectURI:  redirectURI,
			AuthCode:     c.api.Auth,
		},
	}

	err = c.api.callMethod(options)
	return
}

func (api *API) IsAdmin() (out interface{}, err error) {

	options := callMethodOptions{
		Method:  fiber.MethodGet,
		BaseURL: "",
		In:      nil,
		Out:     &out,
		Params:  nil,
	}

	err = api.callMethod(options)
	return
}

func (c *Get) Account(query *models.AccountQuery) (out *models.Account, err error) {
	options := callMethodOptions{
		Method:  fiber.MethodPost,
		BaseURL: mondayAPI,
		In:      query,
		Out:     &ResponseData{},
		Params:  nil,
	}

	if err = c.api.callMethod(options); err != nil {
		return
	}

	out = options.Out.(*ResponseData).Data.Account
	return
}

func (c *Get) Me() (out *models.Me, err error) {
	options := callMethodOptions{
		Method:  fiber.MethodPost,
		BaseURL: mondayAPI,
		In:      meQuery,
		Out:     &ResponseData{},
		Params:  nil,
	}

	if err = c.api.callMethod(options); err != nil {
		return
	}
	out = options.Out.(*ResponseData).Data.Me
	return
}

func (c *Get) Users(query *models.UserQuery) (out []models.User, err error) {
	options := callMethodOptions{
		Method:  fiber.MethodPost,
		BaseURL: mondayAPI,
		In:      query,
		Out:     &ResponseData{},
		Params:  nil,
	}

	if err = c.api.callMethod(options); err != nil {
		return
	}
	out = options.Out.(*ResponseData).Data.Users
	return
}
