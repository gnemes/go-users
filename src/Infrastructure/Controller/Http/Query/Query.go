package queryhttp

type QuerySort struct {
	Field     string
	Direction string
}

// Query model
type Query struct {
	Offset   int32
	Limit    int32
	Filters  map[string]interface{}
	Includes map[string]bool
	Sorts    []QuerySort
	After    *string
	Before   *string
}