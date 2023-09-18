package models

import "time"

type QueryParams struct {
	General      *GeneralParams      `json:"general,omitempty"`
	Users        *UsersParams        `json:"users,omitempty"`
	Boards       *BoardsParams       `json:"boards,omitempty"`
	Workspaces   *WorkspacesParams   `json:"workspaces,omitempty"`
	ActivityLogs *ActivityLogsParams `json:"activity_logs,omitempty"`
	Columns      *ColumnsParams      `json:"columns,omitempty"`
	ColumnValues *ColumnValuesParams `json:"column_values,omitempty"`
	Views        *ViewsParams        `json:"views,omitempty"`
}

type GeneralParams struct {
	IDs         []int `json:"ids,omitempty"`
	Limit       int   `json:"limit,omitempty"`
	Page        int   `json:"page,omitempty"`
	NewestFirst bool  `json:"newest_first,omitempty"`
}

type UsersParams struct {
	Kind                 []string `json:"kind,omitempty"` // all / non_guests / guests / non_pending
	Emails               []string `json:"emails,omitempty"`
	Name                 []string `json:"name,omitempty"`
	NonActive            bool     `json:"non_active,omitempty"` // non_active: true
	RegistrationSequence int      `json:"registration_sequence,omitempty"`
}

type BoardsParams struct {
	OrderBy      string `json:"order_by,omitempty"`   // created_at / used_at
	BoardKind    string `json:"board_kind,omitempty"` // public / private / share
	WorkspaceIds []int  `json:"workspace_ids,omitempty"`
	State        string `json:"state,omitempty"` // all / active / archived / deleted
}

type WorkspacesParams struct {
	Kind    string `json:"kind,omitempty"`     // open / closed
	State   string `json:"state,omitempty"`    // all / active / archived / deleted
	OrderBy string `json:"order_by,omitempty"` // crated_at
}

type ActivityLogsParams struct {
	UserIDs   []int      `json:"user_ids,omitempty"`
	ColumnIDs []int      `json:"column_ids,omitempty"`
	GroupIDs  []int      `json:"group_ids,omitempty"`
	ItemIDs   []int      `json:"item_ids,omitempty"`
	From      *time.Time `json:"from,omitempty"`
	To        *time.Time `json:"to,omitempty"`
}

type ColumnsParams struct {
	Types []string `json:"types,omitempty"` // all types
}

type ColumnValuesParams struct {
	IDs []string `json:"ids,omitempty"`
}

type ViewsParams struct {
	Type string `json:"type,omitempty"`
}
