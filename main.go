package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/shurcooL/githubv4"
	"golang.org/x/oauth2"
)

var (
	repoFlag = flag.String("repo", "", "URL to github repository")
)

func main() {
	flag.Parse()

	split := strings.Split(*repoFlag, "/")
	if len(split) != 3 || (len(split) >= 1 && split[0] != "github.com") {
		fmt.Println("Incorrect repo")
		os.Exit(1)
	}

	owner := split[1]
	name := split[2]

	src := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: os.Getenv("GITHUB_ACCESS_TOKEN")},
	)
	httpClient := oauth2.NewClient(context.Background(), src)
	client := githubv4.NewClient(httpClient)

	var query struct {
		Repository struct {
			Description string
			HomepageURL string

			Releases struct {
				Nodes []struct {
					Description  string
					Name         string
					URL          githubv4.URI
					PublishedAt  githubv4.DateTime
					IsPrerelease githubv4.Boolean
					Tag          struct {
						Name string
					}

					ReleaseAssets struct {
						Nodes []struct {
							Name        string
							Size        int
							DownloadURL githubv4.URI
							URL         githubv4.URI
						}
						TotalCount int
					} `graphql:"releaseAssets(last: 100)"`
				}
				TotalCount int
			} `graphql:"releases(last: 1)"`
		} `graphql:"repository(owner: $owner, name: $name)"`
	}

	variables := map[string]interface{}{
		"owner": githubv4.String(owner),
		"name":  githubv4.String(name),
	}

	if err := client.Query(context.Background(), &query, variables); err != nil {
		panic(err)
	}

	for _, rel := range query.Repository.Releases.Nodes {
		fmt.Printf("%s: prerelease:%v,\n\t%s\n\n",
			rel.Name, rel.IsPrerelease, rel.URL,
		)

		debName := fmt.Sprintf(`emby-server-deb_%s_amd64.deb`, rel.Tag.Name)
		fmt.Printf("Looking for %s...\n", debName)

		for _, asset := range rel.ReleaseAssets.Nodes {
			if asset.Name == debName {
				fmt.Printf("\t%s: %d\n\tDownload URL: %v\n\n",
					asset.Name, asset.Size, asset.DownloadURL)
			}
		}
	}
}
