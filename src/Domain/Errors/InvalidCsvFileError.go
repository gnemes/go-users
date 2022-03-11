package domainerrors

// NotFoundError domainerror
type InvalidCsvFileError struct {
	Err          string
	Message      string
	Code         int32
	ErrorFileUrl string
}

// Error string implementation
func (de *InvalidCsvFileError) Error() string {
	return de.Err
}
