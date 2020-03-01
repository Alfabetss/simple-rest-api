package entity

// Experience entity that related with experience table
type Experience struct {
	ID       int64  `json:"id,omitempty"`
	Company  string `json:"company,omitempty"`
	TalentID int64  `json:"talent_id,omitempty"`
}
