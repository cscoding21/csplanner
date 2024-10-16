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

	let badgeName = $state("")
	let badgeColor = $state("default" as ColorVariant)
	let badgeIndicatorColor = $state("default" as IndicatorColorType)
	
	$effect(() => {
		switch (status) {
			case 'new':
				badgeName = 'New';
				badgeColor = 'red';
				badgeIndicatorColor = 'red';
				break;
			case 'draft':
				badgeName = 'Draft';
				badgeColor = 'yellow';
				badgeIndicatorColor = 'yellow';
				break;
			case 'proposed':
				badgeName = 'Proposed';
				badgeColor = 'blue';
				badgeIndicatorColor = 'blue';
				break;
			case 'approved':				
				badgeName = 'Approved';
				badgeColor = 'purple';
				badgeIndicatorColor = 'purple';
				break;
			case 'rejected':
				badgeName = 'Rejected';
				badgeColor = 'red';
				badgeIndicatorColor = 'red';
				break;
			case 'backlogged':
				badgeName = 'Backlogged';
				badgeColor = 'purple';
				badgeIndicatorColor = 'purple';
				break;
			case 'scheduled':
				badgeName = 'Scheduled';
				badgeColor = 'purple';
				badgeIndicatorColor = 'purple';
				break;
			case 'inflight':
				badgeName = 'In-flight';
				badgeColor = 'blue';
				badgeIndicatorColor = 'blue';
				break;
			case 'complete':
				badgeName = 'Complete';
				badgeColor = 'green';
				badgeIndicatorColor = 'green';
				break;
			case 'deferred':
				badgeName = 'Deferred';
				badgeColor = 'purple';
				badgeIndicatorColor = 'purple';
				break;
			case 'abandoned':
				badgeName = 'Abandoned';
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