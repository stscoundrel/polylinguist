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
				// Should be ignored by language setting.
				github.Language{
					Name:  "CSS",
					Size:  9999999,
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
					Size:  945,
					Color: "",
				},
				// Should be ignored by language setting.
				github.Language{
					Name:  "ASL",
					Size:  9999999,
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
				// Should be combined with C as composite due to alias.
				github.Language{
					Name:  "C++",
					Size:  50,
					Color: "",
				},
			},
		},
		// These repos should be ignored by settings.
		github.Repository{
			Name: "Dummy Repository 5",
			Languages: []github.Language{
				github.Language{
					Name:  "JavaScript",
					Size:  9999,
					Color: "",
				},
			},
		},
		github.Repository{
			Name: "Dummy Repository 6",
			Languages: []github.Language{
				github.Language{
					Name:  "TypeScript",
					Size:  999999,
					Color: "",
				},
				github.Language{
					Name:  "Nim",
					Size:  2000,
					Color: "",
				},
			},
		},
	}

	settings := Settings{
		IgnoredLanguages: []string{"SCSS", "CSS", "ASL", "HTML"},
		IgnoredRepos: []string{
			"Dummy Repository 5",
			"Dummy Repository 6",
		},
		AliasedLanguages: []LanguageAlias{
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

	result := GetTopLanguages(repos, settings)

	// 7 total languages.
	assert.Equal(t, len(result), 7)

	// Go should be the top language.
	assert.Equal(t, "Go", result[0].Name)
	assert.Equal(t, 28.372818839551712, result[0].Percentage)
	assert.Equal(t, 2000, result[0].Size)

	// Composite alias C/C++ should be the least used.
	assert.Equal(t, "C/C++", result[6].Name)
	assert.Equal(t, 1.4186409419775856, result[6].Percentage)
	assert.Equal(t, 100, result[6].Size)

	// Languages with identical share should have identical percentage.
	// Use Java & Kotlin as examples.
	expectedPercentage := 14.257341466874735
	for _, language := range result {
		if language.Name == "Java" || language.Name == "Kotlin" {
			assert.Equal(t, expectedPercentage, language.Percentage)
		}
	}
}
