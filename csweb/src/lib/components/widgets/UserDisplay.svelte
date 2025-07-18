<script lang="ts">
	import { Popover } from 'flowbite-svelte';
	import { UserCard } from '..';
	import { normalizeID } from '$lib/utils/id';
	import type { User } from '$lib/graphql/generated/sdk';

	interface Props {
		user: User;
	}
	let { user }: Props = $props();
</script>

{#if user}
<button class="bg-initial" id={normalizeID(user.email as string)}>{user.firstName} {user.lastName}</button>

<Popover
	class="w-64 text-left text-sm font-light"
	title={user.firstName + " " + user.lastName}
	triggeredBy="#{normalizeID(user.email as string)}"
	trigger="click"
>
	<UserCard {user} />
</Popover>
{/if}
