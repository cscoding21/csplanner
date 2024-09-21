<script lang="ts">
    import { FaceLaughOutline, FaceLaughSolid } from "flowbite-svelte-icons";

    interface Props {
        userID: string
        toggleReaction: Function
        users: string[]
        commentID: string
        size: "xs"|"sm"|"md"|"lg"|"xl"|undefined
    }
    let {
        userID,
        toggleReaction,
        users,
        commentID,
        size
    }:Props = $props()

</script>


<button type="button" onclick={() => toggleReaction(commentID, "laughs_at")} class="flex items-center text-sm text-gray-500 hover:underline dark:text-gray-400 font-medium">
    {#if users?.includes(userID)}
    <span class="text-blue-400">
        <FaceLaughSolid size={size} />
    </span>
    {:else}
    <span class="">
        <span class="hover:text-blue-400">
        <FaceLaughOutline size={size} />
    </span>
    </span>
    {/if}
    <sup class="ml-1 no-underline">{users?.length.valueOf() || ""}</sup>
</button>