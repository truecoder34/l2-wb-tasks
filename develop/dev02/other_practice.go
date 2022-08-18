package main

import "fmt"

func otherP() {
	// mystr := "Welcome to GeeksforGeeks ??????"
	// var mystr2 string = "W"

	// fmt.Printf("%d\n", unsafe.Sizeof(mystr))
	// fmt.Printf("%d\n", unsafe.Sizeof(mystr2))
	// fmt.Printf("%d\n", len(mystr))
	// fmt.Printf("%d\n", len(mystr2))

	// for i, val := range mystr {
	// 	fmt.Printf("%d Bytes = %x Character = %c \n", i, val, val)
	// }

	// address := &mystr
	// fmt.Printf("Address BEFORE : %v\n", &mystr)
	// fmt.Printf("Address BEFORE : %v\n", address)
	// fmt.Printf("Value BEFORE : %v\n", *address)

	// mystr = "ertetWelcome to GeeksforGeeks ??????00000"
	// address = &mystr
	// fmt.Printf("Address AFTER : %v\n", &mystr)
	// fmt.Printf("Address AFTER : %v\n", address)
	// fmt.Printf("Value AFTER : %v\n", *address)

	// fmt.Printf("SUB STRING : %v\n", mystr[3:7])
	// //mystr[3:4]

	// s := "éक्षिaπ囧"
	// for i, rn := range s {
	// 	fmt.Printf("%2v: Bytes = %x 0x%x %v \n", i, rn, rn, string(rn))
	// }
	// fmt.Println(len(s))

	// var r rune = 'δ'
	// s[0] = r

	//. Make new slices.
	s := make([]int, 3, 5)
	fmt.Println(s, len(s), cap(s)) // [0 0 0] 3 5
	s = make([]int, 2)
	fmt.Println(s, len(s), cap(s)) // [0 0] 2 2
	s0 := []int{2, 3, 5}
	fmt.Println(s0, cap(s0)) // [2 3 5] 3
	s1 := append(s0, 7)      // append one element
	fmt.Println(s1, cap(s1)) // [2 3 5 7] 6

	// array
	var numbers [5]int = [5]int{1, 2}
	fmt.Println(numbers)          // [1 2 0 0 0]
	numbers2 := [...]int{1, 2, 3} // длина массива 3
	fmt.Println(numbers2)         // [1 2 3]
	fmt.Println(&numbers2[0])     // [1 2 3]
	fmt.Println(&numbers2[1])     // [1 2 3]
	fmt.Println(&numbers2[2])     // [1 2 3]
}
