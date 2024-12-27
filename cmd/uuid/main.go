package main

import (
	"flag"
	"fmt"
	"strings"

	"github.com/google/uuid"
)

func main() {
	var withoutHyphen bool

	flag.BoolVar(&withoutHyphen, "no-hyphen", false, "with hyphen")
	flag.Parse()

	id := fmt.Sprintf("%s", uuid.New())
	if withoutHyphen {
		id = strings.ReplaceAll(id, "-", "")
	}

	fmt.Print(id)
}
