<script lang="ts">
    import { 
        Badge, 
        Indicator,
        type ColorVariant,
		type IndicatorColorType 
	} from "flowbite-svelte";

	interface Props {
		status: string;
	}
	let { 
		status
	}: Props = $props();

	interface BadgeDisplay {
		name: string;
		color: ColorVariant | undefined;
		indicatorColor: IndicatorColorType | undefined;
	}

    const setElements = (p: string): any => {
		console.log("setElements", p)
		let d:BadgeDisplay = {name: "", color: undefined, indicatorColor: undefined }
		switch (p) {
			case 'new':
                d.name = 'New';
				d.color = 'red';
				d.indicatorColor = 'red';
				return d;
			case 'draft':
                d.name = 'Draft';
				d.color = 'yellow';
				d.indicatorColor = 'yellow';
				return d;
			case 'proposed':
                d.name = 'Proposed';
				d.color = 'blue';
				d.indicatorColor = 'blue';
				return d;
			case 'approved':
                d.name = 'Approved';
				d.color = 'purple';
				d.indicatorColor = 'purple';
				return d;
            case 'rejected':
                d.name = 'Rejected';
				d.color = 'purple';
				d.indicatorColor = 'purple';
			    return d;
            case 'backlogged':
                d.name = 'Backlogged';
				d.color = 'purple';
			    d.indicatorColor = 'purple';
			    return d;
            case 'scheduled':
                d.name = 'Scheduled';
				d.color = 'purple';
				d.indicatorColor = 'purple';
			    return d;
            case 'inflight':
                d.name = 'In-flight';
				d.color = 'purple';
				d.indicatorColor = 'purple';
			    return d;
            case 'complete':
                d.name = 'Complete';
				d.color = 'purple';
				d.indicatorColor = 'purple';
			    return d;
            case 'deferred':
                d.name = 'Deferred';
				d.color = 'purple';
				d.indicatorColor = 'purple';
			    return d;
            case 'abandoned':
                d.name = 'Abandoned';
				d.color = 'purple';
				d.indicatorColor = 'purple';
			    return d;
		}

		return d;
	};
	
	let display = $state(setElements(status));
</script>

{status}
{#key display}
<Badge color={display.color} rounded class="px-2.5 py-1">
    <Indicator
        color={display.indicatorColor}
        size="xs"
        class="mr-1"
    />{display.name}
</Badge>
{/key}