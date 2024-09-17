import type {
	UpdateProject,
	UpdateProjectMilestoneTemplate,
	PageAndFilter,
	UpdateProjectFeature,
	UpdateProjectMilestoneTask,
	Project,
	ProjectResults,
	UpdateProjectValue,
	UpdateProjectCost,
	UpdateProjectDaci,
	UpdateProjectBasics
} from '$lib/graphql/generated/sdk';
import type { ApolloQueryResult } from '@apollo/client/core';
import { getApolloClient } from '$lib/graphql/gqlclient';
import {
	DeleteProjectDocument,
	FindAllProjectTemplatesDocument,
	FindProjectsDocument,
	DeleteProjectFeatureDocument,
	UpdateProjectFeatureDocument,
	DeleteProjectTaskDocument,
	SetProjectMilestonesFromTemplateDocument,
	UpdateProjectDocument,
	UpdateProjectTaskDocument,
	CreateProjectDocument,
	GetProjectDocument
} from '$lib/graphql/generated/sdk';
import { coalesceToType } from '$lib/forms/helpers';
import {
	basicSchema,
	valueSchema,
	costSchema,
	daciSchema,
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
 * create a new project
 * @param input a project graph
 * @returns a result with the created project and operation status
 */
export const createProject = async (input: UpdateProject) => {
	const client = getApolloClient();
	console.log(input);

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
