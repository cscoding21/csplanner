import * as yup from 'yup';

export let loginForm: any = {
    email: "jeph@jmk21.com",
    password: ""
}

export const resetFormValues = (form:any) => {
    form.name = "";
    form.email = "";
}

export const loginSchema = yup.object().shape({
    email: yup.string().email("Email must be in a valid format").required(),
    password: yup.string().required(),
  });