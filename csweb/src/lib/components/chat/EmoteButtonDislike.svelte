<script lang="ts">
    import { ThumbsDownOutline, ThumbsDownSolid } from "flowbite-svelte-icons";

    interface Props {
        userID: string
        toggleReaction: Function
        users: string[]
        commentID: string
        projectID: string
        parentID: string|undefined
        size: "xs"|"sm"|"md"|"lg"|"xl"|undefined
    }
    let {
        userID,
        toggleReaction,
        users,
        commentID,
        projectID,
        parentID,
        size
    }:Props = $props()

</script>


<button type="button" onclick={() => toggleReaction(projectID, parentID, commentID, "dislikes")} class="flex items-center text-sm text-gray-500 hover:underline dark:text-gray-400 font-medium">
    {#if users?.includes(userID)}
    <span class="text-purple-400">
        <ThumbsDownSolid size={size} />
    </span>
    {:else}
    <span class="">
        <span class="hover:text-purple-400">
        <ThumbsDownOutline size={size} />
    </span>
    </span>
    {/if}
    <sup class="ml-1 no-underline">{users?.length.valueOf() || ""}</sup>
</button>