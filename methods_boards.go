package monday

import (
	"errors"
	"github.com/whatcrm/go-monday/models"
)

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

func (c *Get) Boards(page, limit int) (out []models.Board, err error) {

	var query struct {
		Board []models.Board `graphql:"boards ( page: $page limit: $limit ) "`
	}

	variables := map[string]interface{}{
		"limit": limit,
		"page":  page,
		//"ids":           p.IDs,
		//"newest_first":  p.NewestFirst,
		//"board_kind":    p.BoardKind,
		//"state":         p.State,
		//"order_by":      p.OrderBy,
		//"workspace_ids": p.WorkspaceIDs,
	}

	options := makeRequestOptions{
		BaseURL:   mondayAPI,
		Query:     &query,
		Variables: variables,
	}

	err = c.api.makeRequest(options)
	out = query.Board
	return
}

func (c *Get) Board(id ID) (out models.Board, err error) {
	var query struct {
		Board []models.Board `graphql:"boards ( ids: $ids ) "`
	}

	if id == "" {
		return models.Board{}, errors.New("id must be valid")
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

	if query.Board != nil && len(query.Board) > 0 {
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

	if query.Board != nil && len(query.Board) > 0 {
		out = query.Board[0].Groups
	}
	return
}

type ColumnType string

func (c *Mutate) CreateColumn(boardID ID, title string, columnType ColumnType) (out models.Column, err error) {
	var mutation struct {
		Column models.Column `graphql:"create_column ( board_id: $board_id title: $title column_type: $columnType ) "`
	}

	variables := map[string]interface{}{
		"board_id":   boardID,
		"title":      title,
		"columnType": columnType,
	}

	options := makeRequestOptions{
		BaseURL:   mondayAPI,
		Mutation:  &mutation,
		Variables: variables,
	}

	err = c.api.makeRequest(options)

	out = mutation.Column
	return
}

func (c *Get) BoardsByWorkspace(page, limit int, workspaceID []ID) (out []models.Board, err error) {

	var query struct {
		Board []models.Board `graphql:"boards ( page: $page limit: $limit workspace_ids: $wIDs ) "`
	}

	variables := map[string]interface{}{
		"limit": limit,
		"page":  page,
		"wIDs":  workspaceID,
		//"ids":           p.IDs,
		//"newest_first":  p.NewestFirst,
		//"board_kind":    p.BoardKind,
		//"state":         p.State,
		//"order_by":      p.OrderBy,
		//"workspace_ids": p.WorkspaceIDs,
	}

	options := makeRequestOptions{
		BaseURL:   mondayAPI,
		Query:     &query,
		Variables: variables,
	}

	err = c.api.makeRequest(options)
	out = query.Board
	return
}
