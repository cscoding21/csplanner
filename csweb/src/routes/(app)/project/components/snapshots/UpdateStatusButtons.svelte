<script lang="ts">
    import type { Project, ProjectStatusTransition } from "$lib/graphql/generated/sdk";
	import { setProjectStatus } from "$lib/services/project";
    import { addToast } from "$lib/stores/toasts";
    import { Button } from "flowbite-svelte";
    import { callIf, reloadPage } from "$lib/utils/helpers";

    interface Props {
		project: Project;
		update?: Function;
        displayStates?: string[]|undefined
	}
	let { project, update, displayStates }: Props = $props();

    interface StatusPageContent {
        currentTitle: string;
        currentDesc: string;
        buttonDisplay?: string;
        transitionWarning?: string;
        buttonColor: "red" | "yellow" | "green" | "purple" | "blue" | "light" | "dark" | "primary" | "alternative" | undefined;
    }

    const setStatus = async (status: string) => {
		setProjectStatus(project.id as string, status).then((res) => {
			if (res.status?.success) {
				addToast({
					message: 'Project status successfully updated',
					dismissible: true,
					type: 'success'
				});

				callIf(update);
                reloadPage()
			} else {
				addToast({
					message: 'Error setting project status: ' + res.status?.message,
					dismissible: true,
					type: 'error'
				});
			}
		});
	};

    const getNextState = (project:Project, state:string):ProjectStatusTransition|undefined => {
        if(!project || !project.projectStatusBlock) {
            return undefined
        }

        for (let i = 0; i < (project.projectStatusBlock.allowedNextStates?.length || 0); i++) {
            // @ts-expect-error
            const thisState = project.projectStatusBlock?.allowedNextStates[i]

            if (state === thisState.nextState) {
                return thisState
            }
        }


        return undefined
    }

    const getStatusContent = (type:string):StatusPageContent => {
        let def = {
            currentTitle: "Not set",
            currentDesc: "Not set",
            buttonDisplay: "",
            transitionWarning: "",
            buttonColor: undefined as "red" | "yellow" | "green" | "purple" | "blue" | "light" | "dark" | "primary" | "alternative" | undefined
        }

        switch(type) {
            case("new"):
                def.currentTitle = "New Project";
                def.currentDesc = "The project has been created, but no work has been started"
                break;
            case("draft"):
                def.currentTitle = "Draft";
                def.currentDesc = "When in the draft state, all elements of the project can be edited"
                def.buttonDisplay = "Revert to Draft"
                def.transitionWarning = "Moving a project to draft will clear its approval status"
                def.buttonColor = "blue"
                break;
            case("proposed"):
                def.currentTitle = "Proposed";
                def.currentDesc = "The project has been fully characterized and is under consideration for implementation"
                def.buttonDisplay = "Propose this project"
                def.buttonColor = "green"
                break;
            case("approved"):
                def.currentTitle = "Approved";
                def.currentDesc = "The project has been considered and approved for implementation"
                def.buttonDisplay = "Approve this project"
                def.buttonColor = "green"
                break;
            case("rejected"):
                def.currentTitle = "Rejected";
                def.currentDesc = "The project has been considered but cannot be scheduled in its current state"
                def.buttonDisplay = "Reject this project"
                def.buttonColor = "red"
                break;
            case("backlogged"):
                def.currentTitle = "Backlogged";
                def.currentDesc = "The project has been approved, but is not ready for scheduling likely due to resource constraints"
                def.buttonDisplay = "Move to backlog"
                def.buttonColor = "alternative"
                break;
            case("scheduled"):
                def.currentTitle = "Scheduled";
                def.currentDesc = "This project has been scheduled.  It will automatically be promoted to 'in-flight' when the start date has arrived"
                def.buttonDisplay = "Schedule this project"
                def.buttonColor = "green"
                break;
            case("inflight"):
                def.currentTitle = "In-Flight";
                def.currentDesc = "This project has been scheduled and is currently being implemented.  It will be automatically promoted to 'complete' when all tasks have been marked done"
                break;
            case("complete"):
                def.currentTitle = "Complete";
                def.currentDesc = "This project has been completed"
                break;
            case("deferred"):
                def.currentTitle = "Deferred";
                def.currentDesc = "A project that is deferred is taken off the schedule but can be re-scheduled without future approval"
                def.buttonDisplay = "Defer this project"
                def.buttonColor = "alternative"
                break;
            case("abandoned"):
                def.currentTitle = "Abandoned";
                def.currentDesc = "A project that abandoned is removed from the schedule with no intention of coming back to in the future"
                def.buttonDisplay = "Abandon this project"
                def.buttonColor = "red"
                break;
            case("exception"):
                def.currentTitle = "Exception";
                def.currentDesc = "This should not have happened"
                break;

        }

        return def
    } 
</script>


{#if project}

{#if displayStates && displayStates.length > 0}
    {#each displayStates as state, index}
        {@const next = getNextState(project, state)}
            {#if next && next.canEnter}
                {@const statusContent = getStatusContent(next.nextState)}
                <Button pill color={statusContent.buttonColor} class="w-64 m-2" onclick={() => setStatus(next.nextState)}>{statusContent.buttonDisplay}</Button>
            {:else}
            {JSON.stringify(next)}
                {@const statusContent = getStatusContent(next && next.nextState || "")}
                <Button pill disabled color="alternative" size="xs" class="w-64 m-2">{statusContent.buttonDisplay}</Button>
            {/if}
    {/each}
{:else if project.projectStatusBlock.allowedNextStates}
    {#each project.projectStatusBlock.allowedNextStates as next, index}
        {#if next && next.canEnter}
            {@const statusContent = getStatusContent(next.nextState)}
            <Button pill color={statusContent.buttonColor} class="w-64 m-2" onclick={() => setStatus(next.nextState)}>{statusContent.buttonDisplay}</Button>
        {:else}
            {@const statusContent = getStatusContent(next && next.nextState || "")}
            <Button pill disabled color="alternative" size="xs" class="w-64 m-2">{statusContent.buttonDisplay}</Button>
        {/if}
    {/each}
{/if}
{/if}