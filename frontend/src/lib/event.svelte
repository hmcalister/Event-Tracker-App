<script lang="ts">
	// PROPS AND MODELS
	import { models } from "$lib/wailsjs/go/models";
	export let unparsedEventData: models.Event;
	
	// MODAL SETTINGS
	import { getModalStore } from "@skeletonlabs/skeleton";
	import type { ModalSettings } from "@skeletonlabs/skeleton";
	const modalStore = getModalStore();
	function modalDemo(): void {
		const modal: ModalSettings = {
			type: "alert",
			title: "Hello Skeleton",
			body: "This modal example includes a title, body, and image.",
		};
		modalStore.trigger(modal);
	}
	
	// PROGRESS BAR SETTINGS
	import ProgressBar from "@okrad/svelte-progressbar";
	import { cssVarToHexString } from "$lib/cssColorVariableParsing";
	const PROGRESS_BAR_1_COLOR_HEX = cssVarToHexString("--color-secondary-500");
	const PROGRESS_BAR_2_COLOR_HEX = cssVarToHexString("--color-primary-500");
	
	// Parsing props data
	let event = new models.Event(unparsedEventData);
	const currentTime: Date = new Date();
	const EVENT_TIMEOUT_DURATION_SCALE_FACTOR: number = 1000000;
	let previousTriggerDate = new Date(event.StartTime);
	let nextTriggerTime = event.TimeoutDuration / EVENT_TIMEOUT_DURATION_SCALE_FACTOR;

	// Time difference between now and the event start time, in milliseconds
	let timeDifference = Math.min(currentTime.getTime() - event.StartTime.getTime(), nextTriggerTime);

	// Time between interval triggers, in millisecond
	const INTERVAL_PERIOD = 100;
	let timeIncrementInterval = setInterval(function () {
		if (timeDifference / nextTriggerTime >= 1) {
			timeDifference = nextTriggerTime;
			clearInterval(timeIncrementInterval);
			return;
		}
		timeDifference = timeDifference + INTERVAL_PERIOD;
	}, INTERVAL_PERIOD);

	$: backgroundColorClass = (timeDifference / nextTriggerTime < 1) ? 'variant-ghost' : 'variant-ghost-primary';
</script>

<div class="eventGridContainer px-5 pt-2 pb-5 h-full w-full {backgroundColorClass} rounded-lg gap-x-5 items-center">
	<div>
		{#if event.IsRecurring}
			<img src="static/recurringIcon.svg" alt="Recurring Event" class="w-8">
		{/if}
	</div>
	<div class="text-sm">Last Triggered</div>
	<div class="text-sm justify-self-center">Progress</div>
	<div class="text-sm"></div>

	<div>{event.Name}</div>
	<div>{previousTriggerDate.toLocaleString()}</div>
	<div class="justify-self-center">
		<ProgressBar 
			style="radial" 
			thickness={20}
			series={(100 * timeDifference) / nextTriggerTime} 
			valueLabel=" "
			cls="w-16 pt-5 h-auto mx-auto"
			addBackground={false}
			thresholds={[
				{
					till: 99,
					color: PROGRESS_BAR_1_COLOR_HEX,
				},
				{
					till: 100,
					color: PROGRESS_BAR_2_COLOR_HEX,
				}
			]}
		/>
	</div>
	<div><button class="btn variant-filled">Dismiss</div>
</div>

<style>
	.eventGridContainer {
		display: grid;
		grid-template-rows: 1fr 3fr;
		grid-template-columns: 5fr 4fr 3fr 1fr;
	}
</style>
