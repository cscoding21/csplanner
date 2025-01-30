<script lang="ts">
    import type { Snippet } from "svelte";
    import { type Project } from "$lib/graphql/generated/sdk";
	import { setProjectStatus, getProject } from "$lib/services/project";
    import { addToast } from "$lib/stores/toasts";
    import { Button, Modal, Hr } from "flowbite-svelte";
    import { callIf } from "$lib/utils/helpers";
	import { getDefaultProject } from "$lib/forms/project.validation";
	import { SectionHeading } from "$lib/components";
	import { ChevronDoubleRightOutline, InfoCircleOutline } from "flowbite-svelte-icons";

    interface Props {
		id: string;
		update?: Function;
        children: Snippet;
	}
	let { id, update, children }: Props = $props();
    let popupModal = $state(false)
    let content = $state({} as StatusPageContent)

    interface StatusPageContent {
        currentTitle: string;
        currentDesc: string;
        buttonDisplay?: string;
        transitionWarning?: string;
        buttonColor: "none" | "red" | "yellow" | "green" | "purple" | "blue" | "light" | "dark" | "primary" | "alternative" | undefined;
    }


    let project:Project = $state(getDefaultProject() as Project)

    const load = async (): Promise<Project> => {
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
            currentTitle: "Not et",
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
                def.currentDesc = "This is a new project"
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

    let currentStateDef:StatusPageContent = $state(getStatusContent(""))

    const loadPage = async () => {
		load().then(p => {
            project = p

            currentStateDef = getStatusContent(p.projectStatusBlock?.status)
        });
	};
</script>

{#await loadPage()}
    Loading...
{:then promiseData} 

{#if project}

<Button on:click={() => (popupModal = true)}>{@render children()}</Button>

<Modal bind:open={popupModal} size="md" autoclose>
    <div class="">
        <SectionHeading>Current Status: {currentStateDef.currentTitle}</SectionHeading>
        <p>{currentStateDef.currentDesc}</p>
        

        {#if project.projectStatusBlock.allowedNextStates}
        {#each project.projectStatusBlock.allowedNextStates as next, index}
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
            {#if index + 1 < project.projectStatusBlock.allowedNextStates.length}
            <Hr hrClass="h-px mt-2 mb-2 bg-gray-200 border-0 dark:bg-gray-700" />
            {/if}
        {/each}
        {/if}
    </div>
</Modal>
{/if}

    
{/await}