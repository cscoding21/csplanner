<script lang="ts">
    import { onMount } from "svelte"
    import Quill from 'quill';
    import type { Delta } from "quill/core";
    import type { User } from "$lib/graphql/generated/sdk";
    import { normalizeID } from "$lib/utils/id";

    import 'quill/dist/quill.core.css'
    import 'quill/dist/quill.snow.css'
    import 'quill/dist/quill.bubble.css'
    import "quill-mention/autoregister";

    import { handleFileUpload } from "$lib/services/file";
	import { findAllUsers } from "$lib/services/user";
	
    interface Props {
        contents:any;
        attachContext: string;
        quillEditor: any;
        userOpts?: User[]
    }
    let { 
        contents = $bindable(), 
        attachContext, 
        quillEditor = $bindable(),
        userOpts
    }:Props = $props();

    let quill : Quill;

    const editorID = "editor_" + normalizeID(attachContext)

    quillEditor = {
        clear: () => {
            quill.setContents([{ "insert": "\n" }]);
        },
        setContent: (content:Delta) => {
            quill.setContents(content);
        }
    }

    const toolbarOptions = [
        ['bold', 'italic', 'underline', 'strike'],        // toggled buttons
        ['blockquote'],

        [{ 'list': 'ordered'}, { 'list': 'bullet' }],
        [{ 'indent': '-1'}, { 'indent': '+1' }],          // outdent/indent

        [{ 'header': [1, 2, 3, 4, false] }],

        [{ 'color': [] }, { 'background': [] }],          // dropdown with defaults from theme
        [{ 'align': [] }],

        ["link", "image", "video"],

        ['clean']                                         // remove formatting button
    ]
	
    onMount(async () => {
        // debug: 'log',
        
        quill = new Quill("#" + editorID, {
            theme: "bubble",
            modules: {
                mention: mentionOpts,
                toolbar: toolbarOptions,
                
            }
        });

        if (contents) {
            console.log("content", contents.data)
            quill.setContents(JSON.parse(contents), "user");
        }

        quill.on('text-change', (delta, oldDelta, source) => {
            contents = quill.getContents()
        });

        quill.enable(true)
        quill.focus()
    });

    let atValues = $state([{id: "", value: ""}])
    const mentionOpts = {
            allowedChars: /^[A-Za-z\sÅÄÖåäö]*$/,
            mentionDenotationChars: ["@"],
            minChars: 1,
            source: function(searchTerm:any, renderList:any, mentionChar:any) {
                let values = atValues;

                if (searchTerm.length === 0) {
                    renderList(values, searchTerm);
                } else {
                    const matches = [];
                    for (let i = 0; i < values.length; i++) {
                        if (
                            ~values[i].value
                            .toLowerCase()
                            .indexOf(searchTerm.toLowerCase())
                        ) {
                            matches.push(values[i]);
                        }
                    }

                    renderList(matches, searchTerm);
                }
            }
        }

    const loadPage = () => {
        findAllUsers().then(u => {
            atValues = u.results?.map(u => { return  { id: u.email, value: u.firstName + ' ' + u.lastName } } ) as []
        })
    }
</script>


{#await loadPage()}
    Loading...
{:then promiseData} 
<div
    class="w-full rounded-lg text-gray-100">
    <div id="{editorID}"></div>
</div>

{/await}
