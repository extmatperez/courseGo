package main

import "fmt"

func main() {

	var myMapSinMake = map[string]int{}
	myMapConMake := make(map[string]string)

	fmt.Println(myMapSinMake)
	fmt.Printf("%T\n", myMapSinMake)
	fmt.Println(myMapConMake)
	fmt.Printf("%T\n", myMapConMake)

	var students = map[string]int{"Matias": 26, "Juan": 25}
	fmt.Println(students["Benjamin"])
	fmt.Println(students["Matias"])
	fmt.Println(students)

	students["Brenda"] = 19
	students["Marcos"] = 22
	fmt.Println(students)
	students["Matias"] = 22
	fmt.Println(students)
	delete(students, "Marcos")
	fmt.Println(students)

	for key, element := range students {
		fmt.Println("Key:", key, "=>", "Element:", element)
	}

}
