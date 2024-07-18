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
	"strings"
	"time"
)

// Cookie & Authorization will pull the token during run time from aws secrect manager.
const BaseUrl string = "https://cavertech.atlassian.net"
const cookie string = "atlassian.xsrf.token=052f0f6ae563178da122c7f109e7eaaf6bfbbf75_lin"

const Authorization = "Basic bWVnaHZzaGV0dHlAY2F2ZXJ0ZWNoLmNvbTpBVEFUVDN4RmZHRjBZZVNsaDJmNThfNzBFaXZfa1MxZlR6ejBITGRhWnNoVTRTREVLc08zRHBtclNqeXl3YW5FS2FMb2RRa09qbDhfSmxMYlFsNmRwM3Q4UE5ILWMwNFBFYVFzUFFrRWtPNVluWUdVSmxaS3g0MDJEb01zZWY1NDRvRE9zb2FidmdDYmtkWG5MU3h3MXNnMTc5elVrTFkydEdqX1habktYV21LdWk3TnowZlFMejQ9MUZEQUNFOTY="

type IssueTemplate struct {
	SINumber              string `form:SINumber" binding:"required"`
	SILink                string `form:SILink" binding:"required"`
	BIARecord             string `form:BIARecord" binding:"required"`
	ProjectOverview       string `form:ProjectOverview" binding:"required"`
	PlatformName          string `form:PlatformName" binding:"required"`
	LabName               string `form:LabName" binding:"required"`
	GW1Date               string `form:GW1Date" binding:"required"`
	SolutionArchitectName string `form:SolutionArchitectName" binding:"required"`
	ProjectName           string `form:ProjectName" binding:"required"`
	SIRrating             []string
	DataClassification    string `form:DataClassification" binding:"required"`
}

type AARep struct {
	IssueName    string
	DueDate      string
	Description  string
	AssigneeName string
	SIRrating    []string
	Changelog    string
}

type Payload struct {
	Body string `json:"body"`
	// Visibility Visibility `json:"visibility"`
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
func JiraAuth(method, urlExt string, payload []byte) (string, error) {
	req, err := http.NewRequest(method, BaseUrl+urlExt, bytes.NewReader(payload))
	if err != nil {
		return "", fmt.Errorf("failed to fetch url: %w", err)
	}
	req.Header.Add("cookie", cookie)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", Authorization)
	result, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", fmt.Errorf("failed to fetch Jira metadata: %w", err)
	}
	defer result.Body.Close()
	body, _ := io.ReadAll(result.Body)
	return string(body), nil
}

/*
CreateIssue for the given project key.
param : IssueTemplate Struct
returns :

	Jira id, issue number with url to the ticket
	error message
*/
func CreateIssue(p *IssueTemplate) (string, error) {
	urlExt := "/rest/api/2/issue"
	method := "POST"

	// Issue naming convention
	var ticketName string = p.SINumber + "_" + p.LabName + "_" + p.ProjectName

	// Validate date formatting
	var dueDateInput = strings.TrimSpace(p.GW1Date)
	_, err := time.Parse("2006-01-02", dueDateInput)
	if err != nil {
		return "", fmt.Errorf("invalid date format. please use yyyy-mm-dd.: %w", err)
	}

	// Static typed payload
	payload := map[string]interface{}{
		"fields": map[string]interface{}{
			"description": "Project Name: " + p.ProjectName + "\n Platform Name: " + p.PlatformName + "\n Lab Name: " + p.LabName + "\n SI Number: " + p.SINumber + "\n SI Link: " + p.SILink + "\n Data Classification: " + p.DataClassification + "\n BIA Link: " + p.BIARecord + "\n Solution Architect: " + p.SolutionArchitectName + "\n Project Overview: " + p.ProjectOverview + "\n SDA link: ",
			"summary":     ticketName,
			"issuetype": map[string]interface{}{
				"name": "Story",
			},
			"project": map[string]interface{}{
				"key": "A2",
			},
			"customfield_10034": dueDateInput,
			"labels":            p.SIRrating,
		},
	}

	// Encoding payload map[string]interface{} into Json data type
	jsonData, err := json.Marshal(payload)
	if err != nil {
		return "", fmt.Errorf("error marshalling JSON: %w", err)
	}

	result, err := JiraAuth(method, urlExt, jsonData)
	if err != nil {
		return "", fmt.Errorf("failed to create issue: %w", err)
	}

	return result, nil
}

func CreateIssueAA(p *AARep) (string, error) {
	urlExt := "/rest/api/2/issue"
	method := "POST"

	// Validate date formatting
	var dueDateInput = strings.TrimSpace(p.DueDate)
	_, err := time.Parse("2006-01-02", dueDateInput)
	if err != nil {
		return "", fmt.Errorf("invalid date format. please use yyyy-mm-dd.: %w", err)
	}

	// Static typed payload
	payload := map[string]interface{}{
		"fields": map[string]interface{}{
			"description": p.Description + "\n Security Engineer: " + p.AssigneeName,
			"summary":     p.IssueName,
			"issuetype": map[string]interface{}{
				"name": "Story",
			},
			"project": map[string]interface{}{
				"key": "GOV",
			},
			"duedate": p.DueDate,
			"labels":  p.SIRrating,
		},
	}

	// Encoding payload map[string]interface{} into Json data type
	jsonData, err := json.Marshal(payload)
	if err != nil {
		return "", fmt.Errorf("error marshalling JSON: %w", err)
	}

	result, err := JiraAuth(method, urlExt, jsonData)
	if err != nil {
		return "", fmt.Errorf("failed to create issue: %w", err)
	}

	return result, nil

}

func AddComments(IssueKey, JiraTicket string) {
	urlExt := "/rest/api/2/issue/" + IssueKey + "/comment"
	method := "POST"
	payload := Payload{
		Body: "AA-Rep Jira Ticket: " + JiraTicket,
	}

	jsonData, err := json.Marshal(payload)
	if err != nil {
		fmt.Errorf("error marshalling JSON: %w", err)
	}
	JiraAuth(method, urlExt, jsonData)

}
