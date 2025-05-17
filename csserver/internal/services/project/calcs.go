package project

import (
	"csserver/internal/finance"
	"csserver/internal/services/organization"
	"csserver/internal/services/project/ptypes/milestonestatus"
	"csserver/internal/services/resource"
	"fmt"

	"math"
	"strings"

	log "github.com/sirupsen/logrus"
)

func (p *Project) PerformAllCalcs(org organization.Organization, resourceMap map[string]resource.Resource, roleMap map[string]resource.Role) {
	p.AggregateProjectValueLines()
	p.CalculateProjectMilestoneStats()
	p.CalculateProjectTaskStats(org, resourceMap, roleMap)

	p.GetProjectInitialCost()
	p.GetProjectNPV()
	p.GetProjectIRR()

	p.CalculateProjectTeam()
	p.CalculateFeatureStatistics()
}

func (p *Project) AggregateProjectValueLines() {
	if len(p.ProjectValue.ProjectValueLines) == 0 {
		return
	}

	p.ProjectValue.Calculated.YearOneValue = 0
	p.ProjectValue.Calculated.YearTwoValue = 0
	p.ProjectValue.Calculated.YearThreeValue = 0
	p.ProjectValue.Calculated.YearFourValue = 0
	p.ProjectValue.Calculated.YearFiveValue = 0

	for _, l := range p.ProjectValue.ProjectValueLines {
		p.ProjectValue.Calculated.YearOneValue += l.YearOneValue
		p.ProjectValue.Calculated.YearTwoValue += l.YearTwoValue
		p.ProjectValue.Calculated.YearThreeValue += l.YearThreeValue
		p.ProjectValue.Calculated.YearFourValue += l.YearFourValue
		p.ProjectValue.Calculated.YearFiveValue += l.YearFiveValue
	}

	p.ProjectValue.Calculated.FiveYearGross =
		p.ProjectValue.Calculated.YearOneValue +
			p.ProjectValue.Calculated.YearTwoValue +
			p.ProjectValue.Calculated.YearThreeValue +
			p.ProjectValue.Calculated.YearFourValue +
			p.ProjectValue.Calculated.YearFiveValue
}

// GetProjectNPV calclate the NPV for the project based on cost and value entries
func (p *Project) GetProjectNPV() float64 {
	p.AggregateProjectValueLines()

	values := []float64{
		-(p.ProjectCost.Calculated.InitialCost),
		p.ProjectValue.Calculated.YearOneValue,
		p.ProjectValue.Calculated.YearTwoValue,
		p.ProjectValue.Calculated.YearThreeValue,
		p.ProjectValue.Calculated.YearFourValue,
		p.ProjectValue.Calculated.YearFiveValue,
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
	p.AggregateProjectValueLines()

	values := []float64{
		-(p.ProjectCost.Calculated.InitialCost),
		p.ProjectValue.Calculated.YearOneValue,
		p.ProjectValue.Calculated.YearTwoValue,
		p.ProjectValue.Calculated.YearThreeValue,
		p.ProjectValue.Calculated.YearFourValue,
		p.ProjectValue.Calculated.YearFiveValue,
	}

	irr, err := finance.IRR(values)
	if err != nil {
		log.Error(err)
		return 0.0
	}

	p.ProjectValue.Calculated.InternalRateOfReturn = irr

	return irr
}

func (p *Project) CalculateProjectTeam() {
	tm := make(map[string]string)

	for _, d := range p.ProjectDaci.DriverIDs {
		if d != nil {
			tm[*d] = *d
		}
	}

	for _, d := range p.ProjectDaci.ApproverIDs {
		if d != nil {
			tm[*d] = *d
		}
	}

	for _, d := range p.ProjectDaci.ContributorIDs {
		if d != nil {
			tm[*d] = *d
		}
	}

	for _, d := range p.ProjectDaci.InformedIDs {
		if d != nil {
			tm[*d] = *d
		}
	}

	for _, m := range p.ProjectMilestones {
		for _, t := range m.Tasks {
			for _, r := range t.ResourceIDs {
				tm[r] = r
			}
		}
	}

	team := []string{}

	for k := range tm {
		team = append(team, k)
	}

	p.Calculated.Team = team
}

// GetProjectInitialCost calculates the initial cost of the project
func (p *Project) GetProjectInitialCost() (int, float64) {
	cost := 0.0
	hours := 0
	actualizedHours := 0

	if len(p.ProjectMilestones) == 0 {
		return hours, cost
	}

	for _, m := range p.ProjectMilestones {
		if len(m.Tasks) == 0 {
			continue
		}

		for _, t := range m.Tasks {
			cost += t.Calculated.ActualizedCost
			hours += t.HourEstimate
			actualizedHours += t.Calculated.ActualizedHoursToComplete
		}
	}

	if p.ProjectCost == nil {
		p.ProjectCost = &ProjectCost{
			Calculated: ProjectCostCalculatedData{},
		}
	}

	p.ProjectCost.Calculated.HourEstimate = hours
	p.ProjectCost.Calculated.HoursActualized = actualizedHours
	p.ProjectCost.Calculated.InitialCost = cost

	return hours, cost
}

// CalculateFeatureStatistics aggregate numbers related to project features
func (p *Project) CalculateFeatureStatistics() {
	for _, f := range p.ProjectFeatures {
		switch f.Status {
		case "accepted":
			p.Calculated.FeatureStatusAcceptedCount++
		case "proposed":
			p.Calculated.FeatureStatusProposedCount++
		case "removed":
			p.Calculated.FeatureStatusRemovedCount++
		case "done":
			p.Calculated.FeatureStatusDoneCount++
		}
	}
}

// CalculateProjectMilestoneStats iterate the milestones and calculate summary information
func (p *Project) CalculateProjectMilestoneStats() {
	if len(p.ProjectMilestones) == 0 {
		return
	}

	projectUnhealthyTasks := 0

	for i, m := range p.ProjectMilestones {
		totalHours := 0
		completedHours := 0
		removedHours := 0
		hoursRemaining := 0
		completedTasks := 0
		unhealthyTasks := 0
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

				if len(t.ResourceIDs) == 0 || len(t.RequiredSkillID) == 0 || t.HourEstimate == 0 {
					unhealthyTasks++
					projectUnhealthyTasks++
				}
			}
		}

		p.ProjectMilestones[i].Calculated.TotalTasks = len(m.Tasks)
		p.ProjectMilestones[i].Calculated.TotalHours = totalHours
		p.ProjectMilestones[i].Calculated.CompletedTasks = completedTasks
		p.ProjectMilestones[i].Calculated.HoursRemaining = hoursRemaining
		p.ProjectMilestones[i].Calculated.RemovedHours = removedHours
		p.ProjectMilestones[i].Calculated.IsComplete = isComplete
		p.ProjectMilestones[i].Calculated.UnhealthyTasks = unhealthyTasks
		p.ProjectMilestones[i].Calculated.IsInFlight = !isComplete && (hoursRemaining < (totalHours - removedHours))
	}

	p.Calculated.UnhealthyTasks = projectUnhealthyTasks
}

// CalculateProjectTasksStats update the calculated fields in a project task
func (p *Project) CalculateProjectTaskStats(orgSettings organization.Organization, resources map[string]resource.Resource, roleMap map[string]resource.Role) {
	for i, m := range p.ProjectMilestones {
		for j := range m.Tasks {
			CalculateStatsForTask(p.ProjectMilestones[i].Tasks[j], orgSettings, resources, roleMap)
		}
	}
}

func CalculateStatsForTask(task *ProjectMilestoneTask, orgSettings organization.Organization, resourceMap map[string]resource.Resource, roleMap map[string]resource.Role) {
	exceptions := []string{}
	baseHoursEstimate := task.HourEstimate
	commsAdjustedHours := 0
	skillsAdjustedHours := 0
	resourceList := []resource.Resource{}
	hourlyChargeRate := orgSettings.Defaults.GenericBlendedHourlyRate

	if len(task.ResourceIDs) > 0 {
		if len(task.ResourceIDs) > 1 {
			commsAdjustedHours = int(calculateCompoundCoeffieicnt(float64(baseHoursEstimate), orgSettings.Defaults.CommsCoefficient, len(task.ResourceIDs))) - baseHoursEstimate
		}

		for _, r := range task.ResourceIDs {
			res := getResource(resourceMap, r)
			if res == nil {
				exceptions = append(exceptions, fmt.Sprintf("Resource %s not found in lookup table", r))
				continue
			}

			sk := getSkill(task.RequiredSkillID, *res)

			resourceList = append(resourceList, *res)

			if sk == nil {
				exceptions = append(exceptions, fmt.Sprintf("Resource %s does not have skill %s for task %s", res.Name, task.RequiredSkillID, task.Name))
				continue
			}

			co := getProficiencyCoeffieicnt(*sk.Proficiency)
			//thisSkillsAdjustment := -1 * (int(co*float64(baseHoursEstimate)) - baseHoursEstimate)
			thisSkillsAdjustment := (int(co*float64(baseHoursEstimate)) - baseHoursEstimate)

			skillsAdjustedHours += thisSkillsAdjustment
		}
	} else {
		exceptions = append(exceptions, "No resources allocated to task")
	}

	if len(resourceList) > 0 {
		hourlyChargeRate = int(getAverageResourceChargePerHour(orgSettings, resourceList, roleMap))
	}

	task.Calculated.CommsHourAdjustment = commsAdjustedHours
	task.Calculated.SkillsHourAdjustment = skillsAdjustedHours
	task.Calculated.ResourceContention = float64(len(task.ResourceIDs))
	task.Calculated.AverageHourlyRate = float64(hourlyChargeRate)
	task.Calculated.ActualizedHoursToComplete = commsAdjustedHours + skillsAdjustedHours + baseHoursEstimate
	task.Calculated.ActualizedCost = float64(task.Calculated.ActualizedHoursToComplete * hourlyChargeRate)

	task.Calculated.Exceptions = append(task.Calculated.Exceptions, exceptions...)
}

func getAverageResourceChargePerHour(org organization.Organization, resources []resource.Resource, roleMap map[string]resource.Role) float64 {
	rc := float64(len(resources))
	total := 0.0

	for _, r := range resources {
		var tt float64
		if r.Calculated.HourlyCost > 0 {
			tt += r.Calculated.HourlyCost
		} else {
			tt, _ = r.GetHourlyRate(roleMap, float64(org.Defaults.GenericBlendedHourlyRate), org.Defaults.WorkingHoursPerYear)
		}

		total += tt
	}

	return total / rc
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

func getResource(rm map[string]resource.Resource, id string) *resource.Resource {
	res, ok := rm[id]
	if !ok {
		return nil
	}

	return &res
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
