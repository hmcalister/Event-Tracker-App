<script lang="ts">
	import { getModalStore } from "@skeletonlabs/skeleton";
	import type { ModalSettings, ModalComponent, ModalStore } from "@skeletonlabs/skeleton";
    import { models } from "../../wailsjs/go/models";
	const modalStore = getModalStore();

	function modalDemo(): void {
		const modal: ModalSettings = {
			type: "alert",
			title: "Hello Skeleton",
			body: "This modal example includes a title, body, and image.",
		};
		modalStore.trigger(modal);
	}

	export let eventData: models.Event;
</script>

<div
	class="eventGridContainer px-5 py-2 h-full w-full variant-ghost rounded-lg items-center justify-items-center"
>
	<div />
	<div class="text-sm">Last Triggered</div>
	<div class="text-sm">Timeout Duration</div>
	<div class="text-sm">Recurring</div>

	<div class="justify-self-start">{eventData.Name}</div>
	<div>{eventData.StartTime.toLocaleString()}</div>
	<div>
		{eventData.TimeoutDuration}
	</div>
	<div>{eventData.IsRecurring}</div>
</div>

<style>
	.eventGridContainer {
		display: grid;
		grid-template-rows: 1fr 3fr;
		grid-template-columns: 6fr 3fr 3fr 1fr;
	}
</style>
