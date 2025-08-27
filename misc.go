package monday

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"reflect"
	"strings"

	"github.com/hasura/go-graphql-client"
	"golang.org/x/oauth2"
)

func (api *API) setRouter(uri string) (client *graphql.Client) {
	src := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: api.Auth},
	)

	httpClient := oauth2.NewClient(context.Background(), src)

	client = graphql.NewClient(uri, httpClient).WithRequestModifier(func(s *http.Request) {
		s.Header.Set(apiVersion, October2023)
	})
	return
}

func (api *API) makeRequest(options makeRequestOptions) (err error) {
	if options.Query == nil && options.Mutation == nil {
		return errors.New("query and mutation are not found")
	}
	if api.Auth == "" {
		return errors.New("authorization is required")
	}

	l := logOptions{
		Domain: api.Domain,
	}
	l.Variables = options.Variables
	client := api.setRouter(options.BaseURL)

	if options.Query != nil {
		l.Request = reflect.TypeOf(options.Query).Elem().Field(0).Tag.Get("graphql")
		err = client.Query(context.Background(), options.Query, options.Variables)
		l.Response = options.Query
	}

	if options.Mutation != nil {
		l.Request = reflect.TypeOf(options.Mutation).Elem().Field(0).Tag.Get("graphql")
		err = client.Mutate(context.Background(), options.Mutation, options.Variables)
		l.Response = options.Mutation
	}

	if err != nil {
		if strings.Contains(err.Error(), "Internal Server Error") {
			// cutting html body from an error
			err = errors.New("500 Internal Server Error")
		}
		l.Error = err.Error()
		l.Response = nil
	}

	m, _ := json.Marshal(l)

	api.log(string(m))
	return
}

func (api *API) log(message ...interface{}) {
	if api.Debug {
		api.l.Field(api.Domain).Debug(message...)
	}
}
