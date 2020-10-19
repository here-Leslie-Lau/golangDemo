package main

import (
	"fmt"
	"testing"
)

func TestAdd(t *testing.T) {
	add := Add(1, 2)

	if add != 3 {
		t.Error("error!!!")
	} else {
		fmt.Println("right!!!")
	}
}
