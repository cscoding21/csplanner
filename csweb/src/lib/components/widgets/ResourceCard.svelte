<script lang="ts">
    import { Avatar } from 'flowbite-svelte';
    import { getInitialsFromName, formatToCommaSepList } from '$lib/utils/format';
    import type { Resource, Skill } from '$lib/graphql/generated/sdk';
	import { ResourceSkillsList } from '..';

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
		<div class="text-xs">
			<ResourceSkillsList skills={resource.skills as Skill[]} />
		</div>
		<div class="clear-both"></div>
	</div>
</div>
{/if}


