<script lang="ts">
	import { decodeProficiency, deleteResourceSkill, updateResourceSkill } from "$lib/services/resource";
	import { Table, TableHead, TableHeadCell, TableBody, TableBodyRow, TableBodyCell, Button, Alert, Range, Select, type SelectOptionType } from "flowbite-svelte";
	import { TrashBinOutline } from "flowbite-svelte-icons";
	import { AddSkill } from ".";
	import type { Skill } from "$lib/graphql/generated/sdk";
	import { skillForm, skillSchema } from "$lib/forms/skill.validation";
	import { deepCopy } from "$lib/utils/helpers";
	import { addToast } from "$lib/stores/toasts";
	import { findSelectOptsFromList, mergeErrors, parseErrors } from "$lib/forms/helpers";
	import { getList } from "$lib/services/list";
	import { SelectInput } from "$lib/components";
	import SectionSubHeading from "$lib/components/formatting/SectionSubHeading.svelte";

    interface Props {
		parentID: string;
        skills: Skill[]
		update: Function;
        allowAdd?: Boolean
	}
	let { parentID, update, skills, allowAdd }: Props = $props();

    const deleteSkill = async (skillID: string) => {
		deleteResourceSkill(parentID, skillID).then((res) => {
			if (res.success) {
				addToast({
					message: 'Resource skill deleted successfully',
					dismissible: true,
					type: 'success'
				});

				update(parentID)
			} else {
				addToast({
					message: 'Error updating skill: ' + res.message,
					dismissible: true,
					type: 'error'
				});
			}
		});
	};

    const addSkill = () => {
        upsertSkill(newSkill.skillID, newSkill.proficiency)
    }

    const upsertSkill = (sid:string, proficiency:number) => {
		errors = {};

        const usf = {
            skillID: sid,
            proficiency: proficiency,
            resourceID:parentID
        }

		const parsedSkillForm = skillSchema.cast(usf);
		parsedSkillForm.resourceID = parentID;
		skillSchema
			.validate(parsedSkillForm, { abortEarly: false })
			.then(() => {
				updateResourceSkill({
					resourceID: parsedSkillForm.resourceID,
					id: parsedSkillForm.skillID,
					proficiency: parsedSkillForm.proficiency as number
				}).then((res) => {
					if (res.success) {
						addToast({
							message: 'Skill successfully added to resource',
							dismissible: true,
							type: 'success'
						});

						newSkill.id = undefined;
						newSkill.proficiency = 2;

						update()
					} else {
						addToast({
							message: 'Error creating resource: ' + res.message,
							dismissible: true,
							type: 'error'
						});
					}
				});
			})
			.catch((err) => {SelectInput
				errors = mergeErrors(errors, parseErrors(err));

				console.log(errors);
			});
	};

    let availableSkillOpts = $state([] as SelectOptionType<string>[]);
        
    const loadPage = async () => {
		getList('Skills')
			.then((l) => {
				availableSkillOpts = findSelectOptsFromList(l);
			})
			.then(() => {
				skillForm.resourceID = parentID;
			});
	};

	let errors: any = $state({ skillID: '', proficiency: '' });
	let newSkill = $state(deepCopy(skillForm));
    let updateSkills = $derived(deepCopy(skills))
    
</script>

{#await loadPage()}
    Loading...
{:then promiseData} 
{#if skills && skills.length > 0}
<Table>
    <TableHead>
        <TableHeadCell>Skill</TableHeadCell>
        <TableHeadCell>Proficienty</TableHeadCell>
        <TableHeadCell>
            <span class="sr-only">Action</span>
        </TableHeadCell>
    </TableHead>
    <TableBody>
        {#each updateSkills as s, index}
            <TableBodyRow>
                <TableBodyCell>{s.name}</TableBodyCell>
                <TableBodyCell>
                    <Range size="lg" id="range-steps" min="1" max="3" bind:value={updateSkills[index].proficiency as number} onchange={e => upsertSkill(s.id, updateSkills[index].proficiency as number)} step="1" />
                    <br /><small>{decodeProficiency(s.proficiency as number)}</small>
                </TableBodyCell>
                <TableBodyCell tdClass="float-right pt-2">
                    <Button color="dark" onclick={() => deleteSkill(s.id)}>
                        <TrashBinOutline size="sm" />
                    </Button>
                </TableBodyCell>
            </TableBodyRow>
        {/each}
    </TableBody>
    {#if allowAdd}
    <tfoot>
        <tr>
            <th class="px-6 pt-4" colspan={3}><SectionSubHeading>Add skill</SectionSubHeading></th>
    </tr>    
        <tr>
            <th class="px-4">
                <Select
                    bind:value={newSkill.skillID}
                    size="sm"
                    items={availableSkillOpts}
                />
            </th>
            <th class="px-6 py-4">
                <Range size="lg" id="range-steps" min="1" max="3" bind:value={newSkill.proficiency} step="1" />
                <br /><small>{decodeProficiency(newSkill.proficiency)}</small>
            </th>
            <th class="px-6 py-4">
                <Button pill onclick={addSkill}>Add</Button>
            </th>
        </tr>
    </tfoot>
    {/if}
</Table>
{:else}
<Alert>No skills have been added for this resource.</Alert>
{#if allowAdd}
    <hr class="my-4" />
    <AddSkill resourceID={parentID} update={() => update(parentID)} />
{/if}
{/if}


{/await}
