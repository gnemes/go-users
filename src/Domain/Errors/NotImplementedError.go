package domainerrors

// NotImplementedError domainerror
type NotImplementedError struct {
	Err     string
	Message string
	Code    int32
}

// Error string implementation
func (de *NotImplementedError) Error() string {
	return de.Err
}