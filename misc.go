package monday

import (
	"context"
	"errors"
	"github.com/hasura/go-graphql-client"
	"golang.org/x/oauth2"
	"log"
	"net/http"
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

	api.log("setting the router...")
	client := api.setRouter(options.BaseURL)

	api.log("sending the request...")
	if options.Query != nil {
		err = client.Query(context.Background(), options.Query, options.Variables)
	}

	if options.Mutation != nil {
		err = client.Mutate(context.Background(), options.Mutation, options.Variables)
	}

	api.log("finishing the request")
	return
}

func (api *API) log(message ...interface{}) {
	if api.Debug {
		log.Println(message...)
	}
}
