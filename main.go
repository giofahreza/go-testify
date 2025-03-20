package main

import (
	"fmt"
	"go-testify/helpers"
)

func main() {
	total := helpers.Sum(1, 2)
	fmt.Println(total)
}
