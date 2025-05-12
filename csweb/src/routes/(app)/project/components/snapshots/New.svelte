<script lang="ts">
	import type { Project } from "$lib/graphql/generated/sdk";
	import { TrashBinOutline } from "flowbite-svelte-icons";
	import DeleteProject from "../DeleteProject.svelte";
	import { Button, ButtonGroup } from "flowbite-svelte";
	import { setProjectStatus } from "$lib/services/project";
	import { addToast } from "$lib/stores/toasts";
	import { reloadPage } from "$lib/utils/helpers";

    interface Props {
        project:Project
    }
    let { project }:Props = $props()


    const setStatusToDraft = async () => {
		setProjectStatus(project.id as string, "draft").then((res) => {
			if (res.status?.success) {
				addToast({
					message: 'Project status successfully updated',
					dismissible: true,
					type: 'success'
				});

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
</script>

{#if project}
<div class="grid grid-cols-2 gap-4">
    <div class="pr-6">
        <h3 class="text-gray-100 font-semibold">Executive Summary</h3>
        {@html project.projectBasics.description.replaceAll(/[\n]/g, "<br />")}
    </div>
    <div class="text-center p-8">

        <div>
            To begin working on this project, move it to draft.
            <br /><br />
            <Button onclick={setStatusToDraft}>Move to draft >></Button>

        </div>


    <div class="p-8">
        <ButtonGroup>
            <DeleteProject id={project.id as string} name={project.projectBasics?.name}>
                <TrashBinOutline size="sm" class="mr-2" />
                Delete thie project
            </DeleteProject>
        </ButtonGroup>
    </div>

    </div>
</div>
{/if}