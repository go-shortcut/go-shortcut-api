package main

import (
	"fmt"
	"github.com/go-shortcut/go-shortcut-api/pkg/shortcutclient"
	"os"
	"time"
)

func main() {

	token := os.Getenv("SHORTCUT_API_TOKEN")
	if token == "" {
		fmt.Println("SHORTCUT_API_TOKEN environ is required")
		os.Exit(1)
	}
	shortcutClient := shortcutclient.New(token)
	shortcutClient.HTTPClient.Timeout = 30 * time.Second
	shortcutClient.Debug = true
	shortcutClient.UpdateMultipleStories(
		shortcutclient.UpdateMultipleStoriesParams{
			StoryIds: []int64{96000, 96110, 96995},
			LabelsAdd: []shortcutclient.CreateLabelParams{
				{Name: "Live", Description: "Released to the market", Color: "#0000FF"},
			},
			LabelsRemove: []shortcutclient.CreateLabelParams{
				{Name: "DevOps", Description: "URGENT", Color: "#FF0000"},
			},
		})
}
