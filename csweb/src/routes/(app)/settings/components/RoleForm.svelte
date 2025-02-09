<script lang="ts">
	import { SectionHeading, TextInput, NumberInput } from '$lib/components';
	import { Button } from 'flowbite-svelte';
	import { roleSchema, roleForm } from '$lib/forms/resource.validation';
	import { mergeErrors, parseErrors } from '$lib/forms/helpers';
	import { onMount } from 'svelte';
	import { is } from '$lib/utils/check';
	import type { Role, UpdateRole } from '$lib/graphql/generated/sdk';
	import { addToast } from '$lib/stores/toasts';
	import { callIf, deepCopy } from '$lib/utils/helpers';
	import { updateRole } from '$lib/services/resource';

	onMount(async () => {
		if (role && role.id) {
			rf = deepCopy(role) as UpdateRole;
			rf.id = role.id;
		} else {
			rf = deepCopy(roleForm);
		}
	});

	interface Props {
		role: Role | undefined;
		update?: Function;
	}
	let { 
		role = $bindable(), 
		update 
	}: Props = $props();

	let errors: any = $state({});

	const submitForm = () => {
		errors = {};

		const roleFormParsed = roleSchema.cast(rf);
		roleSchema
			.validate(roleFormParsed, { abortEarly: false })
			.then(() => {
				updateRole(roleFormParsed).then((res) => {
					if (res && res.status?.success) {
						addToast({
							message: 'Role updated successfully',
							dismissible: true,
							type: 'success'
						});

						rf = deepCopy(roleForm);
						callIf(update);
					} else {
						addToast({
							message: 'Error updating role: ' + res.status?.message,
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

	let rf = $state({} as UpdateRole);
</script>

<SectionHeading>
	{#if is(role?.id)}
		Update Role
	{:else}
		Add a New Role
	{/if}
</SectionHeading>
<form>
	<div class="grid grid-cols-4 gap-4">
		<div class="col-span-3">
			<TextInput
				bind:value={rf.name}
				placeholder="Role name"
				fieldName="Role Name"
				error={errors.name}
			/>
		</div>
		<div class="col-span-1">
            <NumberInput
				bind:value={rf.hourlyRate as number}
				placeholder="Hourly rate"
				fieldName="Hourly rate"
				error={errors.description}
			/>
		</div>
		<div class="col-span-4">
			<TextInput
				bind:value={rf.description}
				placeholder="Description"
				fieldName="Description"
				error={errors.description}
			/>
		</div>
		<div class="col-span-4">
			<span class="float-right">
				<Button onclick={submitForm}>
					{#if rf.id}
						Update role
					{:else}
						Add role
					{/if}
				</Button>
			</span>
		</div>
	</div>
</form>
