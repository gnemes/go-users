package domainerrors

// NotFoundError domainerror
type NotFoundError struct {
	Err     string
	Message string
	Code    int32
}

// Error string implementation
func (de *NotFoundError) Error() string {
	return de.Err
}
