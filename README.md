# Polylinguist

More exact language usage stats from Github profile.

Features:
- No limits like only "top 10 languages"
- Both public & private repos via access token
- Aliasing similar languages to same stats. For example, combine C & C++ to C/C++
- Ignoring languages. For example, drop markup like SCSS, HTML, Twig etc.
- Ignoring repositories. For example, drop a repository that contains megabytes of generated code, skewing statistics.

## Install

`go get github.com/stscoundrel/polylinguist`


## Usage

Polylinguist exposes a functions for getting Github language stats by username & access token.

With default settings:

```go
package main

import (
    "fmt"

	"github.com/stscoundrel/polylinguist"
)

accessToken := "YOUR_ACCESS_TOKEN"

// Fetches all languages from all repos.
stats, err := polylinguist.GetTopLanguages("YOUR_USERNAME", accessToken)

if err != nil {
    // However you want to deal with your errors.
    // Most likely cause: failed network, failed auth.
}

for index, language := range stats {
    fmt.Printf("%d. %s - %f - %s \n", index+1, language.Name, language.Percentage, language.Color)
}
```

With custom settings:

```go
package main

import (
    "fmt"

	"github.com/stscoundrel/polylinguist"
	"github.com/stscoundrel/polylinguist/stats"
)

accessToken := "YOUR_ACCESS_TOKEN"

// Setup custom settings to skip certain languages & repos
// Also set custom aliases to combine some languages into single statistic.
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

stats, err := polylinguist.GetTopLanguagesWithSettings("YOUR_USERNAME", accessToken, settings)

if err != nil {
    // However you want to deal with your errors.
}

for index, language := range stats {
    fmt.Printf("%d. %s - %f - %s \n", index+1, language.Name, language.Percentage, language.Color)
}
```

Output will be something to the effect of:

```
1. TypeScript - 37.237267 - #3178c6 
2. JavaScript - 28.081825 - #f1e05a 
3. Go - 9.608068 - #00ADD8 
4. Rust - 5.951985 - #dea584 
5. C# - 5.426030 - #178600 
6. PHP - 4.735421 - #4F5D95 
7. Python - 3.653692 - #3572A5 
8. Java - 2.354974 - #b07219 
9. Nim - 1.370613 - #ffc200 
10. Scala - 0.892341 - #c22d40 
11. C/C++ - 0.359692 - #555555 
12. Kotlin - 0.223640 - #A97BFF 
13. F# - 0.076285 - #b845fc 
14. Dockerfile - 0.027167 - #384d54 
15. Shell - 0.001000 - #89e051
```

Up to you what kind of graphic or chart you want to produce with the data. The data includes Github language colors for visual portions.

### As Vercel Cloud Function

For a template on usage as a cloud function on Vercel, see [this repo](https://github.com/stscoundrel/polylinguist-vercel)
