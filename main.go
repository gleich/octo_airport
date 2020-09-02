package main

import (
	"fmt"

	"github.com/Matt-Gleich/octo_airport/pkg/config"
	"github.com/Matt-Gleich/octo_airport/pkg/github"
)

func main() {
	configuration := config.Get()
	v4client := github.Authenticatev4(configuration)
	data := github.GetGraphQLData(v4client, configuration.Username)
	fmt.Println(data)
}
