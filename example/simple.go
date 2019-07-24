package main

import (
	"fmt"

	"github.com/wshaman/go-jsonschema-generator"
)

type EmbeddedType struct {
	Zoo string
}

type Item struct {
	Value string
}

type ExampleBasic struct {
	Foo bool   `json:"foo"`
	Bar string `json:",omitempty,minLength:1"`
	Qux int8
	Baz []string
	EmbeddedType
	List []Item
}

func main() {
	s := &jsonschema.Document{}
	s.Read(&ExampleBasic{})
	fmt.Println(s)
}
