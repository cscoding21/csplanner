import * as yup from 'yup';
import type { UpdateProject } from '$lib/graphql/generated/sdk';

export const basicSchema = yup.object().shape({
	name: yup.string().required('Project name is required'),
	description: yup.string().required('Description is required'),
	ownerID: yup.string().required('Project owner is required'),
	startDate: yup.date().optional(),
});

export const valueSchema = yup.object().shape({
	discountRate: yup
		.number()
		.min(0.0, 'Discount rate cannot be lower than zero')
		.max(20.0, 'Discount rate cannot be higher than 20%')
		.required(),
	isCapitalized: yup.bool().required()
});

export const valueLineSchema = yup.object().shape({
	id: yup.string(),
	projectID: yup.string().required(),
	fundingSource: yup.string().required(),
	valueCategory: yup.string().required(),
	yearOneValue: yup.number().nullable(),
	yearTwoValue: yup.number().nullable(),
	yearThreeValue: yup.number().nullable(),
	yearFourValue: yup.number().nullable(),
	yearFiveValue: yup.number().nullable(),
	description: yup.string().required()
});

export const costSchema = yup.object().shape({
	ongoing: yup.number().min(0.0, 'Ongoing cost cannot be lower than zero').nullable(),
	blendedRate: yup.number().min(0.0, 'Blended rate cannot be lower than zero').nullable()
});

export const daciSchema = yup.object().shape({
	driverIDs: yup.array().nullable(),
	approverIDs: yup.array().nullable(),
	contributorIDs: yup.array().nullable(),
	informedIDs: yup.array().nullable()
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
	requiredSkillID: yup.string().required(),
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
		},
		projectValue: {
			isCapitalized: false,
			discountRate: 7.0
		},
		projectCost: {
			ongoing: 0.0,
			blendedRate: 0.0
		},
		projectDaci: {
			driverIDs: [],
			approverIDs: [],
			contributorIDs: [],
			informedIDs: []
		},
		projectFeatures: [],
		projectMilestones: []
	};
};

export const valueDefaultForm = () => {
	return {
		discountRate: 0.0,
		isCapitalized: false
	}
} 

export const valueLineDefaultForm = (projectID: string): any => {
	return {
		projectID: projectID,
		id: "",
		fundingSource: "",
		valueCategory: "",
		yearOneValue: 0.0,
		yearTwoValue: 0.0,
		yearThreeValue: 0.0,
		yearFourValue: 0.0,
		yearFiveValue: 0.0
	}
} 

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
