package main

import (
	"fmt"
)

type s struct {
	name string
	id int64
}

func main() {
	m := s{
		name: "test",
		id:   1,
	}

	formatters := "\t%+v, %v"
	fmt.Printf("\t%#v", "value-16v", m)
	fmt.Printf(formatters, "value+v", m)
}
