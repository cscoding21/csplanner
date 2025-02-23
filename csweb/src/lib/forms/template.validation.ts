import * as yup from 'yup';

export const templateSchema = yup.object().shape({
    id: yup.string().optional(),
    name: yup.string().required(),
    description: yup.string().required(),
    phases: yup.array().of(yup.object({
        id: yup.string().optional(),
        name: yup.string().required(),
        description: yup.string().required(),
        tasks: yup.array().of(yup.object({
            id: yup.string().optional(),
            name: yup.string().required(),
            description: yup.string().required(),
            requiredSkillID: yup.string().required()
        })
    )}))
});
