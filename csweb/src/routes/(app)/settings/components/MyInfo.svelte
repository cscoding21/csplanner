<script lang="ts">
    import SectionHeading from "$lib/components/formatting/SectionHeading.svelte";
    import { authService } from "$lib/services/auth";

    const as = authService()
    let user = $state(as.currentUser())
</script>


<SectionHeading>My Info</SectionHeading>
<code>{JSON.stringify(user)}</code>
