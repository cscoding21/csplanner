<script lang="ts">
	import type { Project, User } from "$lib/graphql/generated/sdk";
	import { TrashBinOutline } from "flowbite-svelte-icons";
	import DeleteProject from "../DeleteProject.svelte";
	import { Button, ButtonGroup } from "flowbite-svelte";
	import { setProjectStatus } from "$lib/services/project";
	import { addToast } from "$lib/stores/toasts";
	import { reloadPage } from "$lib/utils/helpers";
	import UserCard from "$lib/components/widgets/UserCard.svelte";
	import { SectionSubHeading } from "$lib/components";

    interface Props {
        project:Project
    }
    let { project }:Props = $props()


    const setStatusToApproved = async () => {
		setProjectStatus(project.id as string, "approved").then((res) => {
			if (res.status?.success) {
				addToast({
					message: 'Project approved',
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

    const setStatusToRejected = async () => {
		setProjectStatus(project.id as string, "rejected").then((res) => {
			if (res.status?.success) {
				addToast({
					message: 'Project rejected',
					dismissible: true,
					type: 'success'
				});

                reloadPage()
			} else {
				addToast({
					message: 'Error rejecting project: ' + res.status?.message,
					dismissible: true,
					type: 'error'
				});
			}
		});
	};
</script>

{#if project}

<div class="grid grid-cols-3 gap-16">
    <div>
        <SectionSubHeading>{project.projectBasics.name}</SectionSubHeading>

        <h3 class="text-gray-100 font-semibold">Owner</h3>
        <UserCard user={project.projectBasics.owner as User} />

        <h3 class="text-gray-100 font-semibold mt-6">Executive Summary</h3>
        <p class="py-2 text-gray-200">{@html project.projectBasics.description.replaceAll(/[\n]/g, "<br />")}</p>
    </div>

    <div class="col-span-2">
        <SectionSubHeading>Value Summary</SectionSubHeading>
        <div>

        </div>


        <SectionSubHeading>Cost Summary</SectionSubHeading>
        <div>

        </div>

        <SectionSubHeading>Implementation Details</SectionSubHeading>
        <div>
            
        </div>
    </div>
</div>

{/if}