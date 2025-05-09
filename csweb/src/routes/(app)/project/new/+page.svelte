<script lang="ts">
    import { Alert, Button, P, Toggle, type SelectOptionType } from 'flowbite-svelte';
    import { CSSection, TextInput, SelectInput, TextAreaInput, CSHR, SectionHeading, SectionSubHeading } from '$lib/components';
	import type { Projecttemplate, UpdateNewProject, UpdateProjectBasics } from '$lib/graphql/generated/sdk';
    import { findAllUsers } from '$lib/services/user';
	import { ProjectActionBar } from '../components';
    import { authService } from '$lib/services/auth';
	import { createProject } from '$lib/services/project';
	import { InfoCircleSolid } from 'flowbite-svelte-icons';
    import { newProjectSchema } from '$lib/forms/project.validation';
	import { addToast } from '$lib/stores/toasts';
	import { goto } from '$app/navigation';
	import { mergeErrors, parseErrors } from '$lib/forms/helpers';
	import { findAllProjectTemplates } from '$lib/services/template';

    let errors: any = $state({ name: '', ownerID: '', description: '', isCapitalized: false, templateID: '' });
	let userOpts = $state([] as SelectOptionType<string>[]);
    let newProjectForm = $state({name: '', ownerID: '', description: '', isCapitalized: false, templateID: '' } as UpdateNewProject);
	let templateOpts: SelectOptionType<string>[] = $state([] as SelectOptionType<string>[]);
	let projectTemplates: Projecttemplate[] = $state([] as Projecttemplate[]);
	let currentTemplate: Projecttemplate | undefined = $state();

    let as = authService()
    
    newProjectForm.ownerID = as.currentUser()?.email as string

    const showTemplateDetails = () => {
		currentTemplate = getTemplate(newProjectForm.templateID);
	};

    const getTemplate = (id: string): Projecttemplate => {
		const tmp = projectTemplates.filter((pt) => pt.id === id);

		return tmp[0];
	};

    const createProj = () => {
        errors = {};

		const newProjectFormParsed = newProjectSchema.cast(newProjectForm);
		newProjectSchema
			.validate(newProjectFormParsed, { abortEarly: false })
			.then(async () => {
				createProject(newProjectFormParsed)
					.then((res) => {
						if (res.status?.success) {
                            const id = res.project?.id

                            goto("/project/detail/" + id + "#snapshot", { replaceState: true, invalidateAll: true })
						} else {
							addToast({
								message: 'Error creating new project: ' + res.status?.message,
								dismissible: true,
								type: 'error'
							});
						}
					})
					.catch((err) => {
						addToast({
							message: 'Error creating new project: ' + err,
							dismissible: true,
							type: 'error'
						});
					});
			})
			.catch((err) => {
				errors = mergeErrors(errors, parseErrors(err));
			});
    }

    const loadPage = async () => {
		findAllUsers()
			.then((r) => {
				userOpts = r.results?.map((r) => ({
					name: r.firstName + ' ' + r.lastName,
					value: r.email as string
				})) as SelectOptionType<string>[];

				return r;
			}).then(u => {
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
                    });
	};

</script>

<ProjectActionBar pageDetail="New Project"></ProjectActionBar>

<div class="p-4"> 
    <CSSection>

    <SectionHeading>Create a New Project</SectionHeading>

{#await loadPage()}
    Loading...
{:then promiseData} 
<div class="flex">
<div class="flex-1 px-4">
<SelectInput
		fieldName="Project Milestone Template"
		update={showTemplateDetails}
		bind:value={newProjectForm.templateID}
		error={errors.templateID}
		options={templateOpts}
	/>

<TextInput
    bind:value={newProjectForm.name}
    fieldName="Project Name"
    placeholder="Project name"
    error={errors.name}>
    The name will be used to identify the project
</TextInput>

<SelectInput
    bind:value={newProjectForm.ownerID as string}
    fieldName="Owner"
    error={errors.ownerID}
    options={userOpts}> 
    This is a user in the system as opposed to a resource.  This user will manage the project within csPlanner.
</SelectInput>

<TextAreaInput
    bind:value={newProjectForm.description as string}
    rows={8}
    error={errors.description}
    placeholder="Executive summary"
    fieldName="Executive Summary"
/>

<div class="pb-4 mb-2">
    <Toggle class="mt-3" bind:checked={newProjectForm.isCapitalized}>Capitalized</Toggle>
</div>

<div class="">
    <span class="float-right">
        <Button onclick={createProj}>Create New Project</Button>
    </span>
    <br class="clear-both" />
</div>
</div>

<div class="flex-1 px-4">
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
{/if}
</div>
</div>

{/await}


</CSSection>
</div>