package domain

type Policy struct {
	ID             string       `gorm:"type:uuid;primaryKey" json:"id"`
	PolicyName     string       `gorm:"size:255;not null" json:"policy_name"`
	Description    string       `gorm:"type:text" json:"description,omitempty"`
	FrameworkID    string       `gorm:"type:uuid" json:"framework_id"`
	OrganizationID string       `gorm:"type:uuid" json:"organization_id"`
	Framework      Framework    `gorm:"foreignKey:FrameworkID"`
	Organization   Organization `gorm:"foreignKey:OrganizationID"`
}
