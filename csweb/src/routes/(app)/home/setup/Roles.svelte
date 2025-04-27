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
    let newRole:any = $state({name: "", hourlyRate: 0.0, skills: []})

    const addRoleGroup = (rg:any[]) => {
        roleList = [...new Set([...roleList, ...rg])]
    }

    const saasRoleGroup = [
        {
           name: "Full Stack Engineer",
           hourlyRate: 140.0,
           skills: [
            "Golang",
            "Database",
            "JavaScript",
            "TypeScript"
            ]
        },
        {
           name: "Web Developer",
           hourlyRate: 100.0,
           skills: [
            "Svelte",
            "CSS",
            "JavaScript",
            "HTML"
            ]
        },
        {
           name: "DevOps Engineer/SRE",
           hourlyRate: 140.0,
           skills: [
            "Kubernetes",
            "AWS",
            "Keycloak"
            ]
        },
        {
           name: "Engineering Manager",
           hourlyRate: 180.0,
           skills: [
            "Golang",
            "JavaScript",
            "Leadership",
            "Agile"
            ]
        },
    ]

    const projectRoleGroup = [
        {
           name: "Project Manager",
           hourlyRate: 120.0,
           skills: [
            "Business Analysis",
            "Communications"
            ]
        },
        {
           name: "Product Manager",
           hourlyRate: 100.0,
           skills: [
            "Business Analysis",
            "Product Development"
            ]
        },
        {
           name: "Product Owner",
           hourlyRate: 110.0,
           skills: [
            "Scrum",
            "Agile"
            ]
        },
        {
           name: "Product Designer",
           hourlyRate: 130.0,
           skills: [
            "UI",
            "UX",
            "Product Design"
            ]
        },
    ]

    const autonomyRoleGroup = [
        {
           name: "AI Engineer",
           hourlyRate: 140.0,
           skills: [
            "Golang",
            "Database",
            "JavaScript",
            "TypeScript"
            ]
        },
        {
           name: "Autonomy Engineering Manager",
           hourlyRate: 200.0,
           skills: [
            "Svelte",
            "CSS",
            "JavaScript",
            "HTML"
            ]
        },
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

<SectionSubHeading>Common Roles in Organizations</SectionSubHeading>
<small>Here are some common groups of roles.</small>
<div class="mb-4">
    {@render roleGroup("SaaS Engineering", "saasRoleGroup", saasRoleGroup)}
    {@render roleGroup("Product Management", "projectRoleGroup", projectRoleGroup)}
    {@render roleGroup("Autonomy", "autonomyRoleGroup", autonomyRoleGroup)}
</div>


{#snippet roleGroup(name:string, id:string, group:any)}
<Button id={id} class="m-2" color="alternative" pill onclick={() => addRoleGroup(group)}>{name}</Button>
<Popover triggeredBy={"#" + id} class="w-72 text-sm font-light text-gray-500 bg-white dark:bg-gray-800 dark:border-gray-600 dark:text-gray-400">
    <div class="p-3 space-y-2">
        Add the following roles to your organization:
        <ul class="list-disc pl-4">
        {#each group as role}
            <li>{role.name}</li>
        {/each}
        </ul>
    </div>
</Popover>
{/snippet}