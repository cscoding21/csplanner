<script lang="ts">
	import { Avatar } from 'flowbite-svelte';
	import { getInitialsFromName } from '$lib/utils/format';
	import type { User } from '$lib/graphql/generated/sdk';

	interface Props {
		user: User;
	}
	let { user }: Props = $props();
</script>

{#if user}
<div class="p-2">
	<Avatar src={user.profileImage || ''} class="float-left mr-3 h-16 w-16" rounded
		>{getInitialsFromName(user.firstName + " " + user.lastName)}</Avatar
	>
	<div class="text-left">
		<div>{user.firstName} {user.lastName}</div>
		<div><p>{user.email}</p></div>
		<div class="clear-both"></div>
	</div>
</div>
{/if}
