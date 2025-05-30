<script lang="ts">
	import { SectionSubHeading, UserCard } from "$lib/components";
	import type { Project, User } from "$lib/graphql/generated/sdk";

    interface Props {
        project:Project
    }
    let { project }:Props = $props()


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

        All your error are belong to us!
        
    </div>
</div>

{/if}