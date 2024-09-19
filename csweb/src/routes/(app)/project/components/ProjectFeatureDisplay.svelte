<script lang="ts">
    import { Badge, Indicator, Heading, P, type ColorVariant, type IndicatorColorType } from "flowbite-svelte";
	import type { UpdateProjectFeature } from "$lib/graphql/generated/sdk";

    interface ProjectFeatureDisplayProps {
        feature: UpdateProjectFeature;
    }
    let { feature }:ProjectFeatureDisplayProps = $props()

    const setPriorityElements = (p:any):any => {
        if (!p) {
            return;
        }

        switch(p) {
            case "veryhigh":
                priorityDisplay.name = "Must Have";
                priorityDisplay.color = "red";
                priorityDisplay.indicatorColor = "red";
                return priorityDisplay
            case "high":
                priorityDisplay.name = "High";
                priorityDisplay.color = "yellow";
                priorityDisplay.indicatorColor = "yellow";
                return priorityDisplay
            case "medium":
                priorityDisplay.name = "Medium";
                priorityDisplay.color = "blue";
                priorityDisplay.indicatorColor = "blue";
                return priorityDisplay
            case "low":
                priorityDisplay.name = "Low";
                priorityDisplay.color = "purple";
                priorityDisplay.indicatorColor = "purple";
                return priorityDisplay
        }

        return priorityDisplay
    }

    interface PriorityDisplay {
        name: string;
        color: ColorVariant | undefined
        indicatorColor: IndicatorColorType | undefined
    }

    let priorityDisplay:PriorityDisplay = $state({ name: "", color: undefined, indicatorColor: undefined })

    setPriorityElements(feature.priority)
</script>


<Heading tag="h5" class="text-lg dark:text-white">
    {feature.name}

    <span class="float-right">
        <Badge color={priorityDisplay.color} rounded class="px-2.5 py-1">
            <Indicator color={priorityDisplay.indicatorColor} size="xs" class="mr-1" />{priorityDisplay.name}
        </Badge>
    </span> 
</Heading>
<div class="text-gray-100 my-2 text-xs">Status: <b>{feature.status}</b></div>
<P weight="light" color="text-gray-300 dark:text-gray-200 text-sm">{feature.description}</P>