package main

import "fmt"

func main() {

	firsNum := 0
	secoNum := 0
	thirNum := 0

	for thirNum < 1000 {

		firsNum = firsNum + 1
		secoNum = secoNum + 1
		thirNum = thirNum + 1

		if firsNum <= 10 {
			fmt.Println("v1", firsNum)
			fmt.Println("v2", secoNum)
			fmt.Println("v3", thirNum)
			continue
		}

		if secoNum <= 100 {
			fmt.Println("v2", secoNum)
			fmt.Println("v3", thirNum)
			continue
		}

		fmt.Println("v3", thirNum)

	}

}
