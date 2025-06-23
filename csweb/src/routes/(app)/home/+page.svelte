<script lang="ts">
	import { CSSection, OrgStateChecker } from "$lib/components";
	import type { ActivityResults, Organization, PageAndFilter, ProjectResults } from "$lib/graphql/generated/sdk";
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
	import SectionSubHeading from "$lib/components/formatting/SectionSubHeading.svelte";
	import { findActivity } from "$lib/services/activity";
	import { page } from "$app/state";
	import NoResults from "$lib/components/messages/NoResults.svelte";

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

            return o
        }).then(r => {
            findProjects(getFilters())
                .then(res => {
                    myProjects = res
                })
        }).then(r => {
            findActivity(pageAndFilter).then(a => {
                activity = a
            })
        })   
    }

    let pageAndFilter = $state({} as PageAndFilter)
    let org = $state({} as Organization)
    let myProjects = $state({} as ProjectResults)
    let activity = $state({} as ActivityResults)
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
        {#if myProjects && myProjects.results && myProjects.results?.length > 0}
        <div class="grid grid-cols-3 gap-4">
            <div class="col-span-2 p-4">
                <SectionSubHeading>My Projects</SectionSubHeading>
                <ul>
                {#each myProjects.results as project}
                    <li>
                        <a href="/project/detail/{project?.meta?.id}#snapshot">
                        {project?.data?.projectBasics.name}
                        </a>
                    </li>
                {/each}
                </ul>
            </div>

            <div class="p-4">
                <SectionSubHeading>Activity</SectionSubHeading>
                {#if activity && activity.results}
                <ul>
                    {#each activity.results as a}
                        <li>
                            {a?.summary}
                        </li>
                    {/each}
                </ul>
                {:else}
                <div class="text-center p-6">
                    No activity yet...
                </div>
                {/if}
            </div>
        </div>
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