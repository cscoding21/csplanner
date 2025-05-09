<script lang="ts">
	import type { Project } from "$lib/graphql/generated/sdk";

    interface Props {
        project:Project
    }
    let { project }:Props = $props()
</script>

<div>
    I'm a draft of {project.projectBasics.name}
</div>