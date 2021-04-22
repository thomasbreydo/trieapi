## Overview

The trie system comes with a command-line interface that can:

1. [Add a keyword to the trie](#add)

2. [Delete keywords from the trie](#delete)

3. [Search for a keyword in trie](#search)

4. [Generate completions for a prefix](#complete)

5. [Display the trie](#display)

6. [Clear the trie](#clear)

## Installation

Prerequisites:

- If you don't have Go installed, please follow the instructions on [golang.org/doc/install](https://golang.org/doc/install) before proceeding.

- If you have a command called `trie` installed, please uninstall it before proceeding.

To install the CLI, run the following:

```text
go get -u https://github.com/thomasbreydo/trieapi/cli/trie
```

## Usage

You can use the CLI to do any of the following:

1. [Add a keyword to the trie](#add)

2. [Delete keywords from the trie](#delete)

3. [Search for a keyword in trie](#search)

4. [Generate completions for a prefix](#complete)

5. [Display the trie](#display)

6. [Clear the trie](#clear)

### Add

To add a keyword to the trie, use `add`:

```text
trie add --word <keyword>
```

or

```text
trie add -w <keyword>
```

If the keyword is already in the trie, the CLI outputs `Keyword (<keyword>) present`. 
Otherwise, it outputs `Keyword (<keyword>) added`.

### Delete

To delete a keyword from the trie, use `delete`:

```text
trie delete --word <keyword>
```

or

```text
trie delete -w <keyword>
```

If the keyword isn't in the trie, the CLI outputs `Keyword (<keyword>) missing`.
Otherwise, it outputs `Keyword (<keyword>) deleted`.

### Search

To search if a keyword is present in the trie, use `search`:

```text
trie search --word <keyword>
```

or

```text
trie search -w <keyword>
```

If the keyword is in the trie, the CLI outputs `Keyword (<keyword>) found`. Otherwise, 
it outputs `Keyword (<keyword>) not found`.

### Complete

To get all keywords in the trie that start with a prefix, use `complete`:

```text
trie complete --word <prefix>
```

or

```text
trie complete -w <prefix>
```

The CLI outputs a newline-separated list of keywords.

_Note, if no keywords are found, the command's exit status is set to `1`._

### Display

To display the trie, use `display`:

```text
trie display
```

For example, here is what a trie with the keywords `app, apple, amazon, amazing, already, fix, find` might look like:
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

To delete all keywords from the trie, use `clear`:

```text
trie clear
```

## Edge cases

### The empty string

To set `-w`, `--word` or `--prefix` to be the empty string, do the following:

```text
trie add --word ""
```


### Strings with spaces

When using the shell, spaces must be escaped. For example, 
if you want `-w`, `--word` or `--prefix` to be `a b'`, you must
do the following:

```text
trie add -word "a b'"
```

### Strings with non-ASCII characters

The CLI URL-encodes all keywords to ensure Unicode compatibility.