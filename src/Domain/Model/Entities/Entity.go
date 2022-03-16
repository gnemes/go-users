package entities

type Entity interface {
	Validate() error
}