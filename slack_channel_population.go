package main

import (
	"fmt"
	"github.com/nlopes/slack"
	"os"
	"sort"
)

type ChannelPop struct {
	name  string
	count int
}

func main() {
	slacktoken := os.Getenv("SLACK_TOKEN")
	if slacktoken == "" {
		fmt.Println("You need to set SLACK_TOKEN.")
		os.Exit(1)
	}
	api := slack.New(slacktoken)
	groups, err := api.GetChannels(false)
	if err != nil {
		fmt.Printf("%s\n", err)
		return
	}
	var capture []ChannelPop
	for _, group := range groups {
		chanpop := ChannelPop{name: group.Name, count: group.NumMembers}
		capture = append(capture, chanpop)
	}
	sort.SliceStable(capture, func(i, j int) bool {
		return capture[i].count < capture[j].count
	})
	for _, item := range capture {
		fmt.Println(item.name, item.count)
	}
}
