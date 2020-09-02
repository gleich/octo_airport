package github

import (
	"context"

	"github.com/Matt-Gleich/statuser/v2"
	"github.com/nathan-fiscaletti/consolesize-go"
	"github.com/shurcooL/githubv4"
)

type GraphQLQuery struct {
	User struct {
		Repositories struct {
			Edges []struct {
				Node struct {
					Name      githubv4.String
					UpdatedAt githubv4.DateTime
					Owner     struct {
						URL githubv4.String
					}
					Description githubv4.String
					Issues      struct {
						TotalCount githubv4.Int
					} `graphql:"issues(states: OPEN)"`
					PrimaryLanguage struct {
						Name  githubv4.String
						Color githubv4.String
					}
					ForkCount    githubv4.Int
					IsFork       githubv4.Boolean
					PullRequests struct {
						TotalCount githubv4.Int
					} `graphql:"pullRequests(states: OPEN)"`
					LicenseInfo struct {
						Name githubv4.String
					}
					Stargazers struct {
						TotalCount githubv4.Int
					}
				}
			}
		} `graphql:"repositories(first: $repoAmount, orderBy: {field: UPDATED_AT, direction: DESC}, ownerAffiliations: OWNER)"`
	} `graphql:"user(login: $username)"`
}

func GetGraphQLData(githubv4Client *githubv4.Client, username string) GraphQLQuery {
	_, rows := consolesize.GetConsoleSize()
	// Capping line number at 80
	if rows > 80 {
		rows = 80
	}
	variables := map[string]interface{}{
		"username":   githubv4.String(username),
		"repoAmount": githubv4.Int(rows),
	}
	var data GraphQLQuery
	err := githubv4Client.Query(context.Background(), &data, variables)
	if err != nil {
		statuser.Error("Failed to get general GraphQL data", err, 1)
	}
	return data
}
