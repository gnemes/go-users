package entities

type Error struct {
	ID     string
	Status int
	Code   string
	Title  string
	Detail string
}

func (e *Error) GetID() string {
	return e.ID
}

func (e *Error) Validate() error {
	return nil
}