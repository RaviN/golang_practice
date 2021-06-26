package main

func main() {
	/*
		For Loop
	*/
	/*
		Loop til condition
	*/
	var i int
	for i < 5 {
		println(i)
		i++
		if i == 3 {
			break
		}
	}
	/*
		Loop til condition with post clause
	*/
	for j := 0; j < 5; j++ {
		println(j)
	}
	/*
		looping through slices
	*/
	slice := []int{1, 2, 3}
	for i, v := range slice {
		println(i, v)
	}
	/*
		looping through map
	*/
	wellKnownPorts := map[string]int{"http": 80, "https": 443}
	for _, v := range wellKnownPorts {
		println(v)
	}

	/*
		Panics
	*/
	// panic("something bad happend")

}
