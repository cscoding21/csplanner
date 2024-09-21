<script lang="ts">
	import { Avatar, Button, Dropzone, Hr, Modal } from 'flowbite-svelte';
	import { TextInput } from '$lib/components';
	import { getInitialsFromName } from '$lib/utils/format';
	import { resourceSchema, resourceForm } from '$lib/forms/resource.validation';
	import { getResource, updateResource } from '$lib/services/resource';
	import type { UpdateResource } from '$lib/graphql/generated/sdk';
	import { addToast } from '$lib/stores/toasts';
	import { coalesceToType, mergeErrors, parseErrors } from '$lib/forms/helpers';
	import { deepCopy } from '$lib/utils/helpers';
	import { handleFileUpload } from '$lib/services/file';
	import type { Snippet } from 'svelte';

	let editModalOpen: boolean = $state(false);
	let isUpdate = $state(false);
	let errors = $state({ name: '', role: '' });
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
			formTitle = r.name;
			isUpdate = true;
			rf = coalesceToType<UpdateResource>(r, resourceSchema);
		});
	}

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
</script>

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

	<TextInput
		bind:value={rf.role}
		fieldName="Role"
		placeholder="Resource role"
		error={errors.role}
	/>

	<svelte:fragment slot="footer">
		<div class="place-content-end">
			<Button on:click={updateRes}>Update Resource</Button>
		</div>
	</svelte:fragment>
</Modal>
