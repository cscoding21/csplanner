<script lang="ts">
	import { CSSection, OrgStateChecker, PageHeading } from "$lib/components";
	import type { Organization, PageAndFilter, ProjectResults } from "$lib/graphql/generated/sdk";
	import { authService } from "$lib/services/auth";
	import { getOrganization } from "$lib/services/organization";
	import { findProjects } from "$lib/services/project";

    const as = authService()
    const user = as.currentUser()

    const getFilters = ():PageAndFilter => {
        let pageAndFilter: PageAndFilter = {
			paging: { pageNumber: 1, resultsPerPage: 10 },
            filters: {
                filters: [
                    { key: 'basics.owner_id', value: user?.email as string, operation: 'eq' },
                    { key: 'status.status', value: "new,draft,approved,scheduled,inflight", operation: 'in' }
                ]
            }
		};

		return pageAndFilter
    }

    const loadPage = async () => {
        getOrganization().then(o => {

        }).then(r => {
            findProjects(getFilters())
                .then(res => {
                    myProjects = res
                })
        })   
    }

    let myProjects = $state({} as ProjectResults)
</script>

<div class="p-4">
<PageHeading title={"Welcome " + user?.firstName} />

<CSSection>
    <OrgStateChecker invert={false} stateToCheck="hasSkills">
        I has skills

        {#snippet elseRender()}
            <p>I have no skills</p>
        {/snippet}
    </OrgStateChecker>


    {#await loadPage()}
        Loading...
    {:then promiseData} 

        {#if myProjects && myProjects.results && myProjects.results?.length > 0}
            <ul>
            {#each myProjects.results as project}
                <li>{project?.projectBasics.name}</li>
            {/each}
        </ul>
        {/if}
    {/await}

</CSSection>

</div>