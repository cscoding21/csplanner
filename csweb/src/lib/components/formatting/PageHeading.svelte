<script lang="ts">
    import { Heading } from "flowbite-svelte";

    let { title = $bindable() } = $props()
</script>

<Heading tag="h3" class="px-2 mb-4">
    {title}

    <span class="float-right">
        <slot />
    </span>
</Heading>