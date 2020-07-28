package main

import "fmt"

type keyValue struct {
	key   int
	value int
	next  *keyValue
	taken bool
}

type hashTable struct {
	values [7]keyValue
}

func (ht *hashTable) SeperateChaining(key, value int) {
	node := keyValue{key: key, value: value, next: nil, taken: true}
	idx := key % 7
	elementAtIdx := &ht.values[idx] //err the address was added to elementAtIDx not to the original value
	if elementAtIdx.taken {
		for elementAtIdx.next != nil {
			elementAtIdx = elementAtIdx.next
		}
		elementAtIdx.next = &node
		return
	}
	ht.values[idx] = node
}

func (ht *hashTable) searching(key int) int {
	var elementsAtIdx *keyValue
	elementsAtIdx = &ht.values[key%7]
	if elementsAtIdx.taken != true {
		fmt.Printf("No element with %v as a key exists", key)
		return 0
	}
	// for elementAtIdx.n
	for elementsAtIdx != nil {
		if elementsAtIdx.key == key {
			return elementsAtIdx.value
		}
		elementsAtIdx = elementsAtIdx.next
	}
	fmt.Printf("No element with %v as a key exists", key)
	return 0
}

func main() {
	ht1 := hashTable{}
	var n int
	ht1.SeperateChaining(1, 11)
	ht1.SeperateChaining(2, 22)
	ht1.SeperateChaining(3, 33)
	ht1.SeperateChaining(4, 44)
	ht1.SeperateChaining(5, 55)
	ht1.SeperateChaining(6, 66)
	ht1.SeperateChaining(13, 10)
	ht1.SeperateChaining(20, 8)
	fmt.Println(ht1.values)
	n = ht1.searching(2)
	fmt.Println(n)
	n = ht1.searching(7)
	fmt.Println(n)
	n = ht1.searching(6)
	fmt.Println(n)
	n = ht1.searching(13)
	fmt.Println(n)
	n = ht1.searching(20)
	fmt.Println(n)

}
