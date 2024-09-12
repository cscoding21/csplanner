<script lang="ts">   
    import { Avatar, Heading } from "flowbite-svelte";
    import { formatToCommaSepList, getInitialsFromName } from "$lib/utils/format"

    import type { Resource } from '$lib/graphql/generated/sdk'
    
    interface ResourceCardProps {
        resource: Resource;
    };
    let { resource = $bindable() }:ResourceCardProps = $props()
</script>


<div class="items-center bg-gray-50 rounded-lg shadow sm:flex dark:bg-gray-800 dark:border-gray-700">
    <div class="mx-4 my-2 justify-start">
    <a href="/resource/detail/{resource.id}">
        <Avatar size="lg" src={resource.profileImage || ""}  rounded >{getInitialsFromName(resource.name)}</Avatar>
    </a>
    </div>
    <div class="p-5">
        <Heading tag="h6">
            <a href="/resource/detail/{resource.id}">{resource.name}</a>
        </Heading>
        <span class="text-gray-500 dark:text-gray-400">{resource.role}</span>
        <p class="mb-3 font-light text-gray-500 dark:text-gray-400 text-xs">{formatToCommaSepList(resource.skills?.map(l => l.name) as string[])}</p>
    </div>
</div>