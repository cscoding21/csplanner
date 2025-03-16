<script lang="ts">
	import { page } from '$app/stores';
	import { ButtonGroup, Drawer, Button } from 'flowbite-svelte';
	import { getProject } from '$lib/services/project';
	import { 
		ProjectActionBar, 
		DeleteProject, 
		ProjectSnapshot, 
		ProjectFeatures,
		ProjectBasics,
		ProjectDACI,
		ProjectMilestones,
		ProjectValue,
		ProjectCost,
		ProjectSchedule
	} from '../../components';
	import { CommentList, CSSection } from "$lib/components";
	import { TrashBinOutline, MessagesOutline, ArrowRightToBracketOutline } from 'flowbite-svelte-icons';
	import type { ProjectEnvelope } from '$lib/graphql/generated/sdk';
	import { sineIn } from 'svelte/easing';
	import { ProjectStatusUpdate } from '../../components';
	import { goto } from '$app/navigation';
	import ShowIfStatus from '../../components/ShowIfStatus.svelte';

	const id = $page.params.id;
	let hash = $state($page.url.hash);

	let project: ProjectEnvelope = $state({} as ProjectEnvelope);

	const loadPage = async () => {
		if (!hash) {
			hash = "#snapshot"
		}

		getProject(id).then((p) => {
			project = p;
			return p;
		});
	};

	let chatBarHidden = $state(true);
	let transitionParamsRight = {
		x: 320,
		duration: 200,
		easing: sineIn,
	};
</script>

<svelte:window
	on:hashchange={() => {
		hash = $page.url.hash
	}}
/>

{#await loadPage()}
	<div>Loading...</div>
{:then promiseData}
	<ProjectActionBar pageDetail={project.data?.projectBasics?.name}>
		<ButtonGroup>
			<ProjectStatusUpdate id={id} update={() => { goto("/project/detail/" + id + "#snapshot", { invalidateAll: true }) }}>
				<ArrowRightToBracketOutline size="sm" class="mr-2"  />
				Status
			</ProjectStatusUpdate>
			<ShowIfStatus scope={["new", "draft", "abandoned", "rejected"]} status={project.data?.projectStatusBlock?.status}>
				<DeleteProject id={id} name={project.data?.projectBasics?.name}>
					<TrashBinOutline size="sm" class="mr-2" />
					Delete
				</DeleteProject>
			</ShowIfStatus>
			<Button onclick={() => chatBarHidden = !chatBarHidden}>	
				<MessagesOutline size="sm" class="mr-2" />
				Collaborate
			</Button>
		</ButtonGroup>
	</ProjectActionBar>

	<div class="grid grid-cols-5">
	<div class="flex w-full col-span-5">
		<CSSection>
		{#if hash == "#snapshot"}
			<ProjectSnapshot {id}  />
		{:else if hash == '#basics'}
			<ProjectBasics {id} update={() => console.log("update")} />
		{:else if hash == '#daci'}
			<ProjectDACI {id} update={() => console.log("update")} />
		{:else if hash == '#features'}
			<ProjectFeatures {id} update={() => console.log("update")} />
		{:else if hash == '#milestones'}
			<ProjectMilestones {id} update={() => console.log("update")} />
		{:else if hash == '#value'}
			<ProjectValue {id} update={() => console.log("update")} />
		{:else if hash == '#costs'}
			<ProjectCost {id} update={() => console.log("update")} />
		{:else if hash == '#schedule'}
			<ProjectSchedule {id} update={() => console.log("update")} />
		{/if}
	</CSSection>
	</div>
	
	</div>
{/await}

<Drawer 
	placement="right" 
	width="w-1/3" 
	transitionType="fly" 
	transitionParams={transitionParamsRight} 
	bind:hidden={chatBarHidden} 
	backdrop={false} id="chatbar">
	<CommentList {id} />
</Drawer>
