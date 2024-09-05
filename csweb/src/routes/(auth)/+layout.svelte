<script lang="ts">
    import { onMount } from "svelte";
    import { isAuthenticated } from "$lib/stores/auth";
    import { goto } from '$app/navigation'

    onMount(async () => {
		if(isAuthenticated()) {
			goto("/home")
		}
	});
</script>


<div class="flex flex-row">
  <div class="basis-1/3">&nbsp;</div>
  <div class="basis-1/3"><slot /></div>
  <div class="basis-1/3">&nbsp;</div>
</div>