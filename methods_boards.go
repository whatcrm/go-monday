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

func (c *Get) Board(id ID) (out models.Board, err error) {
	var query struct {
		Board []models.Board `graphql:"boards ( ids: $ids ) "`
	}

	var ids = []ID{id}

	variables := map[string]interface{}{
		"ids": ids,
	}

	options := makeRequestOptions{
		BaseURL:   mondayAPI,
		Query:     &query,
		Variables: variables,
	}

	err = c.api.makeRequest(options)

	if query.Board != nil {
		out = query.Board[0]
	}
	return
}

func (c *Mutate) CreateBoard(name string, kind BoardKind, workspace ID) (out models.Board, err error) {
	var mutation struct {
		Board models.Board `graphql:"create_board ( board_name: $name board_kind: $kind workspace_id: $workspace_id ) "`
	}

	variables := map[string]interface{}{
		"name":         name,
		"kind":         kind,
		"workspace_id": workspace,
	}

	options := makeRequestOptions{
		BaseURL:   mondayAPI,
		Mutation:  &mutation,
		Variables: variables,
	}

	err = c.api.makeRequest(options)

	out = mutation.Board
	return
}

func (c *Get) Groups(id []ID) (out []models.NestedGroup, err error) {

	var query struct {
		Board []models.Board `graphql:"boards ( ids: $ids ) "`
	}

	variables := map[string]interface{}{
		//"limit":         p.Limit,
		//"page":          p.Page,
		"ids": id,
		//"newest_first":  p.NewestFirst,
		//"board_kind":    p.BoardKind,
		//"state":         p.State,
		//"order_by":      p.OrderBy,
		//"workspace_ids": p.WorkspaceIDs,
	}
	//
	options := makeRequestOptions{
		BaseURL:   mondayAPI,
		Query:     &query,
		Variables: variables,
	}

	err = c.api.makeRequest(options)

	if query.Board != nil {
		out = query.Board[0].Groups
	}
	return
}
