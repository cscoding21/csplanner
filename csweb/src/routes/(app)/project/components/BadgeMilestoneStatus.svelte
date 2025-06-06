<script lang="ts">
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
			case 'new':
				badgeName = 'New';
				badgeColor = 'yellow';
				badgeIndicatorColor = 'yellow';
				break;
			case 'accepted':
				badgeName = 'Accepted';
				badgeColor = 'blue';
				badgeIndicatorColor = 'blue';
				break;
			case 'removed':
				badgeName = 'Removed';
				badgeColor = 'red';
				badgeIndicatorColor = 'red';
				break;
			case 'done':				
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