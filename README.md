# gomonetary

[![Build Status](https://travis-ci.com/elliotcourant/gomonetary.svg?branch=master)](https://travis-ci.com/elliotcourant/gomonetary)
[![codecov](https://codecov.io/gh/elliotcourant/gomonetary/branch/master/graph/badge.svg)](https://codecov.io/gh/elliotcourant/gomonetary)

gomonetary is a pure go implementation to parse and format
currency text. Given a locale it can take a float and format
it according to that locale. It can also take a formatted
string from a locale and parse it into a float.

Monetary information for each locale is cached, this allows for better
performance but also guarantees support for all OS's for using the
package.

Currently, only Linux and Mac OS X can generate monetary information for
the cache. This uses the unix command `locale` and the data from the host
operating system to populate the cache.
