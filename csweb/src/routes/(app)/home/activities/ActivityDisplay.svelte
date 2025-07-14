<script lang="ts">
	import type { Activity } from "$lib/graphql/generated/sdk";
	import { Avatar } from "flowbite-svelte";
	import ProjectProjectCreated from "./ProjectProjectCreated.svelte";
	import { formatDateTime, getInitialsFromName } from "$lib/utils/format";
	import CommentCommentCreated from "./CommentCommentCreated.svelte";
	import ProjectProjectUpdated from "./ProjectProjectUpdated.svelte";
	import ResourceResourceCreated from "./ResourceResourceCreated.svelte";
	import ResourceRoleUpdated from "./ResourceRoleUpdated.svelte";
	import ResourceResourceUpdated from "./ResourceResourceUpdated.svelte";
	import ProjectStateUpdated from "./ProjectStateUpdated.svelte";
	import ResourceRoleCreated from "./ResourceRoleCreated.svelte";
	import CommentReplyCreated from "./CommentReplyCreated.svelte";
	import { ClockOutline } from "flowbite-svelte-icons";


    interface Props {
        activity: Activity
    }
    let { activity }:Props = $props()

</script>


<li>
    <a href={activity.link} class="block items-center p-3 hover:bg-gray-100 dark:hover:bg-gray-700 sm:flex">
    <Avatar src={activity.user.profileImage || ''} class="mr-3 h-16 w-16 min-w-16">
        {getInitialsFromName(activity.user.firstName + " " + activity.user.lastName)}
    </Avatar>
    <div class="text-gray-600 dark:text-gray-400 pt-4">


{#if activity.context?.endsWith("project.project.created")}
    <ProjectProjectCreated {activity} />
{:else if activity.context?.endsWith("project.project.updated")}
    <ProjectProjectUpdated {activity} />
{:else if activity.context?.endsWith("project.state.updated")}
    <ProjectStateUpdated {activity} />
{:else if activity.context?.endsWith("comment.comment.created")}
    <CommentCommentCreated {activity} />
{:else if activity.context?.endsWith("comment.reply.created")}
    <CommentReplyCreated {activity} />
{:else if activity.context?.endsWith("resource.resource.created")}
    <ResourceResourceCreated {activity} />
{:else if activity.context?.endsWith("resource.resource.updated")}
    <ResourceResourceUpdated {activity} />
{:else if activity.context?.endsWith("resource.role.created")}
    <ResourceRoleCreated {activity} />
{:else if activity.context?.endsWith("resource.role.updated")}
    <ResourceRoleUpdated {activity} />
{/if}

<br />
        <span class="inline-flex items-center text-xs font-normal text-gray-500 dark:text-gray-400">
        <ClockOutline class="me-1 h-2.5 w-2.5" />
        <time datetime="2022-02-08" title={formatDateTime(activity.activityDate)}>{formatDateTime(activity.activityDate)}</time>
        </span>
    </div>
    </a>
</li>