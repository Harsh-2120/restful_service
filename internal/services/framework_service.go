package services

import (
	"github.com/Aakarsh-Kamboj/rest-service/db"
	"github.com/Aakarsh-Kamboj/rest-service/internal/domain"
	"github.com/Aakarsh-Kamboj/rest-service/internal/responses"
)

func GetFrameworkSummaries() ([]responses.FrameworkSummary, error) {
	var frameworks []domain.Framework
	if err := db.DB.Find(&frameworks).Error; err != nil {
		return nil, err
	}

	var summaries []responses.FrameworkSummary
	for _, framework := range frameworks {
		var totalControls int64
		var compliantControls int64

		// Total Controls for the Framework
		err := db.DB.
			Table("controls").
			Joins("JOIN framework_controls ON controls.id = framework_controls.control_id").
			Where("framework_controls.framework_id = ?", framework.ID).
			Count(&totalControls).Error
		if err != nil {
			return nil, err
		}

		// Compliant Controls for the Framework
		err = db.DB.
			Table("controls").
			Joins("JOIN framework_controls ON controls.id = framework_controls.control_id").
			Where("framework_controls.framework_id = ? AND controls.status = ?", framework.ID, "Compliant").
			Count(&compliantControls).Error
		if err != nil {
			return nil, err
		}

		// Calculate compliance percentage
		var compliance float64
		if totalControls > 0 {
			compliance = (float64(compliantControls) / float64(totalControls)) * 100
		}

		summary := responses.FrameworkSummary{
			ID:                    framework.ID,
			FrameworkName:         framework.FrameworkName,
			NumberOfRequirements:  totalControls,
			NumberOfPolicies:      framework.NumberOfPolicies,
			NumberOfEvidenceTasks: framework.NumberOfEvidenceTasks,
			CompliancePercentage:  compliance,
		}
		summaries = append(summaries, summary)
	}

	return summaries, nil
}
