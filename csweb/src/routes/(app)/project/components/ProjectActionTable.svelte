<script lang="ts">
    import type { ProjectEnvelope, Resource } from "$lib/graphql/generated/sdk";
	import { Table, TableBody, TableBodyCell, TableBodyRow } from "flowbite-svelte";
	import { ResourceList } from "$lib/components";
	import { ProjectStatusBrief, BadgeProjectStatus } from ".";
	import ShowIfStatus from "./ShowIfStatus.svelte";

    interface Props {
        projects:ProjectEnvelope[]
    }
    let { projects }:Props = $props()
</script>

<Table>
    <!-- <TableHead>
        <TableHeadCell>Project</TableHeadCell>
        <TableHeadCell>Status</TableHeadCell>
        <TableHeadCell>Team</TableHeadCell>
        <TableHeadCell></TableHeadCell>
    </TableHead> -->
    <TableBody class="divide-y">
        {#each projects as project}
            <TableBodyRow>
            <TableBodyCell>
            <a href="/project/detail/{project?.meta.id}#snapshot" class="text-md dark:text-gray-300">
                {project.data.projectBasics.name}
            </a>
            </TableBodyCell>
            <TableBodyCell><BadgeProjectStatus status={project?.data?.projectStatusBlock.status} /></TableBodyCell>
            <TableBodyCell><ResourceList size="md" maxSize={4} resources={project?.data.calculated?.teamMembers as Resource[]} /></TableBodyCell>
            <TableBodyCell>
                <ProjectStatusBrief project={project?.data} />
            </TableBodyCell>
            </TableBodyRow>
        {/each}
    </TableBody>
</Table>



