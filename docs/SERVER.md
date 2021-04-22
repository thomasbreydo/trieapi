## Server Overview

### How is the server hosted?

The server is hosted for free on Google Cloud. See [Accessing the Trie](#Accessing-the-Trie)
below to see how the API can be accessed via the CLI or via curl.

### Is the server concurrent?

Yes! The server instance uses `http.ListenAndServe`, which is concurrent.
Behind the scenes, a connection queue maintains the integrity of the
order of requested operations across multiple clients.

### Where is the trie stored?

The trie is stored in one global state managed by the server
instance. The trie is protected from race conditions by a
mutex (`sync.Mutex`).

### How does the CLI interact with the server?

TODO
  
## Accessing the Trie

### Using the CLI

See [docs/CLI.md](CLI.md) for info on how to install and use
a command-line interface to access the API.

### Using curl

TEST
