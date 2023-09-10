<script lang="ts">
	import { Modal } from "@skeletonlabs/skeleton";
	import { initializeStores } from "@skeletonlabs/skeleton";
	import { models } from "$lib/wailsjs/go/models";
	initializeStores();

	import { fade } from "svelte/transition";
	const MODAL_TRANSITION_DURATION = 200;

	import Header from "$lib/Header.svelte";
	import EventDisplay from "$lib/EventDisplay.svelte";
	import { GetAllEvents } from "$lib/wailsjs/go/main/App";

	var allEvents: models.Event[] = [];
	async function AsyncGetEventsDummyFunction() {
		var asyncResponse = await GetAllEvents();
		allEvents = asyncResponse.map((x) => new models.Event(x));
	}
</script>

<Modal transitionIn={fade} transitionInParams={{ duration: MODAL_TRANSITION_DURATION }} transitionOut={fade} transitionOutParams={{ duration: MODAL_TRANSITION_DURATION }} />

<div id="appContainer" class="h-full w-full p-12">
	<Header
		on:eventCreated={(r) => {
			allEvents.push(r.detail);
			allEvents = allEvents;
		}}
	/>

	<div id="body" class="h-full w-full mx-auto flex justify-center items-center pt-12">
		<div class="space-y-5 mx-auto w-full">
			{#await AsyncGetEventsDummyFunction()}
				<h4>Loading Events</h4>
			{:then _}
				{#each allEvents as event (event)}
					<EventDisplay bind:event on:eventDeleted={(r) => {
						allEvents = allEvents.filter((e) => e.ID != event.ID);
					}}/>
				{/each}
			{:catch}
				<h4>Could not get Events from backend!</h4>
			{/await}
		</div>
	</div>
</div>

<style>
</style>
