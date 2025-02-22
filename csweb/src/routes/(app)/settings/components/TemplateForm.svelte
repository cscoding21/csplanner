<script lang="ts">
	import { CSSection, TextAreaInput } from "$lib/components";
	import { TextInput } from "$lib/components";
	import SectionSubHeading from "$lib/components/formatting/SectionSubHeading.svelte";
	import type { Projecttemplate } from "$lib/graphql/generated/sdk";
	import { Accordion, AccordionItem, Button } from "flowbite-svelte";

    interface Props {
		id?: string;
		template: Projecttemplate | undefined;
		update?: Function;
	}
	let { 
		id, 
		template = $bindable(), 
		update 
	}: Props = $props();

    let templateForm = $derived(template) as Projecttemplate

    let errors = $state({name: '', description: ''})

</script>

<CSSection>
    <TextInput
        bind:value={templateForm.name}
        placeholder="Template name"
        fieldName="Template Name"
        error={errors.name}
    />
    <TextAreaInput
        bind:value={templateForm.description}
        rows={4}
        fieldName="Description"
        placeholder="Description"
        error={errors.description}
    />

    <SectionSubHeading>
        Milestone Phases
        <span class="float-right">
            <button onclick={() => templateForm.phases.push({name: 'New Phase', description: '', id: '', order: 100})}>Add Phase</button>
        </span>
    </SectionSubHeading>
    {#each templateForm.phases as phase, i}
    <Accordion>
        <AccordionItem>
            <span slot="header">Milestone Phase: {phase.name}</span>
            <TextInput
                bind:value={templateForm.phases[i].name}
                placeholder="Phase name"
                fieldName="Phase Name"
                error={errors.name}
            />
            <TextAreaInput
                bind:value={templateForm.phases[i].description}
                rows={2}
                fieldName="Description"
                placeholder="Description"
                error={errors.description}
            />
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
                <span slot="header">{task.name}</span>
                    <TextInput
                        bind:value={templateForm.phases[i].tasks[j].name}
                        placeholder="Task name"
                        fieldName="Task Name"
                        error={errors.name}
                    />
                    <TextAreaInput
                        bind:value={templateForm.phases[i].tasks[j].description as string}
                        rows={2}
                        fieldName="Description"
                        placeholder="Description"
                        error={errors.description}
                    />
                </AccordionItem>
                {/each}
            </Accordion>
            {/if}
        </AccordionItem>
      </Accordion>
    {/each}
</CSSection>


