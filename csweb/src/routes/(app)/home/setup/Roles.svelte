<script lang="ts">
	import SectionSubHeading from "$lib/components/formatting/SectionSubHeading.svelte";
	import type { Role, RoleResults, UpdateRole, UpdateSkill } from "$lib/graphql/generated/sdk";
	import { formatCurrency, pluralize } from "$lib/utils/format";
	import { nameToID } from "$lib/utils/id";
	import { Alert, Button, ButtonGroup, Popover, Table, TableBody, TableBodyCell, TableBodyRow, TableHead, TableHeadCell } from "flowbite-svelte";
	import { PlusOutline, TrashBinOutline } from "flowbite-svelte-icons";
    import { roleGroups } from "./roleGroups";
	import RoleModalSetup from "./RoleModalSetup.svelte";
	import { findAllRoles, updateAllRoles } from "$lib/services/resource";
	import { addToast } from "$lib/stores/toasts";
	import { callIf } from "$lib/utils/helpers";

    interface Props {
        onDone: Function
    }
    let { onDone }:Props = $props()

    let roleList:Role[] = $state([] as Role[])

    const addRoleGroup = (rg:any[]) => {
        for (let i = 0; i < rg.length; i++) {
            const r = rg[i]
            if(roleList.some(av => av.id === r.id)) {
                continue
            }

            roleList.push(r)
            roleList = roleList
        }
    }

    const removeRole = (i:number) => {
        roleList.splice(i, 1)
        roleList = roleList
    }

    const saveRoles = () => {
        const updateRoleList = roleList.map(rl => { 
            return { 
                id: rl.id, 
                name: rl.name, 
                description: rl.description, 
                defaultSkills: rl.defaultSkills?.map(s => { return { id: s.id, skillID: s.skillID, proficiency: s.proficiency, parentID: rl.id } as UpdateSkill }), 
                hourlyRate: rl.hourlyRate
            } as UpdateRole } 
        )

		updateAllRoles(updateRoleList).then(res => {
			if (res && res.status?.success) {
				addToast({
					message: 'Roles updated successfully',
					dismissible: true,
					type: 'success'
				});

                callIf(onDone)
			} else {
				addToast({
					message: 'Error updating roles: ' + res.status?.message,
					dismissible: true,
					type: 'error'
				});
			}
		});
	}

    function refresh() {
		load().then(l => {
            roleList = l.results?.map(r => r.data) as Role[]
        });
	}

	const load = async ():Promise<RoleResults> => {
		return await findAllRoles()
			.then(r => {
				return r
			})
			.catch((err) => {
				addToast({
					message: 'Error loading roles (Roles): ' + err,
					dismissible: true,
					type: 'error'
				});

				return err
			});
	};

	const loadPage = async () => {
        refresh()
	};

    loadPage()

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

    <RoleModalSetup update={() => refresh()}>
    <Button color="blue" class="m-2" pill>Create Role</Button>
    </RoleModalSetup>   
</div>


<SectionSubHeading>
    Your Roles
</SectionSubHeading>
{#await loadPage()}
    Loading...
{:then pageData} 
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
                {@render skillEditor(r.id, r)}
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

  <div class="mt-12 text-center">
    <Button onclick={saveRoles}>I'm done with roles for now.  Let's keep going! >></Button>
    </div>
</div>    
{/await}



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

<RoleModalSetup {role} update={() => refresh()}>
    <span id={nameToID(id)} class="m-1">{role.defaultSkills?.length} {pluralize("skill", role.defaultSkills?.length as number)}</span>
</RoleModalSetup>

{/snippet}
