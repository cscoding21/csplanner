<script lang="ts">
	import type { CommentEnvelope } from "$lib/graphql/generated/sdk";
	import ActivityProjectCreated from "./activities/ActivityProjectCreated.svelte";
	import ActivityProjectStateUpdated from "./activities/ActivityProjectStateUpdated.svelte";
	import ActivityProjectUpdated from "./activities/ActivityProjectUpdated.svelte";

    interface Props {
        comment: CommentEnvelope
    }
    let { comment }:Props = $props()

</script>


<div class="p-4 text-yellow-50 text-sm">
{#if comment.data.context?.endsWith("project.project.created")}
    <ActivityProjectCreated {comment} />
{:else if comment.data.context?.endsWith("project.project.updated")}
    <ActivityProjectUpdated {comment} />
{:else if comment.data.context?.endsWith("project.state.updated")}
    <ActivityProjectStateUpdated {comment} />
{/if}
</div>
