import * as yup from 'yup';

export let registerForm: any = {
    name: "",
    email: "jeph@jmk21.com",
    password: "",
    confirmPassword: "",
}

export const resetFormValues = (form:any) => {
    form.name = "";
    form.email = "";
    form.password = "";
    form.confirmPassword = "";
}

export const registerSchema = yup.object().shape({
    name: yup.string().required(),
    email: yup.string().email().required(),
    password: yup.string().min(6).required(),
    confirmPassword: yup.string().oneOf([yup.ref('password'), ""], 'Passwords must match').required(),
  });