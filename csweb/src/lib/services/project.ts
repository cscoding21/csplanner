import type {
	UpdateProject,
	UpdateProjectMilestoneTemplate,
	PageAndFilter,
	UpdateProjectFeature,
	UpdateProjectMilestoneTask,
	ProjectEnvelope,
	ProjectResults,
	Status,
	UpdateProjectValue,
	UpdateProjectCost,
	UpdateProjectDaci,
	UpdateProjectBasics,
	UpdateProjectValueLine,
	CreateProjectResult,
	ProjecttemplateResults,
	ProjectScheduleResult,
	ValidationResult,
	UpdateNewProject
} from '$lib/graphql/generated/sdk';
import { getApolloClient } from '$lib/graphql/gqlclient';
import {
	DeleteProjectDocument,
	FindProjectsDocument,
	DeleteProjectFeatureDocument,
	UpdateProjectFeatureDocument,
	DeleteProjectTaskDocument,
	SetProjectMilestonesFromTemplateDocument,
	UpdateProjectDocument,
	UpdateProjectTaskDocument,
	CreateProjectDocument,
	GetProjectDocument,
	CalculateProjectScheduleDocument,
	CheckProjectStatusDocument,
	SetProjectStatusDocument,
	UpdateProjectValueLineDocument,
	DeleteProjectValueLineDocument
} from '$lib/graphql/generated/sdk';
import { coalesceToType } from '$lib/forms/helpers';
import { safeArray } from '$lib/utils/helpers';
import {
	basicSchema,
	valueSchema,
	costSchema,
	taskSchema
} from '$lib/forms/project.validation';

/**
 * return an list of projects based on passed in filter and paging criteria
 * @param input object with paging and filtering details
 * @returns a paged list of projects
 */
export const findProjects = async (input: PageAndFilter): Promise<ProjectResults> => {
	const client = getApolloClient();

	return client
		.query({ query: FindProjectsDocument, variables: { input } })
		.then((pro) => {
			if (pro) {
				return pro.data.findProjects;
			}
		})
		.catch((err) => {
			return err;
		});
};

/**
 * Get a project from the backend based on the passed in ID
 * @param id the ID of the project to retrieve
 * @returns the full object graph of the specified project
 */
export const getProject = async (id: any): Promise<ProjectEnvelope> => {
	const client = getApolloClient();

	return client
		.query({ query: GetProjectDocument, variables: { id } })
		.then((pro) => {
			if (pro) {
				return pro.data.getProject;
			}
		})
		.catch((err) => {
			return err;
		});
};


/**
 * check to see if a project status can be updated to a given state
 * @param projectID the project ID to check
 * @param newStatus the status of the update dry-run
 * @returns a validation result
 */
export const checkProjectStatus = async (projectID: string, newStatus: string): Promise<ValidationResult> => {
	const client = getApolloClient();

	return client
		.query({ query: CheckProjectStatusDocument, variables: { projectID, newStatus } })
		.then((pro) => {
			if (pro) {
				return pro.data.checkProjectStatus;
			}
		})
		.catch((err) => {
			return err;
		});
};

/**
 * create a new project
 * @param input a project graph
 * @returns a result with the created project and operation status
 */
export const createProject = async (input: UpdateNewProject): Promise<CreateProjectResult> => {
	const client = getApolloClient();

	return client
		.mutate({ mutation: CreateProjectDocument, variables: { project: input } })
		.then((pro) => {
			if (pro) {
				return pro.data.createProject;
			}
		})
		.catch((err) => {
			return err;
		});
};

/**
 * Delete a project from the portfolio
 * @param id the ID of the project to delete
 * @returns a status object indicating whether the operation was successful
 */
export const deleteProject = async (id: string): Promise<Status> => {
	const client = getApolloClient();

	return client
		.mutate({ mutation: DeleteProjectDocument, variables: { id } })
		.then((pro) => {
			if (pro) {
				return pro.data.deleteProject;
			}
		})
		.catch((err) => {
			return err;
		});
};

/**
 * Updates the complete properties of a project
 * @param input the graph of the project to update
 * @returns a status of the operation and the updated project entity
 */
export const updateProject = async (input: UpdateProject): Promise<CreateProjectResult> => {
	const client = getApolloClient();

	return client
		.mutate({ mutation: UpdateProjectDocument, variables: { project: input } })
		.then((pro) => {
			if (pro) {
				return pro.data.updateProject;
			}
		})
		.catch((err) => {
			return err;
		});
};

/**
 * try to update the status of a project
 * @param projectID the project ID to update
 * @param newStatus the new status
 * @returns an update project object
 */
export const setProjectStatus = async (projectID: string, newStatus: string): Promise<CreateProjectResult> => {
	const client = getApolloClient();

	return client
		.mutate({ mutation: SetProjectStatusDocument, variables: { projectID, newStatus } })
		.then((pro) => {
			if (pro) {
				return pro.data.setProjectStatus;
			}
		})
		.catch((err) => {
			return err;
		});
};

/**
 * Updates the "basics" section of a project while leaving other sections in tact
 * @param id the ID of the project to update
 * @param input the basics section of the project
 * @returns a status of the operation and the updated project entity
 */
export const updateProjectBasics = async (
	id: string,
	input: UpdateProjectBasics
): Promise<CreateProjectResult> => {
	const proj = await getProject(id);

	const updateProj = convertProjectToUpdateProject(proj);
	updateProj.projectBasics = input;

	return updateProject(updateProj);
};

/**
 * Updates the "value" section of a project while leaving other sections in tact
 * @param id the ID of the project to update
 * @param input the value section of the project
 * @returns a status of the operation and the updated project entity
 */
export const updateProjectValue = async (
	id: string,
	input: UpdateProjectValue
): Promise<CreateProjectResult> => {
	const proj = await getProject(id);

	const updateProj = convertProjectToUpdateProject(proj);
	updateProj.projectValue = input;

	return updateProject(updateProj);
};

/**
 * Updates the "cost" section of a project while leaving other sections in tact
 * @param id the ID of the project to update
 * @param input the cost section of the project
 * @returns a status of the operation and the updated project entity
 */
export const updateProjectCost = async (
	id: string,
	input: UpdateProjectCost
): Promise<CreateProjectResult> => {
	const proj = await getProject(id);

	const updateProj = convertProjectToUpdateProject(proj);
	updateProj.projectCost = input;

	return updateProject(updateProj);
};

/**
 * Updates the "DACI" section of a project while leaving other sections in tact
 * @param id the ID of the project to update
 * @param input the DACI section of the project
 * @returns a status of the operation and the updated project entity
 */
export const updateProjectDaci = async (
	id: string,
	input: UpdateProjectDaci
): Promise<CreateProjectResult> => {
	const proj = await getProject(id);

	console.log("updateProjectDaci", proj)
	const updateProj = convertProjectToUpdateProject(proj);
	updateProj.projectDaci = input;

	return updateProject(updateProj);
};

/**
 * Copies the items from the selected template into the target project
 * @param input The project id and template ID
 * @returns a status of the operation and the updated project entity
 */
export const setProjectMilestonesFromTemplate = async (
	input: UpdateProjectMilestoneTemplate
): Promise<ProjecttemplateResults> => {
	const client = getApolloClient();

	return client
		.mutate({ mutation: SetProjectMilestonesFromTemplateDocument, variables: { input } })
		.then((pro) => {
			if (pro) {
				return pro.data.setProjectMilestonesFromTemplate;
			}
		})
		.catch((err) => {
			return err;
		});
};

/**
 * Create or update a feature for a given project
 * @param input the project id and feature to add
 * @returns a status of the operation and the updated project entity
 */
export const updateProjectFeature = async (
	input: UpdateProjectFeature
): Promise<CreateProjectResult> => {
	const client = getApolloClient();

	return client
		.mutate({ mutation: UpdateProjectFeatureDocument, variables: { input } })
		.then((pro) => {
			if (pro) {
				return pro.data.updateProjectFeature;
			}
		})
		.catch((err) => {
			return err;
		});
};

/**
 * Remove the specified feature from the target project
 * @param projectID the project ID to update
 * @param featureID the feature ID to delete
 * @returns a status of the operation and the updated project entity
 */
export const deleteProjectFeature = async (
	projectID: string,
	featureID: string
): Promise<CreateProjectResult> => {
	const client = getApolloClient();

	return client
		.mutate({ mutation: DeleteProjectFeatureDocument, variables: { projectID, featureID } })
		.then((pro) => {
			if (pro) {
				return pro.data.deleteProjectFeature;
			}
		})
		.catch((err) => {
			return err;
		});
};


/**
 * Create or update a value line for a given project
 * @param input the project id and value line to add
 * @returns a status of the operation and the updated project entity
 */
export const updateProjectValueLine = async (
	input: UpdateProjectValueLine
): Promise<CreateProjectResult> => {
	const client = getApolloClient();

	return client
		.mutate({ mutation: UpdateProjectValueLineDocument, variables: { input } })
		.then((pro) => {
			if (pro) {
				return pro.data.updateProjectValueLine;
			}
		})
		.catch((err) => {
			return err;
		});
};


/**
 * Remove the specified value line from the target project
 * @param projectID the project ID to update
 * @param valueLineID the value line ID to delete
 * @returns a status of the operation and the updated project entity
 */
export const deleteProjectValueLine = async (
	projectID: string,
	valueLineID: string
): Promise<CreateProjectResult> => {
	const client = getApolloClient();

	return client
		.mutate({ mutation: DeleteProjectValueLineDocument, variables: { projectID, valueLineID } })
		.then((pro) => {
			if (pro) {
				return pro.data.deleteProjectValueLine;
			}
		})
		.catch((err) => {
			return err;
		});
};


/**
 * Create or update a task for a given project and milestone
 * @param input the project, milestone, of the task to update and the task properties
 * @returns a status of the operation and the updated project entity
 */
export const updateProjectTask = async (
	input: UpdateProjectMilestoneTask
): Promise<CreateProjectResult> => {
	const client = getApolloClient();

	input = coalesceToType<UpdateProjectMilestoneTask>(input, taskSchema);

	return client
		.mutate({ mutation: UpdateProjectTaskDocument, variables: { input } })
		.then((pro) => {
			if (pro) {
				return pro.data.updateProjectTask;
			}
		})
		.catch((err) => {
			return err;
		});
};

/**
 * Delete a task from a project milestone
 * @param projectID the ID of the project to update
 * @param milestoneID the milestone frmo which to remove the task
 * @param taskID the ID of the task to delete
 * @returns a status of the operation and the updated project entity
 */
export const deleteProjectTask = async (
	projectID: string,
	milestoneID: string,
	taskID: string
): Promise<CreateProjectResult> => {
	const client = getApolloClient();

	return client
		.mutate({ mutation: DeleteProjectTaskDocument, variables: { projectID, milestoneID, taskID } })
		.then((pro) => {
			if (pro) {
				return pro.data.deleteProjectTask;
			}
		})
		.catch((err) => {
			return err;
		});
};


/**
 * returns a list of all available project milestone templates
 * @returns a list of all project templates
 */
export const calculateProjectSchedule = async (projectID: string, startDate :Date): Promise<ProjectScheduleResult> => {
	const client = getApolloClient();

	return client
		.query({ query: CalculateProjectScheduleDocument, variables: { projectID, startDate }})
		.then((pro) => {
			if (pro) {
				return pro.data.calculateProjectSchedule;
			}
		})
		.catch((err) => {
			return err;
		});
};

/**
 * convert a ProjectEnvelope object to it's updatable counterpart
 * @param project a ProjectEnvelope object to be converted
 * @returns an UpdateProject object
 */
export const convertProjectToUpdateProject = (project: ProjectEnvelope): UpdateProject => {
	const up: UpdateProject = {
		id: project.meta?.id
	};

	console.log("updateProject", coalesceToType<UpdateProjectValue>(project.data?.projectValue, valueSchema))

	up.projectBasics = coalesceToType<UpdateProjectBasics>(project.data?.projectBasics, basicSchema);
	up.projectValue = coalesceToType<UpdateProjectValue>(project.data?.projectValue, valueSchema);
	up.projectCost = coalesceToType<UpdateProjectCost>(project.data?.projectCost, costSchema);

	const df = {
		driverIDs: safeArray(project.data?.projectDaci?.driver?.map(d => d?.id) as string[]),
		approverIDs: safeArray(project.data?.projectDaci?.approver?.map(d => d?.id) as string[]),
		contributorIDs: safeArray(project.data?.projectDaci?.contributor?.map(d => d?.id) as string[]),
		informedIDs: safeArray(project.data?.projectDaci?.informed?.map(d => d?.id) as string[])
	};

	up.projectDaci = df;

	return up;
};
