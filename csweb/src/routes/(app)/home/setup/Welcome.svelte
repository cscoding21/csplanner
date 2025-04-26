<script lang="ts">
	import type { User } from "$lib/graphql/generated/sdk";
	import { Button, P } from "flowbite-svelte";
	import { CheckOutline } from "flowbite-svelte-icons";

    interface Props {
        onDone: Function
        user?: User
    }
    let { onDone, user }:Props = $props()
</script>


<h2 class="text-xl text-center text-gray-50 font-semibold">Howdy, {user?.firstName}, Welcome to csPlanner</h2>
<p class="py-6 text-gray-200">
    In order to get the optimal planning outcomes, we'll need to gather some data about your organization,
    as well as introduce you to some key concepts within csPlanner.
</p>

<ul role="list" class="mt-4 space-y-6">
    <li>
        <h3 class="text-gray-200"><CheckOutline class="float-left mr-2 text-green-400" /> Review Organization Settings</h3>
        <span>We'll take a look at some of the default numbers used to perform costs, value, and scheduling calculations.</span>
    </li>

    <li>
        <h3 class="text-gray-200"><CheckOutline class="float-left mr-2 text-green-400" /> Add Some Skills</h3>
        <span>By organizing work by skills, we're able to provide unique insights into scheduling and staffing needs.</span>
    </li>

    <li>
        <h3 class="text-gray-200"><CheckOutline class="float-left mr-2 text-green-400" /> Add Value Categories</h3>
        <span>To make optimal planning decisions, it's important to know how a project's value impacts your business.</span>
    </li>

    <li>
        <h3 class="text-gray-200"><CheckOutline class="float-left mr-2 text-green-400" /> Add Funding Sources</h3>
        <span>Tracking how a project is funded gives valueable insights into prioritization and viability.</span>
    </li>

    <li>
        <h3 class="text-gray-200"><CheckOutline class="float-left mr-2 text-green-400" /> Add Roles</h3>
        <span>A bit of up-front work organizaing skills and costs into roles can save time down the road.</span>
    </li>

    <li>
        <h3 class="text-gray-200"><CheckOutline class="float-left mr-2 text-green-400" /> Create Resources</h3>
        <span>Resources are the liveblood of your organizaiton and drive capacity.</span>
    </li>
</ul>

<div class="mt-12 text-center">
    <Button onclick={() => onDone()}>I'm ready...let's get started! >></Button>
</div>
