package main

type Pair struct {
	Key   string
	Value int
}

type HashTable struct {
	buckets [][]Pair
	size    int
}

func NewHashTable(size int) *HashTable {
	return &HashTable{
		buckets: make([][]Pair, size),
		size:    size,
	}
}

func (ht *HashTable) hash(key string) int {
	hash := 0
	for _, ch := range key {
		hash += int(ch)
	}
	return hash % ht.size
}

func (ht *HashTable) Put(key string, value int) {
	index := ht.hash(key)

	for i, pair := range ht.buckets[index] {
		if pair.Key == key {
			ht.buckets[index][i].Value = value
			return
		}
	}

	ht.buckets[index] = append(ht.buckets[index], Pair{key, value})
}

func (ht *HashTable) Get(key string) (int, bool) {
	index := ht.hash(key)
	for _, pair := range ht.buckets[index] {
		if pair.Key == key {
			return pair.Value, true
		}
	}

	return 0, false
}
