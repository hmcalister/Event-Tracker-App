<script lang="ts">
	import { models } from "$lib/wailsjs/go/models";
	import { getModalStore } from "@skeletonlabs/skeleton";
	import type { ModalSettings } from "@skeletonlabs/skeleton";
	export let unparsedEventData: models.Event;

	const modalStore = getModalStore();
	function modalDemo(): void {
		const modal: ModalSettings = {
			type: "alert",
			title: "Hello Skeleton",
			body: "This modal example includes a title, body, and image.",
		};
		modalStore.trigger(modal);
	}

	let event = new models.Event(unparsedEventData);
	const currentTime: Date = new Date();
	const EVENT_TIMEOUT_DURATION_SCALE_FACTOR: number = 1000000;
	let previousTriggerDate = new Date(event.StartTime);
	let nextTriggerTime = event.TimeoutDuration / EVENT_TIMEOUT_DURATION_SCALE_FACTOR;

	// Time difference between now and the event start time, in milliseconds
	let timeDifference = currentTime.getTime() - event.StartTime.getTime();
	console.log(timeDifference, nextTriggerTime)

	// Time between interval triggers, in millisecond
	const INTERVAL_PERIOD = 1000
	let timeIncrementInterval = setInterval( function () {
		if (timeDifference / nextTriggerTime >= 1) {
			timeDifference = nextTriggerTime;
			clearInterval(timeIncrementInterval);
			return;
		}
		timeDifference = timeDifference + INTERVAL_PERIOD;
	}, INTERVAL_PERIOD)
</script>

<div class="eventGridContainer px-5 py-2 h-full w-full variant-ghost rounded-lg gap-x-5 items-center">
	<div />
	<div class="text-sm">Last Triggered</div>
	<div class="text-sm">Progress</div>
	<div class="text-sm">Recurring</div>

	<div>{event.Name}</div>
	<div>{previousTriggerDate.toLocaleString()}</div>
	<div>
		{timeDifference / nextTriggerTime}
		<!-- <CircleProgressBar progress={timeDifference / nextTriggerTime} /> -->
	</div>
	<div>{event.IsRecurring}</div>
</div>

<style>
	.eventGridContainer {
		display: grid;
		grid-template-rows: 1fr 3fr;
		grid-template-columns: 6fr 3fr 3fr 1fr;
	}
</style>
