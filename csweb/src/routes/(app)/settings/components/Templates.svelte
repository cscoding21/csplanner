<script lang="ts">
    import { type UpdateProjecttemplate, type Projecttemplate, type ProjecttemplateResults, DeleteProjectTemplateDocument } from "$lib/graphql/generated/sdk";
    import { addToast } from "$lib/stores/toasts";
	import { CSSidebar, CSSidebarItem, SectionSubHeading } from "$lib/components";
	import { deepCopy } from "$lib/utils/helpers";
    import { TemplateForm } from ".";
	import { deleteProjectTemplate, findAllProjectTemplates } from "$lib/services/template";
	import { PlusOutline, TrashBinOutline } from "flowbite-svelte-icons";

    let templates = $state([] as Projecttemplate[])
	let currentTemplate = $state({} as UpdateProjecttemplate)
	let currentTemplateName = $state("")
    let activeUrl:string = $state("")

	const load = async ():Promise<ProjecttemplateResults> => {
		return await findAllProjectTemplates()
			.then(l => {
				return l
			})
			.catch((err) => {
				addToast({
					message: 'Error loading project templates (Templates): ' + err,
					dismissible: true,
					type: 'error'
				});

				return err
			});
	};

    const addNewTemplate = () => {
        templates = [...templates, {
            id: "",
            name: "New template",
            description: '',
            phases: [],
        }]
    }

    const deleteTemp = async (id:string) => {
        return await deleteProjectTemplate(id)
        .then((res) => {
            if (res.success) {
                addToast({
                    message: 'Project template deleted successfully',
                    dismissible: true,
                    type: 'success'
                });

                load()
            } else {
                addToast({
                    message: 'Error deleting project template: ' + res,
                    dismissible: true,
                    type: 'error'
                });
            }
        })
        .catch((err) => {
            addToast({
                message: 'Error deleting project template: ' + err,
                dismissible: true,
                type: 'error'
            });
        });
    }

	const setTemplate = (index:number) => {
		currentTemplateName = templates[index].name
        activeUrl = templates[index].name
		currentTemplate = deepCopy(templates[index])
	}

	const loadPage = async () => {
        load().then(l => {
            templates = l.results as Projecttemplate[]

			setTemplate(0)
        });
	};
</script>


{#await loadPage()}
    Loading...
{:then promiseData} 
<div class="flex">
	<div class="flex-none px-2">
		<SectionSubHeading>
            Templates
            <span class="float-right">
                <button onclick={() => addNewTemplate()}>
                    <PlusOutline size="sm" />
                </button>
            </span>
        </SectionSubHeading>

        <div class="gap-8 lg:flex">
        <CSSidebar>
            {#each templates as template, index}
            <CSSidebarItem label={template.name} onclick={() => setTemplate(index)} active={currentTemplate.name === template.name}>
                <!-- {#snippet icon()}
                <svg class="h-6 w-6 text-gray-400 transition duration-75 group-hover:text-gray-900 dark:text-gray-400 dark:group-hover:text-white" aria-hidden="true" xmlns="http://www.w3.org/2000/svg" width="24" height="24" fill="none" viewBox="0 0 24 24">
                    <path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 7h6l2 4m-8-4v8m0-8V6a1 1 0 0 0-1-1H4a1 1 0 0 0-1 1v9h2m8 0H9m4 0h2m4 0h2v-4m0 0h-5m3.5 5.5a2.5 2.5 0 1 1-5 0 2.5 2.5 0 0 1 5 0Zm-10 0a2.5 2.5 0 1 1-5 0 2.5 2.5 0 0 1 5 0Z" />
                  </svg>    
                {/snippet} -->
            </CSSidebarItem>
            {/each}
        </CSSidebar>
        </div>
	</div>

	<div class="flex-1 px-2">
		<SectionSubHeading>
            Project Template: {currentTemplate.name}
            <span class="float-right">
                <button onclick={() => deleteTemp(currentTemplate.id as string)}>
                    <TrashBinOutline size="sm" />
                </button>
            </span>
        </SectionSubHeading>
		<TemplateForm template={currentTemplate as UpdateProjecttemplate} />
	</div>

</div>
{/await}