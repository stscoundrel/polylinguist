package polylinguist

import (
	"fmt"

	"github.com/stscoundrel/polylinguist/internal/github"
	"github.com/stscoundrel/polylinguist/stats"
)

const DEFAULT_GRAPHQL_URL = "https://api.github.com/graphql"

func getLanguageStats(username string, accessToken string, settings stats.Settings) ([]stats.LanguageStatistic, error) {
	repositories, err := github.GetRepositories(username, DEFAULT_GRAPHQL_URL, accessToken)

	if err != nil {
		fmt.Println(err)
		return []stats.LanguageStatistic{}, err
	}

	return stats.GetTopLanguages(repositories, settings), nil
}

func GetTopLanguages(username string, accessToken string) ([]stats.LanguageStatistic, error) {
	return getLanguageStats(username, accessToken, stats.Settings{})
}

func GetTopLanguagesWithSettings(username string, accessToken string, settings stats.Settings) ([]stats.LanguageStatistic, error) {
	return getLanguageStats(username, accessToken, settings)
}
