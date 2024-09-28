<script lang="ts">
	import { page } from '$app/stores';
	import { getProject } from '$lib/services/project';
	import {
		DeleteProject,
		ProjectActionBar,
		ProjectBasics,
		ProjectFeatures,
		ProjectValue,
		ProjectCost,
		ProjectMilestones,
		ProjectDACI
	} from '../../components';
	import { ButtonGroup, Button, TabItem, Tabs } from 'flowbite-svelte';
	import {
		MessagesOutline,
		TrashBinOutline,
		ScaleBalancedOutline,
		UsersGroupOutline,
		DollarOutline,
		CashOutline,
		RectangleListOutline,
		CalendarMonthOutline
	} from 'flowbite-svelte-icons';
	import { AutoSave, PageHeading } from '$lib/components';

	import type { Project } from '$lib/graphql/generated/sdk';

	const id = $page.params.id;

	let project: Project = $state({} as Project);
	let saver: any = $state({});

	const refresh = async () => {
		getProject(id).then((p) => {
			if (project && saver) {
				setNotSaving(project.updatedAt);
			}

			project = p;
			return project;
		});
	};

	const setSaving = (st: string) => {
		console.log('setSaving', st);
	};

	const setNotSaving = (st: string) => {
		console.log('setNotSaving', st);
	};

	const update = async (e: any) => {
		console.log('project page refresh');
		refresh().then((r) => r);
	};

	const loadPage = async () => {
		refresh().then((r) => r);
	};

	loadPage();
</script>

{#await loadPage}
	<div>Loading...</div>
{:then data}
	<ProjectActionBar pageDetail={project.projectBasics?.name}>
		<ButtonGroup>
			<Button href="/project/detail/{project.id}">
				<MessagesOutline class="mr-2 h-3 w-3" />
				Collaboration
			</Button>
			<DeleteProject id={project.id || ''} name={project.projectBasics?.name}>
				<TrashBinOutline class="mr-2 h-3 w-3" />
				Delete
			</DeleteProject>
		</ButtonGroup>
	</ProjectActionBar>

	<PageHeading title={project?.projectBasics?.name}>
		<AutoSave bind:this={saver} {setSaving} setNotSaving={() => setNotSaving(project.updatedAt)} />
	</PageHeading>

	<Tabs tabStyle="underline">
		<TabItem open title="Basics">
			<div slot="title" class="flex items-center gap-2">
				<ScaleBalancedOutline size="sm" />
				Basics
			</div>
			<div class="p-4"><ProjectBasics {id} {update} /></div>
		</TabItem>

		<TabItem title="Features">
			<div slot="title" class="flex items-center gap-2">
				<RectangleListOutline size="sm" />
				Features
			</div>
			<div class="p-4"><ProjectFeatures {id} {update} /></div>
		</TabItem>

		<TabItem title="Value Prop">
			<div slot="title" class="flex items-center gap-2">
				<DollarOutline size="sm" />
				Value Prop
			</div>
			<div class="p-4"><ProjectValue {id} {update} /></div>
		</TabItem>

		<TabItem title="Costs">
			<div slot="title" class="flex items-center gap-2">
				<CashOutline size="sm" />
				Costs
			</div>
			<div class="p-4"><ProjectCost {id} {update} /></div>
		</TabItem>

		<TabItem title="Milestones">
			<div slot="title" class="flex items-center gap-2">
				<CalendarMonthOutline size="sm" />
				Milestones
			</div>
			<div class="p-4">
				<!-- <ProjectTemplateSelector id={id} update={update} /> -->
				<ProjectMilestones {id} {update} />
			</div>
		</TabItem>

		<TabItem title="Team">
			<div slot="title" class="flex items-center gap-2">
				<UsersGroupOutline size="sm" />
				Team
			</div>
			<div class="p-4"><ProjectDACI {id} {update} /></div>
		</TabItem>
	</Tabs>
{/await}
