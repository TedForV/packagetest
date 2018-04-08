package main

import (
	"fmt"
	"testing"
)

func TestCommon1(t *testing.T) {
	i := -1
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	fmt.Println(i)
}
