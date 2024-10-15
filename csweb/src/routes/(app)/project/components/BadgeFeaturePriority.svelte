<script lang="ts">
    import { 
        Badge, 
        Indicator,
        type ColorVariant,
		type IndicatorColorType } from "flowbite-svelte";
    import type { ProjectFeature } from "$lib/graphql/generated/sdk";

interface Props {
		feature: ProjectFeature;
	}
	let { 
		feature = $bindable(),  
	}: Props = $props();

    const setElements = (p: any): any => {
		if (!p) {
			return;
		}

		switch (p) {
			case 'veryhigh':
				priorityDisplay.name = 'Must Have';
				priorityDisplay.color = 'red';
				priorityDisplay.indicatorColor = 'red';
				return priorityDisplay;
			case 'high':
				priorityDisplay.name = 'High';
				priorityDisplay.color = 'yellow';
				priorityDisplay.indicatorColor = 'yellow';
				return priorityDisplay;
			case 'medium':
				priorityDisplay.name = 'Medium';
				priorityDisplay.color = 'blue';
				priorityDisplay.indicatorColor = 'blue';
				return priorityDisplay;
			case 'low':
				priorityDisplay.name = 'Low';
				priorityDisplay.color = 'purple';
				priorityDisplay.indicatorColor = 'purple';
				return priorityDisplay;
		}

		return priorityDisplay;
	};

	interface PriorityDisplay {
		name: string;
		color: ColorVariant | undefined;
		indicatorColor: IndicatorColorType | undefined;
	}

	let priorityDisplay: PriorityDisplay = $state({
		name: '',
		color: undefined,
		indicatorColor: undefined
	});

	setElements(feature.priority);

</script>


<Badge color={priorityDisplay.color} rounded class="px-2.5 py-1">
    <Indicator
        color={priorityDisplay.indicatorColor}
        size="xs"
        class="mr-1"
    />{priorityDisplay.name}
</Badge>