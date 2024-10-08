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
		ProjectCost
	} from '../../components';
	import { CommentList, Section } from "$lib/components";
	import { TrashBinOutline } from 'flowbite-svelte-icons';
	import type { Project } from '$lib/graphql/generated/sdk';
	import { sineIn } from 'svelte/easing';

	const id = $page.params.id;
	let hash = $state($page.url.hash);

	let project: Project = $state({} as Project);

	const loadPage = async () => {
		if (!hash) {
			hash = "#snapshot"
		}

		getProject(id).then((p) => {
			project = p;
			return p;
		});
	};

	let hidden6 = $state(true);
	let transitionParamsRight = {
		x: 1000,
		duration: 200,
		easing: sineIn
	};
</script>

<svelte:window
	on:hashchange={() => {
		hash = $page.url.hash
	    // console.log($page.url.hash, 'store');
		// console.log(window.location.hash, 'location');
	}}
/>

{#await loadPage()}
	<div>Loading...</div>
{:then promiseData}
	<ProjectActionBar pageDetail={project.projectBasics?.name}>
		<ButtonGroup>
			<DeleteProject id={id || ''} name={project.projectBasics?.name}>
				<TrashBinOutline size="sm" class="mr-2" />
				Delete
			</DeleteProject>
			<Button onclick={() => hidden6 = !hidden6}>	
				Chat
			</Button>
		</ButtonGroup>
	</ProjectActionBar>

	<div class="grid grid-cols-5">
	<div class="flex w-full col-span-5">
		{#if hash == "#snapshot"}
			<Section>
				<ProjectSnapshot {id} />
			</Section>
		{:else if hash == '#basics'}
			<Section>
				<ProjectBasics {id} update={() => console.log("update")} />
			</Section>
		{:else if hash == '#daci'}
			<Section>
				<ProjectDACI {id} update={() => console.log("update")} />
			</Section>
		{:else if hash == '#features'}
			<Section>
				<ProjectFeatures {id} update={() => console.log("update")} />
			</Section>
		{:else if hash == '#milestones'}
			<Section>
				<ProjectMilestones {id} update={() => console.log("update")} />
			</Section>
		{:else if hash == '#value'}
			<Section>
				<ProjectValue {id} update={() => console.log("update")} />
			</Section>
		{:else if hash == '#costs'}
			<Section>
				<ProjectCost {id} update={() => console.log("update")} />
			</Section>
		{/if}
	</div>

	<!-- <div class="ml-4 col-span-2">
		<Section>
			<CommentList {id} />
		</Section>
	</div> -->
	
	</div>
{/await}

<Drawer placement="right" width="w-300" transitionType="fly" transitionParams={transitionParamsRight} bind:hidden={hidden6} backdrop={false} id="sidebar6">
	<CommentList {id} />
</Drawer>
