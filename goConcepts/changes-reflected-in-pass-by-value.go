package main

import "fmt"

type myStruct struct {
	i int
}

// Takes an int and increments it, but it doesn't modify the original value
func modifyInt(i int) {
	i++
	_ = i // This line is added to ignore ineffectual assignment error
}

// Takes a float and increments it, but it doesn't modify the original value
func modifyFloat(f float64) {
	f++
	_ = f // This line is added to ignore ineffectual assignment error
}

// Takes a string and appends "goyal" to it, but it doesn't modify the original value
func modifyString(s string) {
	s = s + "goyal"
	_ = s // This line is added to ignore ineffectual assignment error
}

// Takes an array and sets the first element to 1, but it doesn't modify the original value
func modifyArray(a [10]int) {
	a[0] = 1
}

// Takes a slice, increments the first element and appends 2 to it,
// s[0] gets modified, but new value does not get appended into original slice
func modifySlice(s []int) {
	s[0]++
	s = append(s, 3)
	_ = s // This line is added to ignore ineffectual assignment error
}

// Takes a map and modifies it, original map also gets modified
func modifyMap(m map[int]int) {
	m[1] = 2
	m[2] = 2
}

// Takes a struct and increments the value of the "i" field by 1, but it doesn't modify the original value
func modifyStruct(ms myStruct) {
	ms.i++
}

func main() {
	i := 10
	modifyInt(i)   // pass a copy of i to modifyInt function
	fmt.Println(i) // print the original value of i, which is still 10

	f := 10.0
	modifyFloat(f) // pass a copy of f to modifyFloat function
	fmt.Println(f) // print the original value of f, which is still 10.0

	s := "nitin"
	modifyString(s) // pass a copy of s to modifyString function
	fmt.Println(s)  // print the original value of s, which is still "nitin"

	a := [10]int{}
	modifyArray(a) // pass a copy of a to modifyArray function
	fmt.Println(a) // print the original value of a, which is still an array of all zeroes

	sl := []int{1}
	modifySlice(sl) // pass a copy of sl to modifySlice function
	fmt.Println(sl) // print the half modified value of sl, which is [2]

	m := map[int]int{1: 1}
	modifyMap(m)   // pass a ref of m to modifyMap function
	fmt.Println(m) // print the modified value of m, which is [1:2 2:2]

	ms := myStruct{i: 1}
	modifyStruct(ms) // pass a copy of ms to modifyStruct function
	fmt.Println(ms)  // print the original value of ms, which is still a myStruct with i=1
}
