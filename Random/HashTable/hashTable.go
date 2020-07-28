package main

import "fmt"

type keyValue struct {
	key   int
	value int
	next  *keyValue
}

type hashTable struct {
	values [7]*keyValue
}

func (ht *hashTable) SeperateChaining(key, value int) {
	node := &keyValue{key: key, value: value, next: nil}
	idx := key % 7
	elementAtIdx := ht.values[idx]
	if elementAtIdx != nil {
		elementAtIdx.next = node
		return
	}
	ht.values[idx] = node
}

func main() {
	ht1 := hashTable{}
	ht1.SeperateChaining(1, 11)
	ht1.SeperateChaining(2, 22)
	ht1.SeperateChaining(3, 33)
	ht1.SeperateChaining(4, 44)
	ht1.SeperateChaining(5, 55)
	ht1.SeperateChaining(6, 66)
	ht1.SeperateChaining(6, 66)
	fmt.Println(ht1.values)
	fmt.Println(ht1.values[6].value, ht1.values[6].next)
}
