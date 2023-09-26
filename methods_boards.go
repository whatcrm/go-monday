package monday

import "github.com/whatcrm/go-monday/models"

type BoardParams struct {
	IDs          []int `json:"ids"`
	NewestFirst  bool  `json:"newest_first"`
	Limit        int   `json:"limit"`
	Page         int   `json:"page"`
	BoardKind    `json:"board_kind"`
	State        `json:"state"`
	OrderBy      BoardsOrderBy `json:"order_by"`
	WorkspaceIDs []int         `json:"workspace_ids"`
}

func (c *Get) Boards() (out []models.Board, err error) {

	var query struct {
		Board []models.Board `graphql:"boards"`
	}

	//variables := map[string]interface{}{
	//"limit":         p.Limit,
	//"page":          p.Page,
	//"ids":           p.IDs,
	//"newest_first":  p.NewestFirst,
	//"board_kind":    p.BoardKind,
	//"state":         p.State,
	//"order_by":      p.OrderBy,
	//"workspace_ids": p.WorkspaceIDs,
	//}
	//
	options := makeRequestOptions{
		BaseURL: mondayAPI,
		Query:   &query,
		//Variables: variables,
	}

	err = c.api.makeRequest(options)
	out = query.Board
	return
}
