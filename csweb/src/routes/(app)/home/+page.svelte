<script lang="ts">
	import { CSHR, CSSection, OrgStateChecker } from "$lib/components";
	import type { Activity, ActivityResults, Organization, PageAndFilter, Portfolio, ProjectResults } from "$lib/graphql/generated/sdk";
	import { authService } from "$lib/services/auth";
	import { getOrganization } from "$lib/services/organization";
	import { findProjects, ProjectStatusApproved, ProjectStatusDraft, ProjectStatusInflight, ProjectStatusNew, ProjectStatusProposed, ProjectStatusScheduled } from "$lib/services/project";
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
	import ProjectActionTable from "../project/components/ProjectActionTable.svelte";
	import { getPortfolio } from "$lib/services/portfolio";
	import { Avatar } from "flowbite-svelte";
	import { getInitialsFromName } from "$lib/utils/format";

    const as = authService()
    const user = as.currentUser()

    const statusFilters = [ProjectStatusNew, ProjectStatusDraft, ProjectStatusApproved, ProjectStatusScheduled, ProjectStatusInflight, ProjectStatusProposed]

    const getFilters = ():PageAndFilter => {
        let pageAndFilter: PageAndFilter = {
			paging: { pageNumber: 1, resultsPerPage: 10 },
            filters: {
                filters: [
                    { key: 'data.basics.owner_id', value: user?.email as string, operation: 'eq' },
                    { key: 'data.status.status', value: statusFilters.join(","), operation: 'in' }
                ]
            }
		};

		return pageAndFilter
    }

    const getActivityFilters = ():PageAndFilter => {
        let pf: PageAndFilter = {
			paging: { pageNumber: 1, resultsPerPage: 10 },
            filters: {
                filters: []
            }
		};

		return pf
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
            getPortfolio()
                .then(por => {
                    portfolio = por
                })
        }).then(r => {
            findActivity(getActivityFilters()).then(a => {
                activity = a
            })
        })   
    }

    let pageAndFilter = $state({} as PageAndFilter)
    let activityPageAndFilter = $state({} as PageAndFilter)
    let org = $state({} as Organization)
    let myProjects = $state({} as ProjectResults)
    let portfolio = $state({} as Portfolio)
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
    {#await loadPage()}
        Loading...
    {:then promiseData} 

    <OrgStateChecker invert={false} stateToCheck="isReadyForProjects">
        {#if myProjects && myProjects.results && myProjects.results?.length > 0}
        <div class="grid grid-cols-3 gap-4">
            <CSSection cssClass="col-span-2">
                <SectionSubHeading>My Projects</SectionSubHeading>
                <ProjectActionTable projects={myProjects.results} {portfolio} />
            </CSSection>

            <CSSection>
                <SectionSubHeading>Activity</SectionSubHeading>
                {#if activity && activity.results}
                <ol class="divide-gray-200 mt-3 divide-y dark:divide-gray-700">
                    {#each activity.results as a}
                        {@render activityItem(a as Activity)}
                    {/each}
                </ol>
                {:else}
                <div class="text-center p-6">
                    No activity yet...
                </div>
                {/if}
            </CSSection>
        </div>
        {/if}

        {#snippet elseRender()}
        <CSSection>
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
        </CSSection>    
        {/snippet}
        
    </OrgStateChecker>

    {/await}

</div>



{#snippet activityItem(act:Activity)}
<li>
    <a href={act.link} class="block items-center p-3 hover:bg-gray-100 dark:hover:bg-gray-700 sm:flex">
    <!-- <img class="mb-3 me-3 h-12 w-12 rounded-full sm:mb-0" src="/images/users/jese-leos.png" alt="Jese Leos image" /> -->
    <Avatar src={act.user.profileImage || ''} class="mr-3 h-16 w-16">
        {getInitialsFromName(act.user.firstName + " " + act.user.lastName)}
    </Avatar>
    <div class="text-gray-600 dark:text-gray-400">
        <div class="text-base font-normal">
        <span class="font-medium text-gray-900 dark:text-white">Jese Leos</span> likes <span class="font-medium text-gray-900 dark:text-white">Bonnie Green's</span> post in
        <span class="font-medium text-gray-900 dark:text-white"> How to start with Flowbite library</span>
        </div>
        <div class="text-sm font-normal">"I wanted to share a webinar zeroheight."</div>
        <span class="inline-flex items-center text-xs font-normal text-gray-500 dark:text-gray-400">
        <svg class="me-1 h-2.5 w-2.5" aria-hidden="true" xmlns="http://www.w3.org/2000/svg" fill="currentColor" viewBox="0 0 20 20">
            <path
            d="M10 .5a9.5 9.5 0 1 0 0 19 9.5 9.5 0 0 0 0-19ZM8.374 17.4a7.6 7.6 0 0 1-5.9-7.4c0-.83.137-1.655.406-2.441l.239.019a3.887 3.887 0 0 1 2.082 2.5 4.1 4.1 0 0 0 2.441 2.8c1.148.522 1.389 2.007.732 4.522Zm3.6-8.829a.997.997 0 0 0-.027-.225 5.456 5.456 0 0 0-2.811-3.662c-.832-.527-1.347-.854-1.486-1.89a7.584 7.584 0 0 1 8.364 2.47c-1.387.208-2.14 2.237-2.14 3.307a1.187 1.187 0 0 1-1.9 0Zm1.626 8.053-.671-2.013a1.9 1.9 0 0 1 1.771-1.757l2.032.619a7.553 7.553 0 0 1-3.132 3.151Z"
            />
        </svg>
        Public
        </span>
    </div>
    </a>
</li>
{/snippet}