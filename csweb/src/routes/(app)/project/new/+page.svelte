<script lang="ts">
    import { StepIndicator } from 'flowbite-svelte';
    import { CSSection, PageHeading } from '$lib/components';
	import { ProjectBasics } from '../components';


    let currentStep = 3
    let steps = ['Project Basics', 'Team', 'Implementation Template'];

</script>

<PageHeading title="Add a New Project" />
<div class="p-4"> 
    <CSSection>
<div class="my-4">
  <StepIndicator {currentStep} {steps} color="yellow" />
</div>

<ProjectBasics id=""></ProjectBasics>

</CSSection>

</div>