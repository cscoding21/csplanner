<script lang="ts">
	import { SectionHeading } from "$lib/components";
	import SectionSubHeading from "$lib/components/formatting/SectionSubHeading.svelte";
	import { Badge, Button, Input, P, Popover } from "flowbite-svelte";
	import { CloseCircleSolid } from "flowbite-svelte-icons";

    interface Props {
        onDone: Function
    }
    let { onDone }:Props = $props()
    let skillList:string[] = $state([])
    let newSkill:string = $state("")

    const addSkill = () => {
        if (!newSkill) {
            return
        }

        skillList = [...new Set([...skillList, newSkill])]
        newSkill = ""
    }

    const addSkillGroup = (sg:string[]) => {
        skillList = [...new Set([...skillList, ...sg])]
    }

    const removeSkill = (i:number) => {
        skillList.splice(i, 1)
        skillList = skillList
    }

    const checkEnter = ((event:KeyboardEvent) => {
        if(event.key === "Enter") {
            addSkill()
        }
    })

    const saasSkillGroup = [
        "Backend Development",
        "Frontend Development",
        "Data Architecture",
        "DevOps / SRE",
        "Data Science"
    ]

    const projectSkillGroup = [
        "Project Management",
        "Business Analysis",
        "Agile Development",
        "Scrum",
        "Kanban"
    ]

    const autonomySkillGroup = [
        "State Estimation",
        "Scene Understanding",
        "C++",
        "Mapping",
        "Foundations"
    ]

</script>


<div class="p-4">
<div class="grid grid-cols-2 gap-4">
    <div class="">
        <h2 class="text-xl text-center">Skills</h2>
    
        <section>
    <P class="p-4">
        Skills define the discrete capabilities needed to execute your organization's projects.
        The building block of a project is a task which maps to a single skill.  Allocating resources to 
        a task requires a matching skill by the resource.        
    </P>
    <P class="p-4">
        Each organization can decide on the fidelity of their skill stacks.  For instance, a software development
        department may have a skill of <i>Backend Development</i> to encompass a large area of expertise, or break that
        down into smaller skills (e.g. Golang, Rust, C#, etc). 
    </P>
    <P class="p-4">Skills can be added or modified at any time on the <a class="text-orange-300" href="/settings#lists">List</a> page.  To get started, we'll geet you set up with a broadly scoped
        set of skills.
    </P>
</section>

    </div>
    <div class="">

        <SectionSubHeading>Add Your Skills</SectionSubHeading>
        <div class="mb-4">
            <Input bind:value={newSkill} onclick={addSkill} onkeypress={checkEnter} placeholder="Type a skill name" />
        </div>

        <SectionSubHeading>Common Skill Groups</SectionSubHeading>
        <div class="mb-4">
            <Button id="saasSkillButton" class="m-2" color="alternative" pill onclick={() => addSkillGroup(saasSkillGroup)}>SaaS</Button>
            <Popover triggeredBy="#saasSkillButton" class="w-72 text-sm font-light text-gray-500 bg-white dark:bg-gray-800 dark:border-gray-600 dark:text-gray-400">
                <div class="p-3 space-y-2">
                  Add the following skills to your stack:
                  <ul class="list-disc pl-4">
                    {#each saasSkillGroup as sk}
                        <li>{sk}</li>
                    {/each}
                  </ul>
                </div>
            </Popover>


            <Button id="projectSkillButton" class="m-2" color="alternative" pill onclick={() => addSkillGroup(projectSkillGroup)}>Project Management</Button>
            <Popover triggeredBy="#projectSkillButton" class="w-72 text-sm font-light text-gray-500 bg-white dark:bg-gray-800 dark:border-gray-600 dark:text-gray-400">
                <div class="p-3 space-y-2">
                  Add the following skills to your stack:
                  <ul class="list-disc pl-4">
                    {#each projectSkillGroup as sk}
                        <li>{sk}</li>
                    {/each}
                  </ul>
                </div>
            </Popover>

            <Button id="autonomySkillGroup" class="m-2" color="alternative" pill onclick={() => addSkillGroup(autonomySkillGroup)}>Autonomy</Button>
            <Popover triggeredBy="#autonomySkillGroup" class="w-72 text-sm font-light text-gray-500 bg-white dark:bg-gray-800 dark:border-gray-600 dark:text-gray-400">
                <div class="p-3 space-y-2">
                  Add the following skills to your stack:
                  <ul class="list-disc pl-4">
                    {#each autonomySkillGroup as sk}
                        <li>{sk}</li>
                    {/each}
                  </ul>
                </div>
            </Popover>

        </div>

        <SectionSubHeading>Your Skills</SectionSubHeading>
        <div class="p-4">
            {#each skillList as skill, index}
                <Badge class="m-2" dismissable>
                    {skill}
                <button slot="close-button" onclick={() => removeSkill(index)} type="button" class="inline-flex items-center rounded-full p-0.5 my-0.5 ms-1.5 -me-1.5 text-sm text-white dark:text-primary-80 hover:text-whit dark:hover:text-white" aria-label="Remove">
                    <CloseCircleSolid class="h-4 w-4" />
                    <span class="sr-only">Remove skill</span>
                </button>
                </Badge>
            {/each}
        </div>

        <Button onclick={() => onDone()}>I'm done for now.  Let's keep going! >></Button>
    </div>
</div>
</div>