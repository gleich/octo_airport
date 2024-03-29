package github

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/Matt-Gleich/octo_airport/pkg/config"
	"github.com/Matt-Gleich/statuser"
	"github.com/nathan-fiscaletti/consolesize-go"
	"golang.org/x/oauth2"
)

const query = `
{
	user(login: "$username") {
	repositories(
		first: $termHeight
		orderBy: { field: PUSHED_AT, direction: DESC }
		ownerAffiliations: [COLLABORATOR, OWNER]
	) {
		edges {
		node {
			nameWithOwner
			pushedAt
			description
			issues(states: OPEN) {
			totalCount
			}
			primaryLanguage {
			name
			color
			}
			forkCount
			pullRequests(states: OPEN) {
			totalCount
			}
			licenseInfo {
			name
			}
			stargazers {
			totalCount
			}
			isPrivate
			createdAt
			deployments(
			orderBy: { field: CREATED_AT, direction: ASC }
			first: 1
			) {
			edges {
				node {
				state
				}
			}
			}
			defaultBranchRef {
			target {
				... on Commit {
				checkSuites(first: 10) {
					nodes {
					conclusion
					status
					}
				}
				}
			}
			}
		}
		}
	}
	}
}`

// Get the data for the table
func GetData(configuration config.Outline) DataOutline {
	// Creating http client with the PAT (Personal Access Token)
	src := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: configuration.PAT},
	)
	httpClient := oauth2.NewClient(context.Background(), src)

	// Filling in the variables in the graphql query and generating the json data
	filledQuery := strings.ReplaceAll(query, "$username", configuration.Username)
	_, termHeight := consolesize.GetConsoleSize()
	if termHeight > 100 {
		termHeight = 100
	}
	filledQuery = strings.ReplaceAll(filledQuery, "$termHeight", fmt.Sprintf("%v", termHeight))
	jsonData := map[string]string{"query": filledQuery}

	// Creating the request
	jsonValue, _ := json.Marshal(jsonData)
	request, err := http.NewRequest("POST", "https://api.github.com/graphql", bytes.NewBuffer(jsonValue))
	request.Header.Set("Accept", "application/vnd.github.antiope-preview+json") // Required to use the Checks api currently in preview
	if err != nil {
		statuser.Error("Failed to formulate request to GitHub", err, 1)
	}

	// Making the request with the client
	response, err := httpClient.Do(request)
	if err != nil {
		statuser.Error("Failed to request data from GitHub", err, 1)
	}
	defer response.Body.Close()

	binaryData, _ := ioutil.ReadAll(response.Body)
	var data DataOutline
	err = json.Unmarshal(binaryData, &data)
	if err != nil {
		statuser.Error("Failed to parse data from GitHub", err, 1)
	}
	return data
}
