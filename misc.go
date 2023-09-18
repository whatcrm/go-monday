package monday

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/valyala/fasthttp"
	"log"
	"net/url"
	"reflect"
	"strings"
)

func (api *API) call(method, baseURL string, params *RequestParams) (req *fasthttp.Request, resp *fasthttp.Response) {

	req = fasthttp.AcquireRequest()
	resp = fasthttp.AcquireResponse()

	req.Header.SetContentType(applicationJson)
	req.Header.SetMethod(method)
	req.Header.Set(fasthttp.HeaderAuthorization, "Bearer "+api.Auth)
	req.Header.Set(apiVersion, October2023)

	baseURL = api.buildURL(baseURL, params)
	//baseURL = "https://eokh00hewgqxm3a.m.pipedream.net"
	req.SetRequestURI(baseURL)
	api.log(baseURL)
	return
}

func (api *API) callMethod(options callMethodOptions) (err error) {
	client := fasthttp.Client{
		ReadBufferSize: 10000,
	}

	req, resp := api.call(options.Method, options.BaseURL, options.Params)
	defer fasthttp.ReleaseRequest(req)
	defer fasthttp.ReleaseResponse(resp)

	if options.In != nil || options.In == meQuery {
		api.log("setting the body...")
		req.SetBody([]byte(api.setQueries(options.In, "")))
		api.log("QUERY: ", string(req.Body()))
	}

	api.log("sending the data...")
	if err = client.DoRedirects(req, resp, 10); err != nil {
		fmt.Println("Error: ", err)
		return
	}

	api.log(string(resp.Body()))
	if err = errorCheck(resp.Body(), resp.StatusCode()); err != nil {
		return err
	}

	api.log("getting an answer...")
	if err = json.Unmarshal(resp.Body(), options.Out); err != nil {
		err = errors.New(string(resp.Body()))
		return err
	}

	api.log("returning the struct...")
	return
}

func (api *API) buildURL(baseURL string, params *RequestParams) string {
	api.fixDomain()

	u, err := url.Parse(baseURL)
	if err != nil {
		return "can't parse domain, due to " + err.Error()
	}

	u.Scheme = scheme

	if baseURL != tokenRequest || params == nil {
		return u.String()
	}

	query := u.Query()

	typ := reflect.TypeOf(*params)
	val := reflect.ValueOf(*params)
	for i := 0; i < typ.NumField(); i++ {
		field := typ.Field(i)
		jsonFieldName := field.Tag.Get("json")

		fieldValue := val.Field(i).Interface()
		v, ok := fieldValue.(string)
		if !ok {
			continue
		}

		if v == "" {
			continue
		}

		if jsonFieldName == "subdomain" && strings.Contains(v, "monday.com") {
			v, _ = strings.CutSuffix(v, ".monday.com")
		}
		query.Set(jsonFieldName, v)
	}

	u.RawQuery = query.Encode()

	return u.String()
}

func (api *API) log(message ...interface{}) {
	if api.Debug {
		log.Println(message...)
	}
}

func errorCheck(body []byte, status int) error {
	// TODO handle body on most used errors
	if len(body) == 0 && status == fiber.StatusCreated {
		return nil
	}

	e := ErrorResponse{}
	if err := json.Unmarshal(body, &e); err != nil {
		// if it cannot unmarshal, there is no errors in answer
		return nil
	}

	if e.Error != "" && e.ErrorDescription != "" {
		return fmt.Errorf(e.ErrorDescription)
	}

	if e.Errors == nil {
		return nil
	}
	log.Println("ErrorResponse: ", e)
	return fmt.Errorf(e.Errors[0].Message)
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
