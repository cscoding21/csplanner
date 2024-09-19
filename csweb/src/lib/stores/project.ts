import { writable, get } from 'svelte/store';
import type { ProjectResults, Projecttemplate, PageAndFilter } from '$lib/graphql/generated/sdk';
import { findProjects, findAllProjectTemplates } from '$lib/services/project';

const getInitialValue = (): ProjectResults => {
	const projects: ProjectResults = {
		results: [],
		paging: {},
		filters: {}
	};

	return projects;
};

const getInitialTemplateValue = (): Projecttemplate[] => {
	return [];
};

export const projectStore = writable(getInitialValue());
export const templateStore = writable(getInitialTemplateValue());

export const refreshProjectStore = async (input: PageAndFilter) => {
	const response = await findProjects(input);

	projectStore.set(response);
};

export const refreshTemplateStore = async () => {
	const response = await findAllProjectTemplates();

	templateStore.set(response.results as Projecttemplate[]);
};

export const getTemplateFromStore = async (nameOrID: string) => {
	await refreshTemplateStore();

	const templates = get(templateStore);

	for (let i = 0; i < templates.length; i++) {
		const item = templates[i];

		if (item.id === nameOrID || item.name === nameOrID) {
			return item;
		}
	}

	return null;
};
