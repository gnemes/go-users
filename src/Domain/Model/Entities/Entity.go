package entities

type Entity interface {
	GetID() string
	Validate() error
}