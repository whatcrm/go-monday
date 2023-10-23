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

type Webhook struct {
	BoardID int    `graphql:"board_id"`
	Config  string `graphql:"config"`
	Event   string `graphql:"event"`
	ID      int    `graphql:"id"`
}
type WebhookEventType string

func (m *Mutate) CreateWebhook(boardID, uri string, event WebhookEventType) (err error) {
	var mutation struct {
		Webhook Webhook `graphql:"create_webhook (board_id: $board_id url: $url event: $event)"`
	}

	variables := map[string]interface{}{
		"board_id": boardID,
		"url":      uri,
		"event":    event,
	}

	options := makeRequestOptions{
		BaseURL:   mondayAPI,
		Mutation:  mutation,
		Variables: variables,
	}

	err = m.api.makeRequest(options)
	return
}

func (m *Mutate) DeleteWebhook(id ID) (err error) {
	var mutation struct {
		Webhook Webhook `graphql:"delete_webhook (id: $id)"`
	}

	variables := map[string]interface{}{
		"id": id,
	}

	options := makeRequestOptions{
		BaseURL:   mondayAPI,
		Mutation:  mutation,
		Variables: variables,
	}

	err = m.api.makeRequest(options)
	return
}
