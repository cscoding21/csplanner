<script lang="ts">
	import { SectionHeading, TextInput, NumberInput, SectionSubHeading, SelectInput } from '$lib/components';
	import { Alert, Button, Range, Select, Table, TableBody, TableBodyCell, TableBodyRow, TableHead, TableHeadCell, type SelectOptionType } from 'flowbite-svelte';
	import { roleSchema, roleForm } from '$lib/forms/resource.validation';
	import { findSelectOptsFromList, mergeErrors, parseErrors } from '$lib/forms/helpers';
	import { onMount } from 'svelte';
	import { is } from '$lib/utils/check';
	import type { Role, UpdateRole, UpdateSkill } from '$lib/graphql/generated/sdk';
	import { addToast } from '$lib/stores/toasts';
	import { callIf, deepCopy } from '$lib/utils/helpers';
	import { decodeProficiency, updateRole } from '$lib/services/resource';
	import { TrashBinOutline } from 'flowbite-svelte-icons';
	import { skillForm, skillSchema } from '$lib/forms/skill.validation';
	import { getList } from '$lib/services/list';
	import { newID } from '$lib/utils/id';

	onMount(async () => {
		if (role && role.id) {
			rf = deepCopy(role) as UpdateRole;
		} else {
			rf = deepCopy(roleForm);
			rf.id = newID()
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
		sf.parentID = rf.id
		sf.id = newID()
		const parsedSkillForm = skillSchema.cast(sf);

		skillSchema
			.validate(parsedSkillForm, { abortEarly: false })
			.then(() => {
				rf.defaultSkills = [...rf.defaultSkills as UpdateSkill[], { parentID: rf.id, id: sf.id, skillID: sf.skillID, proficiency: sf.proficiency} as UpdateSkill]
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
				roleFormParsed.defaultSkills[i].parentID = rf.id as string

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

	const getSkillName = (skillID:string):string => {
		const sk = availableSkillOpts.filter(o => o.value === skillID)

		if (sk && sk.length > 0) {
			return sk[0].name as string
		}

		return ""
	}

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
		<SectionSubHeading>Default Skills</SectionSubHeading>

		{#if rf?.defaultSkills && rf?.defaultSkills.length > 0}
		<Table>
			<TableHead>
				<TableHeadCell>Skill</TableHeadCell>
				<TableHeadCell>Proficienty</TableHeadCell>
				<TableHeadCell>
					<span class="sr-only">Action</span>
				</TableHeadCell>
			</TableHead>
			<TableBody>
				{#each rf.defaultSkills as s, index}
					<TableBodyRow>
						<TableBodyCell>{getSkillName(rf.defaultSkills[index].skillID)}</TableBodyCell>
						<TableBodyCell>
							<Range size="lg" id="range-steps" min="1" max="3" bind:value={rf.defaultSkills[index].proficiency as number} step="1" />
							<br /><small>{decodeProficiency(rf.defaultSkills[index].proficiency as number)}</small>
						</TableBodyCell>
						<TableBodyCell class="float-right pt-2">
							<Button color="dark" onclick={() => delSkill(s.id as string)}>
								<TrashBinOutline size="sm" />
							</Button>
						</TableBodyCell>
					</TableBodyRow>
				{/each}
			</TableBody>
			<tfoot>
				<tr>
					<th class="px-6 pt-4" colspan={3}><SectionSubHeading>Add skill</SectionSubHeading></th>
			</tr>    
				<tr>
					<th class="px-4">
						<Select
							bind:value={sf.skillID}
							size="sm"
							items={availableSkillOpts}
						/>
					</th>
					<th class="px-6 py-4">
						<Range size="lg" id="range-steps" min="1" max="3" bind:value={sf.proficiency} step="1" />
						<br /><small>{decodeProficiency(sf.proficiency)}</small>
					</th>
					<th class="px-6 py-4">
						<Button pill onclick={addSkill}>Add</Button>
					</th>
				</tr>
			</tfoot>
		</Table>
		{:else}
			<Alert>No default skills have been added for this role.</Alert>
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

