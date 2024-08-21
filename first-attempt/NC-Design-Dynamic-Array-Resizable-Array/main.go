package main

import "fmt"

func main() {
	value := []interface{}{"Array", 1, "getSize", "getCapacity"}

	DynamicArray(value)
}

func DynamicArray(values []interface{}) {

	fmt.Println(values[0])
	fmt.Println(values[1])

}
