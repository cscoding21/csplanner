<script lang="ts">
    import { Avatar } from 'flowbite-svelte';
    import { getInitialsFromName } from '$lib/utils/format';
    import type { Resource } from '$lib/graphql/generated/sdk';

    interface Props {
        resource: Resource
    }

    let { resource }: Props = $props();
</script>

{#if resource}
<div class="p-2">
	<Avatar src={resource.profileImage || ''} class="float-left mr-3 h-16 w-16" rounded
		>{getInitialsFromName(resource.name)}</Avatar
	>
	<div class="text-left">
		<div>{resource.name}</div>
		<div><p>{resource.role}</p></div>
		<div class="clear-both"></div>
	</div>
</div>
{/if}


