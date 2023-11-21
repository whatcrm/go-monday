package models

import "time"

type Board struct {
	//ActivityLogs    []ActivityLogs `graphql:"activity_logs"`
	BoardFolderID any    `graphql:"board_folder_id"`
	BoardKind     string `graphql:"board_kind"`
	Columns       []Columns
	Communication string // JSON string
	//Creator         Nested
	Description     string
	Groups          []NestedGroup
	ID              string
	ItemTerminology string `graphql:"item_terminology"`
	ItemsCount      int    `graphql:"items_count"`
	Name            string
	//Owners          []Nested
	Permissions string
	State       string
	//Subscribers     []Nested
	Tags []Tag
	//TeamOwners      []Teams `graphql:"team_owners"`
	//TopGroup        Group   `graphql:"top_group"`
	Type      string
	UpdatedAt time.Time `graphql:"updated_at"`
	//Updates         []Update  `graphql:"updates"`
	//Views           []BoardView
	//Workspace       Workspace
	WorkspaceID string `graphql:"workspace_id"`
}

type ActivityLogs struct {
	AccountID string `graphql:"account_id"`
	CreatedAt string `graphql:"created_at"`
	Data      string `graphql:"data"`
	Entity    string `graphql:"entity"`
	Event     string `graphql:"event"`
	ID        string `graphql:"id"`
	UserID    string `graphql:"user_id"`
}

type Columns struct {
	Archived    bool        `graphql:"archived"`
	Description interface{} `graphql:"description"`
	ID          string      `graphql:"id"`
	SettingsStr string      `graphql:"settings_str"`
	Title       string      `graphql:"title"`
	Type        string      `graphql:"type"`
	Width       int         `graphql:"width"`
}

type Group struct {
	Archived bool   `graphql:"archived"`
	ID       string `graphql:"id"`
	Color    string `graphql:"color"`
	Deleted  bool   `graphql:"deleted"`
	Position string `graphql:"position"`
	Title    string `graphql:"title"`
}

type Tag struct {
	Color string `graphql:"color"`
	ID    string `graphql:"id"`
	Name  string `graphql:"name"`
}

type Update struct {
	Assets    []Nested  `graphql:"assets"`
	Body      string    `graphql:"body"`
	CreatedAt time.Time `graphql:"created_at"`
	Creator   Nested    `graphql:"creator"`
	CreatorID string    `graphql:"creator_id"`
	ID        string    `graphql:"id"`
	ItemID    string    `graphql:"item_id"`
	Replies   []Replies `graphql:"replies"`
	TextBody  string    `graphql:"text_body"`
	UpdatedAt time.Time `graphql:"updated_at"`
}

type BoardView struct {
	ID                  string `graphql:"id"`
	Name                string `graphql:"name"`
	SettingsStr         string `graphql:"settings_str"`
	Type                string `graphql:"type"`
	ViewSpecificDataStr string `graphql:"view_specific_data_str"`
}

type Assets struct {
	CreatedAt        time.Time `graphql:"created_at"`
	FileExtension    string    `graphql:"file_extension"`
	FileSize         int       `graphql:"file_size"`
	ID               string    `graphql:"id"`
	Name             string    `graphql:"name"`
	OriginalGeometry string    `graphql:"original_geometry"`
	PublicURL        string    `graphql:"public_url"`
	UploadedBy       Nested    `graphql:"uploaded_by"`
	URL              string    `graphql:"url"`
	URLThumbnail     string    `graphql:"url_thumbnail"`
}

type Replies struct {
	Body      string    `graphql:"body"`
	CreatedAt time.Time `graphql:"created_at"`
	Creator   *Nested   `graphql:"creator"`
	CreatorID string    `graphql:"creator_id"`
	ID        string    `graphql:"id"`
	TextBody  string    `graphql:"text_body"`
	UpdatedAt time.Time `graphql:"updated_at"`
}

type Workspace struct {
	ID                    string            `graphql:"id"`
	Name                  string            `graphql:"name"`
	AccountProduct        AccountProducts   `graphql:"account_product"`
	CreatedAt             time.Time         `graphql:"created_at"`
	Description           string            `graphql:"description"`
	Kind                  string            `graphql:"kind"`
	OwnersSubscribers     []Nested          `graphql:"owners_subscribers"`
	Settings              WorkspaceSettings `graphql:"settings"`
	State                 string            `graphql:"state"`
	TeamOwnersSubscribers []Teams           `graphql:"team_owners_subscribers"`
	UsersSubscribers      []Nested          `graphql:"users_subscribers"`
}

type WorkspaceSettings struct {
	Icon IconSettings `graphql:"icon"`
}

type IconSettings struct {
	Color string `graphql:"color"`
	Image string `graphql:"image"`
}
