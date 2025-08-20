package project

import (
	"context"
	"slices"

	"csserver/internal/common"
	"csserver/internal/services/organization"
	"csserver/internal/services/project/ptypes/milestonestatus"
	"csserver/internal/services/projecttemplate"
	"csserver/internal/services/resource"
	"csserver/internal/utils"

	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
)

// DeleteTaskFromProject if the tasks exists in the project object graph...remove it
func (s *ProjectService) DeleteTaskFromProject(
	ctx context.Context,
	projectID string,
	milestoneID string,
	taskID string,
	resourceMap map[string]resource.Resource,
	roleMap map[string]resource.Role,
	org organization.Organization,
) (common.UpdateResult[*common.BaseModel[Project]], error) {

	result, err := s.GetProjectByID(ctx, projectID)
	if err != nil {
		return common.NewFailingUpdateResult[*common.BaseModel[Project]](nil, err)
	}

	project := result.Data

	updatedProject, deletedTask := deleteTaskFromProjectGraph(project, milestoneID, taskID)
	updatedProject.PerformAllCalcs(org, resourceMap, roleMap)

	s.pubsub.StreamPublish(ctx,
		string(ProjectIdentifier),
		"task",
		"deleted",
		map[string]any{
			"id":           milestoneID,
			"project_id":   projectID,
			"project_name": project.ProjectBasics.Name,
			"name":         deletedTask.Name,
		})

	return s.UpdateProject(ctx, updatedProject)
}

// deleteTaskFromProjectGraph if the milestone exists in the project object graph...remove the specified task
func deleteTaskFromProjectGraph(project Project, milestoneID string, taskID string) (Project, ProjectMilestoneTask) {
	var task ProjectMilestoneTask
	//---find the task in the object graph
	for i, m := range project.ProjectMilestones {
		if *m.ID == milestoneID {
			//---found milestone
			for j, t := range m.Tasks {
				if t.ID != nil && *t.ID == taskID {
					task = *t
					project.ProjectMilestones[i].Tasks = slices.Delete(project.ProjectMilestones[i].Tasks, j, j+1)
					break
				}
			}
		}
	}

	return project, task
}

// UpdateProjectTask update a project task if it exists or create it if it doesn't
func (s *ProjectService) UpdateProjectTask(
	ctx context.Context,
	projectID string,
	milestoneID string,
	task ProjectMilestoneTask,
	resourceMap map[string]resource.Resource,
	roleMap map[string]resource.Role,
	org organization.Organization) (common.UpdateResult[*common.BaseModel[Project]], error) {

	result, err := s.GetProjectByID(ctx, projectID)
	if err != nil {
		return common.NewFailingUpdateResult[*common.BaseModel[Project]](nil, err)
	}

	dg := ""
	op := "created"
	project := result.Data

	//---DO THE UPDATE
	updatedProject, existingTask := UpdateTaskFromProjectGraph(project, milestoneID, task)
	updatedProject.PerformAllCalcs(org, resourceMap, roleMap)

	if existingTask != nil {
		dg = utils.GetDiffGraph(*existingTask, task)
		op = "updated"
	}

	s.pubsub.StreamPublish(ctx,
		string(ProjectIdentifier),
		"task",
		op,
		map[string]any{
			"id":           task.ID,
			"project_id":   projectID,
			"project_name": project.ProjectBasics.Name,
			"name":         task.Name,
			"diffs":        dg,
		})

	return s.UpdateProject(ctx, updatedProject)
}

// UpdateProjectTaskStatus update the status of a given project task
func (s *ProjectService) UpdateProjectTaskStatus(
	ctx context.Context,
	projectID string,
	milestoneID string,
	taskID string,
	status milestonestatus.MilestoneStatus,
	resourceMap map[string]resource.Resource,
	roleMap map[string]resource.Role,
	org organization.Organization) (common.UpdateResult[*common.BaseModel[Project]], error) {

	result, err := s.GetProjectByID(ctx, projectID)
	if err != nil {
		return common.NewFailingUpdateResult[*common.BaseModel[Project]](nil, err)
	}

	dg := ""
	op := "updated"
	project := result.Data
	var existingTask *ProjectMilestoneTask
	var updatedTask *ProjectMilestoneTask

	//---DO THE UPDATE
	for i, m := range project.ProjectMilestones {
		if *m.ID == milestoneID {
			for j, t := range m.Tasks {
				if *t.ID == taskID {
					existingTask, _ = utils.DeepCopy(*project.ProjectMilestones[i].Tasks[j])
					project.ProjectMilestones[i].Tasks[j].Status = status

					updatedTask, _ = utils.DeepCopy(*project.ProjectMilestones[i].Tasks[j])
					break
				}
			}
		}
	}

	if existingTask != nil {
		dg = utils.GetDiffGraph(*existingTask, *updatedTask)
		op = "updated"
	}

	s.pubsub.StreamPublish(ctx,
		string(ProjectIdentifier),
		"task",
		op,
		map[string]any{
			"id":           updatedTask.ID,
			"project_id":   projectID,
			"project_name": project.ProjectBasics.Name,
			"name":         updatedTask.Name,
			"diffs":        dg,
		})

	project.PerformAllCalcs(org, resourceMap, roleMap)

	return s.UpdateProject(ctx, project)
}

// updateTaskFromProjectGraph if the milestone exists in the project object graph...update the specified task.  Otherwise, add the task
func UpdateTaskFromProjectGraph(project Project, milestoneID string, task ProjectMilestoneTask) (Project, *ProjectMilestoneTask) {
	var outTask *ProjectMilestoneTask
	//---DO THE UPDATE
	for i, m := range project.ProjectMilestones {
		log.Infof("Looping project milestone %s", *m.ID)
		if *m.ID == milestoneID {
			log.Infof("Checking tasks")
			if task.ID != nil && len(*task.ID) > 0 {
				log.Infof("Task id = %s...performing edit", *task.ID)
				//---this is an edit
				for j, t := range m.Tasks {
					if t.ID != nil && *t.ID == *task.ID {
						outTask, _ = utils.DeepCopy(*project.ProjectMilestones[i].Tasks[j])
						project.ProjectMilestones[i].Tasks[j] = &task
						break
					}
				}

				//---we're done here
				break
			} else {
				//---this is an add
				log.Infof("Task id = %v...performing insert", task.ID)
				task.ID = utils.ValToRef(uuid.New().String())
				project.ProjectMilestones[i].Tasks = append(project.ProjectMilestones[i].Tasks, &task)

				//---finished here
				break
			}
		}
	}

	return project, outTask
}

// SetProjectMilestonesFromTemplate assign the details of a project plan to a specific project
func (s *ProjectService) SetProjectMilestonesFromTemplate(
	ctx context.Context,
	projectID string,
	template projecttemplate.Projecttemplate,
	resourceMap map[string]resource.Resource,
	roleMap map[string]resource.Role,
	org organization.Organization,
) (common.UpdateResult[*common.BaseModel[Project]], error) {

	project, err := s.GetProjectByID(ctx, projectID)
	if err != nil {
		return common.NewFailingUpdateResult[*common.BaseModel[Project]](nil, err)
	}

	project.Data.ProjectMilestones = make([]*ProjectMilestone, 0)

	for _, p := range template.Phases {
		milestone := ProjectMilestone{
			ID: utils.ValToRef(uuid.New().String()),
			Phase: &ProjectMilestonePhase{
				ID:          p.ID,
				Name:        p.Name,
				Order:       p.PhaseOrder,
				Description: p.Description,
			},
			Tasks: []*ProjectMilestoneTask{},
		}

		for _, t := range p.Tasks {
			task := ProjectMilestoneTask{
				ID:              utils.ValToRef(uuid.New().String()),
				Name:            t.Name,
				Description:     t.Description,
				Status:          milestonestatus.New,
				RequiredSkillID: t.RequiredSkillID,
			}

			milestone.Tasks = append(milestone.Tasks, &task)
		}

		project.Data.ProjectMilestones = append(project.Data.ProjectMilestones, &milestone)
	}

	project.Data.PerformAllCalcs(org, resourceMap, roleMap)
	return s.UpdateProject(ctx, project.Data)
}
