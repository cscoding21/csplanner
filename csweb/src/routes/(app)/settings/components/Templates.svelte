<script lang="ts">
    import { type UpdateProjecttemplate, type Projecttemplate, type ProjecttemplateResults, DeleteProjectTemplateDocument } from "$lib/graphql/generated/sdk";
    import { addToast } from "$lib/stores/toasts";
	import { Sidebar, SidebarWrapper, SidebarGroup, SidebarItem, Button } from "flowbite-svelte";
	import { SectionSubHeading } from "$lib/components";
	import { deepCopy } from "$lib/utils/helpers";
    import { TemplateForm } from ".";
	import { deleteProjectTemplate, findAllProjectTemplates } from "$lib/services/template";
	import { DeleteTableOutline, PlusOutline, TrashBinOutline } from "flowbite-svelte-icons";

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
		<Sidebar {activeUrl}>
			<SidebarWrapper>
			  <SidebarGroup>
				{#each templates as template, index}
					<SidebarItem onclick={() => setTemplate(index)} label={template.name}>
                    </SidebarItem>
				{/each}
			  </SidebarGroup>
			</SidebarWrapper>
		</Sidebar>
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