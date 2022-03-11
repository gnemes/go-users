package domainerrors

// Forbidden forbiddenerror
type ForbiddenError struct {
	Err     string
	Message string
	Code    int32
}

// Error string implementation
func (de *ForbiddenError) Error() string {
	return de.Err
}