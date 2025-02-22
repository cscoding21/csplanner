<script lang="ts">
	import { updateList } from "$lib/services/list";
    import type { ListItem, UpdateList, Projecttemplate, ProjecttemplateResults } from "$lib/graphql/generated/sdk";
    import { addToast } from "$lib/stores/toasts";
	import { Sidebar, SidebarWrapper, SidebarGroup, SidebarItem, Button } from "flowbite-svelte";
	import { SectionSubHeading } from "$lib/components";
	import { deepCopy } from "$lib/utils/helpers";
	import { nameToID } from "$lib/utils/id";
	import { findAllProjectTemplates } from "$lib/services/project";
    import { TemplateForm } from ".";

    let templates = $state([] as Projecttemplate[])
	let currentTemplate = $state({} as UpdateList)
	let currentTemplateName = $state("")
	let newTemplate = $state("")
	let errors = $state({newListItem: ""})
	

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

	const setTemplate = (index:number) => {
		currentTemplateName = templates[index].name
		currentTemplate = deepCopy(templates[index])
	}

	const saveList = () => {
		// @ts-ignore
		delete currentList.name

		updateList(currentTemplate).then(res => {
			if (res && res.status?.success) {
				addToast({
					message: 'Project template updated successfully',
					dismissible: true,
					type: 'success'
				});

				load();
			} else {
				addToast({
					message: 'Error updating project template: ' + res.status?.message,
					dismissible: true,
					type: 'error'
				});
			}
		});
	}

	const addListItem = (item:string) => {
		let li:ListItem = {
			name: item,
			value: nameToID(item),
			sortOrder: 0
		}

		currentTemplate.values.push(li)
		newTemplate = ""
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
		<SectionSubHeading>Templates</SectionSubHeading>
		<Sidebar>
			<SidebarWrapper>
			  <SidebarGroup>
				{#each templates as template, index}
					<SidebarItem onclick={() => setTemplate(index)} label={template.name}></SidebarItem>
				{/each}
			  </SidebarGroup>
			</SidebarWrapper>
		</Sidebar>
	</div>

	<div class="flex-1 px-2">
		<SectionSubHeading>Project Template: {currentTemplateName}</SectionSubHeading>
		<TemplateForm template={currentTemplate} />
		
		<Button onclick={saveList}>Save Template</Button>
	</div>

</div>
{/await}