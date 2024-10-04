<script lang="ts">
    import { NavBar } from "../components";
</script>

<div class="flex w-full h-full">
	<div class=""><NavBar /></div>
    <div class="w-full h-full p-4">
        <slot />
    </div>
</div> 