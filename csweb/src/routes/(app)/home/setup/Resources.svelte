<script lang="ts">
	import { SectionHeading } from "$lib/components";
	import SectionSubHeading from "$lib/components/formatting/SectionSubHeading.svelte";
	import type { UpdateList, UpdateListItem } from "$lib/graphql/generated/sdk";
	import { getList, updateList } from "$lib/services/list";
	import { addToast } from "$lib/stores/toasts";
	import { callIf } from "$lib/utils/helpers";
	import { nameToID } from "$lib/utils/id";
	import { Alert, Badge, Button, Input, P, Popover } from "flowbite-svelte";
	import { CloseCircleSolid } from "flowbite-svelte-icons";

    interface Props {
        onDone: Function
    }
    let { onDone }:Props = $props()

</script>


<h2 class="text-xl text-center text-gray-50 font-semibold">Organization Resources</h2>
    
<section class="mb-6 mt-4">
<p class="py-2 text-gray-200">
    Resources are the staff, hardware, software, etc required to deliver on your project roadmap.
</p>
<p class="py-2 text-gray-200">Resources can be added or modified at any time on the <a class="text-orange-300" href="/resources">Resources</a> page.
</p>
</section>