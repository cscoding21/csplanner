<script lang="ts">
	import { SectionHeading } from "$lib/components";
	import SectionSubHeading from "$lib/components/formatting/SectionSubHeading.svelte";
	import type { UpdateList, UpdateListItem } from "$lib/graphql/generated/sdk";
	import { getList, updateList } from "$lib/services/list";
	import { addToast } from "$lib/stores/toasts";
	import { callIf } from "$lib/utils/helpers";
	import { nameToID } from "$lib/utils/id";
	import { Alert, Badge, Button, Input, P, Popover } from "flowbite-svelte";
	import { CloseCircleSolid } from "flowbite-svelte-icons";

    interface Props {
        onDone: Function
    }
    let { onDone }:Props = $props()

    let roleList:string[] = $state([])
    let roleSkill:any = $state({name: "", hourlyRate: 0.0, skills: []})

    const addRoleGroup = (rg:string[]) => {
        roleList = [...new Set([...roleList, ...rg])]
    }

    const saasSkillGroup = [
        "Backend Development",
        "Frontend Development",
        "Data Architecture",
        "DevOps / SRE",
        "Data Science"
    ]

    const projectSkillGroup = [
        "Project Management",
        "Business Analysis",
        "Agile Development",
        "Scrum",
        "Kanban"
    ]

    const autonomySkillGroup = [
        "State Estimation",
        "Scene Understanding",
        "C++",
        "Mapping",
        "Foundations"
    ]

</script>


<h2 class="text-xl text-center text-gray-50 font-semibold">Organization Roles</h2>
    
<section class="mb-6 mt-4">
<p class="py-2 text-gray-200">
    Roles often map to job titles.  In csPlanner, they are convenience constructs that ease the burden of 
    calculating costs and assigning skills.
</p>
<p class="py-2 text-gray-200">Roles can be added or modified at any time on the <a class="text-orange-300" href="/settings#roles">Roles</a> page.
</p>
</section>

<SectionSubHeading>Common Skill Groups</SectionSubHeading>
<small>Here are some common groups of skills.</small>
<div class="mb-4">
    <Button id="saasSkillButton" class="m-2" color="alternative" pill onclick={() => addRoleGroup(saasSkillGroup)}>SaaS</Button>
    <Popover triggeredBy="#saasSkillButton" class="w-72 text-sm font-light text-gray-500 bg-white dark:bg-gray-800 dark:border-gray-600 dark:text-gray-400">
        <div class="p-3 space-y-2">
            Add the following roles to your stack:
            <ul class="list-disc pl-4">
            {#each saasSkillGroup as r}
                <li>{r}</li>
            {/each}
            </ul>
        </div>
    </Popover>


    <Button id="projectSkillButton" class="m-2" color="alternative" pill onclick={() => addRoleGroup(projectSkillGroup)}>Project Management</Button>
    <Popover triggeredBy="#projectSkillButton" class="w-72 text-sm font-light text-gray-500 bg-white dark:bg-gray-800 dark:border-gray-600 dark:text-gray-400">
        <div class="p-3 space-y-2">
            Add the following skills to your stack:
            <ul class="list-disc pl-4">
            {#each projectSkillGroup as sk}
                <li>{sk}</li>
            {/each}
            </ul>
        </div>
    </Popover>

    <Button id="autonomySkillGroup" class="m-2" color="alternative" pill onclick={() => addRoleGroup(autonomySkillGroup)}>Autonomy</Button>
    <Popover triggeredBy="#autonomySkillGroup" class="w-72 text-sm font-light text-gray-500 bg-white dark:bg-gray-800 dark:border-gray-600 dark:text-gray-400">
        <div class="p-3 space-y-2">
            Add the following skills to your stack:
            <ul class="list-disc pl-4">
            {#each autonomySkillGroup as sk}
                <li>{sk}</li>
            {/each}
            </ul>
        </div>
    </Popover>

</div>
