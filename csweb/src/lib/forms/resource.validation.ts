import * as yup from 'yup';

export const resourceForm : any = {
	id: '',
	type: '',
	name: '',
	userID: '',
	role: '',
	email: '',
	profileImage: ''
};

export const resetFormValues = (form: any) => {
	form.id = '';
	form.name = '';
	form.type = '';
	form.userID = '';
	form.role = '';
	form.email = '';
	form.profileImage = '';
};

export const resourceSchema = yup.object().shape({
	id: yup.string(),
	name: yup.string().required(),
	type: yup.string(),
	userID: yup.string(),
	role: yup.string().required(),
	email: yup.string().email(),
	profileImage: yup.string()
});
