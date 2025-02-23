<script lang="ts">
	import { SelectInput } from '$lib/components';
	import type { Projecttemplate } from '$lib/graphql/generated/sdk';
	import { SectionSubHeading, SectionHeading } from '$lib/components';
	import { setProjectMilestonesFromTemplate } from '$lib/services/project';
	import { CardPlaceholder, Button, P, type SelectOptionType, Alert } from 'flowbite-svelte';
	import { callIf } from '$lib/utils/helpers';
	import { InfoCircleSolid } from 'flowbite-svelte-icons';
	import { findAllProjectTemplates } from '$lib/services/template';

	interface Props {
		id?: string;
		update?: Function;
	}
	let { id, update }: Props = $props();

	let errors = $state({ templateID: '' });
	let templateID: string = $state('');
	let templateOpts: SelectOptionType<string>[] = $state([] as SelectOptionType<string>[]);
	let projectTemplates: Projecttemplate[] = $state([] as Projecttemplate[]);
	let currentTemplate: Projecttemplate | undefined = $state();

	const showTemplateDetails = () => {
		console.log(templateID);
		currentTemplate = getTemplate(templateID);
	};

	const getTemplate = (id: string): Projecttemplate => {
		const tmp = projectTemplates.filter((pt) => pt.id === id);

		return tmp[0];
	};

	const selectTemplate = async () => {
		if(!id) {
			return
		}

		setProjectMilestonesFromTemplate({ projectId: id, templateId: templateID }).then((td) => {
			console.log('template details set');

			callIf(update);
		});
	};

	const loadPage = async () => {
		findAllProjectTemplates()
			.then((r) => {
				projectTemplates = r.results as Projecttemplate[];

				return r;
			})
			.then((r) => {
				templateOpts = r.results?.map((r) => ({
					name: r.name,
					value: r.id as string
				})) as SelectOptionType<string>[];
			})
			.then(() => {});
	};

	loadPage();
</script>

{#await loadPage}
	<CardPlaceholder />
{:then promiseData}
	<SelectInput
		fieldName="Project Milestone Template"
		update={showTemplateDetails}
		bind:value={templateID}
		error={errors.templateID}
		options={templateOpts}
	/>

	{#if currentTemplate !== undefined}
		<SectionHeading>Milestones for '{currentTemplate.name}' template</SectionHeading>
		{#if currentTemplate.description}
		<P class="text-sm mb-6">{currentTemplate.description}</P>
		{/if}

		{#if currentTemplate.phases && currentTemplate.phases.length > 0}
			{#each currentTemplate.phases as phase (phase)}
				<SectionSubHeading>Milestone: {phase.name}</SectionSubHeading>
				<div class="mb-6 text-sm">{phase.description}</div>
				{#if phase.tasks && phase.tasks.length > 0}
					<div class="text-sm text-gray-200 mb-1">Prefilled tasks</div>
					<ul class="list-disc ml-3 space-y-2 mb-4 text-sm">
						{#each phase.tasks as task}
						<li><b>{task.name} ({task.requiredSkillID})</b>: {task.description}</li>
						{/each}
					</ul>
				{:else}
					<Alert border color="blue" class="mt-2 mb-6">
						<InfoCircleSolid slot="icon" class="w-5 h-5" />
						<span class="font-medium">No tasks</span>
						No tasks will be created for this milestone
					</Alert>
				{/if}
			{/each}
		{/if}

		<div class="mt-8">
			<span class="">
				<Button onclick={() => selectTemplate()}>Use This Project Lifecycle</Button>
			</span>
		</div>
	{/if}
{/await}
