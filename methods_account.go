package monday

import (
	"github.com/whatcrm/go-monday/models"
)

func (c *Get) Me() (out models.User, err error) {
	var query struct {
		Me models.User
	}

	options := makeRequestOptions{
		BaseURL: mondayAPI,
		Query:   &query,
	}

	err = c.api.makeRequest(options)
	out = query.Me
	return
}

func (c *Get) Account() (out models.Account, err error) {
	var query struct {
		Account models.Account
	}

	options := makeRequestOptions{
		BaseURL: mondayAPI,
		Query:   &query,
	}

	err = c.api.makeRequest(options)
	out = query.Account
	return
}
