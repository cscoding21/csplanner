<script lang="ts">
    import { onMount } from "svelte"
    import Quill from 'quill';
    import type { Delta } from "quill/core";
    import { normalizeID } from "$lib/utils/id";

    import 'quill/dist/quill.core.css'
    import 'quill/dist/quill.snow.css'
    import 'quill/dist/quill.bubble.css'
    import "quill-mention/autoregister";

    import { handleFileUpload } from "$lib/services/file";
	

    interface Props {
        contents:any;
        attachContext: string;
        quillEditor: any;
    }
    let { 
        contents = $bindable(), 
        attachContext, 
        quillEditor = $bindable()
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
        //mention: mentionOpts
        quill = new Quill("#" + editorID, {
            theme: "bubble",
            modules: {
                toolbar: toolbarOptions,
                
            }
        });

        if (contents) {
            console.log("content", contents)
            quill.setContents(JSON.parse(contents), "user");
        }

        quill.on('text-change', (delta, oldDelta, source) => {
            contents = quill.getContents()
        });

        quill.enable(true)
        quill.focus()
    });

    let atValues = [
        { id: 'jdoe@its-global.vn', value: 'John Doe' },
        { id: 'jadoe@its-global.vn', value: 'Jane Doe' },
        { id: 'jdmith@its-global.vn', value: 'John Smith' },
        { id: 'jasmith@its-global.vn', value: 'Jane Smith' },
    ]

    const mentionOpts = {
            allowedChars: /^[A-Za-z\sÅÄÖåäö]*$/,
            mentionDenotationChars: ["@"],
            mentionContainerClass: "bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-primary-500 focus:border-primary-500 block w-1/3 p-2.5 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-primary-500 dark:focus:border-primary-500",
            listItemClass: "block mb-2 text-sm font-medium text-gray-900 dark:text-white cursor-pointer",
            mentionListClass: "block mb-2 text-sm font-medium text-gray-900 dark:text-white cursor-pointer",
            minChars: 1,
            selectKeys: [9],
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
</script>


<!-- bg-gray-50 border border-gray-200 dark:bg-gray-700 dark:border-gray-600 -->
<div
    class="w-full rounded-lg text-gray-100">
    <div id="{editorID}"></div>
</div>