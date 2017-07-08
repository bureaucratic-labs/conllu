# CoNLL-U [![GoDoc](https://godoc.org/github.com/bureaucratic-labs/conllu?status.svg)](https://godoc.org/github.com/bureaucratic-labs/conllu) [![Build Status](https://travis-ci.org/bureaucratic-labs/conllu.svg?branch=master)](https://travis-ci.org/bureaucratic-labs/conllu)
CoNLL-U format parser for Go language.  
For more info on format, see [Universal Dependencies website](http://universaldependencies.org/format.html) 

# Installation

```bash
$ go get github.com/bureaucratic-labs/conllu
```

# Usage

```go

package main

import (
	"os"
	"fmt"
	"bufio"
	"github.com/bureaucratic-labs/conllu"
)

func main() {
	fd, _ := os.Open('path/to/corpora.conllu')
	rd := bufio.NewReader(fd)

	sentences, _ := conllu.Parse(rd)

	for i := 0; i < count; i++ {
		fmt.Println(fmt.Sprintf("%+v", sentences[i]))
	}
}
```
