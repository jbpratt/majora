package models

type Config struct {
	Title              string        `json:"title"`
	DueDate            string        `json:"due_date"`
	GeneralDescription string        `json:"gen_description"`
	ProjectDescription string        `json:"proj_description"`
	Requirements       []Requirement `json:"requirements"`
	Resources          string        `json:"resources"`
}

type Requirement struct {
	Element   string `json:"element"`
	ElementID string `json:"id"`
	Type      string `json:"type"`
	Points    int    `json:"points"`
	Prompt    string `json:"prompt"`
}
