<script lang="ts">
	import { Chart } from 'flowbite-svelte';
	import { type ApexOptions, default as ApexCharts } from 'apexcharts';

	import type { Project, ProjectValueLine } from '$lib/graphql/generated/sdk';
	import { formatCurrency } from '$lib/utils/format';
	import { deepCopy } from '$lib/utils/helpers';

	interface Props {
		project: Project;
	}
	let { project }: Props = $props();

	let getChartOptions = (p: Project): ApexOptions => {
        let values = []
        let labels = []

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

        return  {
        	series: values,
			chart: {
				width: 240,
				type: 'pie',
			},
			labels: labels,
			fill: {
				colors: ['#5a67d8', '#FDBA8C'],
				opacity: 0.75
			},
			dataLabels: {
				style: {
					fontFamily: 'Inter, sans-serif',
					colors: ['#ffffff']
				},
				
			},
			grid: {
			padding: {
				top: 0,
				bottom: 0,
				left: 0,
				right: 0,
			},
			},
			responsive: [{
				breakpoint: 480,
				options: {
					chart: {
						width: 200
					},
					legend: {
						position: 'bottom'
					}
				}
			}],
			legend: {
				show: false,
				labels: {
					colors: ['#bbb','#bbb','#bbb','#bbb',],
				},
			},
			tooltip: {
				enabled: true,
				y: {
					formatter: function(value) {
						return formatCurrency.format(value);
					}
				}
			}
        };
	};
	let chart = new ApexCharts(document.querySelector('#chartPlaceholder'), getChartOptions(project));
</script>

<Chart options={getChartOptions(project)} {chart} />
<div id="chartPlaceholder"></div>
