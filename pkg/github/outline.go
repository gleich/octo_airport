package github

import "time"

type DataOutline struct {
	Data struct {
		User struct {
			Repositories struct {
				Edges []struct {
					Node struct {
						NameWithOwner string    `json:"nameWithOwner"`
						PushedAt      time.Time `json:"pushedAt"`
						Description   string    `json:"description"`
						Issues        struct {
							TotalCount int `json:"totalCount"`
						} `json:"issues"`
						PrimaryLanguage struct {
							Name  string `json:"name"`
							Color string `json:"color"`
						} `json:"primaryLanguage"`
						ForkCount    int `json:"forkCount"`
						PullRequests struct {
							TotalCount int `json:"totalCount"`
						} `json:"pullRequests"`
						LicenseInfo struct {
							Name string `json:"name"`
						} `json:"licenseInfo"`
						Stargazers struct {
							TotalCount int `json:"totalCount"`
						} `json:"stargazers"`
						IsPrivate   bool      `json:"isPrivate"`
						CreatedAt   time.Time `json:"createdAt"`
						Deployments struct {
							Edges []struct {
								Node struct {
									State string `json:"state"`
								} `json:"node"`
							} `json:"edges"`
						} `json:"deployments"`
						DefaultBranchRef struct {
							Target struct {
								CheckSuites struct {
									Nodes []struct {
										Conclusion interface{} `json:"conclusion"`
										Status     string      `json:"status"`
									} `json:"nodes"`
								} `json:"checkSuites"`
							} `json:"target"`
						} `json:"defaultBranchRef"`
					} `json:"node"`
				} `json:"edges"`
			} `json:"repositories"`
		} `json:"user"`
	} `json:"data"`
}
