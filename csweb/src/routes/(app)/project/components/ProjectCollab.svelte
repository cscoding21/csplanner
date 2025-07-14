<script lang="ts">
	import { CommentList, SectionHeading } from '$lib/components';
	import { getDefaultProject} from '$lib/forms/project.validation';
	import { getProject } from '$lib/services/project';
	import { addToast } from '$lib/stores/toasts';
	import type { ProjectEnvelope } from '$lib/graphql/generated/sdk';
	import { BadgeProjectStatus } from '.';
	import BasicsSummary from './snapshots/BasicsSummary.svelte';
	import SectionSubHeading from '$lib/components/formatting/SectionSubHeading.svelte';
	

	interface Props {
		id: string;
		update: Function;
	}
	let { id , update }: Props = $props();

	let project:ProjectEnvelope = $state(getDefaultProject() as ProjectEnvelope)

	const load = async ():Promise<ProjectEnvelope> => {
		return await getProject(id)
			.then((proj) => {
				project = proj;

				return proj
			})
			.catch((err) => {
				addToast({
					message: 'Error loading project (ProjectCost): ' + err,
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
	Loading...
{:then promiseData}
	{#if project} 
		<SectionHeading>
			Collaborate: {project.data?.projectBasics.name}
			<span class="float-right"><BadgeProjectStatus status={project.data?.projectStatusBlock?.status} /></span>
		</SectionHeading>

        <div class="grid grid-cols-3 gap-16">
        <div>
            <BasicsSummary project={project.data} />
        </div>


        <div class="col-span-2">
            <SectionSubHeading>Conversation</SectionSubHeading>
            <CommentList {id} />
        </div>
        </div>
	{/if}
{/await}

