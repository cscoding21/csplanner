import * as yup from 'yup';

export const skillForm: any = {
	resourceID: '',
	skillID: '',
	proficiency: 2
};

export const resetFormValues = (form: any) => {
	form.skillID = '';
	form.proficiency = 2;
};

export const skillSchema = yup.object().shape({
	resourceID: yup.string().required(),
	skillID: yup.string().required('Skill is required'),
	proficiency: yup.number().min(1).max(3).required('Proficiency is required')
});
