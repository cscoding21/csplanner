<script lang="ts">
    import { Button, Dropdown, Checkbox } from "flowbite-svelte";
    import { ChevronDownOutline } from "flowbite-svelte-icons";

    interface CheckboxOption {
        name: string,
        value: any,
        checked: boolean
    }

    interface Props {
        name: string
        opts: CheckboxOption[]
        group: any,
        change: Function
    }
    let { 
        name, 
        opts = $bindable(), 
        group = $bindable(),
        change = $bindable()
    }:Props = $props()
</script>


<Button color="light">
    <svg xmlns="http://www.w3.org/2000/svg" aria-hidden="true" class="w-4 h-4 mr-2 text-gray-400" fill="currentColor">
      <path fill-rule="evenodd" d="M3 3a1 1 0 011-1h12a1 1 0 011 1v3a1 1 0 01-.293.707L12 11.414V15a1 1 0 01-.293.707l-2 2A1 1 0 018 17v-5.586L3.293 6.707A1 1 0 013 6V3z" clip-rule="evenodd" />
    </svg>
    {name}<ChevronDownOutline />
</Button>
<Dropdown class="w-44 p-3 space-y-3 text-sm">
{#each opts as opt(opt)}
<li>
    <Checkbox bind:group={group} value={opt.value} bind:checked={opt.checked} on:change={e => change(e)}>{opt.name}</Checkbox>
</li>
{/each}
</Dropdown>