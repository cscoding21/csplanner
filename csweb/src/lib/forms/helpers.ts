import type { ObjectSchema, Maybe, AnyObject } from "yup";

function isObjKey<T>(key: any, obj: any): key is keyof T {
    return key in obj;
}

export function coalesceToType<T>(input: any, schema:ObjectSchema<Maybe<AnyObject>>): T {
    let out:T = <T>{}
    const describe = schema.describe()

    Object.keys(input).forEach(key => {
        if(isObjKey<T>(key, describe.fields)) {
            out[key] = input[key]
        }
    })

    return out
}


export const parseErrors = (err: any) =>  {
    return err.inner.reduce((acc: any, err: any) => {
        return { ...acc, [err.path]: err.message };
    }, {});
}


export const mergeErrors = (err1: any, err2: any) => {
    return {
        ...err1,
        ...err2
    }
}