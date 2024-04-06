package main

import (
	"fmt"
	"pertemuan-kedua/utils"
)

// hello function
func hello(name string) {
	fmt.Println("hallo ", name)
}

// get value on index array
func getArray(items []string, i int) {
	fmt.Println(items[i])
}

// looping slice
func loopingArray(items []string) {
	for i, fruit := range items {
		fmt.Println("buah ke : ", i+1, " adalah ", fruit)
	}
}

// change value
func changeValueOfArray(items []string, index int, newValue string) []string {
	items[index] = newValue
	return items
}

// get Slice with range
func getSlice(items []string, start int, end int) {
	fmt.Println(items[start:end])
}

// buble short
func bubleShort(arr []int) {
	n := len(arr)

	for i := 0; i < n; i++ {
		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}

// linier search

func linearSearch(arr []int, x int) string {

	for i, item := range arr {
		if item == x {
			return fmt.Sprintf("angka %d ada di index ke %d", item, i)
		}

	}
	return "angka tidak ditemukan"

}

// main function
func main() {

	var name string = "aufal"
	var fruits = []string{"apel", "anggur", "jeruk", "pepaya"}
	var number = []int{3, 8, 1, 6, 2, 9}

	hello(name)
	getArray(fruits, 2)
	loopingArray(fruits)
	getSlice(fruits, 1, 3)
	fmt.Println(changeValueOfArray(fruits, 2, "pisang"))
	fmt.Println("array yang belum di urutkan : ", number)
	bubleShort(number)
	fmt.Println("array yang sudah di urutkan : ", number)
	fmt.Println(linearSearch(number, 3))
	fmt.Println(utils.Hello("aufal"))
	utils.Pointer()

}
