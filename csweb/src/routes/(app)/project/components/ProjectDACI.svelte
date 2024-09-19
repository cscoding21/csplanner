<script lang="ts">
    import { findAllResources } from '$lib/services/resource'
    import { MultiSelectInput, SectionHeading } from '$lib/components';
	import { getDefaultProject, daciSchema } from '$lib/forms/project.validation'
	import { getProject, updateProjectDaci } from '$lib/services/project'
	import { mergeErrors, parseErrors } from '$lib/forms/helpers'
	import { Button, type SelectOptionType  } from 'flowbite-svelte'
	import type { UpdateProjectDaci } from '$lib/graphql/generated/sdk'
	import { coalesceToType } from '$lib/forms/helpers';
	import { addToast } from '$lib/stores/toasts';
	import { callIf } from '$lib/utils/helpers';


	interface ProjectDACIProps {
        id: string;
        update?: Function;
    }
    let { id, update }:ProjectDACIProps = $props()
	let errors: any = $state({})


	const load = async () => {
		await getProject(id).then(proj => {
			daciForm = coalesceToType<UpdateProjectDaci>(proj.projectDaci, daciSchema)
		}).catch(err => {
			addToast({ 
				message: "Error loading project (ProjectDACI): " + err, 
				dismissible: true, 
				type: "error"}
			)
		})
	}

	const updateDACI = async () => {
		errors = {}

		const projectDaciParsed = daciSchema.cast(daciForm)
		daciSchema.validate(projectDaciParsed, { abortEarly: false })
			.then(async () => {
				await updateProjectDaci(id, projectDaciParsed).then(res => {
						if(res.status?.success) {
							load().then(p => {
								addToast({ 
									message: "Project DACI updated successfully", 
									dismissible: true, 
									type: "success"}
								)

								callIf(update)
							})
						} else {
							addToast({ 
								message: "Error updating project DACI: " + res.status?.message, 
								dismissible: true, 
								type: "error"}
							)
						}
					}).catch(err => {
						addToast({ 
							message: "Error updating project DACI: " + err, 
							dismissible: true, 
							type: "error"}
						)
					})
		}).catch(err => {
			errors = mergeErrors(errors, parseErrors(err))

			console.error(errors);
		})
	}

    const loadPage = async () => {
		findAllResources()
			.then((r) => r)
			.then((r) => {
				resourceOpts = r.results?.map((r) => ({
					name: r.name,
					value: r.id as string
				})) as SelectOptionType<string>[];
			})
			.then(() => {
				load();
			});
	};

	let resourceOpts = $state([] as SelectOptionType<string>[]);
	let daciForm = $state(getDefaultProject().projectDaci);

    loadPage()
</script>


<SectionHeading>DACI</SectionHeading>

{#await loadPage }
	Loading...
{:then promiseData} 
	{#if daciForm}

<MultiSelectInput 
	fieldName="Drivers"
	error={errors.driver}
	options={resourceOpts}
	bind:value={daciForm.driver as string[]} 
	update={() => callIf(update)}
	/>

<MultiSelectInput 
	fieldName="Approvers"
	error={errors.approver}
	options={resourceOpts}
	bind:value={daciForm.approver as string[]} 
	update={() => callIf(update)}
	/>

<MultiSelectInput 
	fieldName="Contributors"
	error={errors.contributor}
	options={resourceOpts}
	bind:value={daciForm.contributor as string[]} 
	update={() => callIf(update)}
	/>

<MultiSelectInput 
	fieldName="Informed"
	error={errors.informed}
	options={resourceOpts}
	bind:value={daciForm.informed as string[]} 
	update={() => callIf(update)}
	/>


<div class="col-span-4">
	<span class="float-right">
		<Button on:click={updateDACI}> Update Team </Button>
	</span>
	<br class="clear-both" />
</div>

	{/if}
{/await}