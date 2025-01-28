<script lang="ts">
	import { Hr } from 'flowbite-svelte';
	import { NoResults, SectionHeading } from '$lib/components';
	import { BadgeProjectStatus, ProjectFeatureDisplay, ProjectFeatureForm } from '.';
	import { getProject } from '$lib/services/project';
	import { addToast } from '$lib/stores/toasts';
	import type { Project } from '$lib/graphql/generated/sdk';
	import { getDefaultProject } from '$lib/forms/project.validation';

	interface Props {
		id: string;
		update: Function;
	}

	let { id, update }: Props = $props();

	let project:Project = $state(getDefaultProject() as Project)

	function refresh() {
		console.log('refreshing')
		load().then(p => {
			project = p

			update();
		})
	}

	const load = async ():Promise<Project> => {
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
	Features: {project.projectBasics.name}
	<span class="float-right"><BadgeProjectStatus status={project.projectBasics.status} /></span>
</SectionHeading>

<div class="mb-8">
	{#if project.projectFeatures && project.projectFeatures.length > 0}
		{#each project.projectFeatures as feature (feature.id)}
			{#if feature && feature.id}
				<ProjectFeatureDisplay {feature} update={refresh} projectID={project.id as string} />
				<Hr hrClass="h-px my-3 bg-gray-200 border-0 dark:bg-gray-700" />
			{/if}
		{/each}
	{:else}
		<NoResults title="No Project Features" newUrl="">
			No project features...create one now.
		</NoResults>
	{/if}
</div>

<ProjectFeatureForm projectID={id} feature={undefined} update={() => refresh()} />

{/if}
{/await}
