package github

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetIssues(t *testing.T) {
	fixtureResponse1, _ := os.Open("../../fixtures/sample_response_1.json")
	fixtureBody1, _ := ioutil.ReadAll(fixtureResponse1)

	fixtureResponse2, _ := os.Open("../../fixtures/sample_response_2.json")
	fixtureBody2, _ := ioutil.ReadAll(fixtureResponse2)

	fixtureResponse3, _ := os.Open("../../fixtures/sample_response_3.json")
	fixtureBody3, _ := ioutil.ReadAll(fixtureResponse3)

	defer fixtureResponse1.Close()
	defer fixtureResponse2.Close()
	defer fixtureResponse3.Close()

	// Manually validated expected queries.
	// Main difference: pagination by endCursor.
	expectedQuery1 := `{"query":"\n\t\t{\n\t\t\tuser(login: \"stscoundrel\") {\n\t\t\t\trepositories(\n\t\t\t\townerAffiliations: OWNER\n\t\t\t\tisFork: false\n\t\t\t\tfirst: 100\n\t\t\t\tafter: null\n\t\t\t\t) {\n\t\t\t\tpageInfo {\n\t\t\t\t\tendCursor\n\t\t\t\t}\n\t\t\t\tnodes {\n\t\t\t\t\tname\n\t\t\t\t\tlanguages(first: 50, orderBy: {field: SIZE, direction: DESC}) {\n\t\t\t\t\tedges {\n\t\t\t\t\t\tsize\n\t\t\t\t\t\tnode {\n\t\t\t\t\t\tcolor\n\t\t\t\t\t\tname\n\t\t\t\t\t\t}\n\t\t\t\t\t}\n\t\t\t\t\t}\n\t\t\t\t}\n\t\t\t\t}\n\t\t\t}\n\t\t}"}`
	expectedQuery2 := `{"query":"\n\t\t{\n\t\t\tuser(login: \"stscoundrel\") {\n\t\t\t\trepositories(\n\t\t\t\townerAffiliations: OWNER\n\t\t\t\tisFork: false\n\t\t\t\tfirst: 100\n\t\t\t\tafter: \"Y3Vyc29yOnYyOpHOHifEvg==\"\n\t\t\t\t) {\n\t\t\t\tpageInfo {\n\t\t\t\t\tendCursor\n\t\t\t\t}\n\t\t\t\tnodes {\n\t\t\t\t\tname\n\t\t\t\t\tlanguages(first: 50, orderBy: {field: SIZE, direction: DESC}) {\n\t\t\t\t\tedges {\n\t\t\t\t\t\tsize\n\t\t\t\t\t\tnode {\n\t\t\t\t\t\tcolor\n\t\t\t\t\t\tname\n\t\t\t\t\t\t}\n\t\t\t\t\t}\n\t\t\t\t\t}\n\t\t\t\t}\n\t\t\t\t}\n\t\t\t}\n\t\t}"}`
	expectedQuery3 := `{"query":"\n\t\t{\n\t\t\tuser(login: \"stscoundrel\") {\n\t\t\t\trepositories(\n\t\t\t\townerAffiliations: OWNER\n\t\t\t\tisFork: false\n\t\t\t\tfirst: 100\n\t\t\t\tafter: \"Y3Vyc29yOnYyOpHOIyAJow==\"\n\t\t\t\t) {\n\t\t\t\tpageInfo {\n\t\t\t\t\tendCursor\n\t\t\t\t}\n\t\t\t\tnodes {\n\t\t\t\t\tname\n\t\t\t\t\tlanguages(first: 50, orderBy: {field: SIZE, direction: DESC}) {\n\t\t\t\t\tedges {\n\t\t\t\t\t\tsize\n\t\t\t\t\t\tnode {\n\t\t\t\t\t\tcolor\n\t\t\t\t\t\tname\n\t\t\t\t\t\t}\n\t\t\t\t\t}\n\t\t\t\t\t}\n\t\t\t\t}\n\t\t\t\t}\n\t\t\t}\n\t\t}"}`

	responses := map[int][]byte{
		1: fixtureBody1,
		2: fixtureBody2,
		3: fixtureBody3,
	}

	queries := map[int]string{
		1: expectedQuery1,
		2: expectedQuery2,
		3: expectedQuery3,
	}

	responseCount := 0

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, request *http.Request) {
		body, _ := io.ReadAll(request.Body)

		responseCount += 1
		assert.Equal(t, queries[responseCount], string(body))
		w.WriteHeader(http.StatusOK)
		w.Write(responses[responseCount])
	}))
	defer server.Close()

	result, _ := GetRepositories("stscoundrel", server.URL, "dummy-api-key")

	fmt.Println(result[89])

	// Should've gathered repos from the tree responses.
	// 1. Full response (=more to be fetched)
	// 2. Non full response (=no more to be fetched)
	// 3. Response that confirms there are no more to be paginated.
	assert.Equal(t, 127, len(result))

	// Randomly picked content.
	expectedRepository := Repository{
		"kven-norwegian-dictionary-builder",
		[]Language{
			Language{"Go", 9566, "#00ADD8"},
		},
	}

	assert.Equal(t, expectedRepository, result[89])

}
