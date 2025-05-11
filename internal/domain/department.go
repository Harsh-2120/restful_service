package domain

type Department struct {
	ID             string       `gorm:"type:uuid;primaryKey;default:gen_random_uuid()" json:"id"`
	DepartmentName string       `gorm:"size:255;not null" json:"department_name"`
	OrganizationID string       `gorm:"type:uuid" json:"organization_id"`
	Organization   Organization `gorm:"foreignKey:OrganizationID"`
}
