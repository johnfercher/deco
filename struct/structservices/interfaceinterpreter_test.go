package structservices

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInterfaceInterpreter_Interpret(t *testing.T) {
	// Arrange
	content := `package example

import "context"

type Example interface {
	Add(ctx context.Context, id string) error
}
`
	interpreter := NewInterfaceInterpreter()

	// Act
	interfaces := interpreter.Interpret(content)

	// Assert
	assert.Len(t, interfaces, 1)
	assert.Equal(t, "example", interfaces[0].PackageName)
	assert.Equal(t, "Example", interfaces[0].Name)
	assert.Equal(t, "context", interfaces[0].Imports[0].Full)
	assert.Equal(t, "context", interfaces[0].Imports[0].Alias)
}
