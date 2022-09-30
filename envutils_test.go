package envutils

import (
	"github.com/go-playground/assert/v2"
	"os"
	"testing"
)

func TestLoadEnv(t *testing.T) {
	LoadEnv()

	actual := os.Getenv("foo")

	expect := "bar"

	assert.Equal(t, expect, actual)
}
