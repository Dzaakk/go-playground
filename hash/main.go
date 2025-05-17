package main

import (
	"fmt"
)

func main() {
	fmt.Println("== Hash Table with Chaining ==")
	ht := NewHashTable(5)
	ht.Put("apple", 100)
	ht.Put("orange", 200)
	ht.Put("grape", 300)
	fmt.Println(ht.Get("orange"))

	fmt.Println("== Hash Table with Open Addressing ==")
	oht := NewOpenAddressingTable(5)
	oht.Put("apple", 100)
	oht.Put("orange", 200)
	oht.Put("grape", 300)
	fmt.Println(oht.Get("grape"))
}
