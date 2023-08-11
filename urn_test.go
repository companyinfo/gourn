package gourn

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

type testCase struct {
	input string
	urn   *URN
	str   string
	err   error
}

var rfc2141TestCases = []testCase{
	{
		"URN:foo:a123,456",
		&URN{
			NID: "foo",
			NSS: "a123,456",
		},
		"urn:foo:a123,456",
		nil,
	},
	{
		"urn:foo:a123,456",
		&URN{
			NID: "foo",
			NSS: "a123,456",
		},
		"urn:foo:a123,456",
		nil,
	},
	{
		"urn:FOO:a123,456",
		&URN{
			NID: "foo",
			NSS: "a123,456",
		},
		"urn:foo:a123,456",
		nil,
	},
	{
		"urn:foo:A123,456",
		&URN{
			NID: "foo",
			NSS: "A123,456",
		},
		"urn:foo:A123,456",
		nil,
	},
	{
		"urn:foo:a123%2C456",
		&URN{
			NID: "foo",
			NSS: "a123%2C456",
		},
		"urn:foo:a123%2C456",
		nil,
	},
	{
		"URN:FOO:a123%2c456",
		&URN{
			NID: "foo",
			NSS: "a123%2c456",
		},
		"urn:foo:a123%2c456",
		nil,
	},
	{
		"urn:a:b",
		&URN{
			NID: "a",
			NSS: "b",
		},
		"urn:a:b",
		nil,
	},
	{
		"urn:a::",
		&URN{
			NID: "a",
			NSS: ":",
		},
		"urn:a::",
		nil,
	},
	{
		"urn:a:-",
		&URN{
			NID: "a",
			NSS: "-",
		},
		"urn:a:-",
		nil,
	},
	{
		"URN:simple:simple",
		&URN{
			NID: "simple",
			NSS: "simple",
		},
		"urn:simple:simple",
		nil,
	},
	{
		"urn:urna:simple",
		&URN{
			NID: "urna",
			NSS: "simple",
		},
		"urn:urna:simple",
		nil,
	},
	{
		"urn:burnout:nss",
		&URN{
			NID: "burnout",
			NSS: "nss",
		},
		"urn:burnout:nss",
		nil,
	},
	{
		"urn:burn:nss",
		&URN{
			NID: "burn",
			NSS: "nss",
		},
		"urn:burn:nss",
		nil,
	},
	{
		"urn:urnurnurn:x",
		&URN{
			NID: "urnurnurn",
			NSS: "x",
		},
		"urn:urnurnurn:x",
		nil,
	},
	{
		"urn:abcdefghilmnopqrstuvzabcdefghilm:x",
		&URN{
			NID: "abcdefghilmnopqrstuvzabcdefghilm",
			NSS: "x",
		},
		"urn:abcdefghilmnopqrstuvzabcdefghilm:x",
		nil,
	},
	{
		"URN:123:x",
		&URN{
			NID: "123",
			NSS: "x",
		},
		"urn:123:x",
		nil,
	},
	{
		"URN:abcd-:x",
		&URN{
			NID: "abcd-",
			NSS: "x",
		},
		"urn:abcd-:x",
		nil,
	},
	{
		"URN:abcd-abcd:x",
		&URN{
			NID: "abcd-abcd",
			NSS: "x",
		},
		"urn:abcd-abcd:x",
		nil,
	},
	{
		"urn:urnx:urn",
		&URN{
			NID: "urnx",
			NSS: "urn",
		},
		"urn:urnx:urn",
		nil,
	},
	{
		"urn:ciao:a:b:c",
		&URN{
			NID: "ciao",
			NSS: "a:b:c",
		},
		"urn:ciao:a:b:c",
		nil,
	},
	{
		"urn:aaa:x:y:",
		&URN{
			NID: "aaa",
			NSS: "x:y:",
		},
		"urn:aaa:x:y:",
		nil,
	},
	{
		"urn:ciao:-",
		&URN{
			NID: "ciao",
			NSS: "-",
		},
		"urn:ciao:-",
		nil,
	},
	{
		"urn:colon:::::nss",
		&URN{
			NID: "colon",
			NSS: "::::nss",
		},
		"urn:colon:::::nss",
		nil,
	},
	{
		"urn:ciao:@!=%2C(xyz)+a,b.*@g=$_'",
		&URN{
			NID: "ciao",
			NSS: "@!=%2C(xyz)+a,b.*@g=$_'",
		},
		"urn:ciao:@!=%2C(xyz)+a,b.*@g=$_'",
		nil,
	},
	{
		"URN:hexes:%25",
		&URN{
			NID: "hexes",
			NSS: "%25",
		},
		"urn:hexes:%25",
		nil,
	},
	{
		"URN:xs:abc%1Dz%2F%3az",
		&URN{
			NID: "xs",
			NSS: "abc%1Dz%2F%3az",
		},
		"urn:xs:abc%1Dz%2F%3az",
		nil,
	},
	{
		"URN:FOO:ABC%FFabc123%2c456",
		&URN{
			NID: "foo",
			NSS: "ABC%FFabc123%2c456",
		},
		"urn:foo:ABC%FFabc123%2c456",
		nil,
	},
	{
		"URN:FOO:ABC%FFabc123%2C456%9A",
		&URN{
			NID: "foo",
			NSS: "ABC%FFabc123%2C456%9A",
		},
		"urn:foo:ABC%FFabc123%2C456%9A",
		nil,
	},
	{
		"urn:ietf:params:scim:schemas:core:2.0:User",
		&URN{
			NID: "ietf",
			NSS: "params:scim:schemas:core:2.0:User",
		},
		"urn:ietf:params:scim:schemas:core:2.0:User",
		nil,
	},
	{
		"urn:ietf:params:scim:schemas:extension:enterprise:2.0:User:meta.lastModified",
		&URN{
			NID: "ietf",
			NSS: "params:scim:schemas:extension:enterprise:2.0:User:meta.lastModified",
		},
		"urn:ietf:params:scim:schemas:extension:enterprise:2.0:User:meta.lastModified",
		nil,
	},
	{
		"These are invalid inputs:",
		nil,
		"",
		ErrInvalidURN,
	},
	{
		"URN:-xxx:x",
		nil,
		"",
		ErrInvalidURN,
	},
	{
		"urn::colon:nss",
		nil,
		"",
		ErrInvalidURN,
	},
	{
		"urn:abcdefghilmnopqrstuvzabcdefghilmn:specificstring",
		nil,
		"",
		ErrInvalidURN,
	},
	{
		"URN:a!?:x",
		nil,
		"",
		ErrInvalidURN,
	},
	{
		"URN:#,:x",
		nil,
		"",
		ErrInvalidURN,
	},
	{
		"urn:urn:NSS",
		nil,
		"",
		ErrInvalidURN,
	},
	{
		"urn:URN:NSS",
		nil,
		"",
		ErrInvalidURN,
	},
	{
		"urn:white space:NSS",
		nil,
		"",
		ErrInvalidURN,
	},
	{
		"urn:concat:no spaces",
		nil,
		"",
		ErrInvalidURN,
	},
	{
		"urn:a:%",
		nil,
		"",
		ErrInvalidURN,
	},
	{
		"urn:",
		nil,
		"",
		ErrInvalidURN,
	},
}

func TestParseRFC2141URN(t *testing.T) {
	for _, tc := range rfc2141TestCases {
		urn, err := Parse(tc.input)

		require.ErrorIs(t, err, tc.err, "Input", tc.input)
		assert.Equal(t, urn, tc.urn, "Input", tc.input)

		if tc.urn != nil {
			assert.Equal(t, urn.String(), tc.str, "Input", tc.input)
		}
	}
}

func TestValidURNJSONMarshaling(t *testing.T) {
	urn := &URN{NID: "api", NSS: "user:123"}

	// Marshal to JSON
	marshaledJSON, err := json.Marshal(urn)
	require.NoError(t, err)

	// Unmarshal from JSON
	var unmarshaledURN URN
	err = json.Unmarshal(marshaledJSON, &unmarshaledURN)
	require.NoError(t, err)

	assert.Equal(t, unmarshaledURN.NID, urn.NID, "Unmarshaled NID is different from original.")
	assert.Equal(t, unmarshaledURN.NSS, urn.NSS, "Unmarshaled NSS is different from original.")
}

func TestInvalidURNJSONMarshaling(t *testing.T) {
	urn := &URN{NID: "urn", NSS: `user:123`}

	// Marshal to JSON
	marshaledJSON, err := json.Marshal(urn)
	require.NoError(t, err)

	// Unmarshal from JSON
	var unmarshaledURN URN
	err = json.Unmarshal(marshaledJSON, &unmarshaledURN)
	assert.ErrorIs(t, err, ErrInvalidURN)

	// Unmarshal JSON syntax error
	var jsonErr *json.SyntaxError
	err = urn.UnmarshalJSON([]byte("urn:api:invalid-json"))

	assert.ErrorAs(t, err, &jsonErr)
}
