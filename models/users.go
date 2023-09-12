package models

import (
	"time"
)

type User struct {
	ID                 int          `json:"id,omitempty"`
	Account            *Account     `json:"account,omitempty"`
	Birthday           time.Time    `json:"birthday,omitempty"`
	CountryCode        string       `json:"country_code,omitempty"`
	CreatedAt          time.Time    `json:"created_at,omitempty"`
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
	JoinDate           time.Time    `json:"join_date,omitempty"`
	LastActivity       time.Time    `json:"last_activity,omitempty"`
	Location           string       `json:"location,omitempty"`
	MobilePhone        string       `json:"mobile_phone,omitempty"`
	OutOfOffice        *OutOfOffice `json:"out_of_office,omitempty"`
	PhotoOriginal      string       `json:"photo_original,omitempty"`
	PhotoSmall         string       `json:"photo_small,omitempty"`
	PhotoThumb         string       `json:"photo_thumb,omitempty"`
	PhotoThumbSmall    string       `json:"photo_thumb_small,omitempty"`
	PhotoTiny          string       `json:"photo_tiny,omitempty"`
	SignUpProductKind  string       `json:"sign_up_product_kind,omitempty"`
	Teams              []Teams      `json:"teams,omitempty"`
	TimeZoneIdentifier string       `json:"time_zone_identifier,omitempty"`
	Title              string       `json:"title,omitempty"`
	URL                string       `json:"url,omitempty"`
	UtcHoursDiff       int          `json:"utc_hours_diff,omitempty"`
}

type OutOfOffice struct {
	Active               bool      `json:"active,omitempty"`
	DisableNotifications bool      `json:"disable_notifications,omitempty"`
	EndDate              time.Time `json:"end_date,omitempty"`
	StartDate            time.Time `json:"start_date,omitempty"`
	Type                 string    `json:"type,omitempty"`
}

type Teams struct {
	ID         int    `json:"id,omitempty"`
	Name       string `json:"name,omitempty"`
	PictureURL string `json:"picture_url,omitempty"`
	Owners     []User `json:"owners,omitempty"`
	Users      []User `json:"users,omitempty"`
}
