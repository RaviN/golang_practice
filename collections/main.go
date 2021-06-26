package main

import "fmt"

func main() {
	/*
		Arrays
	*/
	// size 3 array of int
	var arr [3]int
	arr[0] = 1
	arr[1] = 2
	arr[2] = 3
	fmt.Println(arr)    // [1 2 3]
	fmt.Println(arr[2]) // 3

	// implicit initilization of array
	arr2 := [3]int{1, 2, 3}
	fmt.Println(arr2) // [1 2 3]

	/*
		Slices
	*/
	slice := arr[:]
	arr[1] = 42
	slice[2] = 27
	fmt.Println(arr, slice) // [1 42 27] [1 42 27]

	slice2 := []int{1, 2, 3}
	fmt.Println(slice2) // [1 2 3]
	slice2 = append(slice2, 4, 42, 27)
	fmt.Println(slice2) // [ 1 2 3 4 42 27]

	s3 := slice2[1:]
	s4 := slice2[:2]
	s5 := slice2[1:2]
	fmt.Println(s3, s4, s5) // [2 3 4 42 27] [1 2] [2]

	/*
		Maps
	*/
	m := map[string]int{"foo": 42}
	fmt.Println(m)        // map[foo:42]
	fmt.Println(m["foo"]) // 42

	m["foo"] = 27
	fmt.Println(m) // map[foo:27]

	delete(m, "foo")
	fmt.Println(m) // map[]

	/*
		Struct
	*/
	type user struct {
		ID        int
		FirstName string
		LastName  string
	}

	var u user
	fmt.Println(u) // {0   }

	u.ID = 1
	u.FirstName = "Ravi"
	u.LastName = "N"
	fmt.Println(u.FirstName) // Ravi

	u2 := user{ID: 2, FirstName: "Arthur", LastName: "Dent"}
	fmt.Println(u)  // {1 Ravi N}
	fmt.Println(u2) // {2 Arthur Dent}
}
