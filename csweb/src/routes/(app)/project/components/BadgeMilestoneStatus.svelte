<script lang="ts">
	import { MilestoneStatusAccepted, MilestoneStatusDone, MilestoneStatusNew, MilestoneStatusRemoved } from "$lib/services/project";
    import { 
        Badge, 
        Indicator} from "flowbite-svelte";

interface Props {
		status: string;
	}
	let { 
		status = $bindable(),  
	}: Props = $props();

	let badgeName = $state("")
	let badgeColor = $state("default"  as "red"|"yellow"|"blue"|"green")
	let badgeIndicatorColor = $state("default"  as "red"|"yellow"|"blue"|"green")

	$effect(() => {
		switch (status) {
			case MilestoneStatusNew:
				badgeName = 'New';
				badgeColor = 'yellow';
				badgeIndicatorColor = 'yellow';
				break;
			case MilestoneStatusAccepted:
				badgeName = 'Accepted';
				badgeColor = 'blue';
				badgeIndicatorColor = 'blue';
				break;
			case MilestoneStatusRemoved:
				badgeName = 'Removed';
				badgeColor = 'red';
				badgeIndicatorColor = 'red';
				break;
			case MilestoneStatusDone:				
				badgeName = 'Done';
				badgeColor = 'green';
				badgeIndicatorColor = 'green';
				break;
		}
	})
</script>


<Badge color={badgeColor} rounded class="px-2.5 py-1">
    <Indicator
        color={badgeIndicatorColor}
        size="xs"
        class="mr-1"
    />{badgeName}
</Badge>