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

import (
	"database/sql/driver"
	"fmt"
)

// NullURN represents a URN that may be null.
// NullURN implements the Scanner interface, so
// it can be used as a scan destination:
//
//	var s NullURN
//	err := db.QueryRow("SELECT urn FROM foo WHERE id=?", id).Scan(&s)
//	...
//	if s.Valid {
//	   // use s.URN
//	} else {
//	   // NULL value
//	}
type NullURN struct {
	URN   URN
	Valid bool // Valid is true if URN is not NULL
}

// Scan implements the Scanner interface.
func (nu *NullURN) Scan(value any) error {
	if value == nil {
		nu.URN, nu.Valid = URN{}, false

		return nil
	}

	nu.Valid = true
	urn, err := scan(value)
	if err != nil {
		return err
	}

	nu.URN = *urn

	return nil
}

// Value implements the driver Valuer interface.
func (nu *NullURN) Value() (driver.Value, error) {
	if !nu.Valid {
		return nil, nil
	}

	return nu.URN, nil
}

func scan(v interface{}) (*URN, error) {
	switch value := v.(type) {
	case string:
		urn, err := Parse(value)
		if err != nil {
			return nil, fmt.Errorf("faild to scan URN: %w", err)
		}

		return urn, nil
	default:
		return nil, fmt.Errorf("unexpected type %T", v)
	}
}
