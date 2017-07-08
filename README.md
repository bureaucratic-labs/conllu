# CoNLL-U [![GoDoc](https://godoc.org/github.com/bureaucratic-labs/conllu?status.svg)](https://godoc.org/github.com/bureaucratic-labs/conllu)
CoNLL-U parser for Go 

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
