<script lang="ts">
	import { PieChart } from '$lib/components';
	import type { Project, ProjectValueLine } from '$lib/graphql/generated/sdk';
	import { deepCopy } from '$lib/utils/helpers';

	interface Props {
		project: Project;
	}
	let { project }: Props = $props();

	let values = $state([] as number[])
    let labels = $state([] as string[])

	let getChartData = (p: Project) => {
        const lines = deepCopy (project.projectValue.projectValueLines) as ProjectValueLine[]

        for (let i = 0; i < lines.length; i++) {
            const l = lines[i]
            values.push(0 + 
                (l.yearOneValue || 0) + 
                (l.yearTwoValue || 0) + 
                (l.yearThreeValue || 0)  + 
                (l.yearFourValue || 0)  + 
                (l.yearFiveValue || 0) 
            )
            labels.push(l.valueCategory)
        }
	};

	getChartData(project)
</script>

<PieChart {values} {labels} />
