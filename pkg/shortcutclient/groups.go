package shortcutclient

import (
	"fmt"
	"net/http"
	"time"
)

type Group struct {
	AppURL      string `json:"app_url"`
	Archived    bool   `json:"archived"`
	Color       string `json:"color,omitempty"`
	ColorKey    string `json:"color_key,omitempty"`
	Description string `json:"description"`
	DisplayIcon struct {
		CreatedAt  time.Time `json:"created_at"`
		EntityType string    `json:"entity_type"`
		ID         string    `json:"id"`
		UpdatedAt  time.Time `json:"updated_at"`
		URL        string    `json:"url"`
	} `json:"display_icon,omitempty"`
	EntityType        string   `json:"entity_type"`
	ID                string   `json:"id"`
	MemberIds         []string `json:"member_ids"`
	MentionName       string   `json:"mention_name"`
	Name              string   `json:"name"`
	NumEpicsStarted   int      `json:"num_epics_started"`
	NumStories        int      `json:"num_stories"`
	NumStoriesStarted int      `json:"num_stories_started"`
	WorkflowIds       []int    `json:"workflow_ids"`
}

func (c *Client) ListGroups() ([]Group, error) {
	path := "/groups"

	var groups []Group
	if err := c.get(path, &groups); err != nil {
		return nil, err
	}

	return groups, nil
}

func (c *Client) ListGroupStories(groupId string) ([]*Story, error) {
	path := fmt.Sprintf("/groups/%s/stories", groupId)

	req, err := c.makeRequest(http.MethodGet, path, nil)
	if err != nil {
		return nil, err
	}

	resp, err := c.do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("status code %d", resp.StatusCode)
	}

	var stories []*Story
	if err := c.decode(resp, &stories); err != nil {
		return nil, err
	}

	return stories, nil
}
