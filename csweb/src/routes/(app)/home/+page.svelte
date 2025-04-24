<script lang="ts">
	import { CSSection, OrgStateChecker, PageHeading } from "$lib/components";
	import type { Organization, PageAndFilter, ProjectResults } from "$lib/graphql/generated/sdk";
	import { authService } from "$lib/services/auth";
	import { getOrganization } from "$lib/services/organization";
	import { findProjects } from "$lib/services/project";
	import { Button, ButtonGroup } from "flowbite-svelte";
	import SetupWizard from "./setup/SetupWizard.svelte";
	import Welcome from "./setup/Welcome.svelte";
	import { ArrowLeftOutline, ArrowRightOutline } from "flowbite-svelte-icons";
	import OrgSettings from "./setup/OrgSettings.svelte";
	import Skills from "./setup/Skills.svelte";

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
    let currentStep = $state(1)

    const next = () => {
        ++currentStep
    }
</script>

<div class="p-4">
<PageHeading title={"Welcome " + user?.firstName} />

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
            <h1>Setup wizard</h1>
            <SetupWizard {org} bind:step={currentStep} />

            {#if currentStep == 1}
                <Welcome onDone={next} />
            {:else if currentStep == 2}
                <OrgSettings onDone={next} />
            {:else if currentStep == 3}
                <Skills onDone={next} />
            {:else if currentStep == 4}
                <b>4</b>
            {:else if currentStep == 5}
                <b>5</b>
            {:else if currentStep == 6}
                <b>6</b>
            {:else if currentStep == 7}
                <b>7</b>
            {/if}
        
        
            
            <div class="mt-4">
                <ButtonGroup>
                    <Button onclick={() => { --currentStep }} disabled={currentStep == 1}><ArrowLeftOutline /> Previous</Button>
                    <Button onclick={() => { ++currentStep }} disabled={currentStep == 7}>Next <ArrowRightOutline /></Button>
                </ButtonGroup>
            </div>
        {/snippet}
    </OrgStateChecker>

    {/await}
</CSSection>

</div>