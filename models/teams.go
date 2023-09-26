package models

type Teams struct {
	ID         string   `graphql:"id"`
	Name       string   `graphql:"name"`
	PictureURL string   `graphql:"picture_url"`
	Owners     []Nested `graphql:"owners"`
	Users      []Nested `graphql:"users"`
}
