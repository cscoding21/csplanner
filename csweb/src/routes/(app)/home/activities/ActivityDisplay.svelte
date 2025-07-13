<script lang="ts">
	import type { Activity } from "$lib/graphql/generated/sdk";
	import { Avatar } from "flowbite-svelte";
	import ProjectProjectCreated from "./ProjectProjectCreated.svelte";
	import { formatDateTime, getInitialsFromName } from "$lib/utils/format";
	import CommentCommentCreated from "./CommentCommentCreated.svelte";


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
    Poop...
{:else if activity.context?.endsWith("comment.comment.created")}
    <CommentCommentCreated {activity} />
{/if}


        <span class="inline-flex items-center text-xs font-normal text-gray-500 dark:text-gray-400">
        <svg class="me-1 h-2.5 w-2.5" aria-hidden="true" xmlns="http://www.w3.org/2000/svg" fill="currentColor" viewBox="0 0 20 20">
            <path
            d="M10 .5a9.5 9.5 0 1 0 0 19 9.5 9.5 0 0 0 0-19ZM8.374 17.4a7.6 7.6 0 0 1-5.9-7.4c0-.83.137-1.655.406-2.441l.239.019a3.887 3.887 0 0 1 2.082 2.5 4.1 4.1 0 0 0 2.441 2.8c1.148.522 1.389 2.007.732 4.522Zm3.6-8.829a.997.997 0 0 0-.027-.225 5.456 5.456 0 0 0-2.811-3.662c-.832-.527-1.347-.854-1.486-1.89a7.584 7.584 0 0 1 8.364 2.47c-1.387.208-2.14 2.237-2.14 3.307a1.187 1.187 0 0 1-1.9 0Zm1.626 8.053-.671-2.013a1.9 1.9 0 0 1 1.771-1.757l2.032.619a7.553 7.553 0 0 1-3.132 3.151Z"
            />
        </svg>
        <time datetime="2022-02-08" title={formatDateTime(activity.activityDate)}>{formatDateTime(activity.activityDate)}</time>
        </span>
    </div>
    </a>
</li>