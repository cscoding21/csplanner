<script lang="ts">
    import { onMount } from "svelte"
    import Quill from 'quill';
    import { normalizeID } from "$lib/utils/id";

    import 'quill/dist/quill.core.css'
    import 'quill/dist/quill.snow.css'
    import 'quill/dist/quill.bubble.css'
    import "quill-mention/autoregister";

    interface Props {
        contents:any;
        attachContext: string;
        fieldName?: string;
    }
    let { 
        contents = $bindable(), 
        attachContext, 
        fieldName, 
    }:Props = $props();

    let quill : Quill;

    const editorID = "display_" + normalizeID(attachContext)
	
    onMount(async () => {
        quill = new Quill("#" + editorID, {
            theme: "bubble",
            modules: {}
        });

        if (contents) {
            quill.setContents(JSON.parse(contents), "user");
        }

        quill.enable(false)
        quill.focus()
    });
</script>



<div class="my-1">
    {#if fieldName}
	    <label for="input" class="block text-sm font-medium text-gray-900 dark:text-white">{fieldName}</label>
    {/if}

    <div
        class="w-full rounded-lg text-gray-100">
        <div id="{editorID}"></div>
    </div>
</div>