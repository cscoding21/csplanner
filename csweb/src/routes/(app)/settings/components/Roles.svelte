<script lang="ts">
	import { findAllRoles } from "$lib/services/resource";
    import type { Role, RoleResults } from "$lib/graphql/generated/sdk";
    import { addToast } from "$lib/stores/toasts";
    import { SectionHeading } from "$lib/components";
    import { Table, TableBodyCell, TableBodyRow, TableHead, TableHeadCell, TableBody, ButtonGroup, Button } from "flowbite-svelte";
    import { formatCurrency } from "$lib/utils/format";
    import { DeleteRole, RoleFormModal } from ".";
    import { TrashBinOutline, EditOutline, PlusOutline } from "flowbite-svelte-icons";

    let roles = $state([] as Role[])

    function refresh() {
		load().then(l => {
            roles = l.results as Role[]
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
</script>


{#await loadPage()}
    Loading...
{:then promiseData} 
<SectionHeading>
    Roles
    <span class="float-right">
        <ButtonGroup>
            <RoleFormModal
                    size="md"
                    update={() => refresh()}>
                    <PlusOutline class="mr-2 h-4 w-4" /> Create Role
            </RoleFormModal>
        </ButtonGroup>
    </span>
</SectionHeading>
<Table hoverable={true}>
    <TableHead>
      <TableHeadCell>Role</TableHeadCell>
      <TableHeadCell>Description</TableHeadCell>
      <TableHeadCell>Hourly rate</TableHeadCell>
      <TableHeadCell>Actions</TableHeadCell>
    </TableHead>
    <TableBody tableBodyClass="divide-y">
        {#each roles as role}
      <TableBodyRow>
        <TableBodyCell>{role.name}</TableBodyCell>
        <TableBodyCell>{role.description}</TableBodyCell>
        <TableBodyCell>{formatCurrency.format(role.hourlyRate || 0)}</TableBodyCell>
        <TableBodyCell>
            <ButtonGroup>
                <RoleFormModal
                        role={role}
                        size="md"
                        update={() => refresh()}>
                        <EditOutline size="sm" />
                </RoleFormModal>
            <DeleteRole
                id={role.id || ''}
                size="md"
                name="Delete role"
                update={() => refresh()}
            >
                <TrashBinOutline size="sm" class="" />
            </DeleteRole>
            </ButtonGroup>
        </TableBodyCell>
      </TableBodyRow>
      {/each}
    </TableBody>
  </Table>
{/await}