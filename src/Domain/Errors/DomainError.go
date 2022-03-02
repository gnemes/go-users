package domainerrors

// DomainError domainerror
type DomainError struct {
	Err     string
	Message string
	Code    int32
}

// Error string implementation
func (de *DomainError) Error() string {
	return de.Err
}