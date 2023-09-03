<script lang="ts">
	import { Modal } from "@skeletonlabs/skeleton";
	import { initializeStores } from "@skeletonlabs/skeleton";
	initializeStores();

	import { fade } from "svelte/transition";
	const MODAL_TRANSITION_DURATION = 200;

	import Header from "./header.svelte";
	import Event from "./event.svelte";
	import { GetAllEvents } from "../../wailsjs/go/main/App";
	import { models } from "../../wailsjs/go/models";

	var allEvents: models.Event[] = [];
	async function AsyncGetEventsDummyFunction() {
		allEvents = await GetAllEvents();
	}
</script>

<Modal
	transitionIn={fade}
	transitionInParams={{ duration: MODAL_TRANSITION_DURATION }}
	transitionOut={fade}
	transitionOutParams={{ duration: MODAL_TRANSITION_DURATION }}
/>

<div id="appContainer" class="h-full w-full p-12">
	<Header />

	<div id="body" class="h-full w-full mx-auto flex justify-center items-center pt-12">
		<div class="space-y-5 mx-auto w-full">
			{#await AsyncGetEventsDummyFunction()}
				<h4>Loading Events</h4>
			{:then _} 
				{#each allEvents as event}
					<Event eventData={event} />
				{/each}
			{:catch}
				<h4>Could not get Events from backend!</h4>
			{/await}
		</div>
	</div>
</div>

<style>
</style>
