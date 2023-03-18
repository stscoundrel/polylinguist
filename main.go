package main

import (
	"fmt"
	"strconv"

	"github.com/stscoundrel/polylinguist/internal/github"
)

func main() {
	accessToken := "TODO"

	result, err := github.GetRepositories("stscoundrel", "https://api.github.com/graphql", accessToken)

	if err != nil {
		fmt.Println(err)
		return
	}

	for _, repository := range result {
		fmt.Println(repository.Name + " - " + strconv.Itoa(len(repository.Languages)))
	}
	fmt.Println(strconv.Itoa(len(result)) + " TOTAL REPOS!")

}
