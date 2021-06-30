package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_echo(t *testing.T) {
	said := Echo("hello")
	assert.Equal(t, "hello", said)
}
