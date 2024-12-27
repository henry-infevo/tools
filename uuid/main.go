package main

import (
	_ "embed"
	"flag"
	"fmt"
	"io"
	"strings"

	"github.com/google/uuid"
)

//go:embed doc.go
var doc string

func usage() {
	// Extract the content of the /* ... */ comment in doc.go.
	_, after, _ := strings.Cut(doc, "/*\n")
	doc, _, _ := strings.Cut(after, "*/")
	io.WriteString(flag.CommandLine.Output(), doc)
}

func main() {
	var withoutHyphen bool

	flag.BoolVar(&withoutHyphen, "no-hyphen", false, "with hyphen")
	flag.Usage = usage
	flag.Parse()

	id := fmt.Sprintf("%s", uuid.New())
	if withoutHyphen {
		id = strings.ReplaceAll(id, "-", "")
	}

	fmt.Print(id)
}
