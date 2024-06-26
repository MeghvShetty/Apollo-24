/*
===================================================================================================
JIRA API Package
Jira REST API doc https://developer.atlassian.com/cloud/jira/platform/basic-auth-for-rest-apis/
===================================================================================================
*/
package jira

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// Cookie & Authorization will pull the token during run time from aws secrect manager.
const BaseUrl string = "https://cavertech.atlassian.net"
const cookie string = "cookie token"
const Authorization = "Authorization code"

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
func JiraAuth(method, urlExt string, payload []byte) (string, error) {
	req, err := http.NewRequest(method, BaseUrl+urlExt, bytes.NewReader(payload))
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

/*
CreateIssue for the given project key.
param : nil
returns :

	Jira id, issue number with url to the ticket
	error message
*/
func CreateIssue() {
	urlExt := "/rest/api/2/issue"
	method := "POST"

	// Static typed payload
	payload := map[string]interface{}{
		"fields": map[string]interface{}{
			"description": "This is from Apollo",
			"summary":     "Apollo 24 is here",
			"issuetype": map[string]interface{}{
				"name": "Story",
			},
			"project": map[string]interface{}{
				"key": "A2",
			},
		},
	}

	//Encoding payload map[string]interface{} into Json data type
	jsonData, err := json.Marshal(payload)
	if err != nil {
		fmt.Println("Error marshalling JSON:", err)
		return
	}
	fmt.Println(JiraAuth(method, urlExt, jsonData))
}
