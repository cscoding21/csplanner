<script lang="ts">
	import SectionSubHeading from "$lib/components/formatting/SectionSubHeading.svelte";
	import type { UpdateList, UpdateListItem } from "$lib/graphql/generated/sdk";
	import { getList, updateList } from "$lib/services/list";
	import { addToast } from "$lib/stores/toasts";
	import { callIf } from "$lib/utils/helpers";
	import { nameToID } from "$lib/utils/id";
	import { Alert, Badge, Button, Input, P, Popover } from "flowbite-svelte";
	import { CloseCircleSolid } from "flowbite-svelte-icons";
    import { skillGroups } from "./skillGroups";

    interface Props {
        onDone: Function
    }
    let { onDone }:Props = $props()
    let skillList:string[] = $state([])
    let newSkill:string = $state("")

    const listName = "list:skills"

    const addSkill = () => {
        if (!newSkill || !okToAdd(newSkill)) {
            return
        }

        skillList = [...new Set([...skillList, newSkill])]
        newSkill = ""
    }

    const addSkillGroup = (sg:string[]) => {
        skillList = [...new Set([...skillList, ...sg])]
    }

    const removeSkill = (i:number) => {
        skillList.splice(i, 1)
        skillList = skillList
    }

    const checkEnter = ((event:KeyboardEvent) => {
        if(event.key === "Enter") {
            addSkill()
        }
    })

    const okToAdd = (skill:string):boolean => {
        const sk = skillList.filter(s => skill.toLocaleLowerCase() === s.toLocaleLowerCase())

        return sk.length === 0
    }

    const getSkillList = () => {
        let list:UpdateList = {
            id: listName,
            description: "A list of skills used by the organization",
            values: []
        }

        for(let i = 0; i < skillList.length; i++) {
            const sk = skillList[i]
            const listItem:UpdateListItem = {
                name: sk,
                value: nameToID(sk),
                sortOrder: 0,
            }

            list.values.push(listItem)
        }

        return list
    }

    const saveList = () => {
		const list = getSkillList()

		updateList(list).then(res => {
			if (res && res.status?.success) {
				addToast({
					message: 'Skills updated successfully',
					dismissible: true,
					type: 'success'
				});

                callIf(onDone)
			} else {
				addToast({
					message: 'Error updating skills: ' + res.status?.message,
					dismissible: true,
					type: 'error'
				});
			}
		});
	}

    const getExistingSkills = async () => {
        getList(listName).then(sk => {
            skillList = sk.values.map(s => s.name)
        })
    }

    getExistingSkills()
</script>

<h2 class="text-xl text-center text-gray-50 font-semibold">Organization Skills</h2>
    
<section class="mb-6 mt-4">
<p class="py-2 text-gray-200">
    Skills define the discrete capabilities needed to execute your organization's projects.
    The building block of a project is a task which maps to a single skill.  Allocating resources to 
    a task requires a matching skill by the resource.  Skills can be updated at any time.
</p>
</section>

<SectionSubHeading>Common Skill Groups</SectionSubHeading>
<small>Here are some common groups of skills.</small>
<div class="mb-4">
    {#each skillGroups as sg, index}
        {@render skillGroup(sg.name, sg.id, sg.skills)}
    {/each}
</div>

<SectionSubHeading >Add Your Skills</SectionSubHeading>
<div class="mb-4">
    <Input bind:value={newSkill} onclick={addSkill} onkeypress={checkEnter} placeholder="Type a skill name. <enter> to add" />
</div>

<SectionSubHeading>Your Skills</SectionSubHeading>
<div class="p-4">
    {#if skillList.length > 0}
    {#each skillList as skill, index}
        <Badge color="yellow" class="m-2" dismissable>
            {skill}
        <button slot="close-button" onclick={() => removeSkill(index)} type="button" class="inline-flex items-center rounded-full p-0.5 my-0.5 ms-1.5 -me-1.5 text-sm text-white dark:text-primary-80 hover:text-whit dark:hover:text-white" aria-label="Remove">
            <CloseCircleSolid class="h-4 w-4" />
            <span class="sr-only">Remove skill</span>
        </button>
        </Badge>
    {/each}
    {:else}
        <Alert>No skills yet</Alert>
    {/if}
</div>

<div class="mt-12 text-center">
<Button onclick={saveList}>I'm done with skills for now.  Let's keep going! >></Button>
</div>


{#snippet skillGroup(name:string, id:string, group:any)}
<Button id={id} class="m-2" color="alternative" pill onclick={() => addSkillGroup(group)}>{name}</Button>
<Popover triggeredBy={"#" + id} class="w-72 text-sm font-light text-gray-500 bg-white dark:bg-gray-800 dark:border-gray-600 dark:text-gray-400">
    <div class="p-3 space-y-2">
        Add the following roles to your organization:
        <ul class="list-disc pl-4">
        {#each group as sk}
            <li>{sk}</li>
        {/each}
        </ul>
    </div>
</Popover>
{/snippet}