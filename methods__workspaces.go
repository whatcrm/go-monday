package monday

import "github.com/whatcrm/go-monday/models"

func (c *Get) Workspaces() (out []models.Workspace, err error) {

	var query struct {
		Workspaces []models.Workspace `graphql:"workspaces"`
	}

	options := makeRequestOptions{
		BaseURL: mondayAPI,
		Query:   &query,
	}

	err = c.api.makeRequest(options)
	out = query.Workspaces
	return
}
