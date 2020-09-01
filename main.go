package main

import (
	"fmt"

	"github.com/Matt-Gleich/octo_airport/pkg/config"
)

func main() {
	configuration := config.Get()
	fmt.Println(configuration)
}
