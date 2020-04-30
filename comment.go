package clubhouse

type Comment struct {
	AppURL           string    `json:"app_url"`
	AuthorID         string    `json:"author_id"`
	Comments         []Comment `json:"comments"`
	CreatedAt        string    `json:"created_at"`
	Deleted          bool      `json:"deleted"`
	EntityType       string    `json:"entity_type"`
	ExternalID       string    `json:"external_id"`
	GroupMentionIDs  []string  `json:"group_mention_ids"`
	ID               int64     `json:"id"`
	MemberMentionIDs []string  `json:"member_mention_ids"`
	MentionIDs       []string  `json:"mention_ids"`
	Text             string    `json:"text"`
	UpdatedAt        string    `json:"updated_at"`
}
