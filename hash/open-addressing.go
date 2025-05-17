package main

import "fmt"

type Entry struct {
	Key   string
	Value int
}

type OpenAddressingTable struct {
	table []*Entry
	size  int
}

func NewOpenAddressingTable(size int) *OpenAddressingTable {
	return &OpenAddressingTable{
		table: make([]*Entry, size),
		size:  size,
	}
}

func (ht *OpenAddressingTable) hash(key string) int {
	hash := 0
	for _, ch := range key {
		hash += int(ch)
	}
	return hash % ht.size
}

func (ht *OpenAddressingTable) Put(key string, value int) {
	index := ht.hash(key)
	originalIndex := index

	for {
		if ht.table[index] == nil || ht.table[index].Key == key {
			ht.table[index] = &Entry{key, value}
			return
		}
		index = (index + 1) % ht.size
		if index == originalIndex {
			fmt.Println("Table full")
			return
		}
	}
}

func (ht *OpenAddressingTable) Get(key string) (int, bool) {
	index := ht.hash(key)
	originalIndex := index
	for {
		if ht.table[index] == nil {
			return 0, false
		}
		if ht.table[index].Key == key {
			return ht.table[index].Value, true
		}
		index = (index + 1) % ht.size
		if index == originalIndex {
			break
		}
	}
	return 0, false
}
