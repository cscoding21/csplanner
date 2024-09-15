<script lang="ts">
    import { Chart } from 'flowbite-svelte'; 

    export let data:[]
    export let height:string = "500"

    const ct:"rangeBar"|"line"|"area"|"bar"|"pie"|"donut"|"radialBar"|"scatter"|"bubble"|"heatmap"|"candlestick"|"boxPlot"|"radar"|"polarArea"|"rangeArea"|"treemap"|undefined = "rangeBar"
    const at:"datetime" | "category" | "numeric" | undefined = "datetime"
    
    const getChartOptions = () => {
        const o = {
            series: [{
            data: data,
        }],
        chart: {
            height: height,
            type: ct
        },
        plotOptions: {
            bar: {
                horizontal: true
            }
        },
        xaxis: {
            type: at,
            labels: {
                style: {
                    colors: ["#ffffff"]
                }
            }
        },
        yaxis: {
            labels: {
                style: {
                    colors: ["#ffffff"]
                }
            }
        },
        colors: [
          "#008FFB", "#00E396", "#FEB019", "#FF4560", "#775DD0",
          "#3F51B5", "#546E7A", "#D4526E", "#8D5B4C", "#F86624",
          "#D7263D", "#1B998B", "#2E294E", "#F46036", "#E2C044"
        ],
        };

        return o
    }
</script>


<Chart options={getChartOptions()} />