<script lang="ts">
    import { Breadcrumb, BreadcrumbItem } from 'flowbite-svelte';

    let { pageDetail } = $props() ;
</script>


<div class="grid grid-cols-4 w-full mb-4">
    <div class="p-2 justify-start col-span-2">
        <Breadcrumb aria-label="Solid background breadcrumb example">
            <BreadcrumbItem href="/" home>Home</BreadcrumbItem>
            <BreadcrumbItem href="/resource">Resources</BreadcrumbItem>
            {#if pageDetail}
            <BreadcrumbItem>{pageDetail}</BreadcrumbItem>
            {/if}
        </Breadcrumb>
    </div>
    <div>&nbsp;</div>
    <div class="p-2 justify-self-end">
        <slot />
    </div>
</div>