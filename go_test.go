package main

import (
	"fmt"
	"testing"
)

type Test struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func TestToken(t *testing.T) {
	var a  [6]*Test

	fmt.Printf("%#v", len(a))
}
