package models

type Account struct {
	ID                   string            `json:"id,omitempty"`
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

type AccountQuery struct {
	Params               *QueryParams `json:"query_params,omitempty"`
	ID                   bool         `json:"id,omitempty"`
	CountryCode          bool         `json:"country_code,omitempty"`
	FirstDayOfTheWeek    bool         `json:"first_day_of_the_week,omitempty"`
	Logo                 bool         `json:"logo,omitempty"`
	Name                 bool         `json:"name,omitempty"`
	Plan                 *PlanQuery   `json:"plan,omitempty"`
	Products             bool         `json:"products,omitempty"`
	ShowTimelineWeekends bool         `json:"show_timeline_weekends,omitempty"`
	SignUpProductKind    bool         `json:"sign_up_product_kind,omitempty"`
	Slug                 bool         `json:"slug,omitempty"`
	Tier                 bool         `json:"tier,omitempty"`
}

type Plan struct {
	MaxUsers int    `json:"max_users,omitempty"`
	Period   string `json:"period,omitempty"`
	Tier     string `json:"tier,omitempty"`
	Version  int    `json:"version,omitempty"`
}

type PlanQuery struct {
	MaxUsers bool `json:"max_users,omitempty"`
	Period   bool `json:"period,omitempty"`
	Tier     bool `json:"tier,omitempty"`
	Version  bool `json:"version,omitempty"`
}

type AccountProducts struct {
	ID   string `json:"id,omitempty"`
	Kind string `json:"kind,omitempty"`
}
