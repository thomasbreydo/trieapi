## Testing Overview

Unit tests are written in Go with the `testing` package.
Tests are automatically run by Travis CI whenever changes
are pushed.

## Running tests

To run all tests, run `./test.sh` from main directory.

## Test Descriptions

### Server and Global State Tests

The server is implemented in [../server.go]() so the testing
suite is lives in [../server_test.go]()

### Trie Data Structure Tests

The trie data structure is implemented in [../tries/]() and
unit tests are in files that follow the naming scheme
`<name>_test.go`.