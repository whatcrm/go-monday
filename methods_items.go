package monday

import (
	"github.com/whatcrm/go-monday/models"
	"log"
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

type ItemsResponse struct {
	Cursor string        `graphql:"cursor"`
	Items  []models.Item `graphql:"items"`
}

type ItemsPageByColumnValuesQuery struct {
	ColumnID     string   `json:"column_id"`
	ColumnValues []string `json:"column_values"`
}

func (c *Get) ItemsPageByColumnValues(boardID ID, columns []ItemsPageByColumnValuesQuery) (out []models.Item, err error) {

	var query struct {
		ItemResponse ItemsResponse `graphql:"items_page_by_column_values ( board_id: $board_id columns: $columns ) "`
	}

	log.Printf("%+v\n", columns)

	var variables = map[string]interface{}{
		"board_id": boardID,
		"columns":  columns, // "lead_phone", [\"79123456789\"]
	}
	log.Println(variables)

	options := makeRequestOptions{
		BaseURL:   mondayAPI,
		Query:     &query,
		Variables: variables,
	}

	err = c.api.makeRequest(options)
	out = query.ItemResponse.Items
	return
}

func (c *Mutate) CreateItem(boardID ID, name, groupID string, clim bool) (out models.Item, err error) {

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

	err = c.api.makeRequest(options)
	out = mutation.CreateItem
	return
}

type ID string

func (c *Mutate) CreateItemWithColumnValues(boardID ID, name, groupID string, columnValues JSON, clim bool) (out models.Item, err error) {

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

	err = c.api.makeRequest(options)
	out = mutation.Item
	return
}

func (c *Get) Update(ids []string) (out []models.UpdateReplies, err error) {

	var query struct {
		Update []models.UpdateReplies `graphql:"updates ( ids: $id ) "`
	}

	var mIDs []ID
	for i := range ids {
		mIDs = append(mIDs, ID(ids[i]))
	}

	var variables = map[string]interface{}{
		"id": mIDs,
	}
	log.Println(variables)

	options := makeRequestOptions{
		BaseURL:   mondayAPI,
		Query:     &query,
		Variables: variables,
	}

	err = c.api.makeRequest(options)
	out = query.Update
	return
}

func (c *Mutate) CreateUpdate(itemID, parentID ID, body string) (out models.Update, err error) {

	var mutation struct {
		Update models.Update `graphql:"create_update (item_id: $item_id body: $body parent_id: $parent_id)"`
	}

	variables := map[string]interface{}{
		"item_id":   itemID,
		"body":      body,
		"parent_id": parentID,
	}

	options := makeRequestOptions{
		BaseURL:   mondayAPI,
		Mutation:  &mutation,
		Variables: variables,
	}

	err = c.api.makeRequest(options)
	out = mutation.Update
	return
}

func (c *Mutate) DeleteUpdate(updateID ID) (err error) {

	var mutation struct {
		Update models.Update `graphql:"delete_update (id: $id)"`
	}

	variables := map[string]interface{}{
		"id": updateID,
	}

	options := makeRequestOptions{
		BaseURL:   mondayAPI,
		Mutation:  &mutation,
		Variables: variables,
	}

	err = c.api.makeRequest(options)
	return
}

func (c *Mutate) LikeUpdate(updateID ID) (out models.Update, err error) {

	var mutation struct {
		Update models.Update `graphql:"like_update ( update_id: $update_id )"`
	}

	variables := map[string]interface{}{
		"update_id": updateID,
	}

	options := makeRequestOptions{
		BaseURL:   mondayAPI,
		Mutation:  &mutation,
		Variables: variables,
	}

	err = c.api.makeRequest(options)
	out = mutation.Update
	return
}

func (c *Mutate) CreateOrGetTag(boardID ID, tagName string) (out models.Tag, err error) {
	var mutation struct {
		Tag models.Tag `graphql:"create_or_get_tag (tag_name: $tag_name board_id: $board_id)"`
	}
	variables := map[string]interface{}{
		"board_id": boardID,
		"tag_name": tagName,
	}

	options := makeRequestOptions{
		BaseURL:   mondayAPI,
		Mutation:  &mutation,
		Variables: variables,
	}

	err = c.api.makeRequest(options)
	out = mutation.Tag
	return
}

func (c *Mutate) ChangeColumnValue(boardID, itemID, columnID string, columnValue JSON) (out models.Nested, err error) {
	var mutation struct {
		Nested models.Nested `graphql:"change_column_value (board_id: $board_id item_id: $item_id column_id: $column_id value: $value)"`
	}

	variables := map[string]interface{}{
		"board_id":  ID(boardID),
		"item_id":   ID(itemID),
		"column_id": columnID,
		"value":     columnValue,
	}

	options := makeRequestOptions{
		BaseURL:   mondayAPI,
		Mutation:  &mutation,
		Variables: variables,
	}

	err = c.api.makeRequest(options)
	out = mutation.Nested
	return
}

func (c *Mutate) ChangeMultipleColumnValue(boardID, itemID, columnValues JSON) (out models.Nested, err error) {
	var mutation struct {
		Nested models.Nested `graphql:"change_multiple_column_values (board_id: $board_id item_id: $item_id column_values: $column_values)"`
	}

	variables := map[string]interface{}{
		"board_id":      ID(boardID),
		"item_id":       ID(itemID),
		"column_values": columnValues,
	}

	options := makeRequestOptions{
		BaseURL:   mondayAPI,
		Mutation:  &mutation,
		Variables: variables,
	}

	err = c.api.makeRequest(options)
	out = mutation.Nested
	return
}
