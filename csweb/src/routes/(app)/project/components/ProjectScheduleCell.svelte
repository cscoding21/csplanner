<script lang="ts">
    import { Popover, Button } from "flowbite-svelte";
    import { normalizeGUID } from "$lib/utils/id";
	import { formatDate } from "$lib/utils/format";
	import WeekPopupSummary from "$lib/components/widgets/WeekPopupSummary.svelte";
	import type { ProjectScheduleRowCell } from "$lib/services/schedule";

    interface Props {
        week: ProjectScheduleRowCell
    }
    let { week }: Props = $props();

    let uuid = normalizeGUID(crypto.randomUUID());
</script>


{#if week.activities.length > 0}
{@const cellColor = week.risks.length > 0 ? "yellow" : "green" }
{@const popWidth = week.risks.length > 0 ? "w-[600px]" : "w-64" }
<Button  size="xs" color={cellColor} pill id={"id_" + uuid}>{week.activities.reduce((acc, curr) => acc + (curr.hoursSpent || 0), 0)}</Button>
<Popover class="text-sm font-light {popWidth}" title={"Week ending " + formatDate(week.end)} triggeredBy={"#id_" + uuid}>
    <div class="p-2">
    <ul class="">
        <WeekPopupSummary activities={week.activities} risks={week.risks} />
    </ul>
</div>
</Popover>
{/if}