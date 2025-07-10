import { FindActivityDocument, type ActivityResults, type PageAndFilter } from "$lib/graphql/generated/sdk";
import { getApolloClient } from "$lib/graphql/gqlclient";


export const findActivity = async (input: PageAndFilter):Promise<ActivityResults> => {
    const client = getApolloClient()

    return client.query({ query: FindActivityDocument, variables: { input } , fetchPolicy: "no-cache" }).then((res) => {
        if (res) {
            return res.data.findActivity;
        }
    })
    .catch((err) => {
        return err;
    });
}