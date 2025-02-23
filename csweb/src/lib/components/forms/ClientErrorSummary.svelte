<script lang="ts">
	import { Alert } from "flowbite-svelte";
	import { InfoCircleSolid } from "flowbite-svelte-icons";


interface Props {
    errors?: any;
}
let { 
    errors, 
}: Props = $props();
</script>


{#if errors}

{#if errors.inner && errors.inner.length > 0}
<Alert class="items-start! mb-6" border color="red">
    <span slot="icon">
      <InfoCircleSolid class="w-5 h-5" />
      <span class="sr-only">Info</span>
    </span>
    <p class="font-medium">There was a problem processing the form:</p>
    <ul class="mt-1.5 ms-4 list-disc list-inside">
        {#each errors.inner as err}
            <li>{err.errors}</li>
        {/each}
    </ul>
  </Alert>
{/if}
{/if}