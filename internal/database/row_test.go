package database

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewRow(t *testing.T) {
	result := NewRow(-10, "test", "yooho")

	assert.Equal(t, float64(-10), result.Value, "Value should be -10")
	assert.Len(t, result.Tags, 2, "Row tags should be 2")
	assert.Equal(t, "CAD", result.Currency, "Currency should be CAD")
}
