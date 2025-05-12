package responses

type FrameworkSummary struct {
	ID                    string  `json:"id"`
	FrameworkName         string  `json:"framework_name"`
	NumberOfRequirements  int64   `json:"number_of_requirements"`
	NumberOfPolicies      int     `json:"number_of_policies"`
	NumberOfEvidenceTasks int     `json:"number_of_evidence_tasks"`
	CompliancePercentage  float64 `json:"compliance_percentage"` // e.g., 85.5
}

type EvidenceSummary struct {
	Total          int `json:"total"`
	Uploaded       int `json:"uploaded"`
	NotUploaded    int `json:"not_uploaded"`
	NeedsAttention int `json:"needs_attention"`
}
