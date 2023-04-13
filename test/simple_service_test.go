package test

import (
	"fmt"
	simplegooglewire "go-simple-restful-api/simple-google-wire"
	"testing"
)

func TestSimpleService(t *testing.T) {
	simpleService := simplegooglewire.InitializedService()
	fmt.Println(simpleService.SimpleRepository)
}
