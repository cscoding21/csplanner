<script lang="ts">
    import { 
        Badge, 
        Indicator,
        type ColorVariant,
		type IndicatorColorType } from "flowbite-svelte";

interface Props {
		status: string;
	}
	let { 
		status = $bindable(),  
	}: Props = $props();

	let badgeName = $state("")
	let badgeColor = $state("default" as ColorVariant)
	let badgeIndicatorColor = $state("default" as IndicatorColorType)

	$effect(() => {
		switch (status) {
			case 'proposed':
				badgeName = 'Proposed';
				badgeColor = 'yellow';
				badgeIndicatorColor = 'yellow';
				break;
			case 'accepted':
				badgeName = 'Accepted';
				badgeColor = 'blue';
				badgeIndicatorColor = 'blue';
				break;
			case 'deferred':
				badgeName = 'Deferred';
				badgeColor = 'yellow';
				badgeIndicatorColor = 'yellow';
				break;
			case 'removed':				
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