package usecases

type Presenter interface {
	Present() ([]byte, error)
}