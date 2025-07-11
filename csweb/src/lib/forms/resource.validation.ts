import * as yup from 'yup';

export const resourceForm: any = {
	id: '',
	type: 'human',
	status: 'proposed',
	name: '',
	userID: '',
	roleID: '',
	email: '',
	profileImage: '',
	initialCost: 0.0,
	annualizedCost: 0.0,
	availableHoursPerWeek: 0
};

export const resetFormValues = (form: any) => {
	form.id = '';
	form.name = '';
	form.type = 'human';
	form.status = 'proposed';
	form.userID = '';
	form.roleID = '';
	form.email = '';
	form.profileImage = '';
	form.initialCost = 0.0;
	form.annualizedCost = 0.0;
	form.availableHoursPerWeek = 0;
};

export const resourceSchema = yup.object().shape({
	id: yup.string(),
	name: yup.string().required(),
	type: yup.string().required(),
	status: yup.string().required(),
	userID: yup.string(),
	roleID: yup.string().nullable(),
	email: yup.string().email(),
	profileImage: yup.string().nullable(),
	initialCost: yup.number(),
	annualizedCost: yup.number(),
	availableHoursPerWeek: yup.number()
});


export const roleForm: any = {
	id: '',
	name: '',
	description: '',
	hourlyRate: 0.0,
	defaultSkills: []
};

export const roleSchema = yup.object().shape({
	id: yup.string(),
	name: yup.string().required(),
	description: yup.string(),
	hourlyRate: yup.number(),
	defaultSkills: yup.array().of(yup.object({
		parentID: yup.string(),
		id: yup.string(),
		skillID: yup.string().required(),
		proficiency: yup.number().required()
	}))
});
