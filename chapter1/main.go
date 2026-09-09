package main

import "fmt"

func main() {
	Ex2()
	Ex3()
	Ex4()
	Ex5()
	Ex6()
	Ex8()
	Ex9(3, 4)
	Ex10(10, 0)
	Ex11(10.0, 0.0)
	Ex12()
	Ex13("", 15)
	Ex16(5, 3)
	Ex18(5, 3)

	res, err := Ex19(5, 3, "+")
	if err != nil {
		fmt.Println("Ex19 error:", err)
	} else {
		fmt.Println("Ex19 result:", res)
	}

}
