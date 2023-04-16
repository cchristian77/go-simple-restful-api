package simple_google_wire

import "errors"

/*
Provider
to do a dependency injection, constructor function needs to be created.
In Google Wire, constructor functions is called provider.

Error
If provider return error, Google Wire can handle the errror.
Provider & Injector are  added return error,
*/

type SimpleRepository struct {
	Error bool
}

// provider (constructor) function
// pass parameter (signature) isError to provider function
func NewSimpleRepository(isError bool) *SimpleRepository {
	return &SimpleRepository{
		Error: isError,
	}
}

// SimpleService requires (depends) SimpleRepository
type SimpleService struct {
	*SimpleRepository
}

// provider (constructor) function
func NewSimpleService(repository *SimpleRepository) (*SimpleService, error) {
	if repository.Error {
		return nil, errors.New("failed create service")
	} else {
		return &SimpleService{
			SimpleRepository: repository,
		}, nil
	}
}
