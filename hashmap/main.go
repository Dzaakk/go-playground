package main

import "fmt"

type KeyValue struct {
	Key   string
	Value string
}

const bucketSize = 10

type HashMap struct {
	buckets [][]KeyValue
}

func NewHashMap() *HashMap {
	return &HashMap{buckets: make([][]KeyValue, bucketSize)}
}

func hash(key string) int {
	sum := 0
	for _, c := range key {
		sum += int(c)
	}
	return sum % bucketSize
}

func (h *HashMap) Put(key, value string) {
	index := hash(key)
	bucket := h.buckets[index]
	for i, kv := range bucket {
		if kv.Key == key {
			bucket[i].Value = value
			h.buckets[index] = bucket
			return
		}
	}

	h.buckets[index] = append(bucket, KeyValue{Key: key, Value: value})
}

func (h *HashMap) Get(key string) (string, bool) {
	index := hash(key)
	for _, kv := range h.buckets[index] {
		if kv.Key == key {
			return kv.Value, true
		}
	}

	return "", false
}

func main() {
	m := NewHashMap()
	m.Put("name", "Bob")
	m.Put("age", "100")
	m.Put("game", "Valorant")
	m.Put("name", "Updated Name")   // Case: Update existing key
	m.Put("mega", "Test Collision") // Case: Collision

	for i, bucket := range m.buckets {
		fmt.Printf("Bucket[%d]: %v\n", i, bucket)
	}

}
