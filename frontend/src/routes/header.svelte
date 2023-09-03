<script lang="ts">
	import { Modal, getModalStore } from "@skeletonlabs/skeleton";
	import type { ModalSettings, ModalComponent, ModalStore } from "@skeletonlabs/skeleton";
	const modalStore = getModalStore();

	function modalDemo(): void {
		const modal: ModalSettings = {
			type: "alert",
			title: "Hello Skeleton",
			body: "This modal example includes a title, body, and image.",
		};
		modalStore.trigger(modal);
	}
</script>

<div id="header" class="container h-full space-y-5">
	<div class="container h-full mx-auto flex justify-start">
		<h1 class="h1">Event Tracker App</h1>
		<div class="flex-grow" />
		<!-- <h4 class="text-right">Developed by<br>Hayden McAlister</h4> -->
	</div>

	<div id="navDiv" class="h-full flex justify-center gap-5">
		<button class="w-full btn variant-filled" on:click={modalDemo}
			>New Event</button
		>
	</div>
</div>

<style>
	#navDiv {
		/* border: solid 1px red; */
	}
</style>
