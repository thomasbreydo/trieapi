## Overview

The trie system comes with a command-line interface that can:

1. [Add a keyword to the trie](#add)

2. [Delete a keywords from the trie](#delete)

3. [Search for a keyword in trie](#search)

4. [Generate completions for a prefix](#complete)

5. [Display the trie](#display)

6. [Clear the trie](#clear)

## Installation

Prerequisites:

- If you don't have Go installed, please follow the instructions on [golang.org/doc/install](https://golang.org/doc/install) before proceeding.

- If you have a command called `trie` installed, please uninstall it before proceeding.

To install the CLI, run the following:

```shell
go get -u github.com/thomasbreydo/trieapi/cli/trie
```

## Usage

You can use the CLI to do any of the following:

1. [Add a keyword to the trie](#add)

2. [Delete a keywords from the trie](#delete)

3. [Search for a keyword in trie](#search)

4. [Generate completions for a prefix](#complete)

5. [Display the trie](#display)

6. [Clear the trie](#clear)

### Add

To add a keyword to the trie, use `add`:

```shell
trie add "<keyword>"
```

If the keyword is already in the trie, the CLI outputs `Keyword (<keyword>) present`. 
Otherwise, it outputs `Keyword (<keyword>) added`.

### Delete

To delete a keyword from the trie, use `delete`:

```shell
trie delete "<keyword>"
```

If the keyword isn't in the trie, the CLI outputs `Keyword (<keyword>) missing`.
Otherwise, it outputs `Keyword (<keyword>) deleted`.

### Search

To search if a keyword is present in the trie, use `search`:

```shell
trie search "<keyword>"
```

If the keyword is in the trie, the CLI outputs `Keyword (<keyword>) found`. Otherwise, 
it outputs `Keyword (<keyword>) not found`.

### Complete

To get all words in the trie that start with a prefix, use `complete`:

```shell
trie complete "<prefix>"
```

The CLI outputs a newline-separated list of words.

_Note, if no words are found, the command's exit status is set to `1`._

### Display

To display the trie, use `display`:

```shell
trie display
```

For example, here is what a trie with the words `app, apple, amazon, amazing, already, fix, find` might look like:
```text
app
│ └le
├mazon
│  └ing
└lready
fix
 └nd
```

_Note, if the trie contains the empty string then the first line will be blank:_
```text

app
│ └le
├mazon
│  └ing
└lready
fix
 └nd
```

### Clear

To delete all words from the trie, use `clear`:

```shell
trie clear
```
