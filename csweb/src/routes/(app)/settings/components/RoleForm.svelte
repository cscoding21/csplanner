<script lang="ts">
	import { SectionHeading, TextInput, NumberInput, SectionSubHeading, SelectInput } from '$lib/components';
	import { Alert, Button, Rating, Table, TableBody, TableBodyCell, TableBodyRow, TableHead, TableHeadCell, type SelectOptionType } from 'flowbite-svelte';
	import { roleSchema, roleForm } from '$lib/forms/resource.validation';
	import { findSelectOptsFromList, mergeErrors, parseErrors } from '$lib/forms/helpers';
	import { onMount } from 'svelte';
	import { is } from '$lib/utils/check';
	import type { Role, UpdateRole, UpdateSkill } from '$lib/graphql/generated/sdk';
	import { addToast } from '$lib/stores/toasts';
	import { callIf, deepCopy } from '$lib/utils/helpers';
	import { updateRole } from '$lib/services/resource';
	import { TrashBinOutline } from 'flowbite-svelte-icons';
	import { skillForm, skillSchema } from '$lib/forms/skill.validation';
	import { getList } from '$lib/services/list';

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

	const delSkill = (id:string) => {
		rf.defaultSkills = rf.defaultSkills?.filter(s => s.id != id)
	}

	const addSkill = () => {
		const parsedSkillForm = skillSchema.cast(sf);
		parsedSkillForm.resourceID = rf.id as string
		skillSchema
			.validate(parsedSkillForm, { abortEarly: false })
			.then(() => {
				rf.defaultSkills = [...rf.defaultSkills as UpdateSkill[], { resourceID: rf.id, id: sf.skillID, proficiency: sf.proficiency} as UpdateSkill]
			})
			.catch((err) => {
				alert(err)
				errors = mergeErrors(errors, parseErrors(err));
			});
	}

	const submitForm = () => {
		errors = {};

		const roleFormParsed = roleSchema.cast(rf);

		if(roleFormParsed.defaultSkills) {
			for(let i = 0; i < roleFormParsed.defaultSkills?.length; i++) {
				roleFormParsed.defaultSkills[i].resourceID = rf.id as string

				//@ts-expect-error
				delete roleFormParsed.defaultSkills[i].name
			}
		}

		roleSchema
			.validate(roleFormParsed, { abortEarly: false })
			.then(() => {
				//@ts-expect-error
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
	let sf = $state(deepCopy(skillForm));

	let availableSkillOpts = $state([] as SelectOptionType<string>[]);
	const proficiencyOpts = [
		{ value: 1, name: 'Novice' },
		{ value: 2, name: 'Competent' },
		{ value: 3, name: 'Expert' }
	];

	const loadPage = async () => {
		getList('Skills')
			.then((l) => {
				availableSkillOpts = findSelectOptsFromList(l);
			});
	};
</script>

<SectionHeading>
	{#if is(role?.id)}
		Update Role
	{:else}
		Add a New Role
	{/if}
</SectionHeading>

{#await loadPage()}
	Loading...
{:then promiseData} 

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
		<SectionSubHeading>Default Skills</SectionSubHeading>

		{#if rf?.defaultSkills && rf?.defaultSkills.length > 0}
			<Table>
			<TableBody>
				{#each rf.defaultSkills as s, index}
					<TableBodyRow>
						<TableBodyCell>{s.id}</TableBodyCell>
						<TableBodyCell
							><Rating rating={s.proficiency?.valueOf() as number} total={3} /></TableBodyCell
						>
						<TableBodyCell tdClass="float-right pt-2">
							<Button color="dark" onclick={() => delSkill(s.id)}>
								<TrashBinOutline size="sm" />
							</Button>
						</TableBodyCell>
					</TableBodyRow>
				{/each}
			</TableBody>
			</Table>

			<div class="grid grid-cols-5 gap-5 mt-4">
				<span class="col-span-2">
					<SelectInput
						fieldName="Skill"
						bind:value={sf.skillID}
						error={errors.skillID}
						options={availableSkillOpts}
					/>
				</span>
				<span class="col-span-2">
					<SelectInput
						fieldName="Proficiency"
						bind:value={sf.proficiency}
						error={errors.proficiency}
						options={proficiencyOpts}
					/>
				</span>
				<span class="w-1/5 pt-8">
					<Button onclick={addSkill}>Add</Button>
				</span>
			</div>
		{:else}
			<Alert>No default skills have been added for this role.</Alert>
		{/if}
		</div>



		<div class="col-span-4">
			<span class="float-right">
				<Button onclick={submitForm} color="primary">
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
	
{/await}

