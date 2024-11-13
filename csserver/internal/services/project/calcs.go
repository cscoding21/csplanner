package project

import (
	"csserver/internal/finance"
	"csserver/internal/services/organization"
	"csserver/internal/services/project/ptypes/milestonestatus"
	"csserver/internal/services/resource"

	"math"
	"strings"

	log "github.com/sirupsen/logrus"
)

// GetProjectNPV calclate the NPV for the project based on cost and value entries
func (p *Project) GetProjectNPV() float64 {

	values := []float64{
		-(p.ProjectCost.Calculated.InitialCost),
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

	p.ProjectValue.Calculated.NetPresentValue = npv

	return npv
}

// GetProjectIRR calclate the NPV for the project based on cost and value entries
func (p *Project) GetProjectIRR() float64 {
	values := []float64{
		-(p.ProjectCost.Calculated.InitialCost),
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

	p.ProjectValue.Calculated.InternalRateOfReturn = irr

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

	p.ProjectCost.Calculated.HourEstimate = hours
	p.ProjectCost.Calculated.InitialCost = cost

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

			if t.Status == milestonestatus.Done {
				completedHours += t.HourEstimate
				completedTasks++
			} else if t.Status == milestonestatus.Removed {
				removedHours += t.HourEstimate
			} else {
				hoursRemaining += t.HourEstimate
				isComplete = false
			}
		}

		p.ProjectMilestones[i].Calculated.TotalTasks = len(m.Tasks)
		p.ProjectMilestones[i].Calculated.TotalHours = totalHours
		p.ProjectMilestones[i].Calculated.CompletedTasks = completedTasks
		p.ProjectMilestones[i].Calculated.HoursRemaining = hoursRemaining
		p.ProjectMilestones[i].Calculated.RemovedHours = removedHours
		p.ProjectMilestones[i].Calculated.IsComplete = isComplete
		p.ProjectMilestones[i].Calculated.IsInFlight = !isComplete && (hoursRemaining < (totalHours - removedHours))
	}
}

// CalculateProjectTasksStats update the calculated fields in a project task
func (p *Project) CalculateProjectTasksStats(orgSettings organization.Organization, resources []resource.Resource) {
	for i, m := range p.ProjectMilestones {
		for j := range m.Tasks {
			calculateStatsForTask(p.ProjectMilestones[i].Tasks[j], orgSettings, resources)
		}
	}
}

func calculateStatsForTask(task *ProjectMilestoneTask, orgSettings organization.Organization, resources []resource.Resource) {
	exceptions := []string{}
	adjustedHours := task.HourEstimate

	if len(task.ResourceIDs) > 0 {
		adjustedHours = int(calculateCompoundCoeffieicnt(float64(adjustedHours), orgSettings.Defaults.FocusFactor, len(task.ResourceIDs)))

		for _, r := range task.ResourceIDs {
			res := getResource(r, resources)
			sk := getSkill(task.RequiredSkillID, *res)

			co := getProficiencyCoeffieicnt(*sk.Proficiency)
			adjustedHours = int(calculateCompoundCoeffieicnt(float64(adjustedHours), co, len(task.ResourceIDs)))
		}
	}

	task.Calculated.ResourceContention = float64(len(task.ResourceIDs))
	task.Calculated.ActualizedHoursToComplete = adjustedHours

	task.Calculated.Exceptions = append(task.Calculated.Exceptions, exceptions...)
}

func getProficiencyCoeffieicnt(prof float64) float64 {
	profInt := int(prof)
	switch profInt {
	case (1):
		return 1.25
	case (2):
		return 1.0
	case (3):
		return 0.75
	default:
		return 0.0
	}
}

func getResource(id string, res []resource.Resource) *resource.Resource {
	for _, r := range res {
		if strings.EqualFold(r.ID, id) {
			return &r
		}
	}

	return nil
}

func getSkill(id string, res resource.Resource) *resource.Skill {
	for _, s := range res.Skills {
		if strings.EqualFold(s.ID, id) {
			return s
		}
	}

	return nil
}

func calculateCompoundCoeffieicnt(base float64, drag float64, periods int) float64 {
	return base * math.Pow((1+drag/100), float64(periods))
}

/*
Use the following equation to calculate compound interest:

Compound Interest = P [(1 + i)n – P]

P stands for principal; i stands for interest; n stands for the number of compounding periods.
*/
