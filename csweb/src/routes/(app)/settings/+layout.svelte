<script lang="ts">
	import NavBar from "./components/NavBar.svelte";


</script>

<div class="flex w-full">
	<div class="">
        <NavBar />
    </div>
    <div class="w-full p-4">
        <slot />
    </div>
</div> 