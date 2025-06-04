<script lang="ts">
    import type { Project } from "$lib/graphql/generated/sdk";
	import { setProjectStatus } from "$lib/services/project";
    import { addToast } from "$lib/stores/toasts";
    import { Button } from "flowbite-svelte";
    import { callIf, reloadPage } from "$lib/utils/helpers";

    interface Props {
		project: Project;
		update?: Function;
	}
	let { project, update }: Props = $props();

    interface StatusPageContent {
        currentTitle: string;
        currentDesc: string;
        buttonDisplay?: string;
        transitionWarning?: string;
        buttonColor: "none" | "red" | "yellow" | "green" | "purple" | "blue" | "light" | "dark" | "primary" | "alternative" | undefined;
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

    const getStatusContent = (type:string):StatusPageContent => {
        let def = {
            currentTitle: "Not set",
            currentDesc: "Not set",
            buttonDisplay: "",
            transitionWarning: "",
            buttonColor: "none" as "none" | "red" | "yellow" | "green" | "purple" | "blue" | "light" | "dark" | "primary" | "alternative" | undefined
        }

        switch(type) {
            case("new"):
                def.currentTitle = "New Project";
                def.currentDesc = "The project has been created, but no work has been started"
                break;
            case("draft"):
                def.currentTitle = "Draft";
                def.currentDesc = "When in the draft state, all elements of the project can be edited"
                def.buttonDisplay = "Move to Draft"
                def.transitionWarning = "Moving a project to draft will clear its approval status"
                def.buttonColor = "primary"
                break;
            case("proposed"):
                def.currentTitle = "Proposed";
                def.currentDesc = "The project has been fully characterized and is under consideration for implementation"
                def.buttonDisplay = "Propose this project"
                def.buttonColor = "primary"
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
                def.buttonColor = "yellow"
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
                def.buttonColor = "yellow"
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

{#if project.projectStatusBlock.allowedNextStates}
{#each project.projectStatusBlock.allowedNextStates as next, index}
        {#if next.canEnter}
            {@const statusContent = getStatusContent(next.nextState)}
            <Button pill color={statusContent.buttonColor} class="w-64 m-2" onclick={() => setStatus(next.nextState)}>{statusContent.buttonDisplay}</Button>
        {:else}
            {@const statusContent = getStatusContent(next.nextState)}
            <!-- <h3 class="font-semibold text-red-400">Cannot move to '{statusContent.currentTitle}'</h3>
            <div class="text-sm">
                {statusContent.currentDesc}
            </div>
            {#if next.checkResult && next.checkResult.messages}
            <ul class="mt-4">
            {#each next.checkResult.messages as msg}
                <li class="text-sm text-red-300">
                    <InfoCircleOutline size="sm" class="float-left mr-2" />
                    {msg?.message}
                    <br class="clear-both" />
                </li>
            {/each}
            </ul>
            {/if} -->
            <Button pill disabled color="alternative" size="xs" class="w-64 m-2">{statusContent.buttonDisplay}</Button>
        {/if}
{/each}
{/if}
{/if}