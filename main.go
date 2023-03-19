package main

import (
	"fmt"

	"github.com/stscoundrel/polylinguist/internal/github"
	"github.com/stscoundrel/polylinguist/stats"
)

func main() {
	accessToken := "TODO"

	settings := stats.Settings{
		IgnoredLanguages: []string{"SCSS", "CSS", "ASL", "HTML"},
		IgnoredRepos: []string{
			"old-norwegian-dictionary",
			"old-norwegian-dictionary-rs",
			"old-norwegian-dictionary-go",
		},
		AliasedLanguages: []stats.LanguageAlias{
			{
				Language: "C",
				Alias:    "C/C++",
			},
			{
				Language: "C++",
				Alias:    "C/C++",
			},
		},
	}

	repositories, err := github.GetRepositories("stscoundrel", "https://api.github.com/graphql", accessToken)

	if err != nil {
		fmt.Println(err)
		return
	}

	stats := stats.GetTopLanguages(repositories, settings)

	for index, language := range stats {
		fmt.Printf("%d. %s - %f - %s \n", index+1, language.Name, language.Percentage, language.Color)
	}

}
