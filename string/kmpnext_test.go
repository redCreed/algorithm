package string

import (
	"fmt"
	"testing"
)

func TestKmpNext(t *testing.T) {
	KmpNext("ABA")
}

func TestKmpSearch(t *testing.T) {
	str1 :=  "124 43 224"
	str2 := "22"
	next := KmpNext(str2)

	index := KmpSearch(str1,str2, next)
	fmt.Println("index:", index)
}