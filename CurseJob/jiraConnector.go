package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
	"time"

	"github.com/spf13/viper"
)

type JiraIssue struct {
	ID     string `json:"id"`
	Key    string `json:"key"`
	Fields struct {
		Summary string `json:"summary"`
	} `json:"fields"`
}

type JiraResponse struct {
	Issues []JiraIssue `json:"issues"`
}

func fetchAndStoreJiraProject(projectKey string) error {
	jiraUrl := viper.GetString("jiraUrl")
	issueCount := viper.GetInt("issueInOneRequest")
	threadCount := viper.GetInt("threadCount")
	minTimeSleep := viper.GetInt("minTimeSleep")
	maxTimeSleep := viper.GetInt("maxTimeSleep")

	url := fmt.Sprintf("%s/rest/api/2/search?jql=project=%s&maxResults=%d", jiraUrl, projectKey, issueCount)

	var wg sync.WaitGroup
	issuesChan := make(chan JiraIssue, 100)

	for i := 0; i < threadCount; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			client := &http.Client{}
			currentSleep := minTimeSleep

			for {
				resp, err := client.Get(url)
				if err != nil {
					time.Sleep(time.Duration(currentSleep) * time.Millisecond)
					currentSleep *= 2
					if currentSleep > maxTimeSleep {
						break
					}
					continue
				}

				body, err := ioutil.ReadAll(resp.Body)
				if err != nil {
					resp.Body.Close()
					continue
				}
				resp.Body.Close()

				var jiraResponse JiraResponse
				if err := json.Unmarshal(body, &jiraResponse); err != nil {
					continue
				}

				for _, issue := range jiraResponse.Issues {
					issuesChan <- issue
				}

				if len(jiraResponse.Issues) < issueCount {
					break
				}

				time.Sleep(time.Duration(currentSleep) * time.Millisecond)
				currentSleep = minTimeSleep
			}
		}()
	}

	go func() {
		wg.Wait()
		close(issuesChan)
	}()

	for issue := range issuesChan {
		storeIssueInDB(issue)
	}

	return nil
}

func storeIssueInDB(issue JiraIssue) {
	// Implement your logic to store issues in the database
}

func getProjectsFromDB(limit, page, search string) ([]JiraIssue, PageInfo, error) {
	// Implement your logic to retrieve projects from the database
	return nil, PageInfo{}, nil
}

type PageInfo struct {
	PageCount     int `json:"pageCount"`
	CurrentPage   int `json:"currentPage"`
	ProjectsCount int `json:"projectsCount"`
}
