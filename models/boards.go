package models

import "time"

type Board struct {
	ActivityLogs    []ActivityLogs `json:"activity_logs"`
	BoardFolderID   any            `json:"board_folder_id"`
	BoardKind       string         `json:"board_kind"`
	Columns         []Columns      `json:"columns"`
	Communication   string         `json:"communication"` // JSON string
	Creator         User           `json:"creator"`
	Description     any            `json:"description"`
	Groups          []Group        `json:"groups"`
	ID              string         `json:"id"`
	ItemTerminology string         `json:"item_terminology"`
	ItemsCount      int            `json:"items_count"`
	Name            string         `json:"name"`
	Owners          []User         `json:"owners"`
	Permissions     string         `json:"permissions"`
	State           string         `json:"state"`
	Subscribers     []User         `json:"subscribers"`
	Tags            []Tag          `json:"tags"`
	TeamOwners      []Teams        `json:"team_owners"`
	TopGroup        *Group         `json:"top_group"`
	Type            string         `json:"type"`
	UpdatedAt       time.Time      `json:"updated_at"`
	Updates         []interface{}  `json:"updates"`
	Views           []struct {
		ID                  string `json:"id"`
		Name                string `json:"name"`
		SettingsStr         string `json:"settings_str"`
		Type                string `json:"type"`
		ViewSpecificDataStr string `json:"view_specific_data_str"`
	} `json:"views"`
	Workspace   *Workspace `json:"workspace"`
	WorkspaceID int        `json:"workspace_id"`
}

type ActivityLogs struct {
	AccountID string    `json:"account_id"`
	CreatedAt string    `json:"created_at"`
	Data      time.Time `json:"data"`
	Entity    string    `json:"entity"`
	Event     string    `json:"event"`
	ID        string    `json:"id"`
	UserID    string    `json:"user_id"`
}

type Columns struct {
	Archived    bool        `json:"archived"`
	Description interface{} `json:"description"`
	ID          string      `json:"id"`
	SettingsStr string      `json:"settings_str"`
	Title       string      `json:"title"`
	Type        string      `json:"type"`
	Width       int         `json:"width"`
}

type Group struct {
	Archived bool   `json:"archived"`
	ID       string `json:"id"`
	Color    string `json:"color"`
	Deleted  bool   `json:"deleted"`
	Position string `json:"position"`
	Title    string `json:"title"`
}

type Tag struct {
	Color string `json:"color"`
	ID    int    `json:"id"`
	Name  string `json:"name"`
}

type Update struct {
	Assets    *Assets   `json:"assets"`
	Body      string    `json:"body"`
	CreatedAt time.Time `json:"created_at"`
	Creator   *User     `json:"creator"`
	CreatorID string    `json:"creator_id"`
	ID        int       `json:"id"`
	ItemID    int       `json:"item_id"`
	Replies   Replies   `json:"replies"`
	TextBody  string    `json:"text_body"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Assets struct {
	CreatedAt        time.Time `json:"created_at"`
	FileExtension    string    `json:"file_extension"`
	FileSize         int       `json:"file_size"`
	ID               int       `json:"id"`
	Name             string    `json:"name"`
	OriginalGeometry string    `json:"original_geometry"`
	PublicURL        string    `json:"public_url"`
	UploadedBy       *User     `json:"uploaded_by"`
	URL              string    `json:"url"`
	URLThumbnail     string    `json:"url_thumbnail"`
}

type Replies struct {
	Body      string    `json:"body"`
	CreatedAt time.Time `json:"created_at"`
	Creator   *User     `json:"creator"`
	CreatorID string    `json:"creator_id"`
	ID        int       `json:"id"`
	TextBody  string    `json:"text_body"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Workspace struct {
	ID                    int               `json:"id"`
	Name                  string            `json:"name"`
	AccountProduct        *AccountProducts  `json:"account_product"`
	CreateAt              time.Time         `json:"create_at"`
	Description           string            `json:"description"`
	Kind                  string            `json:"kind"`
	OwnersSubscribers     *User             `json:"owners_subscribers"`
	Settings              WorkspaceSettings `json:"settings"`
	State                 string            `json:"state"`
	TeamOwnersSubscribers []Teams           `json:"team_owners_subscribers"`
	UsersSubscribers      []User            `json:"users_subscribers"`
}

type WorkspaceSettings struct {
	Icon IconSettings `json:"icon"`
}

type IconSettings struct {
	Color string `json:"color"`
	Image string `json:"image"`
}
