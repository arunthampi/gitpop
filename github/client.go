package github

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
)

var linksRegexp = regexp.MustCompile("<([^>]+)>;\\s+rel=\"([^i\"]+)\"")

func FetchGithubReposForUser(userName string) []Repo {
	var allRepos = []Repo{}

	repos, nextLink := callGithubAPI(fmt.Sprintf("https://api.github.com/users/%s/repos", userName))

	if len(repos) > 0 {
		allRepos = append(allRepos, repos...)
	}

	if nextLink != "" {
		for currLink := nextLink; currLink != ""; {
			repos, currLink = callGithubAPI(currLink)
			if len(repos) > 0 {
				allRepos = append(allRepos, repos...)
			}
		}
	}

	return allRepos
}

func callGithubAPI(url string) ([]Repo, string) {
	var nextLink string
	var repos = []Repo{}

	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	links := resp.Header["Link"]
	if links != nil && len(links) > 0 {
		linkStrings := linksRegexp.FindAllStringSubmatch(links[0], -1)

		for _, linkStringMatch := range linkStrings {
			rel := linkStringMatch[2]
			if rel == "next" {
				nextLink = linkStringMatch[1]
			}
		}
	}

	if resp.StatusCode == 200 {
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			panic(err)
		}

		json.Unmarshal(body, &repos)
	}

	return repos, nextLink
}
