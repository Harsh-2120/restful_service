package domain

type Control struct {
	ID              string       `gorm:"type:uuid;primaryKey;default:gen_random_uuid()" json:"id"`
	ControlCode     string       `gorm:"size:50;not null;unique" json:"control_code"`
	ControlName     string       `gorm:"size:255;not null" json:"control_name"`
	ControlDomain   string       `gorm:"size:255" json:"control_domain"`
	Status          string       `gorm:"size:20" json:"status"`
	Assignee        string       `gorm:"size:255" json:"assignee"`
	Description     string       `gorm:"type:text" json:"description,omitempty"`
	ControlQuestion string       `gorm:"type:text" json:"control_question,omitempty"`
	OrganizationID  string       `gorm:"type:uuid" json:"organization_id"`
	Organization    Organization `gorm:"foreignKey:OrganizationID"`
}
