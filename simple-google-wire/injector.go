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

*/

func InitializedService() *SimpleService {
	// inform functions that used as the providers (to create dependency injection)
	wire.Build(
		NewSimpleRepository, NewSimpleService,
	)
	return nil
}
