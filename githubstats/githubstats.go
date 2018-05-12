package githubstats

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type repoMetadata struct {
	Name       string `json:"name"`
	CloneURL   string `json:"clone_url"`
	CommitsURL string `json:"commits_url"`
}

/*
RetrieveRepoMetadata retrieves repo name, clone url, last commit data and author of last commit.
*/
func RetrieveRepoMetadata(repo string) (m map[string]string) {
	//https://api.github.com/users/donvito/repos
	apiRoot := fmt.Sprintf("%s%s", "https://api.github.com/repos/", repo)
	response, err := http.Get(apiRoot)

	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	var _repoMetadata repoMetadata
	err = json.Unmarshal(body, &_repoMetadata)

	parsedCommitsURL := parseCommitsURL(_repoMetadata.CommitsURL)
	lastCommitDate, author := retrieveRepoCommits(parsedCommitsURL)

	m = map[string]string{"RepoName": _repoMetadata.Name, "CloneURL": _repoMetadata.CloneURL, "LastCommitDate": lastCommitDate, "Author": author}

	return

}

func parseCommitsURL(commitsURL string) (parsedCommitsURL string) {

	i := strings.Index(commitsURL, "{/sha}")
	parsedCommitsURL = commitsURL[:i]

	return

}

type repoCommits struct {
	Sha    string `json:"sha"`
	Commit commit `json:"commit"`
}

type commit struct {
	Author author `json:"author"`
}

type author struct {
	Name string `json:"name"`
	Date string `json:"date"`
}

func retrieveRepoCommits(commitsURL string) (latestCommitDate, author string) {

	response, err := http.Get(commitsURL)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)

	// Unmarshal string into structs.
	var repos []repoCommits
	json.Unmarshal(body, &repos)

	//get first element of slice only
	if len(repos) > 0 {
		latestCommitDate = repos[0].Commit.Author.Date
		author = repos[0].Commit.Author.Name
	}

	return

}
