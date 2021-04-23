## Server Overview

### How is the server hosted?

The server is hosted for free on Google Cloud.

### Is the server concurrent?

Yes! The server instance uses `http.ListenAndServe`, which is concurrent.
Behind the scenes, a new goroutine is created for each connection (in
the order that connections are received).

### Where is the trie stored?

The trie is stored in one global state managed by the server
instance. The trie is protected from race conditions by a
mutex (`sync.Mutex`).

### How does the CLI interact with the server?

The CLI uses the server's API to perform operations. For
details on how to use interact with the server's REST endpoint
using curl, see below.
  
## Accessing the Server's Trie

### Using the CLI

See [docs/CLI.md](CLI.md) for info on how to install and use
a command-line interface to access the API.

### Using curl

#### To add a word:

```text
curl -sL https://trieapi.uk.r.appspot.com/api/v1/add/<word>
```

> :warning: Replace `<word>` with a URL-encoded word or 
> leave it blank to add the empty string.

The server responds with

```json
{"modified":true}
```

or

```json
{"modified":false}
```
depending on whether the trie was modified. (If it wasn't,
the word was already present.)

#### To delete a word:

```text
curl -sL https://trieapi.uk.r.appspot.com/api/v1/delete/<word>
```

> :warning: Replace `<word>` with a URL-encoded word or
> leave it blank to delete the empty string.

The server responds with

```json
{"modified":true}
```

or

```json
{"modified":false}
```
depending on whether the trie was modified. (If it wasn't,
the word was not in the trie.)

#### To search for a word:

```text
curl -sL https://trieapi.uk.r.appspot.com/api/v1/search/<word>
```

> :warning: Replace `<word>` with a URL-encoded word or
> leave it blank to search for the empty string.

The server responds with

```json
{"found":true}
```

or

```json
{"found":false}
```
depending on whether the word was found in the trie.

#### To complete a prefix:

```text
curl -sL https://trieapi.uk.r.appspot.com/api/v1/complete/<prefix>
```

> :warning: Replace `<prefix>` with a URL-encoded prefix or
> leave it blank to complete the empty string.

The server responds with a JSON array of completions, for example:

```json
["amazing", "amazon", "amaze"]
```

#### To display the trie:

```text
curl -sL https://trieapi.uk.r.appspot.com/api/v1/display
```

The server responds with a JSON array of every word in
the trie, for example:

```json
["amazing", "water", "amazon", "whale", "piano", "amaze"]
```

#### To clear the trie:

```text
curl -sL https://trieapi.uk.r.appspot.com/api/v1/clear
```

The server responds with

```json
{"modified":true}
```

or

```json
{"modified":false}
```
depending on whether the trie was modified. (If it wasn't,
the trie was already empty.)
