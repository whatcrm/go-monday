package monday

type Get struct {
	api *API
}

type Mutate struct {
	api *API
}

func (api *API) Get() *Get {
	return &Get{api}
}

func (api *API) Mutation() *Mutate {
	return &Mutate{api}
}

func (c *Get) CustomQuery(query interface{}, vars map[string]interface{}) (out interface{}, err error) {
	options := makeRequestOptions{
		BaseURL:   mondayAPI,
		Query:     &query,
		Variables: vars,
	}

	err = c.api.makeRequest(options)
	return
}

func (c *Mutate) CustomMutation(mutation interface{}, vars map[string]interface{}) (out interface{}, err error) {
	options := makeRequestOptions{
		BaseURL:   mondayAPI,
		Mutation:  &mutation,
		Variables: vars,
	}

	err = c.api.makeRequest(options)
	return
}
