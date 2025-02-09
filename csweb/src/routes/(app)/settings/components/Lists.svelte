<script lang="ts">
	import { findAllLists } from "$lib/services/list";
    import type { List, ListResults } from "$lib/graphql/generated/sdk";
    import { addToast } from "$lib/stores/toasts";

    let lists = $state([] as List[])

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

	const loadPage = async () => {
        load().then(l => {
            lists = l.results as List[]
        });
	};
</script>


{#await loadPage()}
    Loading...
{:then promiseData} 
    {#each lists as list}
    <h2>{list.name}</h2>

    <ul class="list-disc">
        {#each list.values as val}
            <ul class="ml-4">{val.name} ({val.value}) ||  {val.sortOrder}</ul>
        {/each}
    </ul>
    {/each}
{/await}