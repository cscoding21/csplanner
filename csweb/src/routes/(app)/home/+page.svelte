<script lang="ts">
	import { CSSection, OrgStateChecker } from "$lib/components";
	import type { Organization, PageAndFilter, ProjectResults } from "$lib/graphql/generated/sdk";
	import { authService } from "$lib/services/auth";
	import { getOrganization } from "$lib/services/organization";
	import { findProjects } from "$lib/services/project";
	import SetupWizard from "./setup/SetupWizard.svelte";
	import Welcome from "./setup/Welcome.svelte";
	import OrgSettings from "./setup/OrgSettings.svelte";
	import Skills from "./setup/Skills.svelte";
	import ValueCats from "./setup/valueCats.svelte";
	import FundingSources from "./setup/FundingSources.svelte";
	import Roles from "./setup/Roles.svelte";
	import Resources from "./setup/Resources.svelte";

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

            currentStep = getStartingStep(org)

            console.log(org)

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
    let currentStep = $state(0)

    const getStartingStep = (o:Organization):number => {
        let out = 1;

        if(o.setup.hasResources) {
            out = 7
        } else if (o.setup.hasRoles) {
            out = 6
        } else if (o.setup.hasFundingSources) {
            out = 5
        } else if (o.setup.hasValueCategories) {
            out = 4
        } else if (o.setup.hasSkills) {
            out = 3
        } 

        return out
    }

    const next = () => {
        ++currentStep
    }
</script>

<div class="p-4">
<CSSection>
    {#await loadPage()}
        Loading...
    {:then promiseData} 

    <OrgStateChecker invert={false} stateToCheck="isReadyForProjects">
        <h1>My Projects</h1>
        {#if myProjects && myProjects.results && myProjects.results?.length > 0}
            <ul>
            {#each myProjects.results as project}
                <li>{project?.data?.projectBasics.name} {project?.meta?.id}</li>
            {/each}
            </ul>
        {/if}

        {#snippet elseRender()}
            <SetupWizard {org} bind:step={currentStep} />

            <div class="p-4 grid grid-cols-4 gap-4">
                <div></div>
                <div class="col-span-2 p-6">
            {#if currentStep == 1}
                <Welcome onDone={next} {user} />
            {:else if currentStep == 2}
                <OrgSettings onDone={next} />
            {:else if currentStep == 3}
                <Skills onDone={next} />
            {:else if currentStep == 4}
                <ValueCats onDone={next} />
            {:else if currentStep == 5}
                <FundingSources onDone={next} />
            {:else if currentStep == 6}
                <Roles onDone={next} />
            {:else if currentStep >= 7}
                <Resources onDone={next} />
            {/if}
            <div></div>
                </div>
            </div>
        {/snippet}
    </OrgStateChecker>

    {/await}
</CSSection>

</div>