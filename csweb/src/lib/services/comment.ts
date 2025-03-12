import { getApolloClient } from "$lib/graphql/gqlclient";
import type { UpdateComment, UpdateCommentReply, UpdateCommentEmote, CommentResults, CommentEnvelope } from "$lib/graphql/generated/sdk";
import { FindProjectCommentsDocument, CreateProjectCommentDocument, CreateProjectCommentReplyDocument, DeleteProjectCommentDocument, ToggleEmoteDocument, UpdateProjectCommentDocument, GetCommentThreadDocument } from "$lib/graphql/generated/sdk";

/**
 * find all comments for a project      
 * @param projectID - the ID of the project
 * @returns a promise of all comments for the project
 */
export const findProjectComments = async (projectID: string):Promise<CommentResults> => {
    const client = getApolloClient()

    return client.query({ query: FindProjectCommentsDocument, variables: { projectID }, fetchPolicy: "no-cache" }).then((res) => {
        if (res) {
            return res.data.findProjectComments;
        }
    })
    .catch((err) => {
        return err;
    });
}

/**
 * add a comment to a project
 * @param input - the comment to add
 * @returns a promise of the created comment
 */
export const addComment = async (input: UpdateComment) => {
    const client = getApolloClient()

    return client.mutate({ mutation: CreateProjectCommentDocument, variables: { comment: input } })
    .then((res) => {
        if (res) {
            return res.data.createProjectComment;
        }
    })
    .catch((err) => {
        return err;
    });
}

/**
 * update a comment
 * @param input - the comment to update
 * @returns a promise of the updated comment
 */
export const updateComment = async (input: UpdateComment) => {
    const client = getApolloClient()

    return client.mutate({ mutation: UpdateProjectCommentDocument, variables: { comment: input } }).then((res) => {
        if (res) {
            return res.data.updateProjectComment;
        }
    })
    .catch((err) => {
        return err;
    });
}

/**
 * add a reply to a comment
 * @param input - the reply to add
 * @returns a promise of the created reply
 */
export const addCommentReply = async (input: UpdateCommentReply) => {
    const client = getApolloClient()

    return client.mutate({ mutation: CreateProjectCommentReplyDocument, variables: { reply: input } }).then((res) => {
        if (res) {
            return res.data.createProjectCommentReply;
        }
    })
    .catch((err) => {
        return err;
    });
}

/**
 * delete a comment
 * @param id - the ID of the comment to delete
 * @returns a promise of the deleted comment
 */
export const deleteComment = async (id: string) => {
    const client = getApolloClient()

    return client.mutate({ mutation: DeleteProjectCommentDocument, variables: { id } })
    .then((res) => {
        if (res) {
            return res.data.deleteProjectComment;
        }
    })
    .catch((err) => {
        return err;
    });
}


/**
 * toggle a comment emote
 * @param input - the emote to toggle
 * @returns a promise of the updated comment
 */
export const toggleCommentEmote = async (input: UpdateCommentEmote) => {
    const client = getApolloClient()

    return client.mutate({ mutation: ToggleEmoteDocument, variables: { input } })
    .then((res) => {
        if (res) {
            return res.data.toggleEmote;
        }
    })
    .catch((err) => {
        return err;
    });
}

/**
 * get a comment thread
 * @param commentID - the ID of the comment
 * @returns a promise of the comment thread
 */ 
export const getCommentThread = async (commentID: String):Promise<CommentEnvelope> => {
    const client = getApolloClient()

    return client.query({ query: GetCommentThreadDocument, variables: { commentID } }).then((res) => {
        if (res) {
            return res.data.getCommentThread;
        }
    })
    .catch((err) => {
        return err;
    });
}