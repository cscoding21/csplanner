<script lang="ts">
    import { type Portfolio } from "$lib/graphql/generated/sdk";
	import { getPortfolio } from "$lib/services/portfolio";
    import { NoResults, CSSection } from "$lib/components";
    import { formatDate } from "$lib/utils/format";
	import { getCSWeeks } from "$lib/utils/cscal";

	let weeks:Date[] = $state([])

	const refresh = async (): Promise<Portfolio> => {
		const res = await getPortfolio();

		weeks = getCSWeeks(new Date(res.begin), new Date(res.end))

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
<CSSection>

{#await loadPage()}
	<div>Loading...</div>
{:then promiseData}
	<h1>{formatDate(portfolio.begin)} - {formatDate(portfolio.end)}</h1>
	<hr />
	<div class="p-4">
	<ul>
		{#each weeks as week}
			<li>{formatDate(week)} - {week.toLocaleDateString('en-US', { weekday: 'long' })}</li>
		{/each}
	</ul>
</div>
	<hr />
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

</CSSection>
</div>
