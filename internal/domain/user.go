package domain

type User struct {
	ID             string       `gorm:"type:uuid;primaryKey" json:"id"`
	Username       string       `gorm:"size:100;unique;not null" json:"username"`
	FullName       string       `gorm:"size:255" json:"full_name"`
	Email          string       `gorm:"size:255;unique;not null" json:"email"`
	Department     string       `gorm:"size:255" json:"department"`
	Role           string       `gorm:"size:50" json:"role"`
	OrganizationID string       `gorm:"type:uuid" json:"organization_id"`
	Organization   Organization `gorm:"foreignKey:OrganizationID"`
}
