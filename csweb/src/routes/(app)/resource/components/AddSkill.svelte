<script lang="ts">
	import { SelectInput } from "$lib/components";
    import { Button } from "flowbite-svelte";
    import { createEventDispatcher, onMount } from "svelte";
    import { skillForm, skillSchema } from "$lib/forms/skill.validation";
    import { mergeErrors, parseErrors } from "$lib/forms/helpers";
    import { updateResourceSkill } from "$lib/services/resource";
	import { refreshListStore, skillListStore } from "$lib/stores/list";
    import { addToast } from "$lib/stores/toasts";

    export let resourceID:string
    let errors: any = {}

    const dispatch = createEventDispatcher();

    function updated() {
        dispatch('updated', {
            text: 'Skill was updated'
        });
    }

    onMount(async () => {  
        refreshListStore()
    })

    const add = () => {
        errors = {}

        const parsedSkillForm = skillSchema.cast(skillForm)
        skillSchema.validate(parsedSkillForm, { abortEarly: false })
			.then(() => {
                updateResourceSkill({ resourceID: parsedSkillForm.resourceID, id: parsedSkillForm.skillID, proficiency: parsedSkillForm.proficiency as number })
                    .then(res => {
                        if(res.data.updateResourceSkill.success) {
                            addToast({ 
									message: "Skill successfully added to resource", 
									dismissible: true, 
									type: "success"}
								)

                            skillForm.id = ""
                            skillForm.proficiency = ""
                            updated()
                        } else {
                            addToast({ 
								message: "Error creating resource: " + res.data.updateResourceSkill.message, 
								dismissible: true, 
								type: "error"}
							)
                        }
                    })
			})
			.catch((err) => {
				errors = mergeErrors(errors, parseErrors(err))

				console.log(errors);
			});
    }
    
    const proficiencyOptions = [
        { value: 1, name: "Novice" },
        { value: 2, name: "Competent" },
        { value: 3, name: "Expert" },
    ]

    skillForm.resourceID = resourceID;

    $: availableSkills = $skillListStore.values; 
</script>


{#await availableSkills}
    <span>Loading...</span>
{:then data } 
<div class="grid grid-cols-5 gap-5">
    <span class="col-span-2">
        <SelectInput 
            fieldName="Skill"
            bind:value={skillForm.skillID} 
            error={errors.skillID}
            options={availableSkills} />
    </span>
    <span class="col-span-2">
        <SelectInput 
            fieldName="Proficiency"
            bind:value={skillForm.proficiency} 
            error={errors.proficiency}
            options={proficiencyOptions} />
    </span>
    <span class="w-1/5 pt-8">
        <Button on:click={add}> Add </Button>
    </span>
</div>
{/await}