<script lang="ts">
	import { getInitialsFromName } from "$lib/utils/format";
	import { Avatar } from "flowbite-svelte";


    interface Props {
        onchange: Function
        avatar?: any
        name: string
    }
    let { onchange, avatar, name }:Props = $props()

    const onFileSelected = (e:any) => {
        if(!e || !e.target || !e.target.files) {
            return 
        }

        let i = e.target.files[0];
        let reader = new FileReader();
        reader.readAsDataURL(i);

        reader.onload = e => {
            // @ts-expect-error
            avatar = e.target.result
        };

        onchange(e)
}

    let fileinput:HTMLElement|undefined = $state();
</script>


<button onclick={() => { fileinput && fileinput.click();}}>
    <Avatar size="xl" src={avatar} class="hover:cursor-pointer">{getInitialsFromName(name)}</Avatar>
</button> 

<input style="display:none" type="file" accept=".jpg, .jpeg, .png" onchange={(e)=>onFileSelected(e)} bind:this={fileinput} >

{#if avatar}
    <div class="text-sm dark:text-gray-300 px-4 mt-4">Click to re-upload</div>
{:else}
    <div class="text-sm dark:text-gray-300 px-4 mt-4">Click to upload</div>
{/if}
