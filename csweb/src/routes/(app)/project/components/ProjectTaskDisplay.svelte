<script lang="ts">
    import { P, Button, ButtonGroup, Modal } from "flowbite-svelte";
    import { EditOutline, TrashBinOutline } from "flowbite-svelte-icons";
    import type { ProjectMilestoneTask, Skill } from "$lib/graphql/generated/sdk";
    import { DeleteProjectTask, ProjectTaskForm } from ".";
    import { ResourceList } from "$lib/components";
    import { callIf } from "$lib/utils/helpers";
    import { formatToCommaSepList, pluralize } from "$lib/utils/format";

    interface Props {
        task: ProjectMilestoneTask
        update: Function
        editClick: Function
        milestoneID: string
        projectID: string
    }
    let { 
        task, 
        update,
        editClick,
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


<div>
    <span class="text-gray-100">{task.name} 
        <small>({task.hourEstimate} {pluralize("hour", task.hourEstimate)})</small>
    </span>
    <P class="mb-3" weight="light" color="text-gray-500 dark:text-gray-100 float-right">
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
    </P>
</div>
<div>
    <div class="text-sm text-gray-400">{task.description}</div>
    <div class="my-2 text-sm text-gray-100">
        Skills Required: {getRequiredSkillsDisplay(task)}
    </div>

    <div class="pl-4">
        <ResourceList maxSize={4} size="sm" resources={task.resources || []} />
    </div>
</div>
<div class="mt-2">
    
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