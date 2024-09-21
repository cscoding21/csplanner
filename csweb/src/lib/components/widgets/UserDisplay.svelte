<script lang="ts">
    import { Popover } from "flowbite-svelte";
    import { getResource } from "$lib/services/resource";
    import { UserCard}  from "..";
    import { normalizeID } from "$lib/utils/id";
	import type { Resource } from "$lib/graphql/generated/sdk";

    interface UserDisplayProps {
        id: string
    }
    let { 
        id
    }:UserDisplayProps = $props()

    let resource:Resource = $state({} as Resource)

    const loadPage = async () => {
        console.log('loadPage', id)
		getResource(id).then(r => {
            console.log("getResource", r)
            resource = r
            return resource
        })
        .then(r => r)
        .catch(err => {
            console.error(err)
        })
	};

	loadPage();
</script>

{#await loadPage}
 NoID
{:then promiseData} 
<button class="bg-initial" id={normalizeID(id)}>{resource.name}</button>

<Popover class="w-64 text-sm font-light text-left" title={resource.name} triggeredBy="#{normalizeID(id)}" trigger="click">
    <UserCard resource={resource} />
</Popover>
{/await}