/*
===================================================================================================
JIRA API Package
Jira REST API doc https://developer.atlassian.com/cloud/jira/platform/basic-auth-for-rest-apis/
===================================================================================================
*/
package jira

import (
	"fmt"
	"io"
	"net/http"
)

// Cookie & Authorization will pull the token during run time from aws secrect manager.
const BaseUrl string = "https://cavertech.atlassian.net"
const cookie string = "cookie token"
const Authorization = "Authorization code"

type JiraIssuePayload struct {
	Fields struct {
		Summary   string `json:"summary"`
		Issuetype struct {
			Name string `json:"name"`
		} `json:"issuetype"`
		Project struct {
			Key string `json:"key"`
		} `json:"project"`
	} `json:"fields"`
}

/*
	Jira auth handle api request use base auth.

Param

	method : GET, POST, UPDATE
	urlExt : rest path
	payload : payload for the api body.

return :

	api results
	error
*/
func JiraAuth(method, urlExt, payload string) (string, error) {
	req, err := http.NewRequest(method, BaseUrl+urlExt, nil)
	if err != nil {
		return "req", fmt.Errorf("failed to fetch url: %w", err)
	}
	req.Header.Add("cookie", cookie)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", Authorization)
	result, err := http.DefaultClient.Do(req)
	if err != nil {
		return "results", fmt.Errorf("failed to fetch Jira metadata: %w", err)
	}
	defer result.Body.Close()
	body, _ := io.ReadAll((result.Body))
	return string(body), nil
}
