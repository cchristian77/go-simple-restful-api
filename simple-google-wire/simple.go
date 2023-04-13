package simple_google_wire

/*
Provider
to do a dependency injection, constructor function needs to be created.
In Google Wire, constructor functions is called provider.
*/

type SimpleRepository struct {
}

// provider (constructor) function
func NewSimpleRepository() *SimpleRepository {
	return &SimpleRepository{}
}

// SimpleService requires (depends) SimpleRepository
type SimpleService struct {
	*SimpleRepository
}

// provider (constructor) function
func NewSimpleService(repository *SimpleRepository) *SimpleService {
	return &SimpleService{
		SimpleRepository: repository,
	}
}
