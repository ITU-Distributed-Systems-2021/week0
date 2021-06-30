package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_echo(t *testing.T) {
	said := echo("hello")
	assert.Equal(t, "hello", said)
}
