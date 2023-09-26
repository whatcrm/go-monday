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

func (c *Get) CustomQuery(query interface{}) (out interface{}, err error) {
	options := makeRequestOptions{
		BaseURL: mondayAPI,
		Query:   &query,
	}

	err = c.api.makeRequest(options)
	return
}

func (m *Mutate) CustomMutation(mutation interface{}) (out interface{}, err error) {
	options := makeRequestOptions{
		BaseURL:  mondayAPI,
		Mutation: &mutation,
	}

	err = m.api.makeRequest(options)
	return
}
