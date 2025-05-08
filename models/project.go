package models

type Project struct {
	ID            int64    `json:"id"`
	ProjectTitle  string   `json:"projectTitle"`
	ClientName    string   `json:"clientName"`
	Category      string   `json:"category"`
	ProjectStills []string `json:"projectStills"`
	Description   string   `json:"description"`
	ProjectLink   string   `json:"projectLink"`
}
