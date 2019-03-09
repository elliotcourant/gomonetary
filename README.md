# gomonetary

[![Build Status](https://travis-ci.com/elliotcourant/gomonetary.svg?branch=master)](https://travis-ci.com/elliotcourant/gomonetary)
[![codecov](https://codecov.io/gh/elliotcourant/gomonetary/branch/master/graph/badge.svg)](https://codecov.io/gh/elliotcourant/gomonetary)

gomonetary is a pure go implementation to parse and format
currency text. Given a locale it can take a float and format
it according to that locale. It can also take a formatted
string from a locale and parse it into a float.

### Cache

Monetary information for each locale is cached, this allows for better
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