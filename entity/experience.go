package entity

// Experience entity that related with experience table
type Experience struct {
	ID       int64  `json:"id"`
	Company  string `json:"company"`
	TalentID int64  `json:"talent_id"`
}
