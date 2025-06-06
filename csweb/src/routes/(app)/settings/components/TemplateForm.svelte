<script lang="ts">
	import { ClientErrorSummary, CSSection, SelectInput, TextAreaInput } from "$lib/components";
	import { TextInput } from "$lib/components";
	import SectionSubHeading from "$lib/components/formatting/SectionSubHeading.svelte";
	import { findSelectOptsFromList } from "$lib/forms/helpers";
	import { templateSchema } from "$lib/forms/template.validation";
	import type { Projecttemplate, UpdateProjecttemplate } from "$lib/graphql/generated/sdk";
	import { getList } from "$lib/services/list";
	import { updateProjectTemplate } from "$lib/services/template";
	import { addToast } from "$lib/stores/toasts";
	import { callIf } from "$lib/utils/helpers";
	import { Accordion, AccordionItem, Button, type SelectOptionType } from "flowbite-svelte";
	import { ToolsOutline } from "flowbite-svelte-icons";
    import { scrollTop } from 'svelte-scrolling'

    interface Props {
		template: UpdateProjecttemplate | undefined;
		update?: Function;
	}
	let { 
		template = $bindable(), 
		update 
	}: Props = $props();


    let errors = $state(null)	

    const saveTemplate = async () => {
        errors = null
        
        templateSchema
            .validate(template, { abortEarly: false })
            .then( async () => {
                updateProjectTemplate(template as UpdateProjecttemplate)
					.then((res) => {
						if (res.status?.success) {
                            addToast({
                                message: 'Project basics updated successfully',
                                dismissible: true,
                                type: 'success'
                            });

                            callIf(update);
						} else {
							addToast({
								message: 'Error updating project basics: ' + res,
								dismissible: true,
								type: 'error'
							});

                            scrollTop()
						}
					})
					.catch((err) => {
						addToast({
							message: 'Error updating project basics: ' + err,
							dismissible: true,
							type: 'error'
						});

                        scrollTop()
					});
			})
			.catch((err) => {
                errors = err
                
                addToast({
                    message: 'Error updating project basics: ' + err,
                    dismissible: true,
                    type: 'error'
                });

                scrollTop()
			});
	}

    const delMilestone = (index:number) => {
        if (!template?.phases || template.phases.length === 0) {
            return
        }

        template?.phases.splice(index, 1)
    }

    const delTask = (mindex: number, tindex:number) => {
        if (!template?.phases || template.phases.length === 0 || !template.phases[mindex].tasks) {
            return
        }
        
        template?.phases[mindex]?.tasks.splice(tindex, 1)
    }

    const loadPage = async () => {
        getList('Skills')
			.then((skillList: any) => {
				skillsOpts = findSelectOptsFromList(skillList);
			});
    }


    let skillsOpts = $state([] as SelectOptionType<string>[]);
    let templateForm = $derived(template) as Projecttemplate
</script>

<CSSection>

<ClientErrorSummary {errors} />


{#await loadPage()}
    Loading...
{:then promiseData} 

<TextInput
    bind:value={templateForm.name}
    placeholder="Template name"
    fieldName="Template Name"
/>
<TextAreaInput
    bind:value={templateForm.description}
    rows={4}
    fieldName="Template Description"
    placeholder="Template Description"
/>

<SectionSubHeading>
    Milestone Phases
    <span class="float-right">
        <button onclick={() => templateForm.phases.push({name: 'New Phase', description: '', id: '', order: 100, tasks: []})}>Add Phase</button>
    </span>
</SectionSubHeading>
{#each templateForm.phases as phase, i}
<Accordion>
    <AccordionItem>
        {#snippet header()}Milestone Phase: {phase.name}{/snippet}
        <TextInput
            bind:value={templateForm.phases[i].name}
            placeholder="Phase name"
            fieldName="Phase Name"
        />
        <TextAreaInput
            bind:value={templateForm.phases[i].description}
            rows={2}
            fieldName="Phase Description"
            placeholder="Phase Description"
        />

        <div class="mb-8">
            <span class="float-right">
                <Button color="red" onclick={() => delMilestone(i)}>Delete Milestone</Button>
            </span>
            <br class="clear-both" />
        </div>

        <SectionSubHeading>
            Phase Tasks
            <span class="float-right">
                <button onclick={() => templateForm.phases[i].tasks?.push({ name: 'New Task', description: '', requiredSkillID: '' })}>Add Task</button>
            </span>
        </SectionSubHeading>
        {#if templateForm.phases[i].tasks && templateForm.phases[i].tasks.length > 0}
        <Accordion flush>
        {#each templateForm.phases[i].tasks as task, j}
            <AccordionItem>
                {#snippet header()}<ToolsOutline size="sm" class="float-left mr-3" /> {task.name}{/snippet}
                <TextInput
                    bind:value={templateForm.phases[i].tasks[j].name}
                    placeholder="Task name"
                    fieldName="Task Name"
                />
                <TextAreaInput
                    bind:value={templateForm.phases[i].tasks[j].description as string}
                    rows={2}
                    fieldName="Task Description"
                    placeholder="TaskDescription"
                />

                <SelectInput
                    bind:value={templateForm.phases[i].tasks[j].requiredSkillID as string}
                    fieldName="Required Skills"
                    bind:options={skillsOpts}
                />

                
                <div class="mb-8">
                    <span class="float-right">
                        <Button pill color="red" onclick={() => delTask(i, j)}>Delete Task</Button>
                    </span>
                    <br class="clear-both" />
                </div>
            </AccordionItem>
            {/each}
        </Accordion>
        {:else}
            <div class="text-center p-6">No Tasks</div>
        {/if}
    </AccordionItem>
    </Accordion>
{/each}

<div class="mt-4">
<span class="float-right">
<Button onclick={saveTemplate}>Save Template</Button>
</span> 
</div>  
    
{/await}

    
</CSSection>


