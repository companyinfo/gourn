package gourn

import (
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

// normalize turns the receiving URN into its normalized version.
// Except the NSS part, the other parts can be in lowercase, because there is <hex> chars in NSS
func (u *URN) normalize() {
	u.NID = strings.ToLower(u.NID)
}
