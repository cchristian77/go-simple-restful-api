//go:build wireinject
// +build wireinject

package simple_google_wire

import "github.com/google/wire"

/*
Injector
Injector is constructor function, the content is configurations (dependencies) which need to be informed to Google Wire.
Google Wire will auto generate dependency injection code, also we need to add comment as the marker,
e.e. go:build wireinject, with this comment the file will be ignored when building the application.

Dependency Injection
After creating Injector & Provider, run command from Google Wire to auto generate dependency injection code.
Command : wire gen module/package_name or wire ./folder_name
Example : wire gen cchristian77/go-simple-restful-api/simple_google_wire or wire gen ./simple-go-wire

Or cd to directory then run command : wire

Injector Signature
in Google Wire, we can send parameters to the injector that will be generated automatically.

Provider Set : to grouping multiple providers

Binding Interface
to bind interface that used in the provider.
to do a binding interface, we need to tell Google Wire the interface will use provider with which data types.
*/

func InitializedService(isError bool) (*SimpleService, error) {
	// inform functions that used as the providers (to create dependency injection)
	wire.Build(
		NewSimpleRepository, NewSimpleService,
	)
	return nil, nil
}

func InitializedDatabase() *DatabaseRepository {
	wire.Build(
		NewDatabasePostgreSQL,
		NewDatabaseMongoDB,
		NewDatabaseRepository,
	)

	return nil
}

// Provider set
// create a set to group all foo providers
var fooSet = wire.NewSet(NewFooRepository, NewFooService)

// Provider set
var barSet = wire.NewSet(NewBarRepository, NewBarService)

func InitializedFooBarService() *FooBarService {
	wire.Build(fooSet, barSet, NewFooBarService)
	return nil
}

// Binding Interface
var helloSet = wire.NewSet(
	NewSayHelloImpl,
	wire.Bind(new(SayHello), new(*SayHelloImpl)), // Bind interface
)

func InitializedHelloService() *HelloService {
	// returns error
	// because NewHelloService requires SayHello interface not SayHelloImpl (not followed the parameter)
	//wire.Build(NewHelloService, NewSayHelloImpl)

	wire.Build(helloSet, NewHelloService)
	return nil
}
