import { type CreateProjectTemplateResult, type ProjecttemplateResults, type Status, type UpdateProjecttemplate, DeleteProjectTemplateDocument, FindAllProjectTemplatesDocument, UpdateProjectTemplateDocument } from "$lib/graphql/generated/sdk";
import { getApolloClient } from "$lib/graphql/gqlclient";


/**
 * returns a list of all available project milestone templates
 * @returns a list of all project templates
 */
export const findAllProjectTemplates = async (): Promise<ProjecttemplateResults> => {
    const client = getApolloClient();

    return client
        .query({ query: FindAllProjectTemplatesDocument })
        .then((pro) => {
            if (pro) {
                return pro.data.findAllProjectTemplates;
            }
        })
        .catch((err) => {
            return err;
        });
};


/**
 * update the properties of a project template
 * @param input a project template object with updated values
 * @returns a status of the operation along witht the new template object
 */
export const updateProjectTemplate = async (input: UpdateProjecttemplate): Promise<CreateProjectTemplateResult> => {
    const client = getApolloClient();

    return client
        .mutate({ mutation: UpdateProjectTemplateDocument, variables: { input } })
        .then((res) => {
            if (res) {
                return res.data.updateProjectTemplate;
            }
        })
        .catch((err) => {
            return err;
        });
};

/**
 * delete a project tempalte from the system
 * @param id the ID of the template to delete
 * @returns a status describing the outcome of the operation
 */
export const deleteProjectTemplate = async (id: string): Promise<Status> => {
    const client = getApolloClient();

    return client
        .mutate({ mutation: DeleteProjectTemplateDocument, variables: { id } })
        .then((res) => {
            if (res) {
                return res.data.deleteProjectTemplate;
            }
        })
        .catch((err) => {
            return err;
        });
};