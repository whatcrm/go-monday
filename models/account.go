package models

type Account struct {
	ID                   int               `json:"id,omitempty"`
	CountryCode          string            `json:"country_code,omitempty"`
	FirstDayOfTheWeek    string            `json:"first_day_of_the_week,omitempty"`
	Logo                 string            `json:"logo,omitempty"`
	Name                 string            `json:"name,omitempty"`
	Plan                 *Plan             `json:"plan,omitempty"`
	Products             []AccountProducts `json:"products,omitempty"`
	ShowTimelineWeekends bool              `json:"show_timeline_weekends,omitempty"`
	SignUpProductKind    string            `json:"sign_up_product_kind,omitempty"`
	Slug                 string            `json:"slug,omitempty"`
	Tier                 string            `json:"tier,omitempty"`
}

type Plan struct {
	MaxUsers int    `json:"max_users,omitempty"`
	Period   string `json:"period,omitempty"`
	Tier     string `json:"tier,omitempty"`
	Version  int    `json:"version,omitempty"`
}

type AccountProducts struct {
	ID   int    `json:"id,omitempty"`
	Kind string `json:"kind,omitempty"`
}
