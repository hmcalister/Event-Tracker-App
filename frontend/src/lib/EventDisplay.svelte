<script lang="ts">
	import { createEventDispatcher } from "svelte";
	import { fade } from "svelte/transition";

	const dispatch = createEventDispatcher();

	// PROPS AND MODELS
	import { models } from "$lib/wailsjs/go/models";
	import { UpdateEvent, DismissEvent, DeleteEvent } from "$lib/wailsjs/go/main/App";
	export let event: models.Event;
	const currentTime: Date = new Date();
	const EVENT_TIMEOUT_DURATION_SCALE_FACTOR: number = 1000000;
	let previousTriggerDate = new Date(event.StartTime);
	let nextTriggerTime = event.TimeoutDuration / EVENT_TIMEOUT_DURATION_SCALE_FACTOR;

	async function updateEvent(r: models.Event) {
		if (r !== undefined) {
			event = new models.Event(await UpdateEvent(r));
		}
	}

	async function dismissEvent() {
		var dismissResult = await DismissEvent(event);
		if (dismissResult !== null) {
			event = new models.Event(dismissResult);
		} else {
			dispatch("eventDeleted");
		}
	}

	async function deleteEvent() {
		var dismissResult = await DeleteEvent(event);
		if (dismissResult !== null) {
			event = new models.Event(dismissResult);
		} else {
			dispatch("eventDeleted");
		}
	}

	// MODAL SETTINGS
	import { getModalStore } from "@skeletonlabs/skeleton";
	import type { ModalSettings, ModalComponent } from "@skeletonlabs/skeleton";
	import EventForm from "./EventForm.svelte";
	const modalStore = getModalStore();
	const eventFormModalComponent: ModalComponent = {
		ref: EventForm,
		props: { event: event },
	};
	function eventFormModal(): void {
		const modal: ModalSettings = {
			type: "component",
			title: "Edit Event",
			component: eventFormModalComponent,
			response: (r: models.Event) => updateEvent(r),
		};
		modalStore.trigger(modal);
	}

	// PROGRESS BAR SETTINGS
	import ProgressBar from "@okrad/svelte-progressbar";
	import { cssVarToHexString } from "$lib/cssColorVariableParsing";
	const PROGRESS_BAR_1_COLOR_HEX = cssVarToHexString("--color-secondary-500");
	const PROGRESS_BAR_2_COLOR_HEX = cssVarToHexString("--color-primary-500");

	// Time difference between now and the event start time, in milliseconds
	let timeDifference = Math.max(Math.min(currentTime.getTime() - event.StartTime.getTime(), nextTriggerTime), 0);

	// Time between interval triggers, in millisecond
	const INTERVAL_PERIOD = 100;
	let timeIncrementInterval = setInterval(function () {
		if (timeDifference <= 0) {
			return;
		}
		if (timeDifference / nextTriggerTime >= 1) {
			timeDifference = nextTriggerTime+1;
			clearInterval(timeIncrementInterval);
			return;
		}
		timeDifference = timeDifference + INTERVAL_PERIOD;
	}, INTERVAL_PERIOD);

	$: backgroundColorClass = timeDifference / nextTriggerTime < 1 ? "variant-ghost" : "variant-ghost-primary";
</script>

{#if event !== null}
	<div transition:fade class="eventGridContainer px-5 pt-2 pb-5 h-full w-full {backgroundColorClass} rounded-lg gap-x-5 items-center">
		<div>
			{#if event.IsRecurring}
				<img src="/recurringIcon.svg" alt="Recurring Event" class="w-8" />
			{/if}
		</div>
		<div class="text-sm">Last Triggered</div>
		<div class="text-sm justify-self-center">Progress</div>
		<div class="text-sm justify-self-end"><img src="/bin.png" alt="Recurring Event" class="w-8 cursor-pointer" on:click={deleteEvent}/></div>

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
						till: 100,
						color: PROGRESS_BAR_1_COLOR_HEX,
					},
					{
						till: 101,
						color: PROGRESS_BAR_2_COLOR_HEX,
					},
				]}
			/>
		</div>
		<div class="flex gap-2">
			<button class="btn variant-filled" on:click={eventFormModal}>Edit</button>
			<button class="btn variant-filled" on:click={dismissEvent}>Dismiss</button>
		</div>
	</div>
{/if}

<style>
	.eventGridContainer {
		display: grid;
		grid-template-rows: 1fr 3fr;
		grid-template-columns: 5fr 4fr 3fr 1fr;
	}
</style>
