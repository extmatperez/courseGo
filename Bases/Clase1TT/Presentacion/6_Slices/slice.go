package main

import "fmt"

func main() {
	primes := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes[1:4]) // Si no ponemos un valor después de los : toma hasta el fin de elementos del slice y viceversa.
	fmt.Printf("%T\n", primes)

	var s = []bool{true, false, false}
	fmt.Println(s[0])

	a := make([]int, 10)
	a[0] = 51
	a = append(a, 5)
	fmt.Println(a, len(a), cap(a))

	nombresArray := [5]string{"Ana", "José", "Daniel", "María", "Carlos"}
	nombresSlice := nombresArray[0:3]
	fmt.Println(nombresArray)
	fmt.Println(nombresSlice)                                             // [Ana José Daniel]
	fmt.Printf("len %d - cap %d\n", len(nombresSlice), cap(nombresSlice)) // len 3 - cap 5

	nombresSlice = append(nombresSlice, "Antonio")
	fmt.Println(nombresSlice)
	fmt.Printf("len %d - cap %d\n", len(nombresSlice), cap(nombresSlice)) // len 4 - cap 5

	nombresSlice = append(nombresSlice, "Daniela")
	fmt.Println(nombresSlice)
	fmt.Printf("len %d - cap %d\n", len(nombresSlice), cap(nombresSlice)) // len 5 - cap 5

	nombresSlice = append(nombresSlice, "Matias")
	fmt.Println(nombresSlice)
	fmt.Printf("len %d - cap %d\n", len(nombresSlice), cap(nombresSlice))

	nombresSlice = append(nombresSlice, "Agustin", "Santiago", "Samira", "Luca", "Rene", "Juan")
	fmt.Println(nombresSlice)
	fmt.Printf("len %d - cap %d\n", len(nombresSlice), cap(nombresSlice))

	fmt.Println(nombresSlice) // [Ana José Daniel Antonio Daniela Carmen]
}
