import * as yup from 'yup';
import type { UpdateProject } from '$lib/graphql/generated/sdk';

export const basicSchema = yup.object().shape({
	name: yup.string().required('Project name is required'),
	description: yup.string().required('Description is required'),
	ownerID: yup.string().required('Project owner is required'),
	status: yup.string().required('Project status is required')
});

export const valueSchema = yup.object().shape({
	fundingSource: yup.string(),
	discountRate: yup
		.number()
		.min(0.0, 'Discount rate cannot be lower than zero')
		.max(20.0, 'Discount rate cannot be higher than 20%')
		.nullable(),
	yearOneValue: yup.number().nullable(),
	yearTwoValue: yup.number().nullable(),
	yearThreeValue: yup.number().nullable(),
	yearFourValue: yup.number().nullable(),
	yearFiveValue: yup.number().nullable()
});

export const costSchema = yup.object().shape({
	ongoing: yup.number().min(0.0, 'Ongoing cost cannot be lower than zero').nullable(),
	blendedRate: yup.number().min(0.0, 'Blended rate cannot be lower than zero').nullable()
});

export const daciSchema = yup.object().shape({
	driver: yup.array().nullable(),
	approver: yup.array().nullable(),
	contributor: yup.array().nullable(),
	informed: yup.array().nullable()
});

export const featureSchema = yup.object().shape({
	projectID: yup.string().required(),
	name: yup.string().required(),
	description: yup.string().required(),
	priority: yup.string().required(),
	status: yup.string().required()
});

export const taskSchema = yup.object().shape({
	projectID: yup.string().required(),
	milestoneID: yup.string().required(),
	id: yup.string(),
	name: yup.string().required(),
	description: yup.string().required(),
	hourEstimate: yup.number().required(),
	requiredSkillIDs: yup.array(),
	resourceIDs: yup.array(),
	status: yup.string().required(),
	startDate: yup.date().nullable(),
	endDate: yup.date().nullable()
});

export const getDefaultProject = (): UpdateProject => {
	return {
		projectBasics: {
			name: '',
			description: '',
			ownerID: '',
			status: 'draft'
		},
		projectValue: {
			fundingSource: '',
			discountRate: 7.0,
			yearOneValue: 0.0,
			yearTwoValue: 0.0,
			yearThreeValue: 0.0,
			yearFourValue: 0.0,
			yearFiveValue: 0.0
		},
		projectCost: {
			ongoing: 0.0,
			blendedRate: 0.0
		},
		projectDaci: {
			driver: [],
			approver: [],
			contributor: [],
			informed: []
		},
		projectFeatures: [],
		projectMilestones: []
	};
};

export const getNewFeature = (projectID: string): any => {
	return {
		projectID: projectID,
		id: '',
		name: '',
		priority: '',
		status: 'proposed',
		description: ''
	};
};
