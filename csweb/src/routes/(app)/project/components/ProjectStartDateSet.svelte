<script lang="ts">
	import { Button, Select, type SelectOptionType } from 'flowbite-svelte';
	import { updateProjectBasics } from '$lib/services/project';
	import { basicSchema } from '$lib/forms/project.validation';
	import { mergeErrors, parseErrors } from '$lib/forms/helpers';
	import type { Project, UpdateProjectBasics } from '$lib/graphql/generated/sdk';
	import { addToast } from '$lib/stores/toasts';
	import { callIf } from '$lib/utils/helpers';
	import { coalesceToType } from '$lib/forms/helpers';
	import { getCSWeeks } from '$lib/utils/cscal';

	interface Props {
		project: Project;
		update?: Function;
	}
	let { project = $bindable(), update }: Props = $props();

    let startDate = $state(new Date(project.projectBasics.startDate).toLocaleDateString())

	const getYearOpts = ():SelectOptionType<number>[] => {
		let opts:SelectOptionType<number>[] = [] 
		const currentYear = new Date().getFullYear()

		opts.push({ name: currentYear.toString(), value: currentYear })

		for(let i = 1; i < 5; i++) {
			opts.push({ name: (currentYear + i).toString(), value: currentYear + i })	
		}

		return opts
	}

	const getWeekOpts = (year:number):SelectOptionType<string>[] => {
		let opts:SelectOptionType<string>[] = []

		const startDate = new Date(year, 0, 1)
		const endDate = new Date(year, 11, 31)
		const weeks = getCSWeeks(startDate, endDate)
		
		for (let i = 0; i < weeks.length; i++) {
		const w = weeks[i]
			opts.push({name: w.toLocaleDateString(), value: w.toLocaleDateString()})
		}

		return opts
	}

	const updateStartDate = async () => {
		errors = {};

		let bf = coalesceToType<UpdateProjectBasics>(project.projectBasics, basicSchema);
		bf.startDate = new Date(startDate)

		const projectBasicsParsed = basicSchema.cast(bf);
		basicSchema
			.validate(projectBasicsParsed, { abortEarly: false })
			.then(async () => {
				if (!project.id) {
					addToast({
						message: 'Error updating project start date: project id not available',
						dismissible: true,
						type: 'error'
					});

					return
				}

				updateProjectBasics(project.id, projectBasicsParsed)
					.then((res) => {
						if (res.status?.success) {
							addToast({
								message: 'Project start date updated successfully',
								dismissible: true,
								type: 'success'
							});

							callIf(update);
						} else {
							addToast({
								message: 'Error updating project start date: ' + res.status?.message,
								dismissible: true,
								type: 'error'
							});
						}
					})
					.catch((err) => {
						addToast({
							message: 'Error updating project start date: ' + err,
							dismissible: true,
							type: 'error'
						});
					});
			})
			.catch((err) => {
				errors = mergeErrors(errors, parseErrors(err));
			});
	};

	let errors: any = $state({ name: '', ownerID: '', description: '' });
	let yearOpts:SelectOptionType<number>[] = $state(getYearOpts())
	let selectedYear = $state(new Date().getFullYear())
	let weekOpts:SelectOptionType<string>[] = $derived(getWeekOpts(selectedYear))
</script>

<div class="flex">
	<div class="flex-1 mr-2"><Select items={yearOpts} bind:value={selectedYear} /></div>
	<div class="flex-1 mr-2"><Select items={weekOpts} bind:value={startDate} /></div>
	<div class="flex-1">
	<span class="float-right">
		<Button onclick={updateStartDate}>Set start date</Button>
	</span>
	<br class="clear-both" />
	</div>
</div>

