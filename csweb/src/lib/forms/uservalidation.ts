import * as yup from 'yup';

export const userForm: any = {
    firstName: '',
    lastName: '',
    profileImage: '',
    email: '',
    id: ''
};

export const userFormSchema = yup.object().shape({
    firstName: yup.string().required(),
    lastName: yup.string().required(),
    profileImage: yup.string(),
    email: yup.string().email().required(),
    id: yup.string().required(),
});
