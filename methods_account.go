package monday

import (
	"github.com/whatcrm/go-monday/models"
)

func (c *Get) Me() (out models.User, err error) {
	var query struct {
		Me models.User `graphql:"me"`
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
		Account models.Account `graphql:"account"`
	}

	options := makeRequestOptions{
		BaseURL: mondayAPI,
		Query:   &query,
	}

	err = c.api.makeRequest(options)
	out = query.Account
	return
}

type Webhooks struct {
	BoardID string `graphql:"board_id"`
	Config  string `graphql:"config"`
	Event   string `graphql:"event"`
	ID      string `graphql:"id"`
}
type WebhookEventType string

func (c *Get) Webhooks(boardID ID) (out []Webhooks, err error) {
	var query struct {
		Webhooks []Webhooks `graphql:"webhooks ( board_id: $board_id ) "`
	}

	variables := map[string]interface{}{
		"board_id": boardID,
	}

	options := makeRequestOptions{
		BaseURL:   mondayAPI,
		Query:     &query,
		Variables: variables,
	}

	err = c.api.makeRequest(options)
	out = query.Webhooks
	return
}

func (c *Mutate) CreateWebhook(boardID ID, uri string, event WebhookEventType) (err error) {
	var mutation struct {
		Webhook Webhooks `graphql:"create_webhook (board_id: $board_id url: $url event: $event)"`
	}

	variables := map[string]interface{}{
		"board_id": boardID,
		"url":      uri,
		"event":    event,
	}

	options := makeRequestOptions{
		BaseURL:   mondayAPI,
		Mutation:  &mutation,
		Variables: variables,
	}

	err = c.api.makeRequest(options)
	return
}

func (c *Mutate) DeleteWebhook(id ID) (err error) {
	var mutation struct {
		Webhook Webhooks `graphql:"delete_webhook (id: $id)"`
	}

	variables := map[string]interface{}{
		"id": id,
	}

	options := makeRequestOptions{
		BaseURL:   mondayAPI,
		Mutation:  mutation,
		Variables: variables,
	}

	err = c.api.makeRequest(options)
	return
}
