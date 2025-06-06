<script lang="ts">
    import { 
        Badge, 
        Indicator} from "flowbite-svelte";

	interface Props {
		priority: string;
	}
	let { 
		priority = $bindable(),  
	}: Props = $props();

	let badgeName = $state("")
	let badgeColor = $state("default" as "red"|"yellow"|"blue"|"purple")
	let badgeIndicatorColor = $state("default" as "red"|"yellow"|"blue"|"purple")

	$effect(() => {
		switch (priority) {
			case 'veryhigh':
				badgeName = 'Must Have';
				badgeColor = 'red';
				badgeIndicatorColor = 'red';
				break;
			case 'high':
				badgeName = 'High';
				badgeColor = 'yellow';
				badgeIndicatorColor = 'yellow';
				break;
			case 'medium':
				badgeName = 'Medium';
				badgeColor = 'blue';
				badgeIndicatorColor = 'blue';
				break;
			case 'low':				
				badgeName = 'Low';
				badgeColor = 'purple';
				badgeIndicatorColor = 'purple';
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