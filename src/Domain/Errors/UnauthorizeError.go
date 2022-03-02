package domainerrors

// UnauthorizeError unauthorize
type UnauthorizeError struct {
	Err     string
	Message string
	Code    int32
}

// Error string implementation
func (de *UnauthorizeError) Error() string {
	return de.Err
}