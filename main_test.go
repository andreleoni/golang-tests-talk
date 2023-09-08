package main

import (
	"golang-tests-talk/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDo(t *testing.T) {
	requestMock := &mocks.Request{}

	requestMock.EXPECT().KeepAlive("https://www.google.com").Return(true)

	assert.Equal(t, true, do(requestMock))
}
