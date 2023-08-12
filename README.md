# GoURN

[![GitHub Workflow Status (branch)](https://img.shields.io/github/actions/workflow/status/companyinfo/gourn/ci.yaml?branch=main)](https://github.com/companyinfo/gourn/actions/workflows/ci.yaml) [![Go Report Card](https://goreportcard.com/badge/github.com/companyinfo/gourn)](https://goreportcard.com/report/github.com/companyinfo/gourn) [![Go Reference](https://pkg.go.dev/badge/github.com/companyinfo/gourn.svg)](https://pkg.go.dev/github.com/companyinfo/gourn)

**GoURN** is a Go package that provides a comprehensive URN (Uniform Resource Name) parsing library. It follows the specifications outlined in [RFC 2141](https://datatracker.ietf.org/doc/html/rfc2141).

## Features

- Parse valid URNs according to [RFC 2141](https://datatracker.ietf.org/doc/html/rfc2141).
- Normalize the parsed URN to ensure consistency.
- Convert parsed URNs to valid string representation.
- Serialize and deserialize URNs to/from JSON format.

## Installation

To use **GoURN** in your Go project, install Go and run `go get`:

```shell
$ go get -u github.com/companyinfo/gourn
```

## Usage
Here's a quick example of how to use **GoURN** to parse and work with URNs:

```go
package main

import (
	"fmt"
	"github.com/companyinfo/gourn"
	"log"
)

func main() {
	urnString := "urn:example:resource"
	urn, err := gourn.Parse(urnString)
	if err != nil {
		log.Print("Unable to parse the URN: ", err)

		return
	}

	fmt.Println("NID:", urn.NID)
	fmt.Println("NSS:", urn.NSS)
	fmt.Println("URN String:", urn)
}
```

This will print:
```shell
NID: example
NSS: resource
URN String: urn:example:resource
```
For more detailed usage examples, please refer to the code and tests in this repository.

## Contribution

Contributions to **GoURN** are welcome! If you find a bug, want to add a feature, or have suggestions for improvements, feel free to open issues or submit pull requests.

## License

Copyright 2023 Company.info

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

       http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
