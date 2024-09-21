<script lang="ts">
    import { onMount } from "svelte"
    import { Button, Dropzone } from "flowbite-svelte";
    import Quill from 'quill';
    import 'quill-mention'

    import 'quill/dist/quill.core.css'
    import 'quill/dist/quill.snow.css'
    import 'quill/dist/quill.bubble.css'
    import "quill-mention/autoregister";

    //import { ImageHandler, VideoHandler, AttachmentHandler } from "quill-upload";
    import MagicUrl from 'quill-magic-url'
    //import { userStore, refreshUserStore } from "$lib/stores/user";
    import { handleFileUpload } from "$lib/services/file";


    interface Props {
        post: Function;
        contents:any;
        attachContext: string;
    }
    let { post, contents, attachContext }:Props = $props();

    let quill : Quill;
    let fileValue:string[] = $state([]);

    // register quill-upload
    //Quill.register("modules/imageHandler", ImageHandler);
    //Quill.register("modules/videoHandler", VideoHandler);
    //Quill.register("modules/attachmentHandler", AttachmentHandler);
    Quill.register('modules/magicUrl', MagicUrl)
	
    onMount(async () => {
        //await refreshUserStore()
        
        // let atValues = $userStore.map(u => { 
        //     let o = { id: u.id, value: u.name };

        //     return o;
        // })

        let atValues = [
            { id: 'jdoe@its-global.vn', value: 'John Doe' },
            { id: 'jadoe@its-global.vn', value: 'Jane Doe' },
            { id: 'jdmith@its-global.vn', value: 'John Smith' },
            { id: 'jasmith@its-global.vn', value: 'Jane Smith' },
        ]

        const _onUpload = function (fd:any, resolve:any) {
            console.log("onUpload")
            console.log(fd)
            /*
            const xhr = new XMLHttpRequest();
            xhr.open("POST", "https://upload.imagekit.io/api/v1/files/upload");
            xhr.setRequestHeader(
                "Authorization",
                "Basic cHJpdmF0ZV9LKzNFRGJnMXRQOXBsejlvOGhkd1J0bkZ0bjQ9Og=="
            );
            xhr.onload = () => {
                if (xhr.status === 200) {
                const response = JSON.parse(xhr.responseText);

                resolve(response.url); // Must resolve as a link to the image
                }
            };

            xhr.send(fd);
            */
        };

        const _onAttachment = function (fd:any, resolve:any) {
            console.log("onAttachment")
            console.log(fd)
            /*
            const xhr = new XMLHttpRequest();
            xhr.open("POST", "https://api.pdf.co/v1/file/upload");
            xhr.setRequestHeader(
                "x-api-key",
                "tainv@its-global.vn_f6eab0d85bf5b00a8df529806894afb788aa"
            );
            xhr.onload = () => {
                if (xhr.status === 200) {
                const response = JSON.parse(xhr.responseText);

                resolve(response.url); // Must resolve as a link to the image
                }
            };

            xhr.send(fd);
            */
        };

        let toolbarOptions = [
            ['bold', 'italic', 'underline', 'strike'],        // toggled buttons
            ['blockquote'],

            [{ 'list': 'ordered'}, { 'list': 'bullet' }],
            [{ 'indent': '-1'}, { 'indent': '+1' }],          // outdent/indent

            [{ 'header': [1, 2, 3, 4, 5, 6, false] }],

            [{ 'color': [] }, { 'background': [] }],          // dropdown with defaults from theme
            [{ 'align': [] }],

            ["link", "image", "video"],

            ['clean']                                         // remove formatting button
        ]

        let keyBindings = {
            custom: {
                key: 'Enter',
                shiftKey: true,
                handler: function(range:any, context:any) {
                    // Handle shift+enter
                    submitEditor()
                }
            },
        }

        quill = new Quill("#editorId", {
            theme: "bubble",
            modules: {
                toolbar: toolbarOptions,
                keyboard: {
                    bindings: keyBindings,
                },
                magicUrl: true,
                mention: {
                    allowedChars: /^[A-Za-z\sÅÄÖåäö]*$/,
                    mentionDenotationChars: ["@"],
                    mentionContainerClass: "bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-primary-500 focus:border-primary-500 block w-1/3 p-2.5 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-primary-500 dark:focus:border-primary-500",
                    listItemClass: "block mb-2 text-sm font-medium text-gray-900 dark:text-white cursor-pointer",
                    mentionListClass: "block mb-2 text-sm font-medium text-gray-900 dark:text-white cursor-pointer",
                    minChars: 1,
                    selectKeys: [13, 9],
                    source: function(searchTerm:any, renderList:any, mentionChar:any) {
                        let values = atValues;

                        if (searchTerm.length === 0) {
                            renderList(values, searchTerm);
                        } else {
                            const matches = [];
                            for (let i = 0; i < values.length; i++)
                            if (
                                ~values[i].value
                                .toLowerCase()
                                .indexOf(searchTerm.toLowerCase())
                            )
                                matches.push(values[i]);
                            renderList(matches, searchTerm);
                        }
                    }
                }
                /*
                imageHandler: {
                    upload: file => {
                        // return a Promise that resolves in a link to the uploaded image
                        return new Promise((resolve) => {
                            const fd = new FormData();
                            fd.append("file", file);
                            fd.append("fileName", file.name);

                            _onUpload(fd, resolve);
                        });
                },
                videoHandler: {
                    upload: file => {
                        // return a Promise that resolves in a link to the uploaded image
                        return new Promise((resolve) => {
                            const fd = new FormData();
                            fd.append("file", file);
                            fd.append("fileName", file.name);

                            _onUpload(fd, resolve);
                        });
                    }
                },
                attachmentHandler: {
                    upload: file => {
                        // return a Promise that resolves in a link to the uploaded image
                        return new Promise((resolve) => {
                            const fd = new FormData();

                            fd.append("file", file);
                            fd.append("name", file.name);

                            _onAttachment(fd, resolve);
                        });
                    }
                }
                
            }*/
        }});

        if (contents) {
            quill.setContents(JSON.parse(contents), "user");
        }

        quill.focus()
    });


    const clearEditor = () => {
        quill.setContents([{ "insert": "\n" }]);
    }

    const submitEditor = async () => {
        let text = JSON.stringify(quill.getContents())
        if(text.length === 0) {
            return
        }

        if (typeof post === 'function') {
            post(text)

            clearEditor()
        }
    }

    const dropHandle = (event:any) => {
        fileValue = [];
        event.preventDefault();
        if (event.dataTransfer.items) {
        [...event.dataTransfer.items].forEach((item, i) => {
            if (item.kind === 'file') {
                const file = item.getAsFile();
                console.log(file)

                handleFileUpload(file, attachContext, "comment_upload", () => { console.log('upload complete ')})

                fileValue.push(file.name);
            }
        });
        } else {
            [...event.dataTransfer.files].forEach((file, i) => {
                fileValue = file.name;
            });
        }
    };

    const handleChange = (event:any) => {
        const files = event.target.files;
        if (files.length > 0) {
            const file = files[0];

            handleFileUpload(file, attachContext, "CommentUpload", () => { console.log('upload complete ')})

            console.log(file)

            fileValue.push(file.name);
            //fileValue = fileValue;
        }
    };

    const showFiles = (files:string[]) => {
        if (files.length === 1) 
            return files[0];

        let concat = '';
        files.map((file:any) => {
            concat += file;
            concat += ',';
            concat += ' ';
        });

        if (concat.length > 40) 
            concat = concat.slice(0, 40);

        concat += '...';
        return concat;
    };
</script>


<form class="mb-6">
    <div
        class="mb-4 w-full bg-gray-50 rounded-lg border border-gray-200 dark:bg-gray-700 dark:border-gray-600">
        <div class="py-2 px-4 bg-gray-50 rounded-t-lg dark:bg-gray-800">
            <label for="comment" class="sr-only">Your comment</label>
            <div id="editorId"></div>
        </div>
        <div class="flex justify-between items-center py-2 px-3 border-t dark:border-gray-600">
            <div class="flex pl-0 space-x-1 sm:pl-2">
                <Dropzone
                    id="dropzone"
                    defaultClass=""
                    on:drop={dropHandle}
                    on:dragover={(event) => {
                        event.preventDefault();
                    }}
                    on:change={handleChange}>
                    <span class="inline-flex justify-center p-2 text-gray-500 rounded cursor-pointer hover:text-gray-900 hover:bg-gray-100 dark:text-gray-400 dark:hover:text-white dark:hover:bg-gray-600">
                        <svg class="w-4 h-4" aria-hidden="true" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 12 20">
                            <path stroke="currentColor" stroke-linejoin="round" stroke-width="2" d="M1 6v8a5 5 0 1 0 10 0V4.5a3.5 3.5 0 1 0-7 0V13a2 2 0 0 0 4 0V6"/>
                        </svg>
                        <span class="sr-only">Attach file</span>
                    </span>
                </Dropzone>
                <!-- <button type="button"
                    class="inline-flex justify-center p-2 text-gray-500 rounded cursor-pointer hover:text-gray-900 hover:bg-gray-100 dark:text-gray-400 dark:hover:text-white dark:hover:bg-gray-600">
                    <svg class="w-4 h-4" aria-hidden="true" xmlns="http://www.w3.org/2000/svg" fill="currentColor" viewBox="0 0 16 20">
                        <path d="M8 0a7.992 7.992 0 0 0-6.583 12.535 1 1 0 0 0 .12.183l.12.146c.112.145.227.285.326.4l5.245 6.374a1 1 0 0 0 1.545-.003l5.092-6.205c.206-.222.4-.455.578-.7l.127-.155a.934.934 0 0 0 .122-.192A8.001 8.001 0 0 0 8 0Zm0 11a3 3 0 1 1 0-6 3 3 0 0 1 0 6Z"/>
                    </svg>
                    <span class="sr-only">Set location</span>
                </button>
                <button type="button"
                    class="inline-flex justify-center p-2 text-gray-500 rounded cursor-pointer hover:text-gray-900 hover:bg-gray-100 dark:text-gray-400 dark:hover:text-white dark:hover:bg-gray-600">
                    <svg class="w-4 h-4" aria-hidden="true" xmlns="http://www.w3.org/2000/svg" fill="currentColor" viewBox="0 0 16 20">
                        <path d="M14.066 0H7v5a2 2 0 0 1-2 2H0v11a1.97 1.97 0 0 0 1.934 2h12.132A1.97 1.97 0 0 0 16 18V2a1.97 1.97 0 0 0-1.934-2ZM10.5 6a1.5 1.5 0 1 1 0 2.999A1.5 1.5 0 0 1 10.5 6Zm2.221 10.515a1 1 0 0 1-.858.485h-8a1 1 0 0 1-.9-1.43L5.6 10.039a.978.978 0 0 1 .936-.57 1 1 0 0 1 .9.632l1.181 2.981.541-1a.945.945 0 0 1 .883-.522 1 1 0 0 1 .879.529l1.832 3.438a1 1 0 0 1-.031.988Z"/>
                        <path d="M5 5V.13a2.96 2.96 0 0 0-1.293.749L.879 3.707A2.98 2.98 0 0 0 .13 5H5Z"/>
                    </svg>
                    <span class="sr-only">Upload image</span>
                </button> -->

                {#if fileValue.length}
                <p>{showFiles(fileValue)}</p>
                {/if}
            </div>
            <button type="button"
                onclick={submitEditor}
                class="inline-flex items-center py-2.5 px-4 text-xs font-medium text-center text-white bg-primary-700 rounded-lg focus:ring-4 focus:ring-primary-200 dark:focus:ring-primary-900 hover:bg-primary-800">
                {#if contents}
                    Edit Comment (shift + enter)
                {:else}
                    Post comment (shift + enter)
                {/if}
            </button>
        </div>
    </div>
</form>

<!-- <div >
    <div id="editorId"></div>
</div>
<div class="mt-4">
    <Button on:click={submitEditor}>Send</Button>
</div> -->

    