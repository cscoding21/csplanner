<script lang="ts">
	import { findAllLists, updateList } from "$lib/services/list";
    import type { List, ListItem, ListResults, UpdateList } from "$lib/graphql/generated/sdk";
    import { addToast } from "$lib/stores/toasts";
	import { Button } from "flowbite-svelte";
	import { CSSidebar, CSSidebarItem, SectionSubHeading } from "$lib/components";
	import { Input, ButtonGroup } from "flowbite-svelte";
	import { PlusOutline } from "flowbite-svelte-icons";
	import { deepCopy } from "$lib/utils/helpers";
	import { nameToID } from "$lib/utils/id";
	import { coalesceToType } from "$lib/forms/helpers";

    let lists = $state([] as List[])
	let currentList = $state({} as UpdateList)
	let currentListName = $state("")
	let newListItem = $state("")
	let errors = $state({newListItem: ""})
	

	const load = async ():Promise<ListResults> => {
		return await findAllLists()
			.then(l => {
				return l
			})
			.catch((err) => {
				addToast({
					message: 'Error loading lists (Lists): ' + err,
					dismissible: true,
					type: 'error'
				});

				return err
			});
	};

	const setList = (index:number) => {
		currentListName = lists[index].name
		currentList = deepCopy(lists[index])
	}

	const saveList = () => {
		// @ts-ignore
		delete currentList.name

		updateList(currentList).then(res => {
			if (res && res.status?.success) {
				addToast({
					message: 'List updated successfully',
					dismissible: true,
					type: 'success'
				});

				load();
			} else {
				addToast({
					message: 'Error updating list: ' + res.status?.message,
					dismissible: true,
					type: 'error'
				});
			}
		});
	}

	const addListItem = (item:string) => {
		let li:ListItem = {
			name: item,
			value: nameToID(item),
			sortOrder: 0
		}

		currentList.values.push(li)
		newListItem = ""
	}

	const loadPage = async () => {
        load().then(l => {
            lists = l.results as List[]

			setList(0)
        });
	};
</script>


{#await loadPage()}
    Loading...
{:then promiseData} 
<div class="flex">
	<div class="flex-none px-2">
		<SectionSubHeading>List</SectionSubHeading>
		<div class="gap-8 lg:flex">
		<CSSidebar>
			{#each lists as list, index}
				<CSSidebarItem onclick={() => setList(index)} label={list.name} active={currentListName === list.name} />
			{/each}
		</CSSidebar>
		</div>
	</div>

	<div class="flex-1 px-2">
		<SectionSubHeading>List Items: {currentListName}</SectionSubHeading>
		<div class="text-sm mb-4">{currentList.description}</div>

		<ul class="mb-2">
			{#each currentList.values as val, index}
				<ul class="mb-4">
					<Input bind:value={currentList.values[index].name} />
					<small>{val.value}</small>
				</ul>
			{/each}
		</ul>
		
		<div class="mt-2 mb-4">
			<ButtonGroup class="w-full">
				<Input placeholder="New list item" bind:value={newListItem} />
				<Button color="primary" class="!p-2.5" type="submit" onclick={() => addListItem(newListItem)}>
					<PlusOutline class="w-5 h-5" />
				</Button>
			</ButtonGroup>
		</div>

		<Button onclick={saveList}>Save List</Button>
	</div>

</div>
{/await}