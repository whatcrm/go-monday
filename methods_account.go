package monday

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"strings"
)

func (c *Get) Token(redirectURI string) (out interface{}, err error) {
	options := callMethodOptions{
		Method:  fiber.MethodGet,
		BaseURL: oAuth,
		In:      "",
		Out:     nil,
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
		In:      "",
		Out:     &out,
		Params:  nil,
	}

	err = api.callMethod(options)
	return
}

func (c *Get) Users(id []int) (out *ResponseData, err error) {
	q := UsersQuery

	if len(id) == 1 {
		q = fmt.Sprintf(UserQueryID, id[0])
	}

	if len(id) > 1 {
		idList := strings.Join(strings.Fields(fmt.Sprint(id)), ", ")
		q = fmt.Sprintf(UserQueryIDs, idList)
	}

	options := callMethodOptions{
		Method:  fiber.MethodPost,
		BaseURL: mondayAPI,
		In:      q,
		Out:     &out,
		Params:  nil,
	}

	err = c.api.callMethod(options)
	return
}
