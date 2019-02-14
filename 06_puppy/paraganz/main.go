package main

import (
	"fmt"
	"io"
	"os"
)

var out io.Writer = os.Stdout

type puppy struct {
	id     string
	breed  string
	colour string
	value  string
}

type storer interface {
	insert(pu puppy)
	read(pu puppy)
	update(pu puppy)
	delete(pu puppy)
}

func main() {
	fmt.Fprintln(out, "Hallo du schöne Welt!")
}
