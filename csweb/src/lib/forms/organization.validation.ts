import * as yup from 'yup';

export const orgForm: any = {
    discountRate: 0.0,
    hoursPerWeek: 0,
    focusFactor: 0.0,
    commsCoefficient: 0.0,
    genericBlendedHourlyRate: 0.0,
    workingHoursPerYear: 0.0
};

export const orgSchema = yup.object().shape({
    discountRate: yup.number().required(),
    hoursPerWeek: yup.number().required(),
    focusFactor: yup.number(),
    commsCoefficient: yup.number().required(),
    genericBlendedHourlyRate: yup.number().required(),
    workingHoursPerYear: yup.number().required()
});
