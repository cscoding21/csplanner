<script lang="ts">
	import { SettingsNavBar } from "./components";


</script>

<div class="flex w-full">
	<div class="">
        <SettingsNavBar />
    </div>
    <div class="w-full p-4">
        <slot />
    </div>
</div> 