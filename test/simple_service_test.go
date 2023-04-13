package test

import (
	"fmt"
	simple_google_wire "go-simple-restful-api/simple-google-wire"
	"testing"
)

func TestSimpleService(t *testing.T) {
	simpleService, err := simple_google_wire.InitializedService()
	fmt.Println(err)
	fmt.Println(simpleService.SimpleRepository)
}
