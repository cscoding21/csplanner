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
			case 'proposed':
                display.name = 'Proposed';
                display.color = 'blue';
                display.indicatorColor = 'blue';
				return display;
			case 'accepted':
                display.name = 'Accepted';
                display.color = 'green';
                display.indicatorColor = 'green';
				return display;
			case 'deferred':
                display.name = 'Deferred';
                display.color = 'yellow';
                display.indicatorColor = 'yellow';
				return display;
			case 'removed':
            display.name = 'Removed';
				display.color = 'red';
				display.indicatorColor = 'red';
				return display;
		}

		return display;
	};

	interface Display {
		name: string;
		color: ColorVariant | undefined;
		indicatorColor: IndicatorColorType | undefined;
	}

	let display: Display = $state({
		name: '',
		color: undefined,
		indicatorColor: undefined
	});

	setElements(feature.status);

</script>


<Badge color={display.color} rounded class="px-2.5 py-1">
    <Indicator
        color={display.indicatorColor}
        size="xs"
        class="mr-1"
    />{display.name}
</Badge>