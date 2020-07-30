package main

import "fmt"

type node struct {
	key   int
	value int
	taken bool
}

type hashTable struct {
	arr         []node
	valuesTaken int
}

func (ht *hashTable) hashFn(key, value int) bool {
	count := 0
	idx := 0
	for count != len(ht.arr) {
		idx = (key + count) % len(ht.arr)
		if ht.arr[idx].taken == false {
			ht.arr[idx] = node{key: key, value: value, taken: true}
			ht.valuesTaken++
			return true
		}

		count++
	}
	return false
}

func (ht *hashTable) linearProbing(key, value int) {

	if float32(ht.valuesTaken/len(ht.arr)) >= 0.75 {
		var tempHT hashTable
		tempHT.arr = append(ht.arr, node{0, 0, false}, node{0, 0, false}, node{0, 0, false}, node{0, 0, false}, node{0, 0, false}, node{0, 0, false}, node{0, 0, false}, node{0, 0, false}, node{0, 0, false}, node{0, 0, false})
		for i, element := range ht.arr {
			// fmt.Println(element)
			tempHT.arr[i] = node{0, 0, false}
			tempHT.hashFn(element.key, element.value)
		}
		ht.arr = tempHT.arr
	}

	// fmt.Println("Aa", ht.arr)
	elementInserted := ht.hashFn(key, value)
	if !elementInserted {
		fmt.Println("Not enough space")
	}
}

func (ht *hashTable) linearProbingSearching(key int) (int, int) {
	if ht.valuesTaken == 0 {
		fmt.Println("Hashtable empty")
		return -1, -1
	}
	idx := key % len(ht.arr)
	counter := 0
	for ht.arr[idx].taken != false || len(ht.arr) != counter {
		if ht.arr[idx].key == key {
			return ht.arr[idx].value, idx
		}
		counter++
	}
	fmt.Println("Element not found")
	return -1, -1
}

func (ht *hashTable) linearProbingDeletion(key int) int {
	_, i := ht.linearProbingSearching(key)
	if i == -1 {
		fmt.Println("Element not found")
		return -1
	}
	// fmt.Println(n)
	v := ht.arr[i].value
	ht.arr[i] = node{0, 0, false}
	return v
}

func main() {
	ht := hashTable{arr: []node{node{0, 0, false}, node{0, 0, false}, node{0, 0, false}, node{0, 0, false}, node{0, 0, false}}}
	ht.linearProbing(9, 9)
	// fmt.Println(ht.arr, ht.valuesTaken)

	ht.linearProbing(3, 3)
	// fmt.Println(ht.arr, ht.valuesTaken)

	ht.linearProbing(8, 8)
	// fmt.Println(ht.arr, ht.valuesTaken)

	ht.linearProbing(5, 1)
	// fmt.Println(ht.arr, ht.valuesTaken)

	ht.linearProbing(6, 1)
	// fmt.Println(ht.arr, ht.valuesTaken)

	ht.linearProbing(7, 1)
	// fmt.Println(ht.arr, ht.valuesTaken)

	ht.linearProbing(10, 1)
	fmt.Println(ht.arr, ht.valuesTaken)
	n, i := ht.linearProbingSearching(10)
	fmt.Println(n, i)
	// n, i = ht.linearProbingSearching(11)
	n = ht.linearProbingDeletion(10)
	fmt.Println(n, ht.arr)
	ht.linearProbing(10, 1)
	fmt.Println(ht.arr)

}
