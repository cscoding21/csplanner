import { GetPortfolioDocument, type Portfolio } from "$lib/graphql/generated/sdk";
import { getApolloClient } from "$lib/graphql/gqlclient";


/**
 * get a list of all projects that are currently scheduled or in-flight
 * @returns a portfolio consisting of all currently scheduled projects
 */
export const getPortfolio = async (): Promise<Portfolio> => {
    const client = getApolloClient();
    return client
        .query({ query: GetPortfolioDocument })
        .then((res) => {
            if (res) {
                return res.data.getPortfolio;
            }
        })
        .catch((err) => {
            return err;
        });
};