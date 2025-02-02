<script lang="ts">
    import { onMount } from "svelte";
    import type { SelectOptionType } from "flowbite-svelte";
    import type { ProjectValueLine } from "$lib/graphql/generated/sdk";
    import { SelectInput, MoneyInput } from "$lib/components"
    import { Heading, Button } from "flowbite-svelte"
    import { getList } from "$lib/services/list";
    import { callIf, deepCopy } from "$lib/utils/helpers";
    import { valueLineDefaultForm, valueLineSchema } from "$lib/forms/project.validation";
    import { findSelectOptsFromList, parseErrors, mergeErrors } from "$lib/forms/helpers";
    import { updateProjectValueLine } from "$lib/services/project";
    import { addToast } from "$lib/stores/toasts";

    let fundingSourceOpts = $state([] as SelectOptionType<string>[]);
    let valueCategoryOpts = $state([] as SelectOptionType<string>[]);

    onMount(async () => {
		if (valueItem && valueItem.id) {
			valueItemForm = deepCopy(valueItem);
			valueItemForm.projectID = projectID;
		} else {
			valueItemForm = valueLineDefaultForm(projectID);
		}
	});

	interface Props {
		id?: string;
		valueItem: ProjectValueLine| undefined;
		projectID: string;
		update?: Function;
	}
	let { 
		id, 
		valueItem = $bindable(), 
		projectID, 
		update 
	}: Props = $props();

    let errors: any = $state({});
    

    const submitForm = () => {
		errors = {};

		const valueItemFormParsed = valueLineSchema.cast(valueItemForm);
		valueLineSchema
			.validate(valueItemFormParsed, { abortEarly: false })
			.then(() => {
				updateProjectValueLine(valueItemFormParsed).then((res) => {
					if (res && res.status?.success) {
						addToast({
							message: 'Project features updated successfully',
							dismissible: true,
							type: 'success'
						});

						valueItemForm = valueLineDefaultForm(projectID);
						callIf(update);
					} else {
						addToast({
							message: 'Error updating project features: ' + res.status?.message,
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

    const loadPage = async () => {
        getList('Funding Source')
            .then((l) => {
                fundingSourceOpts = findSelectOptsFromList(l);
            })
            .then(() => {
                getList('Value Catetgory').then(l => {
                    valueCategoryOpts = findSelectOptsFromList(l);
                });
            });
	};

    let valueItemForm = $state(valueLineDefaultForm(projectID));
</script>


{#await loadPage()}
    Loading...
{:then promiseData}

{#if valueItemForm}
<SelectInput
    bind:value={valueItemForm.fundingSource as string}
    fieldName="Funding Source"
    error={errors.fundingSource}
    options={fundingSourceOpts}
    update={() => callIf(update)}
/>


<SelectInput
    bind:value={valueItemForm.valueCategory as string}
    fieldName="Value Category"
    error={errors.valueCategory}
    options={valueCategoryOpts}
    update={() => callIf(update)}
/>

<Heading tag="h6">Five Year Forecast</Heading>

<MoneyInput
    bind:value={valueItemForm.yearOneValue as number}
    fieldName="Estimated Year One Returns"
    error={errors.yearOneValue}
    update={() => callIf(update)}
/>

<MoneyInput
    bind:value={valueItemForm.yearTwoValue as number}
    fieldName="Estimated Year Two Returns"
    error={errors.yearTwoValue}
    update={() => callIf(update)}
/>

<MoneyInput
    bind:value={valueItemForm.yearThreeValue as number}
    fieldName="Estimated Year Tnree Returns"
    error={errors.yearThreeValue}
    update={() => callIf(update)}
/>

<MoneyInput
    bind:value={valueItemForm.yearFourValue as number}
    fieldName="Estimated Year Four Returns"
    error={errors.yearFourValue}
    update={() => callIf(update)}
/>

<MoneyInput
    bind:value={valueItemForm.yearFiveValue as number}
    fieldName="Estimated Year Five Returns"
    error={errors.yearFiveValue}
    update={() => callIf(update)}
/>

<div class="col-span-4">
    <span class="float-right">
        <Button color="primary" onclick={submitForm}>Update Value</Button>
    </span>
    <br class="clear-both" />
</div>

{/if}
    
{/await}
