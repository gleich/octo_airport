package main

import (
	"fmt"

	"github.com/Matt-Gleich/octo_airport/pkg/config"
	"github.com/Matt-Gleich/octo_airport/pkg/github"
)

func main() {
	configuration := config.Get()
	data := github.GetData(configuration)
	fmt.Println(data)
}
