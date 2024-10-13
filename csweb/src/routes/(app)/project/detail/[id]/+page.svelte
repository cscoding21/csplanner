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
	import { CommentList, CSSection } from "$lib/components";
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

	let chatBarHidden = $state(true);
	let transitionParamsRight = {
		duration: 200,
		easing: sineIn
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
	<ProjectActionBar pageDetail={project.projectBasics?.name}>
		<ButtonGroup>
			<DeleteProject id={id || ''} name={project.projectBasics?.name}>
				<TrashBinOutline size="sm" class="mr-2" />
				Delete
			</DeleteProject>
			<Button onclick={() => chatBarHidden = !chatBarHidden}>	
				Chat
			</Button>
		</ButtonGroup>
	</ProjectActionBar>

	<div class="grid grid-cols-5">
	<div class="flex w-full col-span-5">
		{#if hash == "#snapshot"}
			<CSSection>
				<ProjectSnapshot {id} />
			</CSSection>
		{:else if hash == '#basics'}
			<CSSection>
				<ProjectBasics {id} update={() => console.log("update")} />
			</CSSection>
		{:else if hash == '#daci'}
			<CSSection>
				<ProjectDACI {id} update={() => console.log("update")} />
			</CSSection>
		{:else if hash == '#features'}
			<CSSection>
				<ProjectFeatures {id} update={() => console.log("update")} />
			</CSSection>
		{:else if hash == '#milestones'}
			<CSSection>
				<ProjectMilestones {id} update={() => console.log("update")} />
			</CSSection>
		{:else if hash == '#value'}
			<CSSection>
				<ProjectValue {id} update={() => console.log("update")} />
			</CSSection>
		{:else if hash == '#costs'}
			<CSSection>
				<ProjectCost {id} update={() => console.log("update")} />
			</CSSection>
		{/if}
	</div>

	<!-- <div class="ml-4 col-span-2">
		<Section>
			<CommentList {id} />
		</Section>
	</div> -->
	
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
