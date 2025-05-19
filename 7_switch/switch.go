package main

func main() {
	// i := 2

	// // simple switch
	// switch i {
	// case 1:
	// 	fmt.Println(1)
	// 	break
	// case 2:
	// 	print(2)
	// 	break
	// default:
	// 	print("Infinte")
	// }

	// // Multiple condition switch

	// switch time.Now().Weekday() {
	// case time.Saturday, time.Sunday:
	// 	println("Weekend")
	// default:
	// 	print("Weekday")
	// }

	//  type switch
	whoAmI := func(i interface{}) {
		switch t := i.(type) {
		case int:
			println("Integer")
		case string:
			println("String")
		case bool:
			println("Boolean")
		default:
			println("Other", t)
		}
	}

	whoAmI(55.9)
}
