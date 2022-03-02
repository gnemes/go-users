package domainerrors

// MissingFilterError domainerror
type MissingFilterError struct {
	Err     string
	Message string
	Code    int32
}

// Error string implementation
func (de *MissingFilterError) Error() string {
	return de.Err
}
