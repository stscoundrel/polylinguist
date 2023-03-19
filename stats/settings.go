package stats

type LanguageAlias struct {
	Language string
	Alias    string
	Color    string
}

type Settings struct {
	IgnoredLanguages []string
	IgnoredRepos     []string
	AliasedLanguages []LanguageAlias
}
