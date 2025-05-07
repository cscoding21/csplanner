<script lang="ts">
	import type { Organization, Resource, ResourceResults, UpdateResource} from "$lib/graphql/generated/sdk";
	import { getOrganization } from "$lib/services/organization";
	import { createResource, deleteResource, findAllResources, findAllRoles } from "$lib/services/resource";
	import { addToast } from "$lib/stores/toasts";
	import { newID } from "$lib/utils/id";
	import { Alert, Button, ButtonGroup, Input, Select, Table, TableBody, TableBodyCell, TableBodyRow, TableHead, TableHeadCell, type SelectOptionType } from "flowbite-svelte";
	import { PenNibOutline, TrashBinOutline } from "flowbite-svelte-icons";
	import UpdateResourceModal from "../../resource/components/UpdateResourceModal.svelte";
	import { formatCurrency } from "$lib/utils/format";
	import { goto } from "$app/navigation";

    interface Props {
        onDone: Function
    }
    let { onDone }:Props = $props()

	let resourceList:Resource[] = $state([] as Resource[])
	let newResourceName:string = $state("")
	let newResourceRoleID:string = $state("")
	let org = $state({} as Organization)

	let roleOpts:SelectOptionType<string>[] = $state([] as SelectOptionType<string>[])

	function refresh() {
		load().then(l => {
            resourceList = l.results?.map(r => r.data) as Resource[]
        });
	}

	const removeRes = (id:string) => {
		deleteResource(id).then(res => {
			if (res && res.success) {
				addToast({
					message: 'Resource deleted successfully',
					dismissible: true,
					type: 'success'
				});

				refresh()
			} else {
				addToast({
					message: 'Error deleting resource: ' + res.message,
					dismissible: true,
					type: 'error'
				});
			}
		})
	}

	const addRes = () => {
		console.log("add resource", newResourceName, newResourceRoleID)
		const res:UpdateResource = {
			name: newResourceName,
			status: "inhouse",
			type: "human",
			roleID: newResourceRoleID,
			initialCost: 0,
			annualizedCost: 0,
			availableHoursPerWeek: org.defaults.hoursPerWeek,
			id: newID()
		}

		createResource(res).then(res => {
			if (res && res.status.success) {
				addToast({
					message: 'Resource created successfully',
					dismissible: true,
					type: 'success'
				});

				newResourceName = ""
				newResourceRoleID = ""

				refresh()
			} else {
				addToast({
					message: 'Error creating resource: ' + res.status.message,
					dismissible: true,
					type: 'error'
				});
			}
		})
	}

	const saveResources = () => {
		goto("/")
		onDone()
	}

	const load = async ():Promise<ResourceResults> => {
		return await findAllRoles()
			.then(roles => {
				roleOpts = roles.results?.map(r => { return { name: r.data.name, value: r.data.id} }) as SelectOptionType<string>[]

				return roles
			})
			.then(() => {
				getOrganization().then(o => {
					org = o
				})
			})
			.then(
				res => findAllResources()
				.then(r => {
					return r
				})
				.catch((err) => {
					addToast({
						message: 'Error loading resources (Setup.Resources): ' + err,
						dismissible: true,
						type: 'error'
					});

					return err
				})
			);
	};

	const loadPage = async () => {
        refresh()
	};
</script>


<h2 class="text-xl text-center text-gray-50 font-semibold">Organization Resources</h2>
    
<section class="mb-6 mt-4">
<p class="py-2 text-gray-200">
    Resources are the staff, hardware, software, etc required to deliver on your project roadmap.
</p>
<p class="py-2 text-gray-200">Resources can be added or modified at any time on the <a class="text-orange-300" href="/resources">Resources</a> page.
</p>
</section>

{#await loadPage()}
	Loading...
{:then promiseData}

<div class="p-4">
    
    <Table hoverable={true}>
        <TableHead>
          <TableHeadCell>Name</TableHeadCell>
          <TableHeadCell>Role</TableHeadCell>
		  <TableHeadCell>Hours/Week</TableHeadCell>
		  <TableHeadCell>Hourly Rate</TableHeadCell>
          <TableHeadCell>Actions</TableHeadCell>
        </TableHead>
        <TableBody tableBodyClass="divide-y">
			{#if resourceList.length > 0}
            {#each resourceList as r, index}
          <TableBodyRow>
            <TableBodyCell>{r.name}</TableBodyCell>
            <TableBodyCell>
                {r.role?.name}
            </TableBodyCell>
			<TableBodyCell class="text-right">{r.availableHoursPerWeek}</TableBodyCell>
			<TableBodyCell class="text-right">{formatCurrency.format(r.calculated.hourlyCost as number)}</TableBodyCell>
            <TableBodyCell class="text-right">
                <ButtonGroup>
                    <Button onclick={() => removeRes(r.id as string)} type="button" aria-label="Remove">
                    	<TrashBinOutline size="sm" class="" />
					</Button>   
					<UpdateResourceModal update={() => console.log("update")} id={r.id as string}>
						<PenNibOutline />
					</UpdateResourceModal>
                </ButtonGroup>
            </TableBodyCell>
          </TableBodyRow>
          {/each}
		  {:else}
			<TableBodyRow>
				<TableBodyCell colspan={3}><Alert>No resources yet</Alert></TableBodyCell>
			</TableBodyRow>
  		  {/if} 
        </TableBody>
      </Table>

	  <div class="grid mt-4 grid-cols-5 gap-4 w-full px-2">
		<div class="col-span-2"><Input bind:value={newResourceName} placeholder="Resource name" /></div>
		<div class="col-span-2"><Select bind:value={newResourceRoleID} items={roleOpts} placeholder="Role" /></div>
		<div class="text-center"><Button onclick={addRes}>Add</Button>	</div>
	  </div>
      

  <div class="mt-12 text-center">
    <Button onclick={saveResources}>I'm done with resource for now.  Let's me in!</Button>
  </div>
</div>  
	
{/await}