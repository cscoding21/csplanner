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
	import { DeleteProjectFeature, ProjectFeatureForm, BadgeFeaturePriority, ShowIfStatus, BadgeFeatureStatus } from '.';
	import { callIf } from '$lib/utils/helpers';

	interface Props {
		update?: Function
		projectID: string
		projectStatus: string
		feature: ProjectFeature;
	}
	let { 
		feature = $bindable(), 
		update, 
		projectID, 
		projectStatus
	}: Props = $props();

	let modalState:boolean = $state(false)

</script>

<div class="mb-2 rounded-lg bg-white p-4 shadow-sm dark:bg-gray-900 sm:p-6">
<Heading tag="h5" class="text-lg dark:text-white">
	{feature.name}

	<span class="float-right">
		<ShowIfStatus scope={["new", "draft"]} status={projectStatus}>
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
		</ShowIfStatus>
	</span>
</Heading>
<div class="my-3">
<BadgeFeatureStatus status={feature.status} />
<BadgeFeaturePriority priority={feature.priority} />
</div>
<P weight="light" color="text-gray-300 dark:text-gray-200 text-sm">{feature.description}</P>
</div>


<Modal bind:open={modalState} size="xl">
	<ProjectFeatureForm projectID={projectID} {feature} update={() => { modalState = false; callIf(update)} } />
</Modal>
