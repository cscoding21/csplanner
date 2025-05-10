<script lang="ts">
	import type { Project } from "$lib/graphql/generated/sdk";
    import { ProjectNavBar } from "../components";
</script>

<div class="flex w-full">
	<div class=""><ProjectNavBar project={{} as Project} /></div>
    <div class="w-full p-4">
        <slot />
    </div>
</div> 