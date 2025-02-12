<script lang="ts">
	import { Chart } from 'flowbite-svelte';
	import { type ApexOptions, default as ApexCharts } from 'apexcharts';
	import { formatCurrency } from '$lib/utils/format';
	import { getPrimaryColor, getTertiaryColor, getQuaternaryColor, getWhiteAndDark900Colors } from './chartConfig'

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

	const getChartOptions = (v:number[], l:string[]):ApexOptions => {
		return {
			series: v,
			colors: [getPrimaryColor(), getTertiaryColor(), getQuaternaryColor()],
			chart: {
				height: 240,
				width: "100%",
				type: "pie",
			},
			stroke: {
				colors: [getWhiteAndDark900Colors()],
				lineCap: undefined,
			},
			plotOptions: {
				// pie: {
				// 	labels: {
				// 		show: true,
				// 	},
				// 	size: "100%",
				// 	dataLabels: {
				// 		offset: -25,
				// 	},
				// },
			},
			labels: l,
			dataLabels: {
				enabled: true,
				style: {
					fontFamily: "Inter, sans-serif",
				},
			},
			legend: {
				position: "right",
				fontFamily: "Inter, sans-serif",
				labels: {
					colors: ['#bbb', '#bbb', '#bbb', '#bbb', '#bbb', '#bbb', '#bbb', '#bbb', '#bbb', '#bbb', '#bbb'],
				},
			},
			yaxis: {
				labels: {
					formatter: function (value:any) {
						return formatSelector(value, format);
					},
				},
			},
			xaxis: {
				labels: {
					formatter: function (value:any) {
						return formatSelector(value, format);
					},
				},
				axisTicks: {
					show: false,
				},
				axisBorder: {
					show: false,
				},
			},
		};
	};

	let chart = new ApexCharts(document.querySelector('#pie_chartPlaceholder'), getChartOptions(values, labels));
</script>


<Chart options={getChartOptions(values, labels)} {chart} />
<div id="pie_chartPlaceholder"></div>