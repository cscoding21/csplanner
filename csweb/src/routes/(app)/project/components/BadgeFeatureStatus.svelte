<script lang="ts">
	import { FeatureStatusAccepted, FeatureStatusDone, FeatureStatusProposed, FeatureStatusRemoved } from "$lib/services/project";
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
	let badgeColor = $state("default" as "red"|"yellow"|"blue"|"purple")
	let badgeIndicatorColor = $state("default" as "red"|"yellow"|"blue"|"purple")

	$effect(() => {
		switch (status) {
			case FeatureStatusProposed:
				badgeName = 'Proposed';
				badgeColor = 'yellow';
				badgeIndicatorColor = 'yellow';
				break;
			case FeatureStatusAccepted:
				badgeName = 'Accepted';
				badgeColor = 'blue';
				badgeIndicatorColor = 'blue';
				break;
			case FeatureStatusDone:
				badgeName = 'Done';
				badgeColor = 'yellow';
				badgeIndicatorColor = 'yellow';
				break;
			case FeatureStatusRemoved:				
				badgeName = 'Removed';
				badgeColor = 'red';
				badgeIndicatorColor = 'red';
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