package domainerrors

// BadRequestError domainerror
type BadRequestError struct {
	Err     string
	Message string
	Code    int32
}

// Error string implementation
func (de *BadRequestError) Error() string {
	return de.Err
}
