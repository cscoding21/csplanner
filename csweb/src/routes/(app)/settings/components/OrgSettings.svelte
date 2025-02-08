<script lang="ts">
	import { Button, type SelectOptionType } from 'flowbite-svelte';
	import { NumberInput, SectionHeading } from '$lib/components';
	import { getOrganization, updateOrganization } from '$lib/services/organization';
	import { orgForm, orgSchema } from '$lib/forms/organization.validation';
	import { mergeErrors, parseErrors } from '$lib/forms/helpers';
	import type { UpdateOrganization, Organization } from '$lib/graphql/generated/sdk';
	import { coalesceToType } from '$lib/forms/helpers';
	import { addToast } from '$lib/stores/toasts';
	import { callIf } from '$lib/utils/helpers';
	import { findAllUsers } from '$lib/services/user';

	interface Props {
		update?: Function;
	}
	let { update }: Props = $props();

	const load = async (): Promise<Organization> => {
		return await getOrganization()
			.then((org) => {
				of = coalesceToType<UpdateOrganization>(org.defaults, orgSchema);

				return org;
			})
			.catch((err) => {
				addToast({
					message: 'Error loading organization (OrgSettings): ' + err,
					dismissible: true,
					type: 'error'
				});

				return err;
			});
	};

	const updateOrg = async () => {
		errors = {};

		const orgFormParsed = orgSchema.cast(of);
		orgSchema
			.validate(orgFormParsed, { abortEarly: false })
			.then(async () => {
                console.log("ofp", orgFormParsed)

				updateOrganization  ({ name: "", id: "", defaults: orgFormParsed })
					.then((res) => {
						if (res.status?.success) {
							load().then(() => {
								addToast({
									message: 'Organization settings updated successfully',
									dismissible: true,
									type: 'success'
								});

								callIf(update);
							});
						} else {
							addToast({
								message: 'Error updating organization settings: ' + res.status?.message,
								dismissible: true,
								type: 'error'
							});
						}
					})
					.catch((err) => {
						addToast({
							message: 'Error updating organization settings: ' + err,
							dismissible: true,
							type: 'error'
						});
					});
			})
			.catch((err) => {
				errors = mergeErrors(errors, parseErrors(err));
			});
	};

	const loadPage = async () => {
		findAllUsers()
			.then((r) => {
				userOpts = r.results?.map((r) => ({
					name: r.firstName + ' ' + r.lastName,
					value: r.email as string
				})) as SelectOptionType<string>[];

				return r;
			})
			.then(() => {
				load().then(p => {

				});
			});
	};

	let errors: any = $state({ name: '', ownerID: '', description: '' });
	let userOpts = $state([] as SelectOptionType<string>[]);
	let of = $state({} as typeof orgForm);
</script>

{#await loadPage()}
	Loading...
{:then promiseData}
	{#if of}
		<SectionHeading>
			Organization Settings
		</SectionHeading>
	{/if}

	{#if of}	
		<NumberInput
			bind:value={of.discountRate}
			fieldName="Discount rate"
			placeholder="Discount rate"
			error={errors.discountRate}
		/>

        <NumberInput
			bind:value={of.hoursPerWeek}
			fieldName="Working hours per week"
			placeholder="Working hours per week"
			error={errors.hoursPerWeek}
		/>

        <NumberInput
			bind:value={of.focusFactor}
			fieldName="Focus factor"
			placeholder="Focus factor"
			error={errors.focusFactor}
		/>

        <NumberInput
			bind:value={of.commsCoefficient}
			fieldName="Comms coefficient"
			placeholder="Comms coefficient"
			error={errors.commsCoefficient}
		/>

        <NumberInput
			bind:value={of.genericBlendedHourlyRate}
			fieldName="Generic blended hourly rate"
			placeholder="Generic blended hourly rate"
			error={errors.genericBlendedHourlyRate}
		/>

        <NumberInput
			bind:value={of.workingHoursPerYear}
			fieldName="Workinng houre per year"
			placeholder="Workinng houre per year"
			error={errors.workingHoursPerYear}
		/>

		<div class="col-span-4">
			<span class="float-right">
				<Button onclick={updateOrg}>Update Settings</Button>
			</span>
			<br class="clear-both" />
		</div>
	{/if}
{/await}
