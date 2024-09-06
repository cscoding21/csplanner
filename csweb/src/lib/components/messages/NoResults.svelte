<script lang="ts">
    import { Card, Button } from 'flowbite-svelte';
    import { ArrowRightOutline } from 'flowbite-svelte-icons';

    let { title, newUrl = $bindable() } = $props()
</script>


<Card class="w-full" size="xl">
    <h5 class="mb-2 text-2xl font-bold tracking-tight text-gray-900 dark:text-white">{title}</h5>
    <p class="mb-3 font-normal text-gray-700 dark:text-gray-400 leading-tight"><slot /></p>
    {#if newUrl }
        <Button class="w-fit" href={newUrl as string}>
        Create new <ArrowRightOutline class="w-3.5 h-3.5 ml-2 text-white" />
        </Button>
    {/if}
</Card>