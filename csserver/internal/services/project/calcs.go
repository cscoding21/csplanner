package project

import (
	"csserver/internal/finance"
	"csserver/internal/services/project/ptypes"

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

	p.ProjectValue.NetPresentValue = npv

	return npv
}

// GetProjectIRR calclate the NPV for the project based on cost and value entries
func (p *Project) GetProjectIRR() float64 {
	values := []float64{
		-(p.ProjectCost.InitialCost),
		p.ProjectValue.YearOneValue,
		p.ProjectValue.YearTwoValue,
		p.ProjectValue.YearThreeValue,
		p.ProjectValue.YearFourValue,
		p.ProjectValue.YearFiveValue,
	}

	irr, err := finance.IRR(values)
	if err != nil {
		log.Error(err)
		return 0.0
	}

	p.ProjectValue.InternalRateOfReturn = irr

	return irr
}

// GetProjectInitialCost calculates the initial cost of the project
func (p *Project) GetProjectInitialCost() (int, float64) {
	cost := 0.0
	hours := 0

	if len(p.ProjectMilestones) == 0 {
		return hours, cost
	}

	for _, m := range p.ProjectMilestones {
		if len(m.Tasks) == 0 {
			continue
		}

		for _, t := range m.Tasks {
			hours = hours + t.HourEstimate
			cost = cost + float64(t.HourEstimate)*(*p.ProjectCost.BlendedRate)
		}
	}

	p.ProjectCost.HourEstimate = hours
	p.ProjectCost.InitialCost = cost

	return hours, cost
}

// CalculateProjectMilestoneStats iterate the milestones and calculate summary information
func (p *Project) CalculateProjectMilestoneStats() {
	if len(p.ProjectMilestones) == 0 {
		return
	}

	for i, m := range p.ProjectMilestones {
		totalHours := 0
		completedHours := 0
		removedHours := 0
		hoursRemaining := 0
		completedTasks := 0
		isComplete := true

		for _, t := range m.Tasks {
			totalHours += t.HourEstimate

			if t.Status == ptypes.Done {
				completedHours += t.HourEstimate
				completedTasks++
			} else if t.Status == ptypes.Removed {
				removedHours += t.HourEstimate
			} else {
				hoursRemaining += t.HourEstimate
				isComplete = false
			}
		}

		p.ProjectMilestones[i].TotalHours = totalHours
		p.ProjectMilestones[i].CompletedTasks = completedTasks
		p.ProjectMilestones[i].HoursRemaining = hoursRemaining
		p.ProjectMilestones[i].RemovedHours = removedHours
		p.ProjectMilestones[i].IsComplete = isComplete
		p.ProjectMilestones[i].IsInFlight = !isComplete && (hoursRemaining < (totalHours - removedHours))
	}
}
