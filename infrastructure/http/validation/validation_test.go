package validation

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExistLowerType(t *testing.T) {
	assert.True(t, validateEventType("eventType", "impression", []string{}, nil))
}

func TestExistUpperType(t *testing.T) {
	assert.True(t, validateEventType("eventType", "IMPRESSION", []string{}, nil))
}

func TestNotExistType(t *testing.T) {
	assert.False(t, validateEventType("eventType", "Sample", []string{}, nil))
}

func TestEmptyType(t *testing.T) {
	assert.False(t, validateEventType("eventType", "", []string{}, nil))
}
