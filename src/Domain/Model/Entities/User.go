package entities

type User struct {
	ID         string
	Username   string

	Platform    *Platform
	UserProfile *UserProfile
}

func (e *User) GetID() string {
	return e.ID
}

func (e *User) Validate() error {
	return nil
}