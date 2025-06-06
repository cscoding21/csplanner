<script lang="ts">
	import { Chart } from 'flowbite-svelte';
	import { type ApexOptions, default as ApexCharts } from 'apexcharts';

	import type { Project } from '$lib/graphql/generated/sdk';
	import { formatCurrency } from '$lib/utils/format';

	interface Props {
		project: Project;
	}
	let { project }: Props = $props();

	let getChartOptions = (p: Project): ApexOptions => {
		return {
			colors: ['#1A56DB', '#FDBA8C'],
			series: [
				{
					name: 'Earnings',
					color: '#5a67d8',
					data: [
						{ x: 'Year 1', y: p.projectValue?.calculated?.yearOneValue },
						{ x: 'Year 2', y: p.projectValue?.calculated?.yearTwoValue },
						{ x: 'Year 3', y: p.projectValue?.calculated?.yearThreeValue },
						{ x: 'Year 4', y: p.projectValue?.calculated?.yearFourValue },
						{ x: 'Year 5', y: p.projectValue?.calculated?.yearFiveValue }
					]
				}
			],
			chart: {
				type: 'bar',
				height: '160px',
				fontFamily: 'Inter, sans-serif',
				toolbar: {
					show: false
				}
			},
			plotOptions: {
				bar: {
					horizontal: false,
					columnWidth: '70%',
					borderRadiusApplication: 'end',
					borderRadius: 8
				}
			},
			tooltip: {
				shared: true,
				intersect: false,
				style: {
					fontFamily: 'Inter, sans-serif'
				}
			},
			states: {
				hover: {
					filter: {
						type: 'darken',
					}
				}
			},
			stroke: {
				show: true,
				width: 0,
				colors: ['transparent']
			},
			grid: {
				show: true,
				strokeDashArray: 4,
				padding: {
					left: 2,
					right: 2,
					top: -14
				}
			},
			dataLabels: {
				enabled: false
			},
			legend: {
				show: true
			},
			xaxis: {
				floating: false,
				labels: {
					show: true,
					style: {
						fontFamily: 'Inter, sans-serif',
						cssClass: 'text-xs font-normal fill-gray-500 dark:fill-gray-400'
					}
				},
				axisBorder: {
					show: false
				},
				axisTicks: {
					show: false
				}
			},
			yaxis: {
				show: true,
				labels: {
					formatter: function (value: number) {
						return formatCurrency.format(value);
					},
					style: {
						colors: ['#bbb']
					},
					offsetX: -6
				}
			},
			fill: {
				opacity: 0.75
			}
		};
	};
</script>

<Chart options={getChartOptions(project)} />
<div id="chartPlaceholder"></div>
