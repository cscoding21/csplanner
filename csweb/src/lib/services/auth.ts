import { LoginDocument, type UpdateLogin } from '$lib/graphql/generated/sdk'
import { getApolloClient} from "$lib/graphql/gqlclient"
import { setLoggedInUser } from "$lib/stores/auth"
import { error } from '@sveltejs/kit'


export const login = async (creds: UpdateLogin) => {
    const client = getApolloClient()

    return client.mutate({ mutation: LoginDocument, variables: { creds } }).then(result => {
        setLoggedInUser(result.data.login)

        return result
    }).catch(err => {
        console.error(err)
        
        return error(err)
    });
}