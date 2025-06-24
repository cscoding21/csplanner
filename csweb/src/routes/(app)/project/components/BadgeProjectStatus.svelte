<script lang="ts">
	import { ProjectStatusAbandoned, ProjectStatusApproved, ProjectStatusBacklogged, ProjectStatusComplete, ProjectStatusDeferred, ProjectStatusDraft, ProjectStatusInflight, ProjectStatusNew, ProjectStatusProposed, ProjectStatusRejected, ProjectStatusScheduled } from "$lib/services/project";
    import { 
        Badge, 
        Indicator
	} from "flowbite-svelte";

	interface Props {
		status: string;
	}
	let { 
		status
	}: Props = $props();

	let badgeName = $state("")
	let badgeColor = $state("default" as "red"|"yellow"|"blue"|"purple"|"green")
	let badgeIndicatorColor = $state("default" as "red"|"yellow"|"blue"|"purple"|"green")
	
	$effect(() => {
		switch (status) {
			case ProjectStatusNew:
				badgeName = 'New';
				badgeColor = 'red';
				badgeIndicatorColor = 'red';
				break;
			case ProjectStatusDraft:
				badgeName = 'Draft';
				badgeColor = 'yellow';
				badgeIndicatorColor = 'yellow';
				break;
			case ProjectStatusProposed	:
				badgeName = 'Proposed';
				badgeColor = 'blue';
				badgeIndicatorColor = 'blue';
				break;
			case ProjectStatusApproved:	
				badgeName = 'Approved';
				badgeColor = 'purple';
				badgeIndicatorColor = 'purple';
				break;
			case ProjectStatusRejected:
				badgeName = 'Rejected';
				badgeColor = 'red';
				badgeIndicatorColor = 'red';
				break;
			case ProjectStatusBacklogged:
				badgeName = 'Backlogged';
				badgeColor = 'purple';
				badgeIndicatorColor = 'purple';
				break;
			case ProjectStatusScheduled:
				badgeName = 'Scheduled';
				badgeColor = 'green';
				badgeIndicatorColor = 'green';
				break;
			case ProjectStatusInflight:
				badgeName = 'In-flight';
				badgeColor = 'green';
				badgeIndicatorColor = 'green';
				break;
			case ProjectStatusComplete:
				badgeName = 'Complete';
				badgeColor = 'green';
				badgeIndicatorColor = 'green';
				break;
			case ProjectStatusDeferred:
				badgeName = 'Deferred';
				badgeColor = 'purple';
				badgeIndicatorColor = 'purple';
				break;
			case ProjectStatusAbandoned:
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
