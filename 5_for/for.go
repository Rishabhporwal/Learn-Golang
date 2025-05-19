package main

// for -> only construct in go for looping
func main() {

	// // while loop
	// i := 1
	// for i <= 3 {
	// 	fmt.Println(i)
	// 	i++
	// }

	// // infinite loop

	// for {
	// 	println("1")
	// }

	// // Classic Loop
	// for i := 0; i <= 3; i++ {
	// 	if i == 2 {
	// 		continue
	// 	}
	// 	fmt.Println(i)
	// 	break
	// }

	for i := range 3 {
		println(i)
	}
}
