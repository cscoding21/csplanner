<script lang="ts">
    import { Button, Dropdown, ButtonGroup, MultiSelect, Input } from 'flowbite-svelte';
    import { ChevronDownOutline } from 'flowbite-svelte-icons';

    let searchInput:string = $state("")


    let status:string[] = $state(["new", "approved"])

    let statusOpts = [
        { value: "new", name: "New"},
        { value: "draft", name: "Draft"},
        { value: "proposed", name: "Proposed"},
        { value: "approved", name: "Approved"},
        { value: "rejected", name: "Rejected"},
        { value: "backlogged", name: "Backlogged"},
        { value: "scheduled", name: "Scheduled"},
        { value: "inflight", name: "In-flight"},
        { value: "complete", name: "Complete"},
        { value: "deferred", name: "Deferred"},
        { value: "abandoned", name: "Abandoned"},
    ]

    const statusChange = () => {
        console.log(status)
    }

</script>

<div class="flex">
    <div class="flex-initial w-64 mr-2">
        <Input size="sm" placeholder="Filter by keyword" value={searchInput} />
    </div>

    <ButtonGroup>
        <MultiSelect items={statusOpts} bind:value={status} on:change={statusChange} />


    <!-- <li>
        <Checkbox name="status" bind:status value="draft">Draft (56)</Checkbox>
    </li>
    <li>
        <Checkbox name="status" bind:status value="proposed">Proposed (56)</Checkbox>
    </li>
    <li>
        <Checkbox name="status" value="approved" checked>Approved (97)</Checkbox>
    </li>
    <li>
        <Checkbox name="status" value="rejected" checked>Rejected (97)</Checkbox>
    </li>
    <li>
        <Checkbox name="status" value="backlogged" checked>Backlogged (97)</Checkbox>
    </li>
    <li>
        <Checkbox name="status" value="scheduled" checked>Scheduled (176)</Checkbox>
    </li>
    <li>
        <Checkbox name="status" value="inflight">In-flight (234)</Checkbox>
    </li>
    <li>
        <Checkbox name="status" value="complete" checked>Complete (76)</Checkbox>
    </li>
    <li>
        <Checkbox name="status" value="deferred">Deferred (49)</Checkbox>
    </li>
    <li>
        <Checkbox name="status" value="abandoned">Abandoned (45)</Checkbox>
    </li> -->


    </ButtonGroup>
</div>

