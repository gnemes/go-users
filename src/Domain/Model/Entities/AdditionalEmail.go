package entities

type AdditionalEmail struct {
	ID    string
	Email string
}

func (e *AdditionalEmail) GetID() string {
	return e.ID
}

func (e *AdditionalEmail) Validate() error {
	return nil
}