<script lang="ts">
    import { 
        Badge, 
        Indicator,
        type ColorVariant,
		type IndicatorColorType } from "flowbite-svelte";
    import type { ProjectMilestoneTask } from "$lib/graphql/generated/sdk";

interface Props {
		task: ProjectMilestoneTask;
	}
	let { 
		task = $bindable(),  
	}: Props = $props();

    const setElements = (p: any): any => {
		if (!p) {
			return;
		}

		switch (p) {
			case 'new':
				display.name = 'New';
				display.color = 'yellow';
				display.indicatorColor = 'yellow';
				return display;
			case 'accepted':
                display.name = 'Accepted';
                display.color = 'blue';
                display.indicatorColor = 'blue';
				return display;
			case 'removed':
                display.name = 'Removed';
                display.color = 'dark';
                display.indicatorColor = 'dark';
				return display;
			case 'done':
                display.name = 'Done';
                display.color = 'green';
                display.indicatorColor = 'green';
				return display;
		}

		return display;
	};

	interface PriorityDisplay {
		name: string;
		color: ColorVariant | undefined;
		indicatorColor: IndicatorColorType | undefined;
	}

	let display: PriorityDisplay = $state({
		name: '',
		color: undefined,
		indicatorColor: undefined
	});

	setElements(task.status);

</script>


<Badge color={display.color} rounded class="px-2.5 py-1">
    <Indicator
        color={display.indicatorColor}
        size="xs"
        class="mr-1"
    />{display.name}
</Badge>