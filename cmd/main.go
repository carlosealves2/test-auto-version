package main

import (
	"fmt"

	"github.com/carlosealves2/test-auto-version/internal/service"
)

func main() {
	fmt.Print("iniciando...")
	fmt.Println("sum:", service.Sum(2, 2))
}
