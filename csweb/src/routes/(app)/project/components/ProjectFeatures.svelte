<script lang="ts">
	import { Hr } from 'flowbite-svelte';
	import { NoResults, SectionHeading } from '$lib/components';
	import { BadgeProjectStatus, ProjectFeatureDisplay, ProjectFeatureForm } from '.';
	import { getProject, ProjectStatusDraft, ProjectStatusNew } from '$lib/services/project';
	import { addToast } from '$lib/stores/toasts';
	import type { ProjectEnvelope } from '$lib/graphql/generated/sdk';
	import { getDefaultProject } from '$lib/forms/project.validation';
	import ShowIfStatus from './ShowIfStatus.svelte';

	interface Props {
		id: string;
		update: Function;
	}

	let { id, update }: Props = $props();

	let project:ProjectEnvelope = $state(getDefaultProject() as ProjectEnvelope)

	function refresh() {
		console.log('refreshing')
		load().then(p => {
			project = p

			update();
		})
	}

	const load = async ():Promise<ProjectEnvelope> => {
		return await getProject(id)
			.then((proj) => {
				return proj
			})
			.catch((err) => {
				addToast({
					message: 'Error loading project (ProjectFeature): ' + err,
					dismissible: true,
					type: 'error'
				});

				return err
			});
	};

	const loadPage = async () => {
        load().then(p => {
            project = p
        });
	};

</script>

{#await loadPage()}
	loading...
{:then promiseData} 
	
{#if project}
<SectionHeading>
	Features: {project.data?.projectBasics.name}
	<span class="float-right"><BadgeProjectStatus status={project.data?.projectStatusBlock?.status} /></span>
</SectionHeading>

<div class="mb-8">
	{#if project.data?.projectFeatures && project.data?.projectFeatures.length > 0}
		{#each project.data?.projectFeatures as feature (feature.id)}
			{#if feature && feature.id}
				<ProjectFeatureDisplay {feature} update={refresh} projectStatus={project.data?.projectStatusBlock.status} projectID={project.meta?.id as string} />
			{/if}
		{/each}
	{:else}
		<NoResults title="No Project Features" newUrl="">
			No project features...create one now.
		</NoResults>
	{/if}
</div>

<ShowIfStatus scope={[ProjectStatusNew, ProjectStatusDraft]} status={project.data?.projectStatusBlock?.status}>
	<ProjectFeatureForm projectID={id} feature={undefined} update={() => refresh()} />
</ShowIfStatus>

{/if}
{/await}
