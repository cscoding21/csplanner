import * as yup from 'yup';

export let skillForm: any = {
    resourceID: "",
    skillID: "",
    proficiency: "",
}

export const resetFormValues = (form:any) => {
    form.skillID = "";
    form.proficiency = "";
}

export const skillSchema = yup.object().shape({
    resourceID: yup.string().required(),
    skillID: yup.string().required("Skill is required"),
    proficiency: yup.lazy((value) =>
      value === ''
        ? yup.string().required("Proficiency is required")
        : yup.number().min(1).max(3).required("Proficiency is required")
    ),
  });