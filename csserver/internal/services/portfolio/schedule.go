package portfolio

import (
	"context"
	"csserver/internal/calendar"
	"csserver/internal/services/project"
	"csserver/internal/services/resource"
	"fmt"
	"time"

	"github.com/google/uuid"
)

// ScheduleProject create a schedule for a single project "in a vacuum"
func (service *PortfolioService) ScheduleProject(ctx context.Context, p *project.Project, startDate time.Time) (ProjectSchedule, error) {
	rm, err := service.GetResourceMap(ctx)
	if err != nil {
		return ProjectSchedule{}, err
	}

	return scheduleProject(p, startDate, rm)
}

func scheduleProject(p *project.Project, startDate time.Time, rm map[string]resource.Resource) (ProjectSchedule, error) {
	schedule := ProjectSchedule{
		ProjectID:   p.ID,
		ProjectName: p.ProjectBasics.Name,
	}
	if len(p.ProjectMilestones) == 0 {
		return schedule, nil
	}

	p.ProjectBasics.StartDate = &startDate

	batches := getScheduleItems(p)
	currentWeek := calendar.GetWeek(startDate)
	schedule.Begin = currentWeek.Begin
	weeks := []*ProjectActivityWeek{}

	resourceHoursForWeek := getResourceHoursForWeek(rm, currentWeek)
	exceptionsCreated := make(map[string]string)
	thingCount := 0

	activities := []ProjectActivity{}

	for bi, batch := range batches {
		hoursToCompleteBatch := batch.HoursToComplete

		for hoursToCompleteBatch > 0 {
			//---safety check
			if thingCount > 100 {
				msg := fmt.Sprintf("Task exception: '%v' has no scheduled resources", "schedule:overflow")

				schedule.Exceptions = append(schedule.Exceptions, ScheduleException{
					Scope:   "schedule:overflow",
					Message: "for loop did not terminate",
				})

				exceptionsCreated["schedule:overflow"] = msg

				break
			}

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
			weeks = append(weeks, &ProjectActivityWeek{
				WeekNumber: currentWeek.WeekNumber,
				Year:       currentWeek.Year,
				Begin:      currentWeek.Begin,
				End:        currentWeek.End,
				Activities: activities,
			})

			//---start next week
			currentWeek = calendar.GetNextWeek(currentWeek)
			activities = []ProjectActivity{}
			resourceHoursForWeek = getResourceHoursForWeek(rm, currentWeek)

			thingCount++
		}
	}

	schedule.End = weeks[len(weeks)-1].End
	schedule.ProjectActivityWeeks = weeks

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

func getResourceHoursForWeek(rm map[string]resource.Resource, week calendar.CSWeek) map[string]int {
	outMap := make(map[string]int)

	for k, v := range rm {
		outMap[k] = v.AvailableHoursPerWeek
	}

	return outMap
}
