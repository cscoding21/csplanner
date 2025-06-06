<script lang="ts">
	import SectionSubHeading from "$lib/components/formatting/SectionSubHeading.svelte";
	import type { UpdateList, UpdateListItem } from "$lib/graphql/generated/sdk";
	import { getList, updateList } from "$lib/services/list";
	import { addToast } from "$lib/stores/toasts";
	import { callIf } from "$lib/utils/helpers";
	import { nameToID } from "$lib/utils/id";
	import { Alert, Badge, Button, Input } from "flowbite-svelte";
	import { CloseCircleSolid } from "flowbite-svelte-icons";

    interface Props {
        onDone: Function
    }
    let { onDone }:Props = $props()
    let fsList:any[] = $state([
        { name: "Internal", value: "internal"},
        { name: "External", value: "external"},
    ])
    let newFS:string = $state("")

    const listName = "list:funding-source"

    const addFS = () => {
        if (!newFS || !okToAdd(newFS)) {
            return
        }

        const toAdd = { name: newFS, value: nameToID(newFS) }

        fsList = [...new Set([...fsList, toAdd])]
        newFS = ""
    }

    const removeFS = (i:number) => {
        fsList.splice(i, 1)
        fsList = fsList
    }

    const checkEnter = ((event:KeyboardEvent) => {
        if(event.key === "Enter") {
            addFS()
        }
    })

    const okToAdd = (fs:string):boolean => {
        const sk = fsList.filter(s => fs.toLocaleLowerCase() === s.name.toLocaleLowerCase())

        return sk.length === 0
    }

    const getFSList = () => {
        let list:UpdateList = {
            id: listName,
            description: "A list of funding sources for organization projects",
            values: []
        }

        for(let i = 0; i < fsList.length; i++) {
            const sk = fsList[i]
            const listItem:UpdateListItem = {
                name: sk.name,
                value: sk.value,
                sortOrder: 0,
            }

            list.values.push(listItem)
        }

        return list
    }

    const saveList = () => {
		const list = getFSList()

		updateList(list).then(res => {
			if (res && res.status?.success) {
				addToast({
					message: 'Funding sources updated successfully',
					dismissible: true,
					type: 'success'
				});

                callIf(onDone)
			} else {
				addToast({
					message: 'Error updating funding sources: ' + res.status?.message,
					dismissible: true,
					type: 'error'
				});
			}
		});
	}

    const getExistingFSs = async () => {
        getList(listName).then(sk => {
            if(sk.values && sk.values.length > 0) {
                fsList = sk.values.map(s => { return { name: s.name, value: s.value }})
            }
        })
    }

    getExistingFSs()
</script>

<h2 class="text-xl text-center text-gray-50 font-semibold">Organization Funding Sources</h2>
    
<section class="mb-6 mt-4">
<p class="py-2 text-gray-200">
    Funding sources describe how the work on a given project is accounted.      
</p>
<p class="py-2 text-gray-200">Funding sources can be added or modified at any time on the <a class="text-orange-300" href="/settings#lists">List</a> page.  
    We've pre-filled some high-level funding sources below.
</p>
</section>


<SectionSubHeading >Add Your Funding Sources</SectionSubHeading>
<div class="mb-4">
    <Input bind:value={newFS} onclick={addFS} onkeypress={checkEnter} placeholder="Type a funding source name. <enter> to add" />
</div>

<SectionSubHeading>Your Funding Sources</SectionSubHeading>
<div class="p-4">
    {#if fsList.length > 0}
    {#each fsList as fundsrc, index}
    <Badge dismissable large onclose={() => removeFS(index)}>Default</Badge>
        <Badge color="blue" class="m-2" dismissable>
            {fundsrc.name}
        </Badge>
    {/each}
    {:else}
        <Alert>No funding sources yet</Alert>
    {/if}
</div>

<div class="mt-12 text-center">
<Button onclick={saveList}>I'm done with funding sources for now.  Let's keep going! >></Button>
</div>