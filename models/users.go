package models

import (
	"time"
)

type Me struct {
	ID                 int          `json:"id,omitempty"`
	Account            *Account     `json:"account,omitempty"`
	Birthday           string       `json:"birthday,omitempty"`
	CountryCode        string       `json:"country_code,omitempty"`
	CreatedAt          *time.Time   `json:"created_at,omitempty"`
	CurrentLanguage    string       `json:"current_language,omitempty"`
	Enabled            bool         `json:"enabled,omitempty"`
	IsAdmin            bool         `json:"is_admin,omitempty"`
	IsGuest            bool         `json:"is_guest,omitempty"`
	IsVerified         bool         `json:"is_verified,omitempty"`
	IsViewOnly         bool         `json:"is_view_only,omitempty"`
	IsPending          bool         `json:"is_pending,omitempty"`
	Name               string       `json:"name,omitempty"`
	Phone              string       `json:"phone,omitempty"`
	Email              string       `json:"email,omitempty"`
	JoinDate           any          `json:"join_date,omitempty"`
	LastActivity       *time.Time   `json:"last_activity,omitempty"`
	Location           any          `json:"location,omitempty"`
	MobilePhone        any          `json:"mobile_phone,omitempty"`
	OutOfOffice        *OutOfOffice `json:"out_of_office,omitempty"`
	PhotoOriginal      string       `json:"photo_original,omitempty"`
	PhotoSmall         string       `json:"photo_small,omitempty"`
	PhotoThumb         string       `json:"photo_thumb,omitempty"`
	PhotoThumbSmall    string       `json:"photo_thumb_small,omitempty"`
	PhotoTiny          string       `json:"photo_tiny,omitempty"`
	SignUpProductKind  any          `json:"sign_up_product_kind,omitempty"`
	Teams              []Teams      `json:"teams,omitempty"`
	TimeZoneIdentifier string       `json:"time_zone_identifier,omitempty"`
	Title              string       `json:"title,omitempty"`
	URL                string       `json:"url,omitempty"`
	UtcHoursDiff       int          `json:"utc_hours_diff,omitempty"`
}

type User struct {
	ID                 string       `json:"id,omitempty"`
	Account            *Account     `json:"account,omitempty"`
	Birthday           string       `json:"birthday,omitempty"`
	CountryCode        string       `json:"country_code,omitempty"`
	CreatedAt          *time.Time   `json:"created_at,omitempty"`
	CurrentLanguage    string       `json:"current_language,omitempty"`
	Enabled            bool         `json:"enabled,omitempty"`
	IsAdmin            bool         `json:"is_admin,omitempty"`
	IsGuest            bool         `json:"is_guest,omitempty"`
	IsVerified         bool         `json:"is_verified,omitempty"`
	IsViewOnly         bool         `json:"is_view_only,omitempty"`
	IsPending          bool         `json:"is_pending,omitempty"`
	Name               string       `json:"name,omitempty"`
	Phone              string       `json:"phone,omitempty"`
	Email              string       `json:"email,omitempty"`
	JoinDate           any          `json:"join_date,omitempty"`
	LastActivity       *time.Time   `json:"last_activity,omitempty"`
	Location           any          `json:"location,omitempty"`
	MobilePhone        any          `json:"mobile_phone,omitempty"`
	OutOfOffice        *OutOfOffice `json:"out_of_office,omitempty"`
	PhotoOriginal      string       `json:"photo_original,omitempty"`
	PhotoSmall         string       `json:"photo_small,omitempty"`
	PhotoThumb         string       `json:"photo_thumb,omitempty"`
	PhotoThumbSmall    string       `json:"photo_thumb_small,omitempty"`
	PhotoTiny          string       `json:"photo_tiny,omitempty"`
	SignUpProductKind  any          `json:"sign_up_product_kind,omitempty"`
	Teams              []Teams      `json:"teams,omitempty"`
	TimeZoneIdentifier string       `json:"time_zone_identifier,omitempty"`
	Title              string       `json:"title,omitempty"`
	URL                string       `json:"url,omitempty"`
	UtcHoursDiff       int          `json:"utc_hours_diff,omitempty"`
}

type UserQuery struct {
	Params             *QueryParams      `json:"query_params,omitempty"`
	ID                 bool              `json:"id,omitempty"`
	Account            *AccountQuery     `json:"account,omitempty"`
	Birthday           bool              `json:"birthday,omitempty"`
	CountryCode        bool              `json:"country_code,omitempty"`
	CreatedAt          bool              `json:"created_at,omitempty"`
	CurrentLanguage    bool              `json:"current_language,omitempty"`
	Enabled            bool              `json:"enabled,omitempty"`
	IsAdmin            bool              `json:"is_admin,omitempty"`
	IsGuest            bool              `json:"is_guest,omitempty"`
	IsVerified         bool              `json:"is_verified,omitempty"`
	IsViewOnly         bool              `json:"is_view_only,omitempty"`
	IsPending          bool              `json:"is_pending,omitempty"`
	Name               bool              `json:"name,omitempty"`
	Phone              bool              `json:"phone,omitempty"`
	Email              bool              `json:"email,omitempty"`
	JoinDate           bool              `json:"join_date,omitempty"`
	LastActivity       bool              `json:"last_activity,omitempty"`
	Location           bool              `json:"location,omitempty"`
	MobilePhone        bool              `json:"mobile_phone,omitempty"`
	OutOfOffice        *OutOfOfficeQuery `json:"out_of_office,omitempty"`
	PhotoOriginal      bool              `json:"photo_original,omitempty"`
	PhotoSmall         bool              `json:"photo_small,omitempty"`
	PhotoThumb         bool              `json:"photo_thumb,omitempty"`
	PhotoThumbSmall    bool              `json:"photo_thumb_small,omitempty"`
	PhotoTiny          bool              `json:"photo_tiny,omitempty"`
	SignUpProductKind  bool              `json:"sign_up_product_kind,omitempty"`
	Teams              *TeamsQuery       `json:"teams,omitempty"`
	TimeZoneIdentifier bool              `json:"time_zone_identifier,omitempty"`
	Title              bool              `json:"title,omitempty"`
	URL                bool              `json:"url,omitempty"`
	UtcHoursDiff       bool              `json:"utc_hours_diff,omitempty"`
}

type OutOfOffice struct {
	Active               bool      `json:"active,omitempty"`
	DisableNotifications bool      `json:"disable_notifications,omitempty"`
	EndDate              time.Time `json:"end_date,omitempty"`
	StartDate            time.Time `json:"start_date,omitempty"`
	Type                 string    `json:"type,omitempty"`
}

type OutOfOfficeQuery struct {
	Active               bool `json:"active,omitempty"`
	DisableNotifications bool `json:"disable_notifications,omitempty"`
	EndDate              bool `json:"end_date,omitempty"`
	StartDate            bool `json:"start_date,omitempty"`
	Type                 bool `json:"type,omitempty"`
}

type Teams struct {
	ID         int    `json:"id,omitempty"`
	Name       string `json:"name,omitempty"`
	PictureURL string `json:"picture_url,omitempty"`
	Owners     []User `json:"owners,omitempty"`
	Users      []User `json:"users,omitempty"`
}

type TeamsQuery struct {
	ID         bool       `json:"id,omitempty"`
	Name       bool       `json:"name,omitempty"`
	PictureURL bool       `json:"picture_url,omitempty"`
	Owners     *UserQuery `json:"owners,omitempty"`
	Users      *UserQuery `json:"users,omitempty"`
}
