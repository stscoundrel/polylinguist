package github

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type GraphQLResponse struct {
	Data struct {
		User struct {
			Repositories struct {
				PageInfo struct {
					EndCursor string `json:"endCursor"`
				} `json:"pageInfo"`
				Nodes []struct {
					Name      string `json:"name"`
					Languages struct {
						Edges []struct {
							Size int `json:"size"`
							Node struct {
								Color string `json:"color"`
								Name  string `json:"name"`
							} `json:"node"`
						} `json:"edges"`
					} `json:"languages"`
				} `json:"nodes"`
			} `json:"repositories"`
		} `json:"user"`
	} `json:"data"`
}

type Language struct {
	Name  string
	Size  int
	Color string
}

type Repository struct {
	Name      string
	Languages []Language
}

func post(apiUrl string, accessToken string, query []byte) ([]byte, error) {
	req, err := http.NewRequest("POST", apiUrl, bytes.NewBuffer(query))
	if err != nil {
		return []byte{}, err
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", accessToken))
	req.Header.Set("Content-Type", "application/json")

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return []byte{}, err
	}

	body, parseErr := ioutil.ReadAll(resp.Body)

	if parseErr != nil {
		return []byte{}, parseErr
	}

	return body, nil
}

func getQueryResponse(apiUrl string, accessToken string, query []byte) (GraphQLResponse, error) {
	response, err := post(apiUrl, accessToken, query)

	if err != nil {
		return GraphQLResponse{}, err
	}

	var parsedResult GraphQLResponse

	json.Unmarshal(response, &parsedResult)

	return parsedResult, nil
}

func getQuery(username string, cursor string) []byte {
	queryMap := map[string]string{
		"query": fmt.Sprintf(`
		{
			user(login: "%s") {
				repositories(
				ownerAffiliations: OWNER
				isFork: false
				first: 100
				after: %s
				) {
				pageInfo {
					endCursor
				}
				nodes {
					name
					languages(first: 50, orderBy: {field: SIZE, direction: DESC}) {
					edges {
						size
						node {
						color
						name
						}
					}
					}
				}
				}
			}
		}`,
			username,
			cursor,
		),
	}

	jsonResult, err := json.Marshal(queryMap)

	if err != nil {
		fmt.Printf("There was an error marshaling query to JSON %v", err)
	}

	return jsonResult
}

func GetRepositories(username string, apiUrl string, accessToken string) ([]Repository, error) {
	hasMoreResults := true
	results := []GraphQLResponse{}
	queryCursor := "null"

	for hasMoreResults == true {
		query := getQuery(username, queryCursor)
		result, err := getQueryResponse(apiUrl, accessToken, query)

		if err != nil {
			return []Repository{}, err
		}

		results = append(results, result)

		// API returns null when we're at the end. Clumsy check to catch it without confusing golangs typing.
		if len(result.Data.User.Repositories.PageInfo.EndCursor) == 0 {
			hasMoreResults = false
		}

		// Cursor needs quote wrapping when providing non-null value.
		queryCursor = "\"" + result.Data.User.Repositories.PageInfo.EndCursor + "\""
	}

	repositories := []Repository{}

	for _, result := range results {
		for _, rawRepository := range result.Data.User.Repositories.Nodes {
			repository := Repository{rawRepository.Name, []Language{}}

			for _, edge := range rawRepository.Languages.Edges {
				repository.Languages = append(repository.Languages, Language{edge.Node.Name, edge.Size, edge.Node.Color})
			}

			repositories = append(repositories, repository)
		}
	}

	return repositories, nil
}
