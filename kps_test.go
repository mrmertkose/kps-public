package kps

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetService_ShouldReturnError_WhenEmptyValuesPassed(t *testing.T) {

	var identityNumber int64 = 0
	var year int32 = 0
	name := ""
	surname := ""

	service, err := GetService(identityNumber, name, surname, year)
	assert.Error(t, err)
	assert.Empty(t, service)
}

func TestGetService_ShouldReturnFalse(t *testing.T) {
	var identityNumber int64 = 1
	var year int32 = 1
	name := "name"
	surname := "surname"

	service, err := GetService(identityNumber, name, surname, year)
	assert.NoError(t, err)
	assert.Equal(t, false, service)
}
