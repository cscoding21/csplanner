package portfolio

import (
	"csserver/internal/calendar"
	"csserver/internal/services/project"
	"csserver/internal/services/resource"
	"fmt"
	"time"

	"github.com/google/uuid"
)

// ScheduleProject calculate the dates of a project with its milestones and tasks
func ScheduleProject(p *project.Project, startDate time.Time, rm map[string]resource.Resource) (ProjectSchedule, error) {
	schedule := ProjectSchedule{
		ProjectID:   p.ID,
		ProjectName: p.ProjectBasics.Name,
	}
	if len(p.ProjectMilestones) == 0 {
		return schedule, nil
	}

	endOfWindow := startDate.Add(1 * time.Hour * 24 * 30 * 36)
	dateMap := calendar.GetWeeks(startDate, endOfWindow)

	p.ProjectBasics.StartDate = &startDate

	batches := getScheduleItems(p)
	weeks := []ProjectActivityWeek{}

	itemsToSchedule := batches[0].ScheduleItems
	hoursToCompleteBatch := batches[0].HoursToComplete
	for _, week := range dateMap {
		if len(itemsToSchedule) == 0 {
			break
		}

		resourceHoursForWeek := getResourceHoursForWeek(rm)

		paw := ProjectActivityWeek{Week: week}
		activities := []ProjectActivity{}

		for i, thisItem := range itemsToSchedule {
			if len(thisItem.ResourceIDs) == 0 {
				schedule.Exceptions = append(schedule.Exceptions, fmt.Sprintf("No resources assigned for task '%s'", thisItem.TaskName))
				break
			}
			for _, taskResource := range thisItem.ResourceIDs {
				resHours := resourceHoursForWeek[taskResource]

				if resHours > thisItem.HoursToSchedule {
					resHours = thisItem.HoursToSchedule
				}

				if thisItem.HoursToSchedule > 0 {
					itemsToSchedule[i].HoursToSchedule -= resHours
					itemsToSchedule[i].HoursScheduled += resHours
					hoursToCompleteBatch -= resHours

					activity := ProjectActivity{
						ProjectID:     p.ID,
						ProjectName:   p.ProjectBasics.Name,
						ResourceID:    taskResource,
						ResourceName:  rm[taskResource].Name,
						TaskID:        thisItem.TaskID,
						TaskName:      thisItem.TaskName,
						MilestoneID:   thisItem.MilestoneID,
						MilestoneName: thisItem.MilestoneName,
						HoursSpent:    resHours,
					}

					activities = append(activities, activity)

					if hoursToCompleteBatch == 0 {
						batches = batches[1:]

						if len(batches) > 0 {
							itemsToSchedule = batches[0].ScheduleItems
						}
					}
				}
			}
		}

		if len(activities) > 0 {
			paw.Activities = activities
			weeks = append(weeks, paw)
		}
	}

	schedule.ProjectActivityWeeks = &weeks

	return schedule, nil
}

func getScheduleItems(p *project.Project) []ScheduleBatch {
	sb := []ScheduleBatch{}
	order := 0

	for _, m := range p.ProjectMilestones {
		order++
		batch := ScheduleBatch{
			Order:         order,
			BatchID:       uuid.New(),
			ScheduleItems: []ScheduleWorkspace{},
		}

		hoursToComplete := 0

		for _, t := range m.Tasks {
			workspace := ScheduleWorkspace{
				MilestoneID:     *m.ID,
				MilestoneName:   m.Phase.Name,
				TaskID:          *t.ID,
				TaskName:        t.Name,
				ResourceIDs:     t.ResourceIDs,
				HoursToSchedule: t.Calculated.ActualizedHoursToComplete,
				HoursScheduled:  0,
			}

			hoursToComplete += t.Calculated.ActualizedHoursToComplete

			batch.ScheduleItems = append(batch.ScheduleItems, workspace)
		}

		batch.HoursToComplete = hoursToComplete

		sb = append(sb, batch)
	}

	return sb
}

func getResourceHoursForWeek(rm map[string]resource.Resource) map[string]int {
	outMap := make(map[string]int)

	for k, v := range rm {
		outMap[k] = v.AvailableHoursPerWeek
	}

	return outMap
}
