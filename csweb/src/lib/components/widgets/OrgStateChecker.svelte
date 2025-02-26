<script lang="ts">
	import type { Organization } from "$lib/graphql/generated/sdk";
	import { getOrganization } from "$lib/services/organization";
    import { type Snippet } from "svelte";

    interface Props {
		stateToCheck: "hasSkills"|"hasFundingSources"|"hasValueCategories"|"hasRoles"|"hasTemplates"|"hasResources"|"hasReviewedOrgSettings"|"isReadyForProjects";
		invert?: boolean;
        children: Snippet;
		elseRender?:Snippet;
	}
	let { 
		stateToCheck, 
		invert, 
		children,
		elseRender
	}: Props = $props();

    const checkStatus = ():boolean => {
        if(org.setup) {
            switch(stateToCheck) {
                case("hasSkills"):
                    return org.setup.hasSkills;
                case("hasFundingSources"):
                    return org.setup.hasFundingSources;
                case("hasValueCategories"):
                    return org.setup.hasValueCategories;
                case("hasRoles"):
                    return org.setup.hasRoles;
                case("hasTemplates"):
                    return org.setup.hasTemplates;
                case("hasResources"):
                    return org.setup.hasResources;
                case("hasReviewedOrgSettings"):
                    return org.setup.hasReviewedOrgSettings;
                case("isReadyForProjects"):
                    return org.setup.isReadyForProjects;
                default:
                    return false
            }
        }

        return false
    }

    const loadPage = async () => {
        getOrganization().then(o => {
            org = o
        })  
    }

    let org = $state({} as Organization)
</script>

{#await loadPage()}
    Loading...
{:then promiseData} 
    {#if checkStatus()}
        {@render children()}
    {:else if !checkStatus() && invert }
        {@render children()}
    {:else}
        {#if elseRender}
            {@render elseRender()}
        {/if}
    {/if}
{/await}



<!--
<OrgStateChecker invert={false} stateToCheck="hasSkills">
    I has skills

    {#snippet elseRender()}
        <p>I have no skills</p>
    {/snippet}
</OrgStateChecker>
-->