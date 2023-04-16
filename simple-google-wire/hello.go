package simple_google_wire

type SayHello interface {
	Hello(name string) string
}

type HelloService struct {
	// no need to define attribute's name
	// if the same as the data type's name
	SayHello
}

type SayHelloImpl struct {
}

func (s *SayHelloImpl) Hello(name string) string {
	return "Hello " + name
}

// Best Pratice : return data type is defined the implementation struct, rather than the inteface
func NewSayHelloImpl() *SayHelloImpl {
	return &SayHelloImpl{}
}

func NewHelloService(sayHello SayHello) *HelloService {
	return &HelloService{SayHello: sayHello}
}
