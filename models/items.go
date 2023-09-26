package models

import "time"

type Item struct {
	ID           string        `graphql:"id"`
	Name         string        `graphql:"name"`
	Email        string        `graphql:"email"`
	Creator      Nested        `graphql:"creator"`
	Board        Nested        `graphql:"board"`
	Group        NestedGroup   `graphql:"group"`
	ParentItem   Nested        `graphql:"parent_item"`
	SubItems     []Nested      `graphql:"subitems"`
	CreatedAt    time.Time     `graphql:"created_at"`
	UpdatedAt    time.Time     `graphql:"updated_at"`
	RelativeLink string        `graphql:"relative_link"`
	State        string        `graphql:"state"`
	Subscribers  []Nested      `graphql:"subscribers"`
	Updates      []Update      `graphql:"updates"`
	ColumnValues []ColumnValue `graphql:"column_values"`
	Assets       []Assets      `graphql:"assets"`
}

type NestedUpdate struct {
	ID     int
	ItemID int `graphql:"item_id"`
}

type ColumnValue struct {
	Column `graphql:"column"`
	ID     string `graphql:"id"`
	Text   string `graphql:"text"`
	Type   string `graphql:"type"`
	Value  string `graphql:"value"`
}

type Column struct {
	Archived    bool   `graphql:"archived"`
	Description any    `graphql:"description"` // TODO change to string
	ID          string `graphql:"id"`
	SettingsStr string `graphql:"settings_str"`
	Title       string `graphql:"title"`
	Type        string `graphql:"type"`
	Width       int    `graphql:"width"`
}
