<script lang="ts">
	import { Avatar, Button, Dropzone, Hr, Modal, type SelectOptionType } from 'flowbite-svelte';
	import { TextInput, MoneyInput, SelectInput } from '$lib/components';
	import { getInitialsFromName } from '$lib/utils/format';
	import { resourceSchema, resourceForm } from '$lib/forms/resource.validation';
	import { findAllRoles, getResource, updateResource } from '$lib/services/resource';
	import type { UpdateResource } from '$lib/graphql/generated/sdk';
	import { addToast } from '$lib/stores/toasts';
	import { coalesceToType, mergeErrors, parseErrors } from '$lib/forms/helpers';
	import { deepCopy } from '$lib/utils/helpers';
	import { handleFileUpload } from '$lib/services/file';
	import type { Snippet } from 'svelte';
	import NumberInput from '$lib/components/forms/NumberInput.svelte';

	let editModalOpen: boolean = $state(false);
	let isUpdate = $state(false);
	let errors = $state({ name: '', roleID: '', initialCost: '', annualizedCost: '', type: '', status: '', availableHoursPerWeek: '' });
	let rf = $state(deepCopy(resourceForm));
	let formTitle = $state('New Resource');

	interface Props {
		id?: string | undefined;
		children?: Snippet;
		update: Function;
	}
	let { id, update, children }: Props = $props();

	if (id) {
		getResource(id as string).then((r) => {
			formTitle = r.data?.name || "New resource";
			isUpdate = true;
			rf = coalesceToType<UpdateResource>(r.data, resourceSchema);
			rf.id = r.meta.id
		});
	}

	let typeOpts:SelectOptionType<string>[] = [
		{name: "Human", value: "human"},
		{name: "Software", value: "software"},
		{name: "Equipment", value: "equipment"},
	]

	let statusOpts:SelectOptionType<string>[] = [
		{name: "In House", value: "inhouse"},
		{name: "Proposed", value: "proposed"},
		{name: "Exited", value: "exited"},
	]

	let roleOpts:SelectOptionType<string>[] = $state([] as SelectOptionType<string>[])

	const updateRes = async () => {
		const resourceSchemaParsed = resourceSchema.cast(rf);
		resourceSchema
			.validate(resourceSchemaParsed, { abortEarly: false })
			.then((f) => {
				console.log('validation passed');

				updateResource(f as UpdateResource).then((res) => {
					if (res.status?.success.valueOf()) {
						editModalOpen = false;

						update();

						addToast({
							message: 'Resource ' + rf.name + ' updated successfully',
							dismissible: true,
							type: 'success'
						});
					} else {
						addToast({
							message: 'Error updating resource: ' + res.status?.message,
							dismissible: true,
							type: 'error'
						});
					}
				});
			})
			.catch((err) => {
				errors = mergeErrors(errors, parseErrors(err));

				console.error(errors);
			});
	};

	let fileValue = [];
	const handleChange = (e: any) => {
		const files = e.target.files;
		if (files.length > 0) {
			const file = files[0];
			console.log('handleChange', file);

			handleFileUpload(file, id as string, 'profile_image', handleUploadComplete);

			fileValue.push(file.name);
		}
	};

	const handleUploadComplete = (upload: any) => {
		console.log('upload complete ', upload.file.name);

		rf.profileImage = upload.url;
	};

	const loadPage = async () => {
		findAllRoles()
			.then((l) => {
				roleOpts = l.results?.map(r => { return { name: r.data.name, value: r.meta.id} }) as SelectOptionType<string>[];
			});
	};
</script>

{#await loadPage()}
	Loading...
{:then promiseData} 
<Button onclick={() => (editModalOpen = true)}>
	{#if children}{@render children()}{/if}
</Button>

<Modal title="Update {formTitle}" bind:open={editModalOpen}>
	{#if isUpdate}
		<div class="flex items-center space-x-4 rtl:space-x-reverse">
			<div>
				<Dropzone id="dropzone" defaultClass="float-left mr-4" on:change={handleChange}>
					<Avatar size="xl" src={rf.profileImage || ''} rounded class="hover:cursor-pointer"
						>{getInitialsFromName(rf.name)}</Avatar
					>
				</Dropzone>
				<div class="ml-4 font-medium dark:text-white">Click profile image to re-upload</div>
			</div>
		</div>

		<Hr />
	{/if}

	<TextInput
		bind:value={rf.name}
		fieldName="Name"
		placeholder="Resource name"
		error={errors.name}
	/>

	<div class="flex">
		<div class="flex-1 pr-2">
			<SelectInput
				bind:value={rf.type as string}
				fieldName="Type"
				error={errors.type}
				options={typeOpts}
			/>
		</div>
		<div class="flex-1">
			<SelectInput
				bind:value={rf.status as string}
				fieldName="Status"
				error={errors.status}
				options={statusOpts}
			/>
		</div>
	</div>

	{#if rf.type === "human"}
	<NumberInput
		bind:value={rf.availableHoursPerWeek as number}
		fieldName="Available Hours per Week"
		error={errors.availableHoursPerWeek}
	/>

	<SelectInput
		bind:value={rf.roleID as string}
		fieldName="Role"
		error={errors.roleID}
		options={roleOpts}
	/>
	{/if}

	<div class="flex">
		<div class="flex-1 pr-2">
			<MoneyInput
				bind:value={rf.initialCost}
				fieldName="Initial Cost"
				error={errors.initialCost}
			/>
		</div>	
		<div class="flex-1 pr-2">
			<MoneyInput
				bind:value={rf.annualizedCost}
				fieldName="Annualized Cost"
				error={errors.annualizedCost}
			/>
		</div>
	</div>

	<Hr />
	<div class="float-right">
		<Button color="primary" onclick={updateRes}>Update Resource</Button>
	</div>
</Modal>

{/await}

