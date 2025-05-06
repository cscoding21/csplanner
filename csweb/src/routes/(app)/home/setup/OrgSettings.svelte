<script lang="ts">
    import { Button, P } from 'flowbite-svelte';
	  import { getOrganization } from '$lib/services/organization';
	  import { orgForm, orgSchema } from '$lib/forms/organization.validation';
	  import type { UpdateOrganization, Organization } from '$lib/graphql/generated/sdk';
	  import { coalesceToType } from '$lib/forms/helpers';
	  import { addToast } from '$lib/stores/toasts';
	  import { formatCurrency, formatPercent } from '$lib/utils/format';

    interface Props {
        onDone: Function
    }
    let { onDone }:Props = $props()


    const load = async (): Promise<Organization> => {
		return await getOrganization()
			.then((o) => {
				of = coalesceToType<UpdateOrganization>(o.defaults, orgSchema);

                org = o
				return org;
			})
			.catch((err) => {
				addToast({
					message: 'Error loading organization (OrgSettings): ' + err,
					dismissible: true,
					type: 'error'
				});

				return err;
			});
	};

  let of = $state({} as typeof orgForm);
	let org = $state({} as Organization)
</script>


{#await load()}
    Loading...
{:then promiseData} 
    
<h2 class="text-xl text-center text-gray-50 font-semibold">Organization Settings</h2>
    
<p class="py-6 text-gray-200">
    The following settings are used to calculate costs and value of the projects
    in your portfolio.  These can be modified on the 
    <a class="text-orange-300" href="/settings/#org-settings">Org Settings</a> page at any time.
</p>

<ul role="list" class="divide-y divide-gray-100 dark:divide-gray-700">
    <li class="py-3 sm:py-4">
      <div class="flex items-center">
        <!-- <div class="shrink-0">
          <img class="h-9 w-9" src="/images/e-commerce/imac-front-image.png" alt="Product image" />
        </div> -->
        <div class="ms-4 min-w-0 flex-1 pr-8">
          <p class="truncate font-semibold text-gray-900 dark:text-white">Discount rate</p>
          <p class="text-sm text-gray-500 dark:text-gray-400">The discount rate is in the net present value calculation of a project.</p>
        </div>
        <div class="inline-flex items-center text-base font-semibold text-gray-900 dark:text-white">{formatPercent.format(org.defaults.discountRate * 0.01)}</div>
      </div>
    </li>
    <li class="py-3 sm:py-4">
      <div class="flex items-center">
        <div class="ms-4 min-w-0 flex-1 pr-8">
          <p class="truncate font-semibold text-gray-900 dark:text-white">Focus factor</p>
          <p class="text-sm text-gray-500 dark:text-gray-400">This value is used to estimate additional hours for users that are involved in multiple projects</p>
        </div>
        <div class="inline-flex items-center text-base font-semibold text-gray-900 dark:text-white">{formatPercent.format(org.defaults.focusFactor * 0.01)}</div>
      </div>
    </li>
    <li class="py-3 sm:py-4">
      <div class="flex items-center">
        <div class="ms-4 min-w-0 flex-1 pr-8">
          <p class="truncate font-semibold text-gray-900 dark:text-white">Comms coefficient</p>
          <p class="text-sm text-gray-500 dark:text-gray-400">This value is used to estimate additional hours for multiple users working on a single task</p>
        </div>
        <div class="inline-flex items-center text-base font-semibold text-gray-900 dark:text-white">{formatPercent.format(org.defaults.commsCoefficient * 0.01)}</div>
      </div>
    </li>
    <li class="py-3 sm:py-4">
      <div class="flex items-center">
        <div class="ms-4 min-w-0 flex-1 pr-8">
          <p class="truncate font-semibold text-gray-900 dark:text-white">Generic blended hourly rate</p>
          <p class="text-sm text-gray-500 dark:text-gray-400">This is a default hourly rate that acts as a default when estimating project costs. It can be overridden at the role of individual level</p>
        </div>
        <div class="inline-flex items-center text-base font-semibold text-gray-900 dark:text-white">{formatCurrency.format(org.defaults.genericBlendedHourlyRate)}/hr</div>
      </div>
    </li>
    <li class="py-3 sm:py-4">
      <div class="flex items-center">
        <div class="ms-4 min-w-0 flex-1 pr-8">
          <p class="truncate font-semibold text-gray-900 dark:text-white">Working hours per week</p>
          <p class="text-sm text-gray-500 dark:text-gray-400">This value is used for estimating a work week ***check to see if it can be overridden at the resource level***</p>
        </div>
        <div class="inline-flex items-center text-base font-semibold text-gray-900 dark:text-white">{org.defaults.hoursPerWeek}</div>
      </div>
    </li>
    <li class="py-3 sm:py-4">
      <div class="flex items-center">
        <div class="ms-4 min-w-0 flex-1 pr-8">
          <p class="truncate font-semibold text-gray-900 dark:text-white">Working hours per year</p>
          <p class="text-sm text-gray-500 dark:text-gray-400">The number of working hours in a year is used to calculate an hourly rate for resources that have an annual cost associated with them. This will override orginzation and role level values.</p>
        </div>
        <div class="inline-flex items-center text-base font-semibold text-gray-900 dark:text-white">{org.defaults.workingHoursPerYear}</div>
      </div>
    </li>
  </ul>

  <div class="mt-12 text-center">
    <Button onclick={() => onDone()}>I've got it.  Let's continue! >></Button>
</div>


{/await}