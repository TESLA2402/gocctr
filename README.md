# gocctr
[![License](https://img.shields.io/badge/license-MIT-blue.svg)](https://opensource.org/licenses/MIT)

## Overview

The mytrtool is a command-line tool written in Go that provides functionality similar to the UNIX tr command. It supports various operations like character translation, deletion, and squeezing based on user-defined rules.

## Features

- Translate characters from one set to another
- Support for character ranges and character classes
- Delete specified characters from the input
- Squeeze repeated occurrences of characters into a single instance

## Usage

Get started with mytrtool by providing substitution rules and specifying additional options to customize the translation process.

The following options are supported:

- `-d`: Delete mode. Delete characters specified in the substitution rule from the input.
- `-s`: Squeeze mode. Squeeze multiple occurrences of characters listed in the last operand into a single instance.
- `rule`: The substitution rule. Specify two characters or a character class separated by a space. For example:
  - `A-Z a-z`: Translate uppercase to lowercase.
  - `[:upper:] [:lower:]`: Use character classes for translation.
- The functional requirements for tr are concisely described by it’s man page - give it a go in your local terminal now:
  ```bash
  % man tr
  ```

**Character Classes:**

- `[:alnum:]`: Alphanumeric characters (letters and digits).
- `[:alpha:]`: Alphabetic characters (letters).
- `[:blank:]`: Whitespace characters.
- `[:cntrl:]`: Control characters.
- `[:digit:]`: Numeric characters (digits).
- `[:lower:]`: Lowercase alphabetic characters.
- `[:print:]`: Printable characters.
- `[:punct:]`: Punctuation characters.
- `[:rune:]`: Valid Unicode characters.
- `[:space:]`: Space characters.
- `[:special:]`: Special characters.
- `[:upper:]`: Uppercase alphabetic characters.

## Installation

Clone the repository and use the provided `build.sh` script to build mycurl. Ensure that you have Go installed on your machine.

```bash
git clone https://github.com/TESLA2402/gocctr.git
cd tr
chmod +x build.sh
./build.sh

```

### Examples:

```bash
# If build script implemented

# Translate C -> c
mytrtool C c

# Translate uppercase to lowercase
head -n3 test.txt | mytrtool A-Z a-z

# Delete characters specified in the substitution rule
head -n3 test.txt | mytrtool -d War

# Squeeze multiple occurrences of characters
mytrtool -s AB

# Use character classes for translation
mytrtool "[:upper:]" "[:lower:]"
head -n3 test.txt | mytrtool "[:upper:]" "[:lower:]"

# Test for large amount of input
seq 1 3000 | xargs -Inone cat test.txt | mytrtool "[:upper:]" "[:lower:]" > result.txt

 OR

# Translate C -> c
go run cctr.go C c

# Translate uppercase to lowercase
head -n3 test.txt | go run cctr.go A-Z a-z

# Delete characters specified in the substitution rule
head -n3 test.txt | go run cctr.go -d War

# Squeeze multiple occurrences of characters
go run cctr.go -s AB

# Use character classes for translation
go run cctr.go "[:upper:]" "[:lower:]"
head -n3 test.txt | go run cctr.go "[:upper:]" "[:lower:]"

# Test for large amount of input
seq 1 3000 | xargs -Inone cat test.txt | go run cctr.go "[:upper:]" "[:lower:]" > result.txt
```

