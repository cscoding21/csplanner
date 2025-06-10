<script lang="ts">
    import { 
        SectionHeading
    } from "$lib/components";
    import { BadgeProjectStatus, ShowIfStatus } from ".";
    import type { ProjectEnvelope } from '$lib/graphql/generated/sdk'
    import { getProject } from "$lib/services/project";
    import { getDefaultProject } from "$lib/forms/project.validation";
	import Draft from "./snapshots/Draft.svelte";
	import New from "./snapshots/New.svelte";
    import Proposed from "./snapshots/Proposed.svelte";
    import Exception from "./snapshots/Exception.svelte";
	import Shelved from "./snapshots/Shelved.svelte";
	import Approved from "./snapshots/Approved.svelte";
	import Scheduled from "./snapshots/Scheduled.svelte";
	import InFlight from "./snapshots/InFlight.svelte";

    interface Props {
        id: string
    }
    let { id }:Props = $props()

    let project:ProjectEnvelope = $state(getDefaultProject() as ProjectEnvelope)

    const load = async (): Promise<ProjectEnvelope> => {
		return await getProject(id)
			.then((proj) => {
                project = proj

				return proj;
			})
			.catch((err) => {
				return err;
			});
	};

    const loadPage = async () => {
        load();
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

<ShowIfStatus status={project.data?.projectStatusBlock.status} scope={["proposed"]}>
    <Proposed project={project.data} />
</ShowIfStatus>

<ShowIfStatus status={project.data?.projectStatusBlock.status} scope={["approved"]}>
    <Approved project={project.data} />
</ShowIfStatus>

<ShowIfStatus status={project.data?.projectStatusBlock.status} scope={["scheduled"]}>
    <Scheduled project={project.data} />
</ShowIfStatus>

<ShowIfStatus status={project.data?.projectStatusBlock.status} scope={["inflight"]}>
    <InFlight project={project.data} update={() => load()} />
</ShowIfStatus>



<ShowIfStatus status={project.data?.projectStatusBlock.status} scope={["deferred", "abandoned", "rejected", "backlogged", "complete"]}>
    <Shelved project={project.data} />
</ShowIfStatus>

<ShowIfStatus status={project.data?.projectStatusBlock.status} scope={["exception"]}>
    <Exception project={project.data} />
</ShowIfStatus>

{/if}

{/await}
