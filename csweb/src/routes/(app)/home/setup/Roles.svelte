<script lang="ts">
	import { SectionHeading } from "$lib/components";
	import SectionSubHeading from "$lib/components/formatting/SectionSubHeading.svelte";
	import type { Role } from "$lib/graphql/generated/sdk";
	import { formatCurrency, pluralize } from "$lib/utils/format";
	import { nameToID } from "$lib/utils/id";
	import { Alert, Button, ButtonGroup, Popover, Table, TableBody, TableBodyCell, TableBodyRow, TableHead, TableHeadCell } from "flowbite-svelte";
	import { TrashBinOutline } from "flowbite-svelte-icons";
	import { RoleFormModal } from "../../settings/components";
    import { roleGroups } from "./roleGroups";

    interface Props {
        onDone: Function
    }
    let { onDone }:Props = $props()

    let roleList:Role[] = $state([] as Role[])
    let newRole:any = $state({} as Role)
    let openModal:boolean = $state(false)

    const addRoleGroup = (rg:any[]) => {
        roleList = [...new Set([...roleList, ...rg])]
    }

    const removeRole = (i:number) => {
        roleList.splice(i, 1)
        roleList = roleList
    }

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
    {#each roleGroups as rg, index}
    {@render roleGroup(rg.name, rg.id, rg.roles)}
    {/each}
</div>


<SectionSubHeading>Your Roles</SectionSubHeading>
<div class="p-4">
    {#if roleList.length > 0}
    <Table hoverable={true}>
        <TableHead>
          <TableHeadCell>Role</TableHeadCell>
          <TableHeadCell>Hourly rate</TableHeadCell>
          <TableHeadCell>Skills</TableHeadCell>
          <TableHeadCell>Actions</TableHeadCell>
        </TableHead>
        <TableBody tableBodyClass="divide-y">
            {#each roleList as r, index}
          <TableBodyRow>
            <TableBodyCell>{r.name}</TableBodyCell>
            <TableBodyCell>{formatCurrency.format(r.hourlyRate || 0)}</TableBodyCell>
            <TableBodyCell>
                {@render skillEditor(r.name, r)}
            </TableBodyCell>
            <TableBodyCell>
                <ButtonGroup>
                    <button onclick={() => removeRole(index)} type="button" class="inline-flex items-center rounded-full p-0.5 my-0.5 ms-1.5 -me-1.5 text-sm text-white dark:text-primary-80 hover:text-whit dark:hover:text-white" aria-label="Remove">
                    <TrashBinOutline size="sm" class="" />
                    </button>   
                </ButtonGroup>
            </TableBodyCell>
          </TableBodyRow>
          {/each}
        </TableBody>
      </Table>
      {:else}
      <Alert>No roles yet</Alert>
  {/if} 
    
</div>


{#snippet roleGroup(name:string, id:string, group:any)}
<Button id={"rg_" + id} class="m-2" color="alternative" pill onclick={() => addRoleGroup(group)}>{name}</Button>
<Popover triggeredBy={"#rg_" + id} class="w-72 text-sm font-light text-gray-500 bg-white dark:bg-gray-800 dark:border-gray-600 dark:text-gray-400">
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


{#snippet skillEditor(id:string, role:Role)}

<RoleFormModal size="md" update={() => console.log("update")}>
    <button id={nameToID(id)} class="m-1">{role.defaultSkills?.length} {pluralize("skill", role.defaultSkills?.length as number)}</button>
</RoleFormModal>

{/snippet}
