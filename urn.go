// Copyright 2023 Company.info
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

import "C"
import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"regexp"
	"strings"
)

var (
	// ErrInvalidURN represents an error for invalid URN
	ErrInvalidURN  = errors.New("invalid URN")
	genericPattern = regexp.MustCompile(`^(?i)urn:(?P<nid>[a-z0-9][a-z0-9-]{0,31}):(?P<nss>(?:[-a-z0-9()+,.:=@;$_!*'&~/]|%[0-9a-fA-F]{2})+)$`)
)

// URN represents a Uniform Resource Name.
//
// The general form is: urn:<nid>:<nss>
// Compliant with https://tools.ietf.org/html/rfc2141.
type URN struct {
	NID string // Namespace identifier
	NSS string // Namespace specific string
}

// Parse is responsible to create a URN instance from a byte array matching the correct URN syntax.
func Parse(u string) (*URN, error) {
	matches := genericPattern.FindStringSubmatch(u)
	if matches == nil {
		return nil, ErrInvalidURN
	}

	urn := &URN{
		NID: matches[1],
		NSS: matches[2],
	}

	urn.normalize()

	// To avoid confusion with the "urn:" identifier, the NID "urn" is reserved and MUST NOT be used.
	// And checking by regex in Go is impossible, because the Go doesn't support "negative lookahead" and "positive lookahead".
	if urn.NID == "urn" {
		return nil, ErrInvalidURN
	}

	return urn, nil
}

// String represents the URN struct in a valid URN string.
func (u *URN) String() string {
	return strings.Join([]string{"urn", u.NID, u.NSS}, ":")
}

// MarshalJSON marshals the URN to JSON string.
func (u *URN) MarshalJSON() ([]byte, error) {
	return json.Marshal(u.String())
}

// UnmarshalJSON unmarshals a URN from JSON string.
func (u *URN) UnmarshalJSON(bytes []byte) error {
	var str string
	if err := json.Unmarshal(bytes, &str); err != nil {
		return err
	}

	value, err := Parse(str)
	if err != nil {
		return err
	}

	*u = *value

	return nil
}

// Value implements the database/sql valuer interface
func (u *URN) Value() (driver.Value, error) {
	return u.String(), nil
}

// Scan implements the database/sql scanner interface
func (u *URN) Scan(v interface{}) error {
	urn, err := scan(v)
	if err != nil {
		return err
	}

	*u = *urn

	return nil
}

// normalize turns the receiving URN into its normalized version.
// Except the NSS part, the other parts can be in lowercase, because there is <hex> chars in NSS
func (u *URN) normalize() {
	u.NID = strings.ToLower(u.NID)
}
