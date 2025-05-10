<script lang="ts">
    import { 
        SectionHeading,
        SectionSubHeading
    } from "$lib/components";
    import { BadgeFeaturePriority, BadgeProjectStatus, BadgeFeatureStatus, ShowIfStatus } from ".";
    import { Table, TableBody, TableHead, TableBodyRow, TableHeadCell, TableBodyCell, Popover, Hr } from "flowbite-svelte";
    import { QuestionCircleSolid } from "flowbite-svelte-icons";
    import type { ProjectEnvelope } from '$lib/graphql/generated/sdk'
    import { getProject } from "$lib/services/project";
    import { getDefaultProject } from "$lib/forms/project.validation";
    import { normalizeGUID } from "$lib/utils/id";
	import Draft from "./snapshots/Draft.svelte";
	import New from "./snapshots/New.svelte";

    interface Props {
        id: string
    }
    let { id }:Props = $props()

    let project:ProjectEnvelope = $state(getDefaultProject() as ProjectEnvelope)

    const load = async (): Promise<ProjectEnvelope> => {
		return await getProject(id)
			.then((proj) => {
				return proj;
			})
			.catch((err) => {
				return err;
			});
	};

    const loadPage = async () => {
        load().then(p => {
            project = p
        });
	};

</script>


{#await loadPage()}
	Loading...
{:then promiseData}

{#if project}

<!-- col span 3 -->
<div class="col-span-3">
<SectionHeading>
    Snapshot: {project.data?.projectBasics.name}
    <span class="float-right"><BadgeProjectStatus status={project.data?.projectStatusBlock?.status} /></span>
</SectionHeading>
</div>



<ShowIfStatus status={project.data?.projectStatusBlock.status} scope={["new"]}>
    <New project={project.data} />
</ShowIfStatus>

<ShowIfStatus status={project.data?.projectStatusBlock.status} scope={["draft"]}>
    <Draft project={project.data} />
</ShowIfStatus>


<!-- <div class="col-span-1">
    <SectionSubHeading>Executive Summary</SectionSubHeading>
    <p class="mb-6 text-sm">{project.data?.projectBasics?.description}</p>
</div> -->




<!-- <div class="col-span-2">
    <SectionSubHeading>Features</SectionSubHeading>
    <Table>
        <TableHead>
          <TableHeadCell>Feature</TableHeadCell>
          <TableHeadCell>Priority</TableHeadCell>
          <TableHeadCell>Status</TableHeadCell>
        </TableHead>
        <TableBody tableBodyClass="divide-y">
        {#if project.data?.projectFeatures}
            {#each project.data?.projectFeatures as feature}
          <TableBodyRow>
            <TableBodyCell>
                {feature.name}
                <button id={'feat' + normalizeGUID(feature.id)}>
                    <QuestionCircleSolid class="w-3 h-3 ms-1.5" />
                    <span class="sr-only">Show information</span>
                </button>
                  <Popover triggeredBy={'#feat' + normalizeGUID(feature.id)} class="w-72 text-sm font-light text-gray-500 bg-white dark:bg-gray-800 dark:border-gray-600 dark:text-gray-400" placement="bottom-start">
                    <div class="p-3 space-y-2">
                      {feature.description}
                    </div>
                  </Popover>
            </TableBodyCell>
            <TableBodyCell><BadgeFeaturePriority priority={feature.priority} /></TableBodyCell>
            <TableBodyCell><BadgeFeatureStatus status={feature.status} /></TableBodyCell>
          </TableBodyRow>
          {/each}
        {/if}
        </TableBody>
    </Table>
</div> -->

{/if}

{/await}
