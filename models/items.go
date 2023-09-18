package models

import "time"

type Item struct {
	Assets       *Assets       `json:"assets"`
	Board        *Board        `json:"board"`
	Name         string        `json:"name"`
	ColumnValues []ColumnValue `json:"column_values"`
	CreatedAt    *time.Time    `json:"created_at"`
	Creator      *User         `json:"creator"`
	Email        string        `json:"email"`
	Group        *Group        `json:"group"`
	ID           int           `json:"id"`
	ParentItem   *Item         `json:"parent_item"`
	RelativeLink string        `json:"relative_link"`
	State        string        `json:"state"`
	SubItems     []Item        `json:"subitems"`
	Subscribers  []User        `json:"subscribers"`
	UpdatedAt    *time.Time    `json:"updated_at"`
	Updates      *Update       `json:"updates"`
}

type ItemQuery struct {
	QueryParams  *QueryParams      `json:"query_params"`
	Assets       bool              `json:"assets"`
	Board        bool              `json:"board"`
	Name         bool              `json:"name"`
	ColumnValues *ColumnValueQuery `json:"column_values"`
	CreatedAt    bool              `json:"created_at"`
	Creator      *UserQuery        `json:"creator"`
	Email        bool              `json:"email"`
	Group        *Group            `json:"group"`
	ID           bool              `json:"id"`
	ParentItem   *ItemQuery        `json:"parent_item"`
	RelativeLink bool              `json:"relative_link"`
	State        bool              `json:"state"`
	SubItems     *ItemQuery        `json:"subitems"`
	Subscribers  *UserQuery        `json:"subscribers"`
	UpdatedAt    bool              `json:"updated_at"`
	Updates      *UpdateQuery      `json:"updates"`
}

type ColumnValue struct {
	AdditionalInfo string `json:"additional_info"` // JSON
	Description    string `json:"description"`
	ID             string `json:"id"`
	Text           string `json:"text"`
	Title          string `json:"title"`
	Type           string `json:"type"`
	Value          string `json:"value"` // JSON: phone, changed_at, countryShortName
}

type ColumnValueJSON struct {
	Phone            string    `json:"phone"`
	ChangedAt        time.Time `json:"changed_at"`
	CountryShortName string    `json:"countryShortName"`
}

type ColumnValueQuery struct {
	QueryParams    *QueryParams `json:"query_params"`
	AdditionalInfo bool         `json:"additional_info"` // JSON
	Description    bool         `json:"description"`
	ID             bool         `json:"id"`
	Text           bool         `json:"text"`
	Title          bool         `json:"title"`
	Type           bool         `json:"type"`
	Value          bool         `json:"value"` // JSON: phone, changed_at, countryShortName
}
