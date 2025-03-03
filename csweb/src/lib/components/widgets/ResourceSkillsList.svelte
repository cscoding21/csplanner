<script lang="ts">
    import type  { Skill } from "$lib/graphql/generated/sdk";
	import { decodeProficiency } from "$lib/services/resource";
	import { safeInt } from "$lib/utils/helpers";

    interface Props {
        skills: Skill[]
    }
    let { skills }:Props = $props()

    console.log("skills", skills)

    let sortedSkills = skills.slice().sort((s1:Skill, s2:Skill) => {
        return (safeInt(s1.proficiency) < safeInt(s2.proficiency)) ? 1 : (safeInt(s1.proficiency) > safeInt(s2.proficiency)) ? -1 : 0
    });

    console.log("sortedSkills", sortedSkills)
</script>

{#if sortedSkills}
    {#each sortedSkills as s, index}
        {#if s.proficiency === 3}
        <span class="font-semibold text-gray-800 dark:text-gray-300" title={decodeProficiency(3)}>{s.name}</span>
        {:else if s.proficiency === 2}
        <span class="font-medium text-gray-500 dark:text-gray-400" title={decodeProficiency(2)}>{s.name}</span>
        {:else if s.proficiency === 1}
        <span class="font-normal text-gray-300 dark:text-gray-400" title={decodeProficiency(1)}>{s.name}</span>
        {/if}

        {#if index < sortedSkills.length - 1}
            <span> , </span>
        {/if}
    {/each}
{/if}

