package gourn

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestValidNullURN(t *testing.T) {
	var nullURN NullURN

	// Test valid Nil URN
	err := nullURN.Scan(nil)
	require.NoError(t, err)
	assert.False(t, nullURN.Valid)

	nilVal, err := nullURN.Value()
	require.NoError(t, err)
	assert.Nil(t, nilVal)

	// Test valid URN
	urn := "urn:ns:z:x"
	err = nullURN.Scan(urn)
	require.NoError(t, err)
	assert.True(t, nullURN.Valid)

	urnVal, err := nullURN.Value()
	require.NoError(t, err)
	assert.Equal(t, urnVal, nullURN.URN)
}

func TestInvalidNullURN(t *testing.T) {
	var nullURN NullURN
	err := nullURN.Scan("urn:")

	require.Error(t, err)
	assert.ErrorIs(t, errors.Unwrap(err), ErrInvalidURN)
}
