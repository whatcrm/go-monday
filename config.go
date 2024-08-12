package monday

import "github.com/hasura/go-graphql-client"

const (
	apiVersion  = "API-Version"
	October2023 = "2024-01"

	mondayAPI    = "https://api.monday.com/v2"
	oAuth        = "https://auth.monday.com/oauth2/authorize"
	tokenRequest = "https://auth.monday.com/oauth2/token"
)

// JSON example: "{\"lead_phone\":\"+1111 RU\"}"
type JSON string

// BoardKind example: public/private/share
type BoardKind string

// BoardsOrderBy example: created_at / used_at
type BoardsOrderBy string

// State example: all / active / archived / deleted the default is active.
type State string

type makeRequestOptions struct {
	// BaseURL is an url from constants above.
	BaseURL string
	// Query is a struct that works as Get request
	Query interface{}
	// Mutation is a struct that works as Post/Put/Delete request
	Mutation interface{}
	// Variables is a map with all variables for queries/mutations
	Variables map[string]interface{}
	// Options for the client
	Options graphql.Option
	// Params
}

type logOptions struct {
	Domain    string                 `json:"domain"`
	Request   interface{}            `json:"request"`
	Variables map[string]interface{} `json:"variables"`
	Response  interface{}            `json:"response"`
	Error     string                 `json:"error,omitempty"`
}
