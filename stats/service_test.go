package stats

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stscoundrel/polylinguist/internal/github"
)

func TestGetTopLanguages(t *testing.T) {
	repos := []github.Repository{
		github.Repository{
			Name: "Dummy Repository 1",
			Languages: []github.Language{
				github.Language{
					Name:  "TypeScript",
					Size:  995,
					Color: "",
				},
				github.Language{
					Name:  "JavaScript",
					Size:  999,
					Color: "",
				},
				github.Language{
					Name:  "C",
					Size:  50,
					Color: "",
				},
			},
		},
		github.Repository{
			Name: "Dummy Repository 2",
			Languages: []github.Language{
				github.Language{
					Name:  "Go",
					Size:  1000,
					Color: "",
				},
				github.Language{
					Name:  "Rust",
					Size:  995,
					Color: "",
				},
			},
		},
		github.Repository{
			Name: "Dummy Repository 3",
			Languages: []github.Language{
				github.Language{
					Name:  "Kotlin",
					Size:  1005,
					Color: "",
				},
				github.Language{
					Name:  "Java",
					Size:  1005,
					Color: "",
				},
			},
		},
		github.Repository{
			Name: "Dummy Repository 4",
			Languages: []github.Language{
				github.Language{
					Name:  "Go",
					Size:  1000,
					Color: "",
				},
			},
		},
	}

	result := GetTopLanguages(repos)

	// 7 total languages.
	assert.Equal(t, len(result), 7)

	// Go should be the top language.
	assert.Equal(t, "Go", result[0].Name)
	assert.Equal(t, 28.372818839551712, result[0].Percentage)

	// C should be the least used.
	assert.Equal(t, "C", result[6].Name)
	assert.Equal(t, 0.7093204709887928, result[6].Percentage)

	// Languages with identical share should have identical percentage.
	// Use Java & Kotlin as examples.
	expectedPercentage := 14.257341466874735
	for _, language := range result {
		if language.Name == "Java" || language.Name == "Kotlin" {
			assert.Equal(t, expectedPercentage, language.Percentage)
		}
	}
}
