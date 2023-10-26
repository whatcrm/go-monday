package monday

import (
	"github.com/whatcrm/go-monday/models"
)

func (c *Get) Assets(id []ID) (out []models.Assets, err error) {
	var query struct {
		Assets []models.Assets `graphql:"assets ( ids: $id )"`
	}

	variables := map[string]interface{}{
		"id": id,
	}

	options := makeRequestOptions{
		BaseURL:   mondayAPI,
		Query:     &query,
		Variables: variables,
	}

	err = c.api.makeRequest(options)
	out = query.Assets
	return
}
