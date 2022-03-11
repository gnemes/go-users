package httpError

type HttpError struct {
	Status int
	Code   string
	Title  string
	Detail string
}