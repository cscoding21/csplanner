<script lang="ts">
    import { type Portfolio } from "$lib/graphql/generated/sdk";
	import { getPortfolio } from "$lib/services/portfolio";
    import { NoResults } from "$lib/components";
    import { formatDate } from "$lib/utils/format";



const refresh = async (): Promise<Portfolio> => {
		const res = await getPortfolio();

		return res;
	};

let portfolio = $state({} as Portfolio);
	const loadPage = async () => {
		refresh().then((r) => {
			portfolio = r as Portfolio;
		});
	};
</script>


<div class="p-4">


{#await loadPage()}
	<div>Loading...</div>
{:then promiseData}
	{#if portfolio.schedule != null}
		<div class="">
			{#each portfolio.schedule as s(s.projectID)}
				<div><b>{s.project.projectBasics.name}</b>: {formatDate(s.begin)} - {formatDate(s.end)}</div>
			{/each}
		</div>
	{:else}
		<NoResults title="No Resources" newUrl="/resource/new">
			No projects have been scheduled.  
		</NoResults>
	{/if}
{/await}

</div>
