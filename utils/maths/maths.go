package main

import (
	"os"
	"fmt"
	"strconv"
)

func get_mcm(a int, b int) int {
	return (a * b) / get_mcd(a, b)
}

func get_mcd(a int, b int) int {
	if b == 0 {
		return a
	}
	return get_mcd(b, a % b)
}

func main() {
	a, _ := strconv.Atoi(os.Args[1])
	b, _ := strconv.Atoi(os.Args[2])

	fmt.Println("M.C.D.:", get_mcd(a, b))
	fmt.Println("m.c.m.:", get_mcm(a, b))
}

