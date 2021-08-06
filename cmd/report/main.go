package main

import (
	"errors"
	"fmt"
	"github.com/go-shortcut/go-shortcut-api/pkg/shortcutclient"
	"os"
	"sort"
	"strconv"
	"time"
)

func main() {

	token := os.Getenv("CLUBHOUSE_API_TOKEN")
	if token == "" {
		panic(errors.New("CLUBHOUSE_API_TOKEN environ is required"))
	}
	shortcutClient := shortcutclient.New(token)
	shortcutClient.HTTPClient.Timeout = 30 * time.Second
	//shortcutClient.URL = "https://api.shortcut.com/api/v3"

	// Get Epics information
	epics, err := shortcutClient.ListEpics()
	if err != nil {
		panic(err)
	}
	getEpic := map[int64]shortcutclient.Epic{}
	for _, e := range epics {
		getEpic[e.ID] = e
	}

	// Get Members information
	members, err := shortcutClient.ListMembers()
	if err != nil {
		panic(err)
	}
	getMember := map[string]shortcutclient.Member{}
	for _, m := range members {
		getMember[m.ID] = m
	}
	// Get WorkFlow states information
	workFlows, err := shortcutClient.ListWorkflows()
	if err != nil {
		panic(err)
	}
	getWorkFlowState := map[int64]shortcutclient.WorkflowState{}
	for _, wf := range workFlows {
		for _, epicState := range wf.States {
			getWorkFlowState[epicState.ID] = epicState

		}
	}
	//

	if len(os.Args) > 1 {
		firstCmdArg := os.Args[1]
		_, err := strconv.Atoi(firstCmdArg)
		if err != nil {
			panic(err)

		}
		stories, err := shortcutClient.ListStoriesForProject(firstCmdArg)
		if err != nil {
			panic(err)
		}

		// Sorting stories by id
		sort.Slice(stories, func(i, j int) bool {
			return stories[i].EpicID < stories[j].EpicID
		})

		fmt.Printf("%-50s %-20s %-150s\n", "Story URL", "Owners", "Story Name")
		var LastEpicID int64 = -1
		for _, s := range stories {
			if !s.Completed && !s.Archived {
				if s.EpicID != LastEpicID {
					LastEpicID = s.EpicID
					if LastEpicID == 0 {
						fmt.Println("# Undefined Epic")
					} else {
						fmt.Println("# EpicID: " + strconv.FormatInt(LastEpicID, 10) + " " + getEpic[LastEpicID].Name)

					}
				}

				owners := "owner: "
				if len(s.OwnerIDs) == 0 {
					owners = "[NO_OWNER]"
				}
				for _, oid := range s.OwnerIDs {
					owners += getMember[oid].Profile.MentionName + ", "
				}
				fmt.Printf("%-50s %-30.30s %-30.30s  %-100.100s\n", s.AppURL, getWorkFlowState[s.WorkflowStateID].Name, owners, s.Name)

			}

		}

	} else {
		projects, err := shortcutClient.ListProjects()
		if err != nil {
			panic(err)
		}

		// Sorting projects by id
		sort.Slice(projects, func(i, j int) bool {
			return projects[i].ID < projects[j].ID
		})

		fmt.Printf("%6s %-50s %s\n", "p.ID", "p.AppURL", "p.Name")
		for _, p := range projects {
			fmt.Printf("%6d %-50s %s\n", p.ID, p.AppURL, p.Name)
		}
	}

}
