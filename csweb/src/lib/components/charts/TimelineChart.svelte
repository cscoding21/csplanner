<script lang="ts">
	import { Chart } from 'flowbite-svelte';
	import { type ApexOptions, default as ApexCharts } from 'apexcharts';

	interface Props {
		data: [];
		height: number;
	}
	let { data = $bindable(), height }: Props = $props();

	const getChartOptions = (): ApexOptions => {
		const o: ApexOptions = {
			series: [
				{
					data: data
				}
			],
			chart: {
				height: height,
				type: 'rangeBar'
			},
			plotOptions: {
				bar: {
					horizontal: true
				}
			},
			xaxis: {
				type: 'datetime',
				labels: {
					style: {
						colors: ['#ffffff']
					}
				}
			},
			yaxis: {
				labels: {
					style: {
						colors: ['#ffffff']
					}
				}
			},
			colors: [
				'#008FFB',
				'#00E396',
				'#FEB019',
				'#FF4560',
				'#775DD0',
				'#3F51B5',
				'#546E7A',
				'#D4526E',
				'#8D5B4C',
				'#F86624',
				'#D7263D',
				'#1B998B',
				'#2E294E',
				'#F46036',
				'#E2C044'
			]
		};

		return o;
	};
</script>

<Chart options={getChartOptions() as ApexOptions} />
<div id="chartPlaceholder"></div>
