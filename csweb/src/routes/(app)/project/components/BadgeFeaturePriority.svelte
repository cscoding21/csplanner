<script lang="ts">
    import { 
        Badge, 
        Indicator,
        type ColorVariant,
		type IndicatorColorType } from "flowbite-svelte";

	interface Props {
		priority: string;
	}
	let { 
		priority = $bindable(),  
	}: Props = $props();

	let badgeName = $state("")
	let badgeColor = $state("default" as ColorVariant)
	let badgeIndicatorColor = $state("default" as IndicatorColorType)

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