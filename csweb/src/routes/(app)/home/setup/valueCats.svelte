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
    let vcList:any[] = $state([
            {value: "revenue", name: "Revenue increase"},
            {value: "tax-benefit", name: "Tax write-off"},
            {value: "risk-mitigation", name: "Risk mitigation"},
            {value: "cost-reduction", name: "Cost reduction"}
    ])
    let newVC:string = $state("")

    const listName = "list:value-cats"

    const addVC = () => {
        if (!newVC || !okToAdd(newVC)) {
            return
        }

        const toAdd = { name: newVC, value: nameToID(newVC) }

        vcList = [...new Set([...vcList, toAdd])]
        newVC = ""
    }

    const removeVC = (i:number) => {
        vcList.splice(i, 1)
        vcList = vcList
    }

    const checkEnter = ((event:KeyboardEvent) => {
        if(event.key === "Enter") {
            addVC()
        }
    })

    const okToAdd = (vc:string):boolean => {
        const sk = vcList.filter(s => vc.toLocaleLowerCase() === s.name.toLocaleLowerCase())

        return sk.length === 0
    }

    const getVCList = () => {
        let list:UpdateList = {
            id: listName,
            description: "A list of value categories used by the organization",
            values: []
        }

        for(let i = 0; i < vcList.length; i++) {
            const sk = vcList[i]
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
		const list = getVCList()

		updateList(list).then(res => {
			if (res && res.status?.success) {
				addToast({
					message: 'Value categories updated successfully',
					dismissible: true,
					type: 'success'
				});

                callIf(onDone)
			} else {
				addToast({
					message: 'Error updating value categories: ' + res.status?.message,
					dismissible: true,
					type: 'error'
				});
			}
		});
	}

    const getExistingVCs = async () => {
        getList(listName).then(sk => {
            if(sk.values && sk.values.length > 0) {
                vcList = sk.values.map(s => { return { name: s.name, value: s.value }})
            }
        })
    }

    getExistingVCs()
</script>

<h2 class="text-xl text-center text-gray-50 font-semibold">Organization Value Categories</h2>
    
<section class="mb-6 mt-4">
<p class="py-2 text-gray-200">
    Value categories describe how an organizatioin benefits from the work that it does.      
</p>
<p class="py-2 text-gray-200">Value categories can be added or modified at any time on the <a class="text-orange-300" href="/settings#lists">List</a> page.  
    We've pre-filled some high-level value categories below.
</p>
</section>


<SectionSubHeading >Add Your Value Categories</SectionSubHeading>
<div class="mb-4">
    <Input bind:value={newVC} onclick={addVC} onkeypress={checkEnter} placeholder="Type a value category name. <enter> to add" />
</div>

<SectionSubHeading>Your Value Categories</SectionSubHeading>
<div class="p-4">
    {#if vcList.length > 0}
    {#each vcList as valcat, index}
        <Badge color="green" class="m-2" dismissable onclose={() => removeVC(index)}>
            {valcat.name}
        </Badge>
    {/each}
    {:else}
        <Alert>No value categories yet</Alert>
    {/if}
</div>

<div class="mt-12 text-center">
<Button onclick={saveList}>I'm done with value categories for now.  Let's keep going! >></Button>
</div>