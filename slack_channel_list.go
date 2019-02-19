package main

import (
	"fmt"
	"github.com/nlopes/slack"
	"os"
)

func main() {
	slacktoken := os.Getenv("SLACK_TOKEN")
	if slacktoken == "" {
		fmt.Println("You need to set SLACK_TOKEN.")
		os.Exit(1)
	}
	api := slack.New(slacktoken)
	// If you set debugging, it will log all requests to the console
	// Useful when encountering issues
	// api.SetDebug(true)
	//groups, err := api.GetGroups(false)
	groups, err := api.GetChannels(false)
	if err != nil {
		fmt.Printf("%s\n", err)
		return
	}
	for _, group := range groups {
		fmt.Printf("ID: %s, Name: %s\n", group.ID, group.Name)
	}
}
