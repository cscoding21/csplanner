<script lang="ts">
    import type { Project } from "$lib/graphql/generated/sdk";
    import { getDefaultProject } from "$lib/forms/project.validation";
	import { checkProjectStatus, getProject } from "$lib/services/project";
    import { addToast } from "$lib/stores/toasts";
    import { Button, Modal,  } from "flowbite-svelte";

    interface Props {
		id: string;
		update?: Function;
	}
	let { id, update }: Props = $props();

    let setModal = $state(false)

	let project:Project = $state(getDefaultProject() as Project)

    const load = async (): Promise<Project> => {
		return await getProject(id)
			.then(proj => {
				return proj;
			})
			.catch(err => {
				addToast({
					message: 'Error loading project (ProjectStatusUpdate): ' + err,
					dismissible: true,
					type: 'error'
				});

				return err;
			});
	};

    const loadPage = async () => {
        load().then(p => {
            project = p
        }).then(p => {
            //checkProjectStatus(id, )
        });
	};
</script>



{#await loadPage()}
    Loading...
{:then promiseData} 
    {#if project}
    <Button color="dark" onclick={() => (setModal = true)}>
        Status
    </Button>
    
    <Modal bind:open={setModal} size="xs" autoclose>
        <div class="text-center">
            <h3 class="mb-5 text-lg font-normal text-gray-500 dark:text-gray-400">
                Are you sure you want to delete the task, "{name}"?
            </h3>
            <Button color="red" class="mr-2" onclick={() => console.log(id)}>Yes, I'm sure</Button>
            <Button color="alternative">No, cancel</Button>
        </div>
    </Modal>
    {/if}
{/await}
