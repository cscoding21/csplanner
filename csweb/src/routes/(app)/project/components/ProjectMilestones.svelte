<script lang="ts">
    import type { ProjectMilestone , Skill } from "$lib/graphql/generated/sdk";
    import { P, Button, Hr, Modal, A , type SelectOptionType } from "flowbite-svelte";
    import { ChevronDoubleRightOutline, ChevronRightOutline, EditOutline, TrashBinOutline } from "flowbite-svelte-icons";
    import { SectionHeading, UserList } from "$lib/components";
    import { SelectInput } from "$lib/components";
    import { getTemplateFromStore } from "$lib/stores/project";
    import { findAllProjectTemplates, setProjectMilestonesFromTemplate } from "$lib/services/project";
    import { is } from "$lib/utils/check";
    import { getDefaultProject } from "$lib/forms/project.validation";
    import { ProjectTaskForm, DeleteProjectTask } from ".";
    import { formatToCommaSepList } from "$lib/utils/format";
    import { getProject } from "$lib/services/project";
    import { addToast } from "$lib/stores/toasts";

    function refresh() {
        load(id, currentMilestone.id)

        modalState = []
    }

    interface ProjectMilestonesProps {
        id: string;
        update?: Function;
    }
    let { 
        id, 
        update 
    }:ProjectMilestonesProps = $props()

    let errors = $state({})
    let modalState:boolean[] = $state([]);

	const load = async (id:string, milestoneID:string) => {
		await getProject(id).then(proj => {
			// @ts-ignore
			project = proj

            setCurrentMilestone(milestoneID)
		}).catch(err => {
			console.log('Error on get ', err)
            addToast({ 
				message: "Error loading project (ProjectMilestones): " + err, 
				dismissible: true, 
				type: "error"}
			)
		})
	}

    //---Legacy
    let template = $state("");


    const showTemplateDetails = async () => {
        getTemplateFromStore(template).then(tem => {
            templateDetails = tem
            console.log(tem)
        })
    }

    const selectTemplate = async () => {
        console.log("setting template " + template + " to project " + id)

        setProjectMilestonesFromTemplate({ projectId: id, templateId: template}).then(td => {
            templateDetails = td
        })

        load(id, "")
    }


    const getRequiredSkillsDisplay = (task:any):string => {
        if (!is(task.skills)) {
            return "Not set";
        }

        return formatToCommaSepList(task.skills.map((s:Skill) => s.name))
    }

    const setCurrentMilestone = (milestoneID:any):ProjectMilestone => {
        if (project && project.projectMilestones && project.projectMilestones.length > 0) {
            return currentMilestone;
        }

        if (!milestoneID) {
            currentMilestone = project.projectMilestones[0] as ProjectMilestone;
        }

        for (let i = 0; i < project.projectMilestones.length; i++) {
            const m = project.projectMilestones[i];

            if (m.id == milestoneID) {
                currentMilestone = m;
                break;
            }
        }

        return currentMilestone;
    }

    const loadPage = async () => {
		findAllProjectTemplates()
			.then((r) => r)
			.then((r) => {
				templateOpts = r.results?.map((r) => ({
					name: r.name,
					value: r.id as string
				})) as SelectOptionType<string>[];
			})
			.then(() => {
				load(id, "")
			});
	};

    let project = $state(getDefaultProject());
    let templateDetails = $state()
    let templateOpts =  $state([] as SelectOptionType<string>[]);
    let currentMilestone = $state(setCurrentMilestone(""))
</script>


<SectionHeading>Implementation Plan & Milestones</SectionHeading>

{#await loadPage}
	Loading...
{:then promiseData} 
    {#if project.projectMilestones && project.projectMilestones.length > 0}
    <div class="grid grid-cols-4 gap-8">
        <div class="col-span-1">
            <ol class="relative text-gray-500 border-s border-gray-200 dark:border-gray-700 dark:text-gray-400">      
                {#each project.projectMilestones as milestone(milestone)}           
                <li class="mb-10 ms-6">   
                    {#if currentMilestone.id === milestone?.id}          
                        <span class="absolute flex items-center justify-center w-8 h-8 bg-green-200 rounded-full -start-4 ring-4 ring-white dark:ring-gray-900 dark:bg-green-700 dark:text-gray-50">
                            <ChevronDoubleRightOutline size="xs" />
                        </span>
                    {:else}
                    <span class="absolute flex items-center justify-center w-8 h-8 bg-gray-100 rounded-full -start-4 ring-4 ring-white dark:ring-gray-900 dark:bg-gray-700">
                            <ChevronRightOutline size="xs" />
                        </span>
                    {/if}
                    <button onclick={(e) => setCurrentMilestone(milestone?.id)}>
                        <h3 class="font-medium leading-tight">{milestone?.phase.name}</h3>
                    </button>
                    <p class="text-xs">{milestone?.phase.description}</p>
                </li>
                {/each}
            </ol>
        </div>
        <div class="col-span-3">
            {#if currentMilestone.tasks}
            {#each currentMilestone.tasks as task(task)}
                <div>
                    <span class="text-gray-100">{task.name}</span>
                    <P class="mb-3" weight="light" color="text-gray-500 dark:text-gray-100 float-right">{task.hourEstimate} hours</P>         
                </div>
                <div>
                    <div class="text-gray-400 text-sm">{task.description}</div>
                    <div class="text-gray-100 my-2 text-sm">Skills Required: {getRequiredSkillsDisplay(task)}</div>
                    
                    <div class="pl-4"><UserList maxSize={4} size="xs" resourceIDs={task.resourceIDs || []} /></div>
                </div>
                <div class="mt-2">
                    <Button size="xs" color="dark" on:click={() => { modalState[task.id] = true } }>
                        <EditOutline size="xs" class="mr-2" />
                        Edit
                    </Button>
                    <DeleteProjectTask 
                        name={task.name}
                        size="xs"
                        projectID={id} 
                        milestoneID={currentMilestone.id} 
                        on:updated={refresh}
                        id={task.id}>
                        <TrashBinOutline size="xs" class="mr-2" />
                        Delete
                    </DeleteProjectTask>
                </div>

                <Modal bind:open={modalState[task.id]} size="xl" autoclose={false}>
                    <ProjectTaskForm 
                        task={task} 
                        milestoneID={currentMilestone.id} 
                        projectID={id} 
                        on:updated={refresh} /> 
                </Modal>
                <br class="clear-both" />
            {/each}
            {/if}
            <Hr />
            <ProjectTaskForm 
                milestoneID={currentMilestone.id}
                projectID={id} 
                on:updated={refresh} /> 
        </div>
    </div>
    
{:else}
    <SelectInput 
        fieldName="Project Milestone Template"
        on:updated={showTemplateDetails}
        bind:value={template} 
        error={errors.template}
        options={templateOpts} />


    {#if templateDetails !== undefined}
        <SectionHeading>Milestones for {templateDetails.name} Template</SectionHeading>

        {#if templateDetails.phases && templateDetails.phases.length > 0}
        {#each templateDetails.phases as phase(phase)}
            <SectionHeading>{phase.name}</SectionHeading>
            <P class="mb-3" weight="light" color="text-gray-500 dark:text-gray-400">{phase.description}</P>
        {/each}
        {/if}

        <div class="mt-8">
            <span class="">
                <Button on:click={selectTemplate}> Use This Project Lifecycle </Button>
            </span>
        </div>
    {/if}
{/if}
{/await}