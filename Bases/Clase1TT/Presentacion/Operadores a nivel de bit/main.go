package main

import "fmt"

func main() {
	var a uint = 60 /* 60 = 0011 1100 */
	var b uint = 13 /* 13 = 0000 1101 */
	var c uint = 0

	c = a & b /* 12 = 0000 1100 */
	fmt.Printf("Conjunción - El valor de c es %d\n", c)

	c = a | b /* 61 = 0011 1101 */
	fmt.Printf("Disyunción - El valor de c es %d\n", c)

	c = a ^ b /* 49 = 0011 0001 */
	fmt.Printf("Disyunción exclusiva - El valor de c es %d\n", c)

	c = a << 2 /* 240 = 1111 0000 */
	fmt.Printf("Corrimiento de bits a la izquierda - El valor de c es %d\n", c)

	c = a >> 2 /* 15 = 0000 1111 */
	fmt.Printf("Corrimiento de bits a la derecha - El valor de c es %d\n", c)
}
