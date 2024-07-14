package main

import "github.com/The-Manchester-Project/Apollo-24/router"

func main() {
	// NewIssue := &jira.IssueTemplate{
	// 	SINumber:              "JB007",
	// 	SILink:                "google.com/hjhd/jskhd/dh",
	// 	BIARecord:             "facebook.xom",
	// 	ProjectOverview:       "Man is a man and only a man will know what a man is !!",
	// 	PlatformName:          "Value added",
	// 	LabName:               "Smarty pants",
	// 	GW1Date:               "2024-12-08",
	// 	SolutionArchitectName: "Megh Shetty",
	// 	ProjectName:           "Template value",
	// 	SIRrating:             []string{"SIRC"},
	// 	DataClassification:    "Confidential",
	// }

	// fmt.Println(jira.CreateIssue(NewIssue))
	router.WebServer()
}
