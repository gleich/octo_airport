package main

import (
	"context"
	"fmt"

	"github.com/Matt-Gleich/octo_airport/pkg/config"
	"github.com/Matt-Gleich/octo_airport/pkg/github"
	"github.com/Matt-Gleich/statuser/v2"
	"github.com/shurcooL/githubv4"
)

func main() {
	configuration := config.Get()
	client := github.Authenticate(configuration)
	var query struct {
		Viewer struct {
			Login     githubv4.String
			CreatedAt githubv4.DateTime
		}
	}
	err := client.Query(context.Background(), &query, nil)
	if err != nil {
		statuser.Error("Failed :(", err, 1)
	}
	fmt.Println(query.Viewer.CreatedAt)
}
