<script lang="ts">
	import { SelectInput } from '$lib/components';
	import type { Projecttemplate } from '$lib/graphql/generated/sdk';
	import { SectionHeading } from '$lib/components';
	import { findAllProjectTemplates, setProjectMilestonesFromTemplate } from '$lib/services/project';
	import { CardPlaceholder, Button, P, type SelectOptionType } from 'flowbite-svelte';
	import { callIf } from '$lib/utils/helpers';

	interface Props {
		id: string;
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
		<SectionHeading>Milestones for {showTemplateDetails.name} Template</SectionHeading>

		{#if currentTemplate.phases && currentTemplate.phases.length > 0}
			{#each currentTemplate.phases as phase (phase)}
				<SectionHeading>{phase.name}</SectionHeading>
				<P class="mb-3" weight="light" color="text-gray-500 dark:text-gray-400"
					>{phase.description}</P
				>
			{/each}
		{/if}

		<div class="mt-8">
			<span class="">
				<Button onclick={() => selectTemplate()}>Use This Project Lifecycle</Button>
			</span>
		</div>
	{/if}
{/await}
