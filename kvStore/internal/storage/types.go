package storage

type Storage interface {
	//add or update key-val pair
	Set(key string, value string) error
	//retrieve val
	Get(key string) (string, error)
	//remove key-val pair
	Delete(key string) error
}

type NotFoundError struct{}

func (e *NotFoundError) Error() string {
	return "key not found"
}
