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
		// return {
		// 	colors: ['#1A56DB', '#FDBA8C'],
		// 	series: [
		// 		{
		// 			name: 'Earnings',
		// 			color: '#5a67d8',
		// 			data: [
		// 				{ x: 'Year 1', y: p.projectValue?.calculated?.yearOneValue },
		// 				{ x: 'Year 2', y: p.projectValue?.calculated?.yearTwoValue },
		// 				{ x: 'Year 3', y: p.projectValue?.calculated?.yearThreeValue },
		// 				{ x: 'Year 4', y: p.projectValue?.calculated?.yearFourValue },
		// 				{ x: 'Year 5', y: p.projectValue?.calculated?.yearFiveValue }
		// 			]
		// 		}
		// 	],
		// 	chart: {
		// 		type: 'bar',
		// 		height: '160px',
		// 		fontFamily: 'Inter, sans-serif',
		// 		toolbar: {
		// 			show: false
		// 		}
		// 	},
		// 	plotOptions: {
		// 		bar: {
		// 			horizontal: false,
		// 			columnWidth: '70%',
		// 			borderRadiusApplication: 'end',
		// 			borderRadius: 8
		// 		}
		// 	},
		// 	tooltip: {
		// 		shared: true,
		// 		intersect: false,
		// 		style: {
		// 			fontFamily: 'Inter, sans-serif'
		// 		}
		// 	},
		// 	states: {
		// 		hover: {
		// 			filter: {
		// 				type: 'darken',
		// 				value: 1
		// 			}
		// 		}
		// 	},
		// 	stroke: {
		// 		show: true,
		// 		width: 0,
		// 		colors: ['transparent']
		// 	},
		// 	grid: {
		// 		show: true,
		// 		strokeDashArray: 4,
		// 		padding: {
		// 			left: 2,
		// 			right: 2,
		// 			top: -14
		// 		}
		// 	},
		// 	dataLabels: {
		// 		enabled: false
		// 	},
		// 	legend: {
		// 		show: true
		// 	},
		// 	xaxis: {
		// 		floating: false,
		// 		labels: {
		// 			show: true,
		// 			style: {
		// 				fontFamily: 'Inter, sans-serif',
		// 				cssClass: 'text-xs font-normal fill-gray-500 dark:fill-gray-400'
		// 			}
		// 		},
		// 		axisBorder: {
		// 			show: false
		// 		},
		// 		axisTicks: {
		// 			show: false
		// 		}
		// 	},
		// 	yaxis: {
		// 		show: true,
		// 		labels: {
		// 			formatter: function (value: number) {
		// 				return formatCurrency.format(value);
		// 			},
		// 			style: {
		// 				colors: ['#bbb']
		// 			},
		// 			offsetX: -6
		// 		}
		// 	},
		// 	fill: {
		// 		opacity: 0.75
		// 	}
		// };

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
				width: 320,
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
				}
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
				show: true,
				labels: {
					colors: ['#bbb'],
				}
			}
        };
	};
	let chart = new ApexCharts(document.querySelector('#chartPlaceholder'), getChartOptions(project));
</script>

<Chart options={getChartOptions(project)} {chart} />
<div id="chartPlaceholder"></div>
