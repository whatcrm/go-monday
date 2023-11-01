package models

type Account struct {
	ID                   string
	CountryCode          string `graphql:"country_code"`
	FirstDayOfTheWeek    string `graphql:"first_day_of_the_week"`
	Logo                 string
	Name                 string
	Plan                 Plan              `graphql:"plan"`
	Products             []AccountProducts `graphql:"products"`
	ShowTimelineWeekends bool              `graphql:"show_timeline_weekends"`
	SignUpProductKind    string            `graphql:"sign_up_product_kind"`
	Slug                 string
	Tier                 string
}

type Plan struct {
	MaxUsers int    `graphql:"max_users"`
	Period   string `graphql:"period"`
	Tier     string `graphql:"tier"`
	Version  int    `graphql:"version"`
}

type AccountProducts struct {
	ID   string `graphql:"id"`
	Kind string `graphql:"kind"`
}

type Nested struct {
	ID   string `graphql:"id"`
	Name string `graphql:"name"`
}

type NestedInt struct {
	ID   int
	Name string
}

type NestedGroup struct {
	ID      string `graphql:"id"`
	Deleted bool   `graphql:"deleted"`
	Title   string `graphql:"title"`
}
