package monday

import (
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"log"
	"net/url"
	"strings"
)

func (api *API) getAgent(method, baseURL string, params *RequestParams) (*fiber.Agent, *fiber.Request) {
	a := fiber.AcquireAgent()
	req := a.Request()
	req.Header.SetMethod(method)

	req.Header.SetContentType(fiber.MIMEApplicationJSON)
	req.Header.SetCanonical([]byte("Authorization"), []byte("Bearer "+api.Auth))

	a.MaxRedirectsCount(1)

	req.SetRequestURI(api.buildURL(baseURL, params))
	if params != nil {
		req.SetRequestURI(api.buildURL(baseURL, params))
	}
	return a, req
}

func (api *API) callMethod(options callMethodOptions) (err error) {

	a, req := api.getAgent(options.Method, options.BaseURL, options.Params)

	if options.In != nil {
		api.log("marshaling the data...")
		req, err = marshal(options.In, req)
		if err != nil {
			return
		}
	}

	api.log("sending the data...")
	if err = a.Parse(); err != nil {
		log.Println(err)
		return
	}

	api.log("getting the answer...")
	status, body, errs := a.Bytes()
	if errs != nil {
		log.Println("Errs: ", errs)
		err = errs[0]
		return
	}

	if err = errorCheck(body, status); err != nil {
		return
	}
	api.log("errorCheck passed")

	if err = json.Unmarshal(body, options.Out); err != nil {
		return fiber.NewError(400, string(body))
	}
	api.log("unmarshal passed")

	err = statusChecker(status)
	return
}

func statusChecker(status int) error {
	switch status {
	case 400:
		return fiber.ErrBadRequest
	case 401:
		return fiber.ErrUnauthorized
	case 402:
		return fiber.ErrPaymentRequired
	case 403:
		return fiber.ErrForbidden
	case 404:
		return fiber.ErrNotFound
	case 201:
		return fiber.NewError(201, "Created")
	case 204:
		return fiber.NewError(204, "No content")
	case 200, 202, 302, 301:
		return nil
	default:
		return fiber.NewError(status, "unknown status")
	}
}

func marshal(data interface{}, req *fiber.Request) (*fiber.Request, error) {
	m, err := json.Marshal(&data)
	if err != nil {
		return req, err
	}

	req.SetBody(m)
	return req, nil
}

func (api *API) buildURL(baseURL string, params *RequestParams) string {
	api.fixDomain()

	u, err := url.Parse(baseURL)
	if err != nil {
		return "can't parse domain, due to " + err.Error()
	}

	u.Scheme = scheme
	query := u.Query()

	if baseURL == oAuth {
		query.Set("client_id", api.ClientID)
		query.Set("redirect_uri", params.RedirectURI)
		query.Set("scope", params.Scope)
		query.Set("state", params.State)
		query.Set("app_version_id", params.AppVersionID)
	}

	//query.Set(Auth, api.Auth)

	u.RawQuery = query.Encode()
	log.Println(u.String())
	return u.String()
}

func (api *API) log(message ...interface{}) {
	if api.Debug {
		log.Println(message...)
	}
}

func errorCheck(body []byte, status int) error {
	// TODO handle body on most used errors
	return nil
}

func (api *API) fixDomain() {
	api.Domain = strings.Trim(api.Domain, " ")
	api.Domain = strings.Trim(api.Domain, "/")
	api.Domain = strings.Trim(api.Domain, ".")
	if strings.Contains(api.Domain, "https://") || strings.Contains(api.Domain, "http://") {
		uri, _ := url.Parse(api.Domain)
		api.Domain = uri.Hostname()
	}
	api.Domain = strings.ReplaceAll(api.Domain, "//", "/")
}
