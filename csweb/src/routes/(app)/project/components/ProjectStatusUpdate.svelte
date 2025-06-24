<script lang="ts">
    import type { Snippet } from "svelte";
    import type { ProjectEnvelope } from "$lib/graphql/generated/sdk";
	import { setProjectStatus, getProject, ProjectStatusException, ProjectStatusAbandoned, ProjectStatusApproved, ProjectStatusBacklogged, ProjectStatusComplete, ProjectStatusDeferred, ProjectStatusDraft, ProjectStatusInflight, ProjectStatusNew, ProjectStatusProposed, ProjectStatusRejected, ProjectStatusScheduled } from "$lib/services/project";
    import { addToast } from "$lib/stores/toasts";
    import { Button, Modal, Hr } from "flowbite-svelte";
    import { callIf, reloadPage } from "$lib/utils/helpers";
	import { getDefaultProject } from "$lib/forms/project.validation";
	import { CSHR, SectionHeading } from "$lib/components";
	import { ChevronDoubleRightOutline, InfoCircleOutline } from "flowbite-svelte-icons";

    interface Props {
		id: string;
		update?: Function;
        children: Snippet;
	}
	let { id, update, children }: Props = $props();
    let popupModal = $state(false)

    interface StatusPageContent {
        currentTitle: string;
        currentDesc: string;
        buttonDisplay?: string;
        transitionWarning?: string;
        buttonColor: "red" | "yellow" | "green" | "purple" | "blue" | "light" | "dark" | "primary" | "alternative" | undefined;
    }


    let project:ProjectEnvelope = $state(getDefaultProject() as ProjectEnvelope)

    const load = async (): Promise<ProjectEnvelope> => {
		return await getProject(id)
			.then((proj) => {
				return proj;
			})
			.catch((err) => {
				addToast({
					message: 'Error loading project (ProjectStatusUpdate): ' + err,
					dismissible: true,
					type: 'error'
				});

				return err;
			});
	};

    const setStatus = async (status: string) => {
		setProjectStatus(id, status).then((res) => {
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
            buttonColor: undefined as "red" | "yellow" | "green" | "purple" | "blue" | "light" | "dark" | "primary" | "alternative" | undefined
        }

        switch(type) {
            case(ProjectStatusNew):
                def.currentTitle = "New Project";
                def.currentDesc = "The project has been created, but no work has been started"
                break;
            case(ProjectStatusDraft):
                def.currentTitle = "Draft";
                def.currentDesc = "When in the draft state, all elements of the project can be edited"
                def.buttonDisplay = "Move to Draft"
                def.transitionWarning = "Moving a project to draft will clear its approval status"
                def.buttonColor = "primary"
                break;
            case(ProjectStatusProposed):
                def.currentTitle = "Proposed";
                def.currentDesc = "The project has been fully characterized and is under consideration for implementation"
                def.buttonDisplay = "Propose this project"
                def.buttonColor = "primary"
                break;
            case(ProjectStatusApproved):
                def.currentTitle = "Approved";
                def.currentDesc = "The project has been considered and approved for implementation"
                def.buttonDisplay = "Approve this project"
                def.buttonColor = "green"
                break;
            case(ProjectStatusRejected):
                def.currentTitle = "Rejected";
                def.currentDesc = "The project has been considered but cannot be scheduled in its current state"
                def.buttonDisplay = "Reject this project"
                def.buttonColor = "red"
                break;
            case(ProjectStatusBacklogged):
                def.currentTitle = "Backlogged";
                def.currentDesc = "The project has been approved, but is not ready for scheduling likely due to resource constraints"
                def.buttonDisplay = "Move to backlog"
                def.buttonColor = "yellow"
                break;
            case(ProjectStatusScheduled):
                def.currentTitle = "Scheduled";
                def.currentDesc = "This project has been scheduled.  It will automatically be promoted to 'in-flight' when the start date has arrived"
                def.buttonDisplay = "Schedule this project"
                def.buttonColor = "green"
                break;
            case(ProjectStatusInflight):
                def.currentTitle = "In-Flight";
                def.currentDesc = "This project has been scheduled and is currently being implemented.  It will be automatically promoted to 'complete' when all tasks have been marked done"
                break;
            case(ProjectStatusComplete):
                def.currentTitle = "Complete";
                def.currentDesc = "This project has been completed"
                break;
            case(ProjectStatusDeferred):
                def.currentTitle = "Deferred";
                def.currentDesc = "A project that is deferred is taken off the schedule but can be re-scheduled without future approval"
                def.buttonDisplay = "Defer this project"
                def.buttonColor = "yellow"
                break;
            case(ProjectStatusAbandoned):
                def.currentTitle = "Abandoned";
                def.currentDesc = "A project that abandoned is removed from the schedule with no intention of coming back to in the future"
                def.buttonDisplay = "Abandon this project"
                def.buttonColor = "red"
                break;
            case(ProjectStatusException):
                def.currentTitle = "Exception";
                def.currentDesc = "This should not have happened"
                break;

        }

        return def
    } 

    let currentStateDef:StatusPageContent = $state(getStatusContent(""))

    const loadPage = async () => {
		load().then(p => {
            project = p

            currentStateDef = getStatusContent(p.data?.projectStatusBlock?.status)
        });
	};
</script>

{#await loadPage()}
    Loading...
{:then promiseData} 

{#if project}

<Button onclick={() => (popupModal = true)}>{@render children()}</Button>

<Modal bind:open={popupModal} size="md" autoclose>
    <div class="">
        <SectionHeading>Current Status: {currentStateDef.currentTitle}</SectionHeading>
        <p>{currentStateDef.currentDesc}</p>
        

        {#if project.data?.projectStatusBlock.allowedNextStates}
        {#each project.data?.projectStatusBlock.allowedNextStates as next, index}
            <div class="text-left p-4">  
                {#if next.canEnter}
                    {@const statusContent = getStatusContent(next.nextState)}
                    <h3 class="font-semibold">Move to '{statusContent.currentTitle}'</h3>
                    {#if statusContent.transitionWarning}
                        <div class="text-yellow-300 text-sm">
                            <InfoCircleOutline size="sm" class="float-left mr-2" />
                            {statusContent.transitionWarning}
                            <br class="clear-both" />
                        </div>
                    {/if}
                    <div class="text-sm">
                        {statusContent.currentDesc}
                    </div>
                    <div class="text-right">
                        <Button pill color={statusContent.buttonColor} size="xs" class="w-48 mt-4" onclick={() => setStatus(next.nextState)}>{statusContent.buttonDisplay} <ChevronDoubleRightOutline size="xs"  class="w-5 h-5 ms-2"  /></Button>
                    </div>
                {:else}
                    {@const statusContent = getStatusContent(next.nextState)}
                    <h3 class="font-semibold text-red-400">Cannot move to '{statusContent.currentTitle}'</h3>
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
                    {/if}
                    <div class="text-right">
                        <Button pill disabled color="alternative" size="xs" class="w-48 mt-4" onclick={() => alert("no can move")}>{statusContent.buttonDisplay}</Button>
                    </div>
                {/if}
            </div>
            {#if index + 1 < project.data?.projectStatusBlock.allowedNextStates.length}
            <CSHR />
            {/if}
        {/each}
        {/if}
    </div>
</Modal>
{/if}

    
{/await}