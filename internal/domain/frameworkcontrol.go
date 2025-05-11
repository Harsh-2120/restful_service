package domain

type FrameworkControl struct {
	ID          string `gorm:"type:uuid;primaryKey" json:"id"`
	FrameworkID string `gorm:"type:uuid" json:"framework_id"`
	ControlID   string `gorm:"type:uuid" json:"control_id"`
}
