package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/google/go-github/v52/github"
	"golang.org/x/oauth2"
)

func main() {
	if len(os.Args) < 3 {
		log.Fatalf("Usage: %s <GitHub API token> <organization>", os.Args[0])
	}

	apiToken := os.Args[1]
	organization := os.Args[2]

	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: apiToken},
	)
	tc := oauth2.NewClient(ctx, ts)

	client := github.NewClient(tc)

	opt := &github.RepositoryListByOrgOptions{
		ListOptions: github.ListOptions{PerPage: 10},
	}

	var allRepos []*github.Repository
	for {
		repos, resp, err := client.Repositories.ListByOrg(ctx, organization, opt)
		if err != nil {
			log.Fatalf("Error fetching repositories: %v", err)
		}
		allRepos = append(allRepos, repos...)
		if resp.NextPage == 0 {
			break
		}
		opt.Page = resp.NextPage
	}

	commitAuthors := make(map[string]map[string]struct{})
	oneWeekAgo := time.Now().Add(-7 * 24 * time.Hour)

	for _, repo := range allRepos {
		branches, _, err := client.Repositories.ListBranches(ctx, organization, *repo.Name, &github.BranchListOptions{})
		if err != nil {
			log.Fatalf("Error fetching branches for repository '%s': %v", *repo.Name, err)
		}

		for _, branch := range branches {
			commitOpt := &github.CommitsListOptions{
				Since:       oneWeekAgo,
				SHA:         branch.GetName(),
				ListOptions: github.ListOptions{PerPage: 10},
			}

			for {
				commits, resp, err := client.Repositories.ListCommits(ctx, organization, *repo.Name, commitOpt)
				if err != nil {
					fmt.Printf("Error fetching commits for repository '%s' and branch '%s': %v", *repo.Name, *branch.Name, err)
					break
				}

				for _, commit := range commits {
					if commit.Author != nil && commit.Author.Login != nil {
						author := *commit.Author.Login
						if _, exists := commitAuthors[author]; !exists {
							commitAuthors[author] = make(map[string]struct{})
						}
						commitAuthors[author][*repo.Name] = struct{}{}
					}
				}

				if resp.NextPage == 0 {
					break
				}
				commitOpt.Page = resp.NextPage
			}
		}
	}

	fmt.Println("Users who have committed across all repositories in the organization in the last week:")
	for author, repos := range commitAuthors {
		repoList := make([]string, 0, len(repos))
		for repo := range repos {
			repoList = append(repoList, repo)
		}
		fmt.Printf("- %s, Repositories: %s\n", strings.TrimSpace(author), strings.Join(repoList, ", "))
	}
}
