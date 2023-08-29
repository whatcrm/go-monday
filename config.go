package monday

const (
	scheme = "https"
	oAuth  = "https://auth.monday.com/oauth2/authorize"
)

// Methods

type callMethodOptions struct {
	// Method is a request's method
	Method string
	// BaseURL is a url from constants above.
	BaseURL string
	// In is a struct, which will be marshalled to Request Body
	In interface{}
	// Out is a struct, which will be unmarshalled
	Out interface{}
	// Params is a URL Parameters
	Params *RequestParams
}
