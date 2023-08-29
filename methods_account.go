package monday

import (
	"github.com/gofiber/fiber/v2"
)

func (api *API) Authorize() (out interface{}, err error) {
	options := callMethodOptions{
		Method:  fiber.MethodGet,
		BaseURL: oAuth,
		In:      nil,
		Out:     nil,
		Params: &RequestParams{
			ClientID: api.ClientID,
		},
	}

	err = api.callMethod(options)
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
