# gomonetary

[![Build Status](https://travis-ci.com/elliotcourant/gomonetary.svg?branch=master)](https://travis-ci.com/elliotcourant/gomonetary)
[![codecov](https://codecov.io/gh/elliotcourant/gomonetary/branch/master/graph/badge.svg)](https://codecov.io/gh/elliotcourant/gomonetary)

gomonetary is a pure go implementation to parse and format
currency text. Given a locale it can take a float and format
it according to that locale. It can also take a formatted
string from a locale and parse it into a float.

## Example

```go
package main

import (
	"fmt"
	"github.com/elliotcourant/gomonetary"
)

func main() {
	/*
                        Formatting
	*/
	usDollars, _ := monetary.Format(123.56, "en_US")
	fmt.Println("Formatted Value:", usDollars)
	// Output: Formatted Value: $123.56

	ruExample, _ := monetary.Format(5438.98, "ru_RU")
	fmt.Println("Formatted Value:", ruExample)
	// Output: Formatted Value: 5.438,98 руб.



	/*
                        Parsing
	*/
	usParsed, _ := monetary.Parse(usDollars, "en_US")
	fmt.Println("Parsed Value:", usParsed)
	// Output: Parsed Value: 123.56

	ruParsed, _ := monetary.Parse(ruExample, "ru_RU")
	fmt.Println("Parsed Value:", ruParsed)
	// Output: Parsed Value: 5438.98
}
```

## Cache

Monetary information for each locale is cached ([generated.go](generated.go) - cache file), this allows for better
performance but also guarantees support for all OS's for using the
package.

Currently, only Linux and Mac OS X can generate monetary information for
the cache. This uses the unix command `locale` and the data from the host
operating system to populate the cache.

To regenerate the cache on a supported OS run: 

```bash
make metadata
```

This will replace the current cache (I will make it additive in the future)
with all of the locale information on your system. Note: On Linux you might
need to add other locales because they are not typically included be default.
Mac OS usually has a couple hundred installed by default.