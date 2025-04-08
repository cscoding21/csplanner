<script lang="ts">
	import { SectionHeading } from "$lib/components";
	import { Button } from "flowbite-svelte";
	import { CheckOutline } from "flowbite-svelte-icons";

    interface Props {
        onDone: Function
    }
    let { onDone }:Props = $props()
</script>


<div class="p-4">
<div class="grid grid-cols-2 gap-4">
    <div class="">
        <SectionHeading>Welcome to csPlanner</SectionHeading>
        <div>In order to get the optimal planning outcomes, we'll need to gather some data about your organization,
            as well as introduce you to some key concepts within csPlanner.</div>
        
        <ul class="mt-4 space-y-4">
            <li>
                <h3 class="text-gray-200"><CheckOutline class="float-left mr-2 text-green-400" /> Review Organization Settings</h3>
                <small>We'll take a look at some of the default numbers used to perform costs, value, and scheduling calculations.</small>
            </li>

            <li>
                <h3 class="text-gray-200"><CheckOutline class="float-left mr-2 text-green-400" /> Add Some Skills</h3>
                <small>By organizing work by skills, we're able to provide unique insights into scheduling and staffing needs.</small>
            </li>

            <li>
                <h3 class="text-gray-200"><CheckOutline class="float-left mr-2 text-green-400" /> Add Value Categories</h3>
                <small>To make optimal planning decisions, it's important to know how a project's value impacts your business.</small>
            </li>

            <li>
                <h3 class="text-gray-200"><CheckOutline class="float-left mr-2 text-green-400" /> Add Funding Sources</h3>
                <small>Tracking how a project is funded gives valueable insights into prioritization and viability.</small>
            </li>

            <li>
                <h3 class="text-gray-200"><CheckOutline class="float-left mr-2 text-green-400" /> Add Roles</h3>
                <small>A bit of up-front work organizaing skills and costs into rolews can save time down the road.</small>
            </li>

            <li>
                <h3 class="text-gray-200"><CheckOutline class="float-left mr-2 text-green-400" /> Create Resources</h3>
                <small>Resources are the liveblood of your organizaiton and drive capacity.</small>
            </li>

        </ul>

    </div>
    <div class="">
        <Button onclick={() => onDone()}>I'm ready...let's get started! >></Button>
    </div>
</div>
</div>