<script lang="ts">
    import { Button , Hr, Modal } from "flowbite-svelte";
    import { NoResults, SectionHeading } from "$lib/components";
    import { EditOutline, TrashBinOutline } from "flowbite-svelte-icons";
    import { getDefaultProject } from "$lib/forms/project.validation";
	import { ProjectFeatureDisplay, DeleteProjectFeature, ProjectFeatureForm} from ".";
    import { getProject } from "$lib/services/project";
    import { addToast } from "$lib/stores/toasts";

    interface Props {
        id: string;
        update: Function;
    }

    let { id, update }:Props = $props()

    function refresh() {
        //---close all modals
        modalState = []

        load()
        update()
    }

    let modalState:boolean[] = $state([])

	const load = async () => {
		await getProject(id).then(proj => {
			project = proj
		}).catch(err => {
            addToast({ 
				message: "Error loading project (ProjectFeature): " + err, 
				dismissible: true, 
				type: "error"}
			)
		})
	}

    // @ts-ignore
	let project:Project = $state(getDefaultProject())

    const loadPage = async () => {
		load().then(r => r)
	};

    loadPage()
</script>


<SectionHeading>Features</SectionHeading>

<div class="mb-8">
{#if project.projectFeatures && project.projectFeatures.length > 0}

    {#each project.projectFeatures as feature(feature.id)}
        {#if feature && feature.id }
        <ProjectFeatureDisplay feature={feature} />
        <div class="mt-2 mb-4">
            <Button size="xs" color="dark" on:click={() => { modalState[feature.id] = true } }>
                <EditOutline size="xs" class="mr-2" />
                Edit
            </Button>
            <DeleteProjectFeature
                id={feature.id || ""}
                name={feature.name}
                projectID={id} 
                size="xs"
                update={refresh}>
                    <TrashBinOutline size="xs" class="mr-2" />
                    Delete
            </DeleteProjectFeature>
        </div>

        <Modal bind:open={modalState[feature.id]} size="xl">
            <ProjectFeatureForm
                projectID={id} 
                feature={feature}
                update={refresh} />
        </Modal>
        <Hr hrClass="h-px mt-2 mb-6 bg-gray-200 border-0 dark:bg-gray-700" />
        {/if}
    {/each}

{:else}
    <NoResults title="No Project Features" newUrl="">
        No project features...create one now.
    </NoResults>
{/if}
</div>

<ProjectFeatureForm
    projectID={id} 
    feature={undefined}
    update={refresh} />