package plugin_test

import (
	"testing"

	. "github.com/aunem/transpose/pkg/plugin"
	"github.com/stretchr/testify/assert"
)

func TestResolvePlugin(*testing.T) {
	path, err := ResolvePlugin(name, pkg string)
	assert.Nil(t, err)
}

func TestBuildPlugin(*testing.T) {
	path, err := ResolvePlugin(name, pkg string)
	assert.Nil(t, err)
	path, err := BuildPlugin(name, path, MiddlewareType)
	assert.Nil(t, err)
}
