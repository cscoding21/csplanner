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
			error={errors.discountRate}>
            The discount rate is in the net present value calculation of a project.
        </NumberInput>

        <NumberInput
			bind:value={of.hoursPerWeek}
			fieldName="Working hours per week"
			placeholder="Working hours per week"
			error={errors.hoursPerWeek}> 
            This value is used for estimating a work week ***check to see if it can be overridden at the resource level***
        </NumberInput>

        <NumberInput
			bind:value={of.focusFactor}
			fieldName="Focus factor"
			placeholder="Focus factor"
			error={errors.focusFactor}> 
            This value is used to estimate additional hours for users that are involved in multiple projects
        </NumberInput>

        <NumberInput
			bind:value={of.commsCoefficient}
			fieldName="Comms coefficient"
			placeholder="Comms coefficient"
			error={errors.commsCoefficient}> 
            This value is used to estimate additional hours for multiple users working on a single task
        </NumberInput>

        <NumberInput
			bind:value={of.genericBlendedHourlyRate}
			fieldName="Generic blended hourly rate"
			placeholder="Generic blended hourly rate"
			error={errors.genericBlendedHourlyRate}> 
            This is a default hourly rate that acts as a default when estimating project costs.  It can be overridden
            at the role of individual level
        </NumberInput>

        <NumberInput
			bind:value={of.workingHoursPerYear}
			fieldName="Workinng hours per year"
			placeholder="Workinng hours per year"
			error={errors.workingHoursPerYear}> 
            The number of working hours in a year is used to calculate an hourly rate for resources that have an annual cost 
            associated with them.  This will override orginzation and role level values.
        </NumberInput>

		<div class="col-span-4">
			<span class="float-right">
				<Button onclick={updateOrg}>Update Settings</Button>
			</span>
			<br class="clear-both" />
		</div>
	{/if}
{/await}
