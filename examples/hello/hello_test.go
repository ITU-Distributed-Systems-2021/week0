package hello

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func Test_echo(t *testing.T) {
	assert.Equal(t, "hello", echo("hello"))
}
