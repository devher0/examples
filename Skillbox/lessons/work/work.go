package main

import (
	"fmt"
)

func main() {

	shiftDuration := 480
	customerTime := 2
	cashierTime := 4

	clientsNumber := shiftDuration / (customerTime + cashierTime)

	fmt.Println("За смену длиной", shiftDuration, "минут кассир успеет обслужить", clientsNumber, "клиентов.")
}
