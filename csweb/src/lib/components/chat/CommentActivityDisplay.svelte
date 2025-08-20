<script lang="ts">
	import type { CommentEnvelope } from "$lib/graphql/generated/sdk";
	import ActivityProjectStateUpdated from "./activities/ActivityProjectStateUpdated.svelte";
	import ActivityObjectCreated from "./activities/ActivityObjectCreated.svelte";
	import ActivityObjectUpdated from "./activities/ActivityObjectUpdated.svelte";
	import ActivityObjectDeleted from "./activities/ActivityObjectDeleted.svelte";

    interface Props {
        comment: CommentEnvelope
    }
    let { comment }:Props = $props()

    const getObjectFromContext = (c:string):string => {
        if(!c) {
            return ""
        }

        const parts = c.split(".")

        if(!parts || parts.length !== 3) {
            return "unknown object"
        }

        return parts[1]
    }

    let ctx = $derived(getObjectFromContext(comment.data.context as string))
</script>


<div class="p-4 text-yellow-50 text-sm">
{#if comment.data.context?.endsWith("project.state.created")}
    <ActivityProjectStateUpdated {comment} />
{:else if comment.data.context?.endsWith(".created")}
    <ActivityObjectCreated object={ctx} {comment} />
{:else if comment.data.context?.endsWith(".updated")}
    <ActivityObjectUpdated object={ctx} {comment} />
{:else if comment.data.context?.endsWith(".deleted")}
    <ActivityObjectDeleted object={ctx} {comment} />
{/if}
</div>
