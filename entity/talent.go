package entity

// Talent entity for talent table
type Talent struct {
	ID   int64  `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}
