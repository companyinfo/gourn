// Copyright 2023-2025 Company.info
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

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
