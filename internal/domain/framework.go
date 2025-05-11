package domain

type Framework struct {
	ID                     string       `gorm:"type:uuid;primaryKey;default:gen_random_uuid()" json:"id"`
	FrameworkName          string       `gorm:"size:255;not null" json:"framework_name"`
	NumberOfPolicies       int          `json:"number_of_policies"`
	NumberOfEvidenceTasks  int          `json:"number_of_evidence_tasks"`
	NumberOfAutomatedTests int          `json:"number_of_automated_tests"`
	IsCustom               bool         `json:"is_custom"`
	OrganizationID         string       `gorm:"type:uuid" json:"organization_id"`
	Organization           Organization `gorm:"foreignKey:OrganizationID"`
}
