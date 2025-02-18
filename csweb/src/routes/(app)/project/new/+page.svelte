<script lang="ts">
    import { Toggle, type SelectOptionType } from 'flowbite-svelte';
    import { CSSection, PageHeading, TextInput, SelectInput, TextAreaInput, CSHR } from '$lib/components';
    import { getDefaultProject } from '$lib/forms/project.validation';
	import type { UpdateProjectBasics } from '$lib/graphql/generated/sdk';
    import { findAllUsers } from '$lib/services/user';
	import { ProjectTemplateSelector } from '../components';

    let errors: any = $state({ name: '', ownerID: '', description: '', isCapitalized: false });
	let userOpts = $state([] as SelectOptionType<string>[]);
    let basicsForm = $state(getDefaultProject().projectBasics as UpdateProjectBasics);
    

    const loadPage = async () => {
		findAllUsers()
			.then((r) => {
				userOpts = r.results?.map((r) => ({
					name: r.firstName + ' ' + r.lastName,
					value: r.email as string
				})) as SelectOptionType<string>[];

				return r;
			});
	};

</script>

<PageHeading title="Create a New Project" />
<div class="p-4"> 
    <CSSection>

{#await loadPage()}
    Loading...
{:then promiseData} 
<div class="max-w-lg">
<TextInput
    bind:value={basicsForm.name}
    fieldName="Project Name"
    placeholder="Project name"
    error={errors.name}>
    The name will be used to identify the project
</TextInput>

<SelectInput
    bind:value={basicsForm.ownerID as string}
    fieldName="Owner"
    error={errors.ownerID}
    options={userOpts}> 
    This is a user in the system as opposed to a resource.  This user will manage the project within csPlanner.
</SelectInput>

<TextAreaInput
    bind:value={basicsForm.description as string}
    rows={8}
    error={errors.description}
    placeholder="Executive summary"
    fieldName="Executive Summary"
/>

<div class="pb-4 mb-2">
    <Toggle class="mt-3" bind:checked={basicsForm.isCapitalized}>Capitalized</Toggle>
</div>
</div>

<CSHR />


{/await}


</CSSection>
</div>