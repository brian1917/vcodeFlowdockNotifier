package main

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/brian1917/vcodeapi"
)

func main() {

	// Parse config file
	config := parseConfig(os.Args[1])

	// Build the endpoint
	endpoint := "https://api.flowdock.com/flows/" + config.FlowdockOrg + "/" + config.FlowdockFlow + "/messages"

	// Get build list and most recent build
	buildList, err := vcodeapi.ParseBuildList(config.CredsFile, config.AppID)
	if err != nil {
		log.Fatal(err)
	}
	recentBuild := buildList[len(buildList)-1].BuildID

	// Get metadata from detailed report for most recent build
	detReportMeta, err := vcodeapi.ParseBuildMetaData(config.CredsFile, recentBuild)
	if err != nil {
		log.Fatal(err)
	}

	// Set JSON payload based on IncludePolicyStatus config parameter
	var content string
	if config.IncludePolicyStatus == true {
		content = fmt.Sprintf("Veracode Scan Complete. \n **Application:** %s \n **Build:** %s \n **Policy:** %s \n **Policy Status:** %s",
			detReportMeta.AppName, recentBuild, detReportMeta.PolicyName, detReportMeta.PolicyComplianceStatus)
	} else {
		content = fmt.Sprintf("Veracode Scan Complete. \n **Application:** %s \n **Build:** %s", detReportMeta.AppName, recentBuild)
	}
	var jsonStr = []byte(`{
		"event": "message",
		"content":"` + content + `"}`)

	// Check to see if Flowdock Notification is required
	if (config.OnlyNotifyOnNotPass == true && detReportMeta.PolicyComplianceStatus != "Pass") || config.OnlyNotifyOnNotPass == false {

		// Create HTTP request and set header
		req, err := http.NewRequest("POST", endpoint, bytes.NewBuffer(jsonStr))
		if err != nil {
			log.Fatal(err)
		}
		req.Header.Set("Content-Type", "application/json")
		req.SetBasicAuth(config.FlowdockToken, "")

		// Create HTTP client and send request
		client := http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			log.Fatal(err)
		}
		if resp.StatusCode != 200 {
			log.Fatal(resp.Status)
		}
	}
}
