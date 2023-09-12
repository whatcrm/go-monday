package monday

const (
	scheme          = "https"
	applicationJson = "application/json"
	mondayAPI       = "https://api.monday.com/v2"
	oAuth           = "https://auth.monday.com/oauth2/authorize"
	tokenRequest    = "https://auth.monday.com/oauth2/token"
)

const (
	UsersQuery = `{
       "query": "{ users  { id account { id country_code first_day_of_the_week logo name plan { max_users period tier version } products { id kind } show_timeline_weekends sign_up_product_kind slug tier } birthday country_code created_at current_language enabled is_admin is_guest is_verified is_view_only is_pending name phone email join_date last_activity location mobile_phone out_of_office { active disable_notifications end_date start_date type } photo_original photo_small photo_thumb photo_thumb_small photo_tiny sign_up_product_kind teams { id name picture_url owners { id name } users { id name } } time_zone_identifier title url utc_hours_diff } }"
}`
	UserQueryID = `{
       "query": "{ users(ids: %d ) { id account { id country_code first_day_of_the_week logo name plan { max_users period tier version } products { id kind } show_timeline_weekends sign_up_product_kind slug tier } birthday country_code created_at current_language enabled is_admin is_guest is_verified is_view_only is_pending name phone email join_date last_activity location mobile_phone out_of_office { active disable_notifications end_date start_date type } photo_original photo_small photo_thumb photo_thumb_small photo_tiny sign_up_product_kind teams { id name picture_url owners { id name } users { id name }  } time_zone_identifier title url utc_hours_diff } }"
}`
	UserQueryIDs = `{
       "query": "{ users(ids: %s ) { id account { id country_code first_day_of_the_week logo name plan { max_users period tier version } products { id kind } show_timeline_weekends sign_up_product_kind slug tier } birthday country_code created_at current_language enabled is_admin is_guest is_verified is_view_only is_pending name phone email join_date last_activity location mobile_phone out_of_office { active disable_notifications end_date start_date type } photo_original photo_small photo_thumb photo_thumb_small photo_tiny sign_up_product_kind teams { id name picture_url owners { id name } users { id name } } time_zone_identifier title url utc_hours_diff } }"
}`
)

// Methods

type callMethodOptions struct {
	// Method is a request's method
	Method string
	// BaseURL is a url from constants above.
	BaseURL string
	// In is a struct, which will be marshalled to Request Body
	In string
	// Out is a struct, which will be unmarshalled
	Out interface{}
	// Params is a URL Parameters
	Params *RequestParams
}
