<script lang="ts">
	import {
		Heading,
		P,
		Button,
		ButtonGroup,
		Modal,
	} from 'flowbite-svelte';
	import { EditOutline, TrashBinOutline } from 'flowbite-svelte-icons';
	import type { ProjectFeature } from '$lib/graphql/generated/sdk';
	import { DeleteProjectFeature, ProjectFeatureForm, BadgeFeaturePriority } from '.';
	import { callIf } from '$lib/utils/helpers';

	interface Props {
		update?: Function
		projectID: string
		feature: ProjectFeature;
	}
	let { 
		feature = $bindable(), 
		update, 
		projectID 
	}: Props = $props();

	let modalState:boolean = $state(false)

</script>

<Heading tag="h5" class="text-lg dark:text-white">
	{feature.name}

	<span class="float-right">
		<BadgeFeaturePriority {feature} />

		<ButtonGroup>
		<Button
			color="dark"
			class=""
			onclick={() => {
				modalState = true;
			}}
		>
			<EditOutline size="sm" />
		</Button>
		<DeleteProjectFeature
			id={feature.id || ''}
			name={feature.name}
			projectID={projectID}
			size="sm"
			update={() => callIf(update)}
		>
			<TrashBinOutline size="sm" class="" />
		</DeleteProjectFeature>
		</ButtonGroup>
	</span>
</Heading>
<div class="my-2 text-xs text-gray-100">Status: <b>{feature.status}</b></div>
<P weight="light" color="text-gray-300 dark:text-gray-200 text-sm">{feature.description}</P>


<Modal bind:open={modalState} size="xl">
	<ProjectFeatureForm projectID={projectID} {feature} update={() => { modalState = false; callIf(update)} } />
</Modal>
