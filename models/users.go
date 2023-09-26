package models

import (
	"time"
)

type User struct {
	ID                 string
	Name               string
	Account            Account `graphql:"account"`
	CountryCode        string  `graphql:"country_code"`
	Birthday           string
	CreatedAt          time.Time `graphql:"created_at"`
	CurrentLanguage    string    `graphql:"current_language"`
	Enabled            bool
	IsAdmin            bool `graphql:"is_admin"`
	IsGuest            bool `graphql:"is_guest"`
	IsVerified         bool `graphql:"is_verified"`
	IsViewOnly         bool `graphql:"is_view_only"`
	IsPending          bool `graphql:"is_pending"`
	Phone              string
	Email              string
	JoinDate           any       `graphql:"join_date"`
	LastActivity       time.Time `graphql:"last_activity"`
	Location           any
	MobilePhone        any         `graphql:"mobile_phone"`
	OutOfOffice        OutOfOffice `graphql:"out_of_office"`
	PhotoOriginal      string      `graphql:"photo_original"`
	PhotoSmall         string      `graphql:"photo_small"`
	PhotoThumb         string      `graphql:"photo_thumb"`
	PhotoThumbSmall    string      `graphql:"photo_thumb_small"`
	PhotoTiny          string      `graphql:"photo_tiny"`
	SignUpProductKind  any         `graphql:"sign_up_product_kind"`
	Teams              []Teams     `graphql:"teams"`
	TimeZoneIdentifier string      `graphql:"time_zone_identifier"`
	Title              string      `graphql:"title"`
	URL                string      `graphql:"url"`
	UtcHoursDiff       int         `graphql:"utc_hours_diff"`
}

type OutOfOffice struct {
	Active               bool      `graphql:"active"`
	DisableNotifications bool      `graphql:"disable_notifications"`
	EndDate              time.Time `graphql:"end_date"`
	StartDate            time.Time `graphql:"start_date"`
	Type                 string    `graphql:"type"`
}

type Notification struct {
	ID   string `graphql:"id"`
	Text string `graphql:"text"`
}

type NotificationTargetType string
