<script lang="ts">
    import { FaceLaughOutline, FaceLaughSolid } from "flowbite-svelte-icons";

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


<button type="button" onclick={() => toggleReaction(projectID, parentID, commentID, "laughs_at")} class="flex items-center text-sm text-gray-500 hover:underline dark:text-gray-400 font-medium">
    {#if users?.includes(userID)}
    <span class="text-yellow-200">
        <FaceLaughSolid size={size} />
    </span>
    {:else}
    <span class="">
        <span class="hover:text-yellow-200">
        <FaceLaughOutline size={size} />
    </span>
    </span>
    {/if}
    <sup class="ml-1 no-underline">{users?.length.valueOf() || ""}</sup>
</button>