<script lang="ts">
	import { TextInput, SectionHeading } from "$lib/components";
    import type { UpdateUser, User } from "$lib/graphql/generated/sdk";
    import { authService } from "$lib/services/auth";
    import { Dropzone, Avatar, Hr, Button } from "flowbite-svelte";
    import { getInitialsFromName } from "$lib/utils/format";
    import { handleFileUpload } from "$lib/services/file";
	import { updateUser } from "$lib/services/user";
    import { userFormSchema } from "$lib/forms/uservalidation"
    import { deepCopy } from "$lib/utils/helpers";
    import { addToast } from "$lib/stores/toasts";
    import { mergeErrors, parseErrors } from "$lib/forms/helpers";

    let errors = $state({ firstName: "", lastName: ""})
    let uf = $state({} as UpdateUser)

    let fileValue = [];
	const handleChange = (e: any) => {
		const files = e.target.files;
		if (files.length > 0) {
			const file = files[0];
			console.log('handleChange', file);

			handleFileUpload(file, uf?.id as string, 'user_image', handleUploadComplete);

			fileValue.push(file.name);
		}
	};

	const handleUploadComplete = (upload: any) => {
		console.log('upload complete ', upload.file.name);
        console.log("upload", upload)

		uf.profileImage = upload.url;
	};

    const saveUser = () => {
        const userFormParsed = userFormSchema.cast(uf);

        userFormSchema
			.validate(userFormParsed, { abortEarly: false })
			.then(() => {
                if(!userFormParsed.profileImage) {
                    userFormParsed.profileImage = ''
                } 
				updateUser(userFormParsed).then((res) => {
					if (res && res.status?.success) {
						addToast({
							message: 'User info updated successfully',
							dismissible: true,
							type: 'success'
						});

						uf = deepCopy(res.user);
					} else {
						addToast({
							message: 'Error updating user info: ' + res.status?.message,
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
    }

    const loadPage = async () => {
        const as = authService()
        const user = await as.currentUser() as User

        uf = {
            firstName: user?.firstName,
            lastName: user?.lastName,
            id: user?.id,
            email: user?.email,
            profileImage: user?.profileImage
        }
    }
</script>


<SectionHeading>My Info</SectionHeading>

{#await loadPage()}
    Loading...
{:then promiseData} 

<div class="flex">
    <div class="flex-none pr-6 items-center space-x-4 rtl:space-x-reverse">
        <div>
            <Dropzone id="dropzone" defaultClass="" on:change={handleChange}>
                <Avatar size="xl" src={uf.profileImage || ''} rounded class="hover:cursor-pointer"
                    >{getInitialsFromName(uf.firstName + " " + uf.lastName)}</Avatar
                >
            </Dropzone>
        </div>
        <small class="text-sm dark:text-white mt-2">Click to re-upload</small>
    </div>
    <div class="flex-1">
        <h3 class="mb-4">
        Email: <b>{uf.email}</b>
        </h3>

        <h3 class="mb-4">
        User ID: <b>{uf.id}</b>
        </h3>
    </div>
</div>

<Hr />

<TextInput
    bind:value={uf.firstName as string}
    placeholder="First name"
    fieldName="First Name"
    error={errors.firstName}
/>

<TextInput
    bind:value={uf.lastName as string}
    placeholder="Las name"
    fieldName="Last Name"
    error={errors.lastName}
/>

<div class="col-span-4">
    <span class="float-right">
        <Button onclick={saveUser}>
            Update User
        </Button>
    </span>
</div>
    
{/await}
