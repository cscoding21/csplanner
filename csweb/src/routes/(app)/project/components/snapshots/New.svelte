<script lang="ts">
	import type { Project, User } from "$lib/graphql/generated/sdk";
	import { TrashBinOutline } from "flowbite-svelte-icons";
	import DeleteProject from "../DeleteProject.svelte";
	import { Button, ButtonGroup } from "flowbite-svelte";
	import { setProjectStatus } from "$lib/services/project";
	import { addToast } from "$lib/stores/toasts";
	import { reloadPage } from "$lib/utils/helpers";
	import { ProjectBasics } from "..";
	import UserCard from "$lib/components/widgets/UserCard.svelte";

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

<div class="p-4 grid grid-cols-4 gap-4">
<div></div>
<div class="col-span-2 p-6">

<h2 class="text-xl text-left text-gray-50 font-semibold mb-4">{project.projectBasics.name}</h2>

<h3 class="text-gray-100 font-semibold">Owner</h3>
<UserCard user={project.projectBasics.owner as User} />

<h3 class="text-gray-100 font-semibold mt-6">Executive Summary</h3>
<p class="py-2 text-gray-200">{@html project.projectBasics.description.replaceAll(/[\n]/g, "<br />")}</p>

<div class="mt-6 text-center">
    <div>
        Move this project to "draft" to begin fleshing out details, costs, and benefit.
    </div>

    <Button class="mt-4" onclick={setStatusToDraft}>Move to draft</Button>

    <div class="p-8 mt-2">
        <div class="mb-4 text-sm">Is this a duplicate?  If so, you can delete it now.</div>
        <ButtonGroup>
            <DeleteProject id={project.id as string} name={project.projectBasics?.name}>
                <TrashBinOutline color="red" size="xs" class="mr-2" />
                Delete thie project
            </DeleteProject>
        </ButtonGroup>
    </div>
</div>

<div></div>
</div>
</div>

{/if}