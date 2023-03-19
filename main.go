package main

import (
	"fmt"

	"github.com/stscoundrel/polylinguist/internal/github"
	"github.com/stscoundrel/polylinguist/stats"
)

func main() {
	accessToken := "TODO"

	repositories, err := github.GetRepositories("stscoundrel", "https://api.github.com/graphql", accessToken)

	if err != nil {
		fmt.Println(err)
		return
	}

	stats := stats.GetTopLanguages(repositories)

	if err != nil {
		fmt.Println(err)
		return
	}

	for index, language := range stats {
		fmt.Printf("%d. %s - %f - %s \n", index+1, language.Name, language.Percentage, language.Color)
	}

}
