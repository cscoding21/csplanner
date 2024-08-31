package project

import (
	"csserver/internal/finance"

	log "github.com/sirupsen/logrus"
)

// GetProjectNPV calclate the NPV for the project based on cost and value entries
func (p *Project) GetProjectNPV() float64 {

	values := []float64{
		-(p.ProjectCost.InitialCost),
		p.ProjectValue.YearOneValue,
		p.ProjectValue.YearTwoValue,
		p.ProjectValue.YearThreeValue,
		p.ProjectValue.YearFourValue,
		p.ProjectValue.YearFiveValue,
	}

	npv, err := finance.NPV((p.ProjectValue.DiscountRate / 100.0), values)
	if err != nil {
		log.Error(err)
		return 0.0
	}

	return npv
}

// GetProjectIRR calclate the NPV for the project based on cost and value entries
func (p *Project) GetProjectIRR(project Project) float64 {
	values := []float64{
		-(project.ProjectCost.InitialCost),
		project.ProjectValue.YearOneValue,
		project.ProjectValue.YearTwoValue,
		project.ProjectValue.YearThreeValue,
		project.ProjectValue.YearFourValue,
		project.ProjectValue.YearFiveValue,
	}

	irr, err := finance.IRR(values)
	if err != nil {
		log.Error(err)
		return 0.0
	}

	return irr
}

// GetProjectInitialCost calculates the initial cost of the project
func (p *Project) GetProjectInitialCost(project Project) (int, float64) {
	cost := 0.0
	hours := 0

	if project.ProjectMilestones == nil || len(project.ProjectMilestones) == 0 {
		return hours, cost
	}

	for _, m := range project.ProjectMilestones {
		if m.Tasks == nil || len(m.Tasks) == 0 {
			continue
		}

		for _, t := range m.Tasks {
			hours = hours + t.HourEstimate
			cost = cost + float64(t.HourEstimate)*(*project.ProjectCost.BlendedRate)
		}
	}

	return hours, cost
}
