<script lang="ts">
    import { Avatar, Rating, Card, Alert, Table, TableBody, TableBodyRow, TableHead, TableHeadCell, TableBodyCell, Timeline, TimelineItem, Button, ButtonGroup, Dropzone, Modal, Hr } from "flowbite-svelte";
    import { TrashBinOutline, PenOutline } from "flowbite-svelte-icons";
    import { page } from "$app/stores";
    import { PageHeading, TextInput } from "$lib/components";
	import { ResourceActionBar, AddSkill, DeleteResource } from "../../components";
    import { getResource, deleteResourceSkill, updateResource } from "$lib/services/resource";
	import { formatDate } from "$lib/utils/format";
    import { getInitialsFromName } from "$lib/utils/format";
    import { resourceForm, resourceSchema } from "$lib/forms/resource.validation";
    import { handleFileUpload } from "$lib/services/file";
    import { coalesceToType, mergeErrors, parseErrors } from "$lib/forms/helpers";
    import type { UpdateResource, Resource } from "$lib/graphql/generated/sdk";
    import { addToast } from "$lib/stores/toasts";
	import { onMount } from "svelte";
	import { deepCopy } from "$lib/utils/helpers";
    
    const id = $page.params.id;

    let editModalOpen:boolean = $state(false)
    let errors = $state({name: "", role: ""})
    let rf = $state(deepCopy(resourceForm))

    const refresh = async () => {
        resourcePromise = getResource(id).then(r => {
            rf = coalesceToType<UpdateResource>(r, resourceSchema)

            return r
        })

        return resourcePromise
    }

    const deleteSkill = async (skillID:string) => {
        deleteResourceSkill(id, skillID).then(res => {
            if(res.success) {
                addToast({ 
                    message: "Resource skill deleted successfully", 
                    dismissible: true, 
                    type: "success"}
                )

                resourcePromise = refresh()
            } else {
                addToast({ 
                    message: "Error updating resource: " + res.message, 
                    dismissible: true, 
                    type: "error"}
                )
            }
        })
    }

    const updateRes = async () => {
        const resourceSchemaParsed = resourceSchema.cast(rf)
        resourceSchema.validate(resourceSchemaParsed, { abortEarly: false })
			.then((f) => {
				console.log('validation passed')

                updateResource(f as UpdateResource)
                    .then((res) => {
                        if(res.status?.success.valueOf()) {
                            editModalOpen = false

                            addToast({ 
                                message: "Resource " + rf.name + " updated successfully", 
                                dismissible: true, 
                                type: "success"}
                            )

                            refresh()
                        } else {
                            addToast({ 
								message: "Error updating resource: " + res.status?.message, 
								dismissible: true, 
								type: "error"}
							)
                        }
                    })
			})
			.catch((err) => {
				errors = mergeErrors(errors, parseErrors(err))

				console.error(errors);
			});
    }

    let fileValue = [];
    const handleChange = (e:any) => {
        const files = e.target.files;
        if (files.length > 0) {
            const file = files[0];

            handleFileUpload(file, id, "profile_image", handleUploadComplete)

            fileValue.push(file.name);
        }
    };

    const handleUploadComplete = (upload:any) => {
        console.log("upload complete ", upload.file.name)

        rf.profileImage = upload.url
    }

    let resourcePromise:Promise<Resource> = $state(refresh())

    onMount(async () => {
        refresh()
    })
</script>


{#await resourcePromise }
    <div>Loading...</div>
{:then promiseData}
{#if promiseData}
<ResourceActionBar pageDetail={promiseData.name}>
    {#if !promiseData.isBot}
    <ButtonGroup>
        <DeleteResource id={promiseData.id || ""} name={promiseData.name}>
            <TrashBinOutline class="w-3 h-3 mr-2" />
            Delete
        </DeleteResource>
    </ButtonGroup>
    {/if}
</ResourceActionBar>


<PageHeading title={promiseData.name} />

<div class="grid grid-cols-2">

<div class="mr-4">
<Card padding="sm" size="xl">
    <div class="flex justify-end">
        {#if !promiseData.isBot}
        <button onclick={() => { editModalOpen =  true }}><PenOutline /></button>
        {/if}
    </div>
    <div class="flex flex-col items-center pb-4">
      <Avatar size="lg" src={promiseData.profileImage as string}  rounded >{getInitialsFromName(promiseData.name)}</Avatar>
      <h5 class="mb-1 text-xl font-medium text-gray-900 dark:text-white">{promiseData.name}</h5>
      <span class="text-sm text-gray-500 dark:text-gray-400">{promiseData.role}</span>
    </div>
    <hr class="mt-2 mb-4" />
    <div>
        <ul class="list">
            <li>
                <span>Role</span>
                <span class="flex-auto float-right font-semibold">{promiseData.role}</span>
            </li>
            <li>
                <span>User</span>
                {#if promiseData.user?.id}
                <span class="flex-auto float-right font-semibold">{promiseData.user?.email}</span>
                {:else}
                <span class="flex-auto float-right font-semibold">No User Account</span>
                {/if}
            </li>
            <li>
                <span>Created Date</span>
                <span class="flex-auto float-right font-semibold">{formatDate(promiseData.createdAt)}</span>
            </li>
        </ul>
    </div>

    <caption class="mt-4 text-lg font-semibold text-left text-gray-900 bg-white dark:text-white dark:bg-gray-800">
    User skills
    <p class="mt-1 mb-4 text-sm font-normal text-gray-500 dark:text-gray-400">Identifying user skills allows <b>analyzer</b> to properly estimate and propose project resourcing strategies.</p>
    </caption>

    
    {#if promiseData.skills && promiseData.skills.length > 0 }
    <Table>
        <TableHead>
            <TableHeadCell>Skill</TableHeadCell>
            <TableHeadCell>Proficienty</TableHeadCell>
            <TableHeadCell>
                <span class="sr-only">Action</span>
            </TableHeadCell>
        </TableHead>
        <TableBody>
    {#each promiseData.skills as s(s)}
        <TableBodyRow>
            <TableBodyCell>{s.name}</TableBodyCell>
            <TableBodyCell><Rating rating={s.proficiency?.valueOf()} total={3} /></TableBodyCell>
            <TableBodyCell tdClass="float-right pt-2">
                <Button outline color="dark" on:click={() => deleteSkill(s.id)}>
                    <TrashBinOutline size="xs" />
                </Button>
            </TableBodyCell>
        </TableBodyRow>
    {/each}
</TableBody>
</Table>
    {:else}
        <Alert>No skills have been added for this resource.</Alert>
    {/if}
    
    {#if !promiseData.isBot}
    <hr class="my-4" />
    <AddSkill resourceID={id} on:updated={refresh} />
    {/if}
  </Card>
</div>

<div>
    <Card size="lg">
        <!-- <TimelineChart  height="350"/> -->
        <!-- <ResourceAllocation resourceId={id} /> -->
    </Card>
</div>
</div>

<Modal title="Update {promiseData.name}" bind:open={editModalOpen}>
    <div class="flex items-center space-x-4 rtl:space-x-reverse">
        <div>
            <Dropzone
                    id="dropzone"
                    defaultClass="float-left mr-4"
                    on:success={(e) => console.log(e)}
                    on:change={handleChange}>
            <Avatar size="xl" src={rf.profileImage || ""} rounded class="hover:cursor-pointer">{getInitialsFromName(promiseData.name)}</Avatar>
            </Dropzone>
            <div class="ml-4 font-medium dark:text-white">
                Click profile image to re-upload 
            </div>
        </div>
    </div>

    <Hr />
    <TextInput bind:value={rf.name} 
        fieldName="Name" 
        placeholder="Resource name"
        error={errors.name} />

    <TextInput bind:value={rf.role} 
        fieldName="Role" 
        placeholder="Resource role"
        error={errors.role} />

    <svelte:fragment slot="footer">
        <div class="place-content-end">
        <Button on:click={updateRes}> Update Resource </Button>
        </div>  
    </svelte:fragment>
  </Modal>
  {/if}
{/await}