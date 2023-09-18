package monday

import (
	"fmt"
	"github.com/whatcrm/go-monday/models"
	"reflect"
	"strconv"
	"strings"
)

func (api *API) setQueries(queries interface{}, query string) string {
	if queries == meQuery {
		return meQuery
	}

	if query == "" {
		query = `{ "query": " { `
	}

	switch q := queries.(type) {
	case *models.AccountQuery:
		query = api.buildQuery(q, query, "account")
	case *models.PlanQuery:
		return api.buildQuery(q, query, "plan")
	case *models.UserQuery:
		query = api.buildQuery(q, query, "users")
	case *models.OutOfOfficeQuery:
		query = api.buildQuery(q, query, "out_of_office")
	case *models.TeamsQuery:
		query = api.buildQuery(q, query, "teams")
	case *models.ItemQuery:
		query = api.buildQuery(q, query, "items")
	default:
		return "Unknown type"
	}

	query += `} " }`
	return query
}

func (api *API) buildQuery(q interface{}, query, parentQueryName string) string {
	query += parentQueryName + " { "

	typ := reflect.TypeOf(q)
	if typ.Kind() == reflect.Ptr {
		typ = typ.Elem() // Dereference the pointer to get the underlying struct type.
	}

	val := reflect.ValueOf(q)
	if val.Kind() == reflect.Ptr {
		val = val.Elem() // Dereference the pointer to get the struct value.
	}

	// getting json and value from interface{}
	for i := 0; i < typ.NumField(); i++ {
		jsonName, _ := strings.CutSuffix(typ.Field(i).Tag.Get("json"), ",omitempty")

		if !val.Field(i).CanInterface() {
			continue
		}

		fieldValue := val.Field(i).Interface()

		if isStruct(fieldValue) && jsonName == "query_params" {
			query = api.buildParams(fieldValue, query)
			continue
		}

		if isStruct(fieldValue) {
			query = api.buildQuery(fieldValue, query, jsonName)
			continue
		}

		v, ok := fieldValue.(bool)
		// if ok , false, last iteration
		if ok && !v && i != typ.NumField()-1 {
			continue
		}

		// if ok, true, products
		if ok && v && jsonName == "products" {
			query += jsonName + " { id kind } "
			continue
		}

		// if ok, true
		if ok && v {
			query += jsonName + " "
		}

		if i == typ.NumField()-1 {
			query += "} "
		}
		continue
	}

	return query
}

func (api *API) buildParams(p interface{}, query string) string {
	//INFO :
	// general: ids , limit, page, newest_first
	// users: kind, emails,name, non_active, registration_sequence
	// boards: order_by, board_kind, workspace_ids, state
	// workspaces: kind, state, order_by
	// teams, tags: ids
	// updates: limit, page

	// activity_logs: user_ids, column_ids, group_ids, item_ids, from, to
	// columns: types
	// groups: ids
	// team_owners, updates, team_owners_subscribers, users_subscribers	: limit, page
	// views: ids, types
	//

	if p == nil {
		return query
	}

	params, ok := p.(*models.QueryParams)
	if !ok {
		return query
	}
	query, _ = strings.CutSuffix(query, "{ ")

	if params.General != nil {
		addQueryParam(&query, "ids", strings.Join(strings.Fields(fmt.Sprint(params.General.IDs)), ", "))
		addQueryParam(&query, "limit", strconv.Itoa(params.General.Limit))
		addQueryParam(&query, "page", strconv.Itoa(params.General.Page))
		if params.General.NewestFirst {
			addQueryParam(&query, "newest_first", "true")
		}
	}

	if params.Users != nil {
		addQueryParam(&query, "name", arrayToString(params.Users.Name))
		addQueryParam(&query, "kind", arrayToString(params.Users.Kind))
		addQueryParam(&query, "emails", arrayToString(params.Users.Emails))
		addQueryParam(&query, "registration_sequence", strconv.Itoa(params.Users.RegistrationSequence))
		if params.Users.NonActive {
			addQueryParam(&query, "non_active", "true")
		}
	}

	if params.Columns != nil {
		addQueryParam(&query, "type", arrayToString(params.Columns.Types))
	}

	if params.ColumnValues != nil {
		addQueryParam(&query, "ids", arrayToString(params.ColumnValues.IDs))
	}

	query += ") { "
	return query
}

func arrayToString(s []string) (result string) {
	// it will be transformed to []byte later on, therefore we need \" in string
	// do not remove, unless necessary

	result = "["
	for i, v := range s {
		result += fmt.Sprintf("\\\"") + v + fmt.Sprintf("\\\"")
		if i != len(s)-1 {
			result += fmt.Sprintf(" , ")
		}
	}
	result += "]"
	fmt.Println("StringArray: ", result)
	return
}

func addQueryParam(query *string, name, value string) {
	if value == "[]" || value == "0" {
		return
	}
	*query += fmt.Sprintf("( %s: %s ", name, value)
	return
}

func isStruct(fieldValue any) bool {
	if fieldValue == nil {
		return false
	}

	// If it's a struct, check if it's a non-zero value.
	val := reflect.ValueOf(fieldValue)
	if val.IsZero() {
		return false
	} // Check if the value is not a zero value.

	if reflect.TypeOf(fieldValue).Kind() == reflect.Struct ||
		(reflect.TypeOf(fieldValue).Kind() == reflect.Ptr &&
			reflect.TypeOf(fieldValue).Elem().Kind() == reflect.Struct) {
		return true
	}
	return false
}
