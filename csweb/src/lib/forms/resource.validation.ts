import * as yup from 'yup';

export const resourceForm: any = {
	id: '',
	type: 'human',
	status: 'proposed',
	name: '',
	userID: '',
	role: '',
	email: '',
	profileImage: '',
	initialCost: 0.0,
	annualizedCost: 0.0
};

export const resetFormValues = (form: any) => {
	form.id = '';
	form.name = '';
	form.type = 'human';
	form.status = 'proposed';
	form.userID = '';
	form.role = '';
	form.email = '';
	form.profileImage = '';
	form.initialCost = 0.0;
	form.annualizedCost = 0.0;
};

export const resourceSchema = yup.object().shape({
	id: yup.string(),
	name: yup.string().required(),
	type: yup.string().required(),
	status: yup.string().required(),
	userID: yup.string(),
	role: yup.string().required(),
	email: yup.string().email(),
	profileImage: yup.string(),
	initialCost: yup.number(),
	annualizedCost: yup.number()
});
