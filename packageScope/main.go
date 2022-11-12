// Golang program to illustrate the usage of
// fmt.Printf() function

// Including the main package
package main

import (
	"examples/packageScope/nestedPackage"
	"fmt"
)

func main() {

	fmt.Println(nestedPackage.AirplaneName)

}
