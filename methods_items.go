package monday

import (
	"github.com/whatcrm/go-monday/models"
)

func (c *Get) Items(ids []ID, newestFirst, excludeNonActive bool, limit, page int) (out []models.Item, err error) {
	var query struct {
		Items []models.Item `graphql:"items ( limit: $limit page: $page ids: $ids newest_first: $newest_first exclude_nonactive: $exclude_nonactive ) "`
	}

	if limit < 25 {
		limit = 25
	}

	if page < 1 {
		page = 1
	}
	variables := map[string]interface{}{
		"limit":             25,
		"page":              1,
		"ids":               ids,
		"newest_first":      newestFirst,
		"exclude_nonactive": excludeNonActive,
	}

	options := makeRequestOptions{
		BaseURL:   mondayAPI,
		Query:     &query,
		Variables: variables,
	}

	err = c.api.makeRequest(options)
	out = query.Items
	return
}

func (m *Mutate) CreateItem(boardID ID, name, groupID string, clim bool) (out models.Item, err error) {

	var mutation struct {
		CreateItem models.Item `graphql:"create_item ( item_name: $item_name board_id: $board_id group_id: $group_id create_labels_if_missing: $clim )"`
	}

	variables := map[string]interface{}{
		"item_name": name,
		"board_id":  boardID,
		"group_id":  groupID,
		"clim":      clim,
	}

	options := makeRequestOptions{
		BaseURL:   mondayAPI,
		Mutation:  &mutation,
		Variables: variables,
	}

	err = m.api.makeRequest(options)
	out = mutation.CreateItem
	return
}

type ID string

func (m *Mutate) CreateItemWithColumnValues(boardID ID, name, groupID string, columnValues JSON, clim bool) (out models.Item, err error) {

	var mutation struct {
		Item models.Item `graphql:"create_item ( item_name: $item_name board_id: $board_id group_id: $group_id column_values: $column_values create_labels_if_missing: $clim )"`
	}
	variables := map[string]interface{}{
		"item_name":     name,
		"board_id":      boardID,
		"group_id":      groupID,
		"column_values": columnValues,
		"clim":          clim, // create_labels_if_missing
	}

	options := makeRequestOptions{
		BaseURL:   mondayAPI,
		Mutation:  &mutation,
		Variables: variables,
	}

	err = m.api.makeRequest(options)
	out = mutation.Item
	return
}

func (m *Mutate) CreateUpdate(itemID ID, body string) (out models.Update, err error) {

	var mutation struct {
		Update models.Update `graphql:"create_update (item_id: $item_id body: $body)"`
	}

	variables := map[string]interface{}{
		"item_id": itemID,
		"body":    body,
	}

	options := makeRequestOptions{
		BaseURL:   mondayAPI,
		Mutation:  &mutation,
		Variables: variables,
	}

	err = m.api.makeRequest(options)
	out = mutation.Update
	return
}
