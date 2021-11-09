package event

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTypeExist(t *testing.T) {
	assert.True(t, TypeExist("Click"))
}

func TestTypeExistLowerCase(t *testing.T) {
	assert.False(t, TypeExist("visible"))
}

func TestTypeExistUpperCase(t *testing.T) {
	assert.False(t, TypeExist("VISIBLE"))
}

func TestTypeExistDontExist(t *testing.T) {
	assert.False(t, TypeExist("sample"))
}

func TestTypeExistEmpty(t *testing.T) {
	assert.False(t, TypeExist(""))
}

func TestTypeFromString(t *testing.T) {
	value, err := TypeFromString("Visible")
	if err != nil {
		panic("Can't get string from numeric value")
	}
	assert.Equal(t, Visible, value)
}

func TestTypeFromStringLowerCase(t *testing.T) {
	_, err := TypeFromString("visible")
	assert.NotNil(t, err)
	assert.Equal(t, "This event type doesn't exist", err.Error())
}

func TestTypeFromStringUpperCase(t *testing.T) {
	_, err := TypeFromString("VISIBLE")
	assert.NotNil(t, err)
	assert.Equal(t, "This event type doesn't exist", err.Error())
}
func TestTypeFromStringDontExist(t *testing.T) {
	_, err := TypeFromString("sample")
	assert.NotNil(t, err)
	assert.Equal(t, "This event type doesn't exist", err.Error())
}

func TestTypeFromStringEmpty(t *testing.T) {
	_, err := TypeFromString("")
	assert.NotNil(t, err)
	assert.Equal(t, "This event type doesn't exist", err.Error())
}

func TestTypeString(t *testing.T) {
	assert.Equal(t, "Impression", Impression.String())
}

func TestTypeStringDontExist(t *testing.T) {
	assert.Equal(t, "Unknown", Type(10).String())
}
