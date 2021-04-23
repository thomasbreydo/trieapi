## A Trie System

[![Build Status](https://www.travis-ci.com/thomasbreydo/trieapi.svg?token=LeqszHew9vryiRmRN9P8&branch=main)](https://www.travis-ci.com/thomasbreydo/trieapi)

This repository contains

- A unicode-safe implementation of the trie data structure

- A concurrent API server (see [docs/SERVER.md](docs/SERVER.md))

- A CLI to access the API (see [docs/CLI.md](docs/CLI.md))

- Unit tests that are automatically run by Travis CI
  (see [docs/TESTING.md](docs/TESTING.md))

## Repository structure

- [`tries/`](tries): Go subpackage that implements the trie data structure

  - Source code in files called `<name>.go`
  
  - Unit tests in files called `<name>_test.go`

- [`cli/trie/`](cli/trie): Go subpackage that implements the CLI

- [`docs/`](docs): Documentation folder

  - [Server README](docs/SERVER.md)
    
  - [CLI README](docs/CLI.md)
    
  - [Testing README](docs/TESTING.md)

- [`server.go`](server.go): server source code
  
- [`server_test.go`](server_test.go): local server tests

- `*.sh`: convenience scripts

  - [`deploy.sh`](deploy.sh): deploy changes to Google Cloud
  
  - [`run.sh`](run.sh): get an instance of the API running locally 
  
  - [`test.sh`](test.sh): run all unit tests