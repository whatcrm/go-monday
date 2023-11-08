package monday

import "github.com/whatcrm/go-monday/models"

func (c *Get) Users() (out []models.User, err error) {
	var query struct {
		Users []models.User
	}

	options := makeRequestOptions{
		BaseURL: mondayAPI,
		Query:   &query,
	}

	err = c.api.makeRequest(options)
	out = query.Users
	return
}

func (c *Get) User(ids []ID) (out []models.User, err error) {
	var query struct {
		Users []models.User `graphql:"users ( ids: $user_ids ) "`
	}

	variables := map[string]interface{}{
		"user_ids": ids,
	}

	options := makeRequestOptions{
		BaseURL:   mondayAPI,
		Query:     &query,
		Variables: variables,
	}

	err = c.api.makeRequest(options)
	out = query.Users
	return
}

func (c *Get) Teams(ids []int) (out []models.Teams, err error) {
	var query struct {
		Teams []models.Teams `graphql:"teams ( ids: $team_ids ) "`
	}

	variables := map[string]interface{}{
		"team_ids": ids,
	}

	options := makeRequestOptions{
		BaseURL:   mondayAPI,
		Query:     &query,
		Variables: variables,
	}

	err = c.api.makeRequest(options)
	out = query.Teams
	return
}

func (m *Mutate) CreateNotification(text string, userID, targetID ID, targetType models.NotificationTargetType) (out models.Notification, err error) {
	var mutation struct {
		Notification models.Notification `graphql:"create_notification ( text: $text user_id: $user_id target_id: $target_id target_type: $target_type )"`
	}

	variables := map[string]interface{}{
		"text":        text,
		"user_id":     userID,
		"target_id":   targetID,
		"target_type": targetType,
	}

	options := makeRequestOptions{
		BaseURL:   mondayAPI,
		Mutation:  &mutation,
		Variables: variables,
	}

	err = m.api.makeRequest(options)
	out = mutation.Notification
	return
}
