<script lang="ts">
	import { Chart } from 'flowbite-svelte';
	import { type ApexOptions, default as ApexCharts } from 'apexcharts';
	import { formatCurrency } from '$lib/utils/format';

    interface Props {
		values: number[];
        labels: string[];
        format?: "none"|"currency"|undefined;
	}
	let { values = $bindable(), labels = $bindable(), format }: Props = $props();

    const formatSelector = (val:any, f:"none"|"currency"|undefined) => {
        switch(f) {
            case("currency"):
                return formatCurrency.format(val)
            default:
                return val

        }
    }

    let getChartOptions = (v:number[], l:string[]): ApexOptions => {
        return  {
        	series: v,
			chart: {
				width: '60%',
				type: 'pie',
			},
			labels: l,
			// fill: {
			// 	colors: ['#5a67d8', '#FDBA8C'],
			// 	opacity: 0.75
			// },
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
				show: true,
				labels: {
					colors: ['#bbb', '#bbb', '#bbb', '#bbb', '#bbb', '#bbb', '#bbb', '#bbb', '#bbb', '#bbb', '#bbb'],
				},
			},
			tooltip: {
				enabled: true,
				y: {
					formatter: function(value) {
						return formatSelector(value, format);
					}
				}
			}
        };
	};
	let chart = new ApexCharts(document.querySelector('#pie_chartPlaceholder'), getChartOptions(values, labels));
</script>


<Chart options={getChartOptions(values, labels)} {chart} />
<div id="pie_chartPlaceholder"></div>