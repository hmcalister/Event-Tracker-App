<script lang="ts">
    import { createEventDispatcher } from "svelte";
	import { models } from "$lib/wailsjs/go/models";
	import { CreateEvent } from "$lib/wailsjs/go/main/App";

	const dispatch = createEventDispatcher();

	// MODAL SETTINGS
	import { getModalStore } from "@skeletonlabs/skeleton";
	import type { ModalSettings, ModalComponent } from "@skeletonlabs/skeleton";
	import EventForm from "./EventForm.svelte";
	const modalStore = getModalStore();

	async function createEvent(r: models.Event) {
		if (r !== undefined) {
			dispatch('eventCreated', new models.Event(await CreateEvent(r)));
		}
	}

	const eventFormModalComponent: ModalComponent = {
		ref: EventForm,
		props: { event: new models.Event({
			"Name": "",
			"StartTime": new Date(),
			"TimeoutDuration": 0,
			"IsRecurring": false,
		}) },
	};
	function eventFormModal(): void {
		const modal: ModalSettings = {
			type: "component",
			title: "Create Event",
			component: eventFormModalComponent,
			response: (r: models.Event) => createEvent(r),
		};
		modalStore.trigger(modal);
	}
</script>

<div id="header" class="h-full space-y-5">
	<div class="h-full mx-auto flex justify-start">
		<h1 class="h1">Event Tracker App</h1>
		<div class="flex-grow" />
		<!-- <h4 class="text-right">Developed by<br>Hayden McAlister</h4> -->
	</div>

	<div id="navDiv" class="h-full mx-auto flex justify-center gap-5">
		<button class="w-full btn variant-filled-primary" on:click={eventFormModal}>New Event</button>
	</div>
</div>

<style>
</style>
