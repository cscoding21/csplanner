<script lang="ts">
	import { SelectInput } from '$lib/components';
	import { Button, type SelectOptionType } from 'flowbite-svelte';
	import { skillForm, skillSchema } from '$lib/forms/skill.validation';
	import { mergeErrors, parseErrors, findSelectOptsFromList } from '$lib/forms/helpers';
	import { updateResourceSkill } from '$lib/services/resource';
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

						skillForm.id = '';
						skillForm.proficiency = '';

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
	const proficiencyOpts = [
		{ value: 1, name: 'Novice' },
		{ value: 2, name: 'Competent' },
		{ value: 3, name: 'Expert' }
	];

	const loadPage = async () => {
		getList('Skills')
			.then((l) => {
				availableSkillOpts = findSelectOptsFromList(l);
			})
			.then(() => {
				skillForm.resourceID = resourceID;
			});
	};

	loadPage();
</script>

{#await loadPage}
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
			<SelectInput
				fieldName="Proficiency"
				bind:value={sf.proficiency}
				error={errors.proficiency}
				options={proficiencyOpts}
			/>
		</span>
		<span class="w-1/5 pt-8">
			<Button on:click={add}>Add</Button>
		</span>
	</div>
{/await}
