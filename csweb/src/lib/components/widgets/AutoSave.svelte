<script lang="ts">
	import { onMount } from 'svelte';
	import { Badge, Indicator } from 'flowbite-svelte';

	let lastSaveTime:string = $state("");
	let isSaving: boolean = $state(false);

	const SECONDS_TO_UPDATE_EDIT_TIME = 10000;

	interface Props {
		setNotSaving: (st: string) => void;
		setSaving: (st: string) => void;
	}
	let { setNotSaving = $bindable(), setSaving = $bindable() }: Props = $props();

	// export const saver = {
	// 	setSaving(st: string) {
	// 		isSaving = true;
	// 		lastSaveTime = st;
	// 		deltaDisplay(st);
	// 	},
	// 	setNotSaving(st: string) {
	// 		isSaving = false;
	// 		lastSaveTime = st;
	// 		deltaDisplay(st);
	// 	}
	// };

	const deltaDisplay = (st: string) => {
		if (st === null || st === undefined) {
			return '';
		}

		const last = Date.parse(st);
		const now = new Date();
		const diff = now.getTime() - last;

		const days = Math.floor(diff / (1000 * 60 * 60 * 24));
		const hours = Math.floor((diff % (1000 * 60 * 60 * 24)) / (1000 * 60 * 60));
		const minutes = Math.floor((diff % (1000 * 60 * 60)) / (1000 * 60));
		const seconds = Math.floor((diff % (1000 * 60)) / 1000);

		let display = '';

		if (days > 0) {
			let s = days > 1 ? 's' : '';
			display = days + ' day' + s + ' ago';
		} else if (hours > 0) {
			let s = hours > 1 ? 's' : '';
			display = hours + ' hour' + s + ' ago';
		} else if (minutes > 0 && minutes < 3) {
			let ms = minutes > 1 ? 's' : '';
			let ss = seconds > 1 ? 's' : '';
			display = minutes + ' min' + ms + ' ' + seconds + ' sec' + ss + ' ago';
		} else if (minutes >= 3) {
			display = minutes + ' minutes ago';
		} else {
			display = 'mere seconds ago';
		}

		return display;
	};

	onMount(async () => {
		setInterval(() => {
			timeDisplay = deltaDisplay(lastSaveTime);
		}, SECONDS_TO_UPDATE_EDIT_TIME);
	});

	let timeDisplay = $state('');
</script>

<div>
	{#if isSaving}
		<Badge color="yellow" rounded class="px-2.5 py-1.5">
			<Indicator color="yellow" size="md" class="mr-1" />
			Saving...
		</Badge>
	{:else}
		<Badge color="green" rounded class="px-2.5 py-1.5">
			<Indicator color="green" size="md" class="mr-1" />
			Updated {timeDisplay}
		</Badge>
	{/if}
</div>
