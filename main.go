package main

import (
	"fmt"
	"os"
	"sort"

	"github.com/zerobotlabs/gitpop/github"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("usage: gitpop <github username>\n")
		os.Exit(1)
	}

	userName := os.Args[1]
	repos := github.FetchGithubReposForUser(userName)

	if len(repos) > 0 {
		sorter := github.Sorter{SortParam: "stars", Repos: repos}
		sort.Sort(sorter)
		fmt.Printf("No.,Name,URL,Stars,Watchers,Forks,Issues\n")
		for i, repo := range sorter.Repos {
			fmt.Printf("%d,%s,%s,%d,%d,%d,%d\n", i+1, repo.Name, repo.Url, repo.Stars, repo.Watchers, repo.Forks, repo.Issues)
		}
	} else {
		fmt.Printf("no repos found for Github user: %s\n", userName)
		os.Exit(1)
	}
}
