package schedule

import (
	"context"
	"csserver/internal/calendar"
	"csserver/internal/hashstructure"
	"csserver/internal/services/project"
	"csserver/internal/services/project/ptypes/milestonestatus"
	"csserver/internal/services/resource"
	"fmt"
	"time"

	"github.com/google/uuid"
)

// ScheduleProject create a schedule for a single project "in a vacuum"
func (service *ScheduleService) CalculateProjectSchedule(ctx context.Context, p *project.Project, startDate time.Time, rm map[string]resource.Resource) (Schedule, error) {
	ps := Schedule{
		ProjectID:   p.ID,
		ProjectName: p.ProjectBasics.Name,
	}

	exceptions := validateProjectForScheduling(*p, rm)
	if len(exceptions) > 0 {
		ps.Exceptions = append(ps.Exceptions, exceptions...)
		return ps, nil
	}

	ram := NewResourceAllocationMapFromResourceMap(rm)

	return ScheduleProjectAlgo(p, startDate, ram)
}

func ScheduleProjectAlgo(p *project.Project, startDate time.Time, ram ResourceAllocationMap) (Schedule, error) {
	projectID := p.ID
	schedule := Schedule{
		ProjectID:   projectID,
		ProjectName: p.ProjectBasics.Name,
	}
	if len(p.ProjectMilestones) == 0 {
		return schedule, nil
	}

	p.ProjectBasics.StartDate = &startDate

	batches := GetScheduleItems(p)
	currentWeek := calendar.GetWeek(startDate)
	schedule.Begin = currentWeek.Begin
	weeks := []*ProjectActivityWeek{}

	//resourceHoursForWeek := GetResourceHoursForWeek(rm, currentWeek)
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

					// hoursToAllocate := resourceHoursForWeek[resourceID]
					resourceToAllocate := ram.GetResource(currentWeek, projectID, resourceID)
					if resourceToAllocate == nil {
						//---TODO: handle this
					}

					hoursToAllocate := resourceToAllocate.HoursAvailableForProject

					if hoursToAllocate > batch.ScheduleItems[ti].HoursToSchedule {
						hoursToAllocate = batch.ScheduleItems[ti].HoursToSchedule
					}

					if hoursToAllocate <= 0 {
						continue
					}

					batches[bi].ScheduleItems[ti].HoursToSchedule -= hoursToAllocate
					batches[bi].ScheduleItems[ti].HoursScheduled += hoursToAllocate
					//resourceHoursForWeek[resourceID] -= hoursToAllocate
					ram.ReduceResourceProjectHours(currentWeek, projectID, resourceID, hoursToAllocate)

					activity := ProjectActivity{
						ProjectID:     p.ID,
						ProjectName:   p.ProjectBasics.Name,
						ResourceID:    resourceID,
						ResourceName:  resourceToAllocate.ResourceName,
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
			//resourceHoursForWeek = GetResourceHoursForWeek(rm, currentWeek)

			thingCount++
		}
	}

	schedule.ProjectActivityWeeks = weeks

	if len(weeks) > 0 {
		schedule.End = weeks[len(weeks)-1].End
	}

	hash, err := hashstructure.Hash(schedule, hashstructure.FormatV2, nil)
	if err != nil {
		//---TODO: figure this out
	}

	schedule.Hash = hash

	return schedule, nil
}

func GetScheduleItems(p *project.Project) []ScheduleBatch {
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

func newScheduleException(scope, msg string, t ScheduleExceptionType, level int) ScheduleException {
	return ScheduleException{
		Type:    t,
		Level:   level,
		Scope:   scope,
		Message: msg,
	}
}

func validateProjectForScheduling(p project.Project, rm map[string]resource.Resource) []ScheduleException {
	out := []ScheduleException{}

	for _, m := range p.ProjectMilestones {
		for _, t := range m.Tasks {
			//---no need to validate removed tasks
			if t.Status == milestonestatus.Removed {
				continue
			}

			//---tasks require a skill and resources to be assigned
			if len(t.RequiredSkillID) == 0 {
				out = append(out, newScheduleException(
					*t.ID,
					fmt.Sprintf("Task '%s' does not have a skill assigned", t.Name),
					TaskAssignedResourceMissingSkill,
					1,
				))
			}
			if len(t.ResourceIDs) == 0 {
				out = append(out, newScheduleException(
					*t.ID,
					fmt.Sprintf("Task '%s' has no assigned resources", t.Name),
					TaskLacksAssignedResource,
					1))
			}

			for _, rss := range t.ResourceIDs {
				resource := rm[rss]
				skill := resource.GetSkill(t.RequiredSkillID)

				if skill == nil {
					out = append(out, newScheduleException(
						*t.ID,
						fmt.Sprintf("Task '%s' resoure, %s, lacks required skill %s", t.Name, resource.Name, t.RequiredSkillID),
						TaskAssignedResourceMissingSkill,
						1))
				}
			}
		}
	}

	return out
}