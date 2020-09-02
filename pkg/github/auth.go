package github

import (
	"context"

	"github.com/Matt-Gleich/octo_airport/pkg/config"
	"github.com/shurcooL/githubv4"
	"golang.org/x/oauth2"
)

func Authenticatev4(configuration config.Outline) *githubv4.Client {
	src := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: configuration.PAT},
	)
	httpClient := oauth2.NewClient(context.Background(), src)
	return githubv4.NewClient(httpClient)
}
