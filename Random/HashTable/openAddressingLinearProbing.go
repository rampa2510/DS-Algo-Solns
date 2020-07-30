package main

import "fmt"

// some notes
// This is the program for open addressing collision resolution technique
// here every element is added to the hashtable and none of them are linked
// the insertion happens as follow first we check whether the index returned by the hash function has an element in arr or not
// if it is occupied then increment it till we get an index which isnot occupied in arr. If i is incremented till len(arr)
// then the arr is full and exit if the hash function returns index which is not occupied directly add at the index then
// my implemenetation
// the hash table consists of nodes which have 3 parts key,value and taken -> taken indicates whether the index is occupied or not
// hash function here -> (key+i)%len(arr) where i is
// So for insertion first I have checked whether the arr is consumed for 75% if it is then append 10 new elements and run the whole hash function for
// already existing elements and then append the new key,value pair
// the other two methods are same in quadratic probing the counter is squaared
// in double hahsing the hasing taks as follow we take 2 hash functions = h1(x) + i*h2(x)
type node struct {
	key   int
	value int
	taken bool
}

type hashTable struct {
	arr         []node
	valuesTaken int
}

func (ht *hashTable) hashFnAndAppend(key, value int) bool {
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
			tempHT.hashFnAndAppend(element.key, element.value)
		}
		ht.arr = tempHT.arr
	}

	// fmt.Println("Aa", ht.arr)
	elementInserted := ht.hashFnAndAppend(key, value)
	if !elementInserted {
		fmt.Println("Not enough space")
	}
}

// remember the for loop condition
func (ht *hashTable) linearProbingSearching(key int) (int, int) {
	if ht.valuesTaken == 0 {
		fmt.Println("Hashtable empty")
		return -1, -1
	}
	idx := key % len(ht.arr)
	counter := 0
	for (ht.arr[idx].taken == true || (ht.arr[idx].value != -1 && ht.arr[idx].taken == false)) && len(ht.arr) != counter {
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
	ht.arr[i] = node{-1, -1, false}
	ht.valuesTaken--
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
	// fmt.Println(ht.arr, ht.valuesTaken)
	n, i := ht.linearProbingSearching(10)
	// fmt.Println(n, i)
	// n, i = ht.linearProbingSearching(11)
	n = ht.linearProbingDeletion(10)
	// fmt.Println(n, ht.arr)
	ht.linearProbing(11, 1)
	// fmt.Println(ht.arr)
	n, i = ht.linearProbingSearching(11)
	fmt.Println(n, i, ht.arr)
}
