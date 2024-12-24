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

	p.ProjectBasics.StartDate = &startDate

	batches := getScheduleItems(p)
	startWeek := calendar.GetWeek(startDate)
	weeks := []ProjectActivityWeek{}

	resourceHoursForWeek := getResourceHoursForWeek(rm)
	exceptionsCreated := make(map[string]string)

	paw := ProjectActivityWeek{Week: startWeek}
	activities := []ProjectActivity{}

	for bi, batch := range batches {
		hoursToCompleteBatch := batch.HoursToComplete
		//itemsToSchedule := batch.ScheduleItems

		for hoursToCompleteBatch > 0 {
			for ti, task := range batch.ScheduleItems {
				if len(task.ResourceIDs) == 0 {
					_, ok := exceptionsCreated[task.TaskID]
					if ok {
						continue
					}

					msg := fmt.Sprintf("Task exception: '%v' has no scheduled resources", task.TaskName)

					//---this task cannot be scheduled.  remove from hours to schedule and
					//	 log and exception
					hoursToCompleteBatch -= task.HoursToSchedule
					schedule.Exceptions = append(schedule.Exceptions, ScheduleException{
						Scope:   task.TaskID,
						Message: msg,
					})

					exceptionsCreated[task.TaskID] = msg
				}
				for _, resourceID := range task.ResourceIDs {
					if batches[bi].ScheduleItems[ti].HoursToSchedule == batches[bi].ScheduleItems[ti].HoursScheduled {
						//---task is complete...continue
						continue
					}

					hoursToAllocate := resourceHoursForWeek[resourceID]

					if hoursToAllocate > batch.ScheduleItems[ti].HoursToSchedule {
						hoursToAllocate = batch.ScheduleItems[ti].HoursToSchedule
					}

					if hoursToAllocate <= 0 {
						continue
					}

					batches[bi].ScheduleItems[ti].HoursToSchedule -= hoursToAllocate
					batches[bi].ScheduleItems[ti].HoursScheduled += hoursToAllocate
					resourceHoursForWeek[resourceID] -= hoursToAllocate

					activity := ProjectActivity{
						ProjectID:     p.ID,
						ProjectName:   p.ProjectBasics.Name,
						ResourceID:    resourceID,
						ResourceName:  rm[resourceID].Name,
						TaskID:        task.TaskID,
						TaskName:      task.TaskName,
						MilestoneID:   task.MilestoneID,
						MilestoneName: task.MilestoneName,
						HoursSpent:    hoursToAllocate,
					}

					activities = append(activities, activity)
					hoursToCompleteBatch -= hoursToAllocate
				}
			}

			//---log activities for week
			paw.Activities = activities
			weeks = append(weeks, paw)

			//---start next week
			nextWeek := calendar.GetNextWeek(paw.Week)
			paw = ProjectActivityWeek{Week: nextWeek}
			activities = []ProjectActivity{}
			resourceHoursForWeek = getResourceHoursForWeek(rm)
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
