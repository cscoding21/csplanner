<script lang="ts">
    import { ThumbsUpOutline, ThumbsUpSolid } from "flowbite-svelte-icons";

    interface Props {
        userID: string
        toggleReaction: Function
        users: string[]
        commentID: string
        projectID: string
        size: "xs"|"sm"|"md"|"lg"|"xl"|undefined
    }
    let {
        userID,
        toggleReaction,
        users,
        commentID,
        projectID,
        size
    }:Props = $props()

</script>


<button type="button" onclick={() => toggleReaction(projectID, commentID, "likes")} class="flex items-center text-sm text-gray-500 hover:underline dark:text-gray-400 font-medium">
    {#if users?.includes(userID)}
    <span class="text-green-400">
        <ThumbsUpSolid size={size} />
    </span>
    {:else}
    <span class="">
        <span class="hover:text-green-400">
        <ThumbsUpOutline size={size} />
    </span>
    </span>
    {/if}
    <sup class="ml-1 no-underline">{users?.length.valueOf() || ""}</sup>
</button>