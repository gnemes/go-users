package middleware

import (
	"net/http"
	"strings"
	"strconv"

	context "github.com/gnemes/go-users/Domain/Services/Context"
	controllerhttp "github.com/gnemes/go-users/Infrastructure/Controller/Http"
	domainerrors "github.com/gnemes/go-users/Domain/Errors"
	logger "github.com/gnemes/go-users/Domain/Services/Logger"
	queryhttp "github.com/gnemes/go-users/Infrastructure/Controller/Http/Query"
)

const (
	defaultQueryLimit  = 20
	defaultQueryOffset = 0
)

type RequestQueryParserMiddleware struct {
	Logger          logger.Logger
	ErrorController *controllerhttp.Error
	Context         *context.Context
}

func (m *RequestQueryParserMiddleware) Execute(next func(w http.ResponseWriter, r *http.Request)) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		m.Logger.Debugf("Middleware / RequestQueryParserMiddleware()")
		defer m.Logger.Debugf("Middleware / RequestQueryParserMiddleware() ending...")

		// Init vars
		var limits []string
		var offsets []string
		var afters []string
		var befores []string
		var filters map[string]interface{}
		var includes map[string]bool
		var sorts []queryhttp.QuerySort

		// Get URL from request
		url := r.URL

		filters = make(map[string]interface{})
		includes = make(map[string]bool)
		sorts = make([]queryhttp.QuerySort, 0)

		// Looping parameters
		params := url.Query()
		for key, v := range params {
			k := strings.ToLower(key)

			switch k {
			case "page[limit]":
				limits = v
			case "page[offset]":
				offsets = v
			case "page[after]":
				afters = v
			case "page[before]":
				befores = v
			case "include":
				if len(v) > 0 {
					for _, value := range v {
						i := strings.Split(value, ",")
						if len(i) > 0 {
							for _, item := range i {
								includes[item] = true
							}
						}
					}
				}
			case "sort":
				if len(v) > 0 {
					for _, value := range v {
						i := strings.Split(value, ",")
						if len(i) > 0 {
							for _, item := range i {
								direction := "ASC"
								if strings.HasPrefix(item, "-") {
									direction = "DESC"
									item = strings.TrimLeft(item, "-")
								}
								sorts = append(sorts, queryhttp.QuerySort{
									Field:     item,
									Direction: direction,
								})
							}
						}
					}
				}
			default:
				if strings.HasPrefix(k, "filter[") && strings.HasSuffix(k, "]") {
					newkey := strings.TrimPrefix(strings.TrimSuffix(k, "]"), "filter[")
					filters[newkey] = v
				} else {
					m.Logger.Debugf("Unrecognized key set to query: %s", key)
				}
			}
		}

		// Limit
		var limit int32
		limit = defaultQueryLimit
		if limits != nil && len(limits) > 0 {
			l, err := strconv.ParseInt(limits[0], 10, 32)
			if err != nil {
				m.Logger.Errorf("Invalid limit params.")
				m.ErrorController.WriteHttpError(&domainerrors.BadRequestError{Err: "Invalid limit params"}, w)
				return
			}
			limit = int32(l)
		}

		// Offset
		var offset int32
		offset = defaultQueryOffset
		if offsets != nil && len(offsets) > 0 {
			o, err := strconv.ParseInt(offsets[0], 10, 32)
			if err != nil {
				m.Logger.Errorf("Invalid offset params.")
				m.ErrorController.WriteHttpError(&domainerrors.BadRequestError{Err: "Invalid offset params"}, w)
				return
			}
			offset = int32(o)
		}

		// After
		var after *string
		if afters != nil && len(afters) > 0 {
			after = &afters[0]
		}

		// After
		var before *string
		if befores != nil && len(befores) > 0 {
			before = &befores[0]
		}

		// Query instance
		query := &queryhttp.Query{
			Filters:  filters,
			Includes: includes,
			Offset:   offset,
			Limit:    limit,
			Sorts:    sorts,
			After:    after,
			Before:   before,
		}

		m.Context.Add("Query", query)

		next(w, r)
	}
}