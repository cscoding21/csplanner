<script lang="ts">
    import { P, Button, ButtonGroup, Modal, Heading } from "flowbite-svelte";
    import { EditOutline, TrashBinOutline } from "flowbite-svelte-icons";
    import type { ProjectMilestoneTask, Skill } from "$lib/graphql/generated/sdk";
    import { DeleteProjectTask, ProjectTaskForm, BadgeMilestoneStatus, ShowIfStatus } from ".";
    import { ResourceList } from "$lib/components";
    import { callIf } from "$lib/utils/helpers";
    import { formatToCommaSepList, pluralize } from "$lib/utils/format";

    interface Props {
        task: ProjectMilestoneTask
        update: Function
        editClick: Function
        projectStatus: string
        milestoneID: string
        projectID: string
    }
    let { 
        task = $bindable(), 
        update,
        editClick,
        projectStatus,
        milestoneID,
        projectID
    }:Props = $props()

    let modalState:boolean = $state(false)

    const getRequiredSkillsDisplay = (task: any): string => {
		if (task.skills == null || task.skills.length === 0) {
			return 'Not set';
		}

		return formatToCommaSepList(task.skills.map((s: Skill) => s.name));
	};

</script>

<div class="mb-2 rounded-lg bg-white p-4 shadow-sm dark:bg-gray-900 sm:p-6">
<div>
    <Heading tag="h5" class="text-lg dark:text-white">
        {task.name}

        <span class="float-right">
            <ShowIfStatus scope={["new", "draft"]} status={projectStatus}>
            <ButtonGroup>
            <Button
                size="sm"
                color="dark"
                onclick={() => {
                    editClick()
                }}
            >
                <EditOutline size="sm" class="" />
            </Button>
            <DeleteProjectTask
                name={task.name}
                size="sm"
                projectID={projectID}
                milestoneID={milestoneID}
                update={() => callIf(update)}
                id={task.id}
            >
                <TrashBinOutline size="sm" class="" />
            </DeleteProjectTask>
        </ButtonGroup>
    </ShowIfStatus>
        </span>
    
    </Heading>
    <div class="my-2">
    <small>({task.hourEstimate} {pluralize("hour", task.hourEstimate)})</small>
    <BadgeMilestoneStatus status={task.status} />
    </div>
</div>
<div>
    <div class="text-sm text-gray-400">{task.description}</div>
    <div class="my-2 text-xs text-gray-100">
        Skill Required: {getRequiredSkillsDisplay(task)}
    </div>

    <div class="pl-4">
        <ResourceList maxSize={4} size="sm" resources={task.resources || []} />
    </div>
</div>

</div>
<br class="clear-both" />

<Modal bind:open={modalState} size="xl" autoclose={false}>
    <ProjectTaskForm
        {task}
        milestoneID={milestoneID}
        projectID={projectID}
        update={() => { modalState = false; callIf(update) }}
    />
</Modal>