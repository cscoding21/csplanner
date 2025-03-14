<script lang="ts">
	import { CSSection, OrgStateChecker, PageHeading } from "$lib/components";
	import type { Organization, PageAndFilter, ProjectResults } from "$lib/graphql/generated/sdk";
	import { authService } from "$lib/services/auth";
	import { getOrganization } from "$lib/services/organization";
	import { findProjects } from "$lib/services/project";
	import SetupWizard from "./setup/SetupWizard.svelte";

    const as = authService()
    const user = as.currentUser()

    const getFilters = ():PageAndFilter => {
        let pageAndFilter: PageAndFilter = {
			paging: { pageNumber: 1, resultsPerPage: 10 },
            filters: {
                filters: [
                    { key: 'data.basics.owner_id', value: user?.email as string, operation: 'eq' },
                    { key: 'data.status.status', value: "new,draft,approved,scheduled,inflight", operation: 'in' }
                ]
            }
		};

		return pageAndFilter
    }

    const loadPage = async () => {
        getOrganization().then(o => {
            org = o

            return o
        }).then(r => {
            findProjects(getFilters())
                .then(res => {
                    myProjects = res
                })
        })   
    }

    let org = $state({} as Organization)
    let myProjects = $state({} as ProjectResults)
</script>

<div class="p-4">
<PageHeading title={"Welcome " + user?.firstName} />

<CSSection>
    


    {#await loadPage()}
        Loading...
    {:then promiseData} 
    <SetupWizard {org} />

    <OrgStateChecker invert={false} stateToCheck="isReadyForProjects">
        {#if myProjects && myProjects.results && myProjects.results?.length > 0}
            <ul>
            {#each myProjects.results as project}
                <li>{project?.data?.projectBasics.name} {project?.meta?.id}</li>
            {/each}
            </ul>
        {/if}

        {#snippet elseRender()}
            <p>I have no projects</p>
        {/snippet}
    </OrgStateChecker>

        
    {/await}

</CSSection>

</div>