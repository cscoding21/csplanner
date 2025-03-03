<script lang="ts">
	import { SelectInput } from '$lib/components';
	import { Button, Label, Range, type SelectOptionType } from 'flowbite-svelte';
	import { skillForm, skillSchema } from '$lib/forms/skill.validation';
	import { mergeErrors, parseErrors, findSelectOptsFromList } from '$lib/forms/helpers';
	import { decodeProficiency, updateResourceSkill } from '$lib/services/resource';
	import { addToast } from '$lib/stores/toasts';
	import { callIf, deepCopy } from '$lib/utils/helpers';
	import { getList } from '$lib/services/list';

	interface Props {
		resourceID: string;
		update?: Function;
	}
	let { resourceID, update }: Props = $props();

	let errors: any = $state({ skillID: '', proficiency: '' });
	let sf = $state(deepCopy(skillForm));

	const add = () => {
		errors = {};

		const parsedSkillForm = skillSchema.cast(sf);
		parsedSkillForm.resourceID = resourceID;
		skillSchema
			.validate(parsedSkillForm, { abortEarly: false })
			.then(() => {
				updateResourceSkill({
					resourceID: parsedSkillForm.resourceID,
					id: parsedSkillForm.skillID,
					proficiency: parsedSkillForm.proficiency as number
				}).then((res) => {
					if (res.success) {
						addToast({
							message: 'Skill successfully added to resource',
							dismissible: true,
							type: 'success'
						});

						sf.id = '';
						skillForm.sf = 2;

						callIf(update);
					} else {
						addToast({
							message: 'Error creating resource: ' + res.message,
							dismissible: true,
							type: 'error'
						});
					}
				});
			})
			.catch((err) => {
				errors = mergeErrors(errors, parseErrors(err));

				console.log(errors);
			});
	};

	let availableSkillOpts = $state([] as SelectOptionType<string>[]);
	// const proficiencyOpts = [
	// 	{ value: 1, name: decodeProficiency(1) },
	// 	{ value: 2, name: decodeProficiency(2) },
	// 	{ value: 3, name: decodeProficiency(3) }
	// ];

	const loadPage = async () => {
		getList('Skills')
			.then((l) => {
				availableSkillOpts = findSelectOptsFromList(l);
			})
			.then(() => {
				skillForm.resourceID = resourceID;
			});
	};
</script>

{#await loadPage()}
	<span>Loading...</span>
{:then data}
	<div class="grid grid-cols-5 gap-5">
		<span class="col-span-2">
			<SelectInput
				fieldName="Skill"
				bind:value={sf.skillID}
				error={errors.skillID}
				options={availableSkillOpts}
			/>
		</span>
		<span class="col-span-2">
			<!-- <SelectInput
				fieldName="Proficiency"
				bind:value={sf.proficiency}
				error={errors.proficiency}
				options={proficiencyOpts}
			/> -->
			<Label>Proficiency</Label>
			<Range size="lg" id="range-steps" min="1" max="3" bind:value={sf.proficiency} step="1" />
			<br /><small>{decodeProficiency(sf.proficiency)}</small>
		</span>
		<span class="w-1/5 pt-8">
			<Button onclick={add}>Add</Button>
		</span>
	</div>
{/await}
