package test

import (
	"github.com/stretchr/testify/assert"
	simple_google_wire "go-simple-restful-api/simple-google-wire"
	"testing"
)

func TestSimpleServiceError(t *testing.T) {
	simpleService, err := simple_google_wire.InitializedService(true)
	assert.NotNil(t, err)
	assert.Nil(t, simpleService)
}

func TestSimpleServiceSuccess(t *testing.T) {
	simpleService, err := simple_google_wire.InitializedService(false)
	assert.Nil(t, err)
	assert.NotNil(t, simpleService)
}
