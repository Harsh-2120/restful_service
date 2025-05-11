package domain

import "time"

type EvidenceTask struct {
	ID             string       `gorm:"type:uuid;primaryKey;default:gen_random_uuid()" json:"id"`
	EvidenceName   string       `gorm:"size:255;not null" json:"evidence_name"`
	Status         string       `gorm:"type:enum('Pending','In Progress','Completed','Reviewed')" json:"status"`
	Assignee       string       `gorm:"size:255" json:"assignee"`
	Department     string       `gorm:"size:255" json:"department"`
	DueDate        time.Time    `json:"due_date"`
	UploadedDate   time.Time    `json:"uploaded_date"`
	FrameworkID    string       `gorm:"type:uuid" json:"framework_id"`
	OrganizationID string       `gorm:"type:uuid" json:"organization_id"`
	Framework      Framework    `gorm:"foreignKey:FrameworkID"`
	Organization   Organization `gorm:"foreignKey:OrganizationID"`
}
