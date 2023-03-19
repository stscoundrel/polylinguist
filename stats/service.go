package stats

import (
	"sort"

	"github.com/stscoundrel/polylinguist/internal/github"
)

type LanguageStatistic struct {
	Name       string
	Percentage float64
	Color      string
}

func filterRepositories(repositories []github.Repository, settings Settings) []github.Repository {
	filtered := []github.Repository{}

	for _, repository := range repositories {
		if !isInSlice(repository.Name, settings.IgnoredRepos) {
			filteredLanguages := []github.Language{}

			for _, language := range repository.Languages {
				if !isInSlice(language.Name, settings.IgnoredLanguages) {
					filteredLanguages = append(filteredLanguages, language)
				}
			}

			repository.Languages = filteredLanguages
			filtered = append(filtered, repository)
		}
	}

	return filtered
}

func getLanguagesMap(repositories []github.Repository, settings Settings) map[string]int {
	languagesBySize := map[string]int{}

	for _, repository := range repositories {
		for _, language := range repository.Languages {
			languagesBySize[language.Name] = languagesBySize[language.Name] + language.Size
		}
	}

	return languagesBySize
}

func getColorsMap(repositories []github.Repository, settings Settings) map[string]string {
	colors := map[string]string{}

	for _, repository := range repositories {
		for _, language := range repository.Languages {
			colors[language.Name] = language.Color
		}
	}

	return colors
}

func getTotalSize(sizeMapping map[string]int) int {
	size := 0

	for _, languageSize := range sizeMapping {
		size += languageSize
	}

	return size
}

func getLanguagesStats(sizeMapping map[string]int, colorMapping map[string]string) []LanguageStatistic {
	stats := []LanguageStatistic{}

	totalSize := getTotalSize(sizeMapping)

	for language, size := range sizeMapping {
		stats = append(stats, LanguageStatistic{
			Name:       language,
			Percentage: (float64(size) / float64(totalSize)) * 100.0,
			Color:      colorMapping[language],
		})
	}

	return stats
}

func GetTopLanguages(repositories []github.Repository, settings Settings) []LanguageStatistic {
	repos := filterRepositories(repositories, settings)

	bySize := getLanguagesMap(repos, settings)
	byColor := getColorsMap(repos, settings)
	stats := getLanguagesStats(bySize, byColor)

	sort.Slice(stats, func(i, j int) bool {
		return stats[i].Percentage > stats[j].Percentage
	})

	return stats
}
