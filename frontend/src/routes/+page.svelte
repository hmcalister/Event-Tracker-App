<script lang="ts">
	import { Modal } from "@skeletonlabs/skeleton";
	import { initializeStores } from "@skeletonlabs/skeleton";
	initializeStores();

	import { fade } from "svelte/transition";
	const MODAL_TRANSITION_DURATION = 200;

	import Header from "./header.svelte";
</script>

<div class="container h-full mx-auto p-12">
<Modal
	transitionIn={fade}
	transitionInParams={{ duration: MODAL_TRANSITION_DURATION }}
	transitionOut={fade}
	transitionOutParams={{ duration: MODAL_TRANSITION_DURATION }}
/>
	<Header />

	<div
		id="body"
		class="container h-full mx-auto flex justify-center items-center pt-12"
	>
		<div class="space-y-5">
			<p>Start by exploring:</p>
			<ul>
				<li>
					<code class="code">/src/routes/+layout.svelte</code> - barebones
					layout
				</li>
				<li>
					<code class="code">/src/app.postcss</code> - app wide css
				</li>
				<li>
					<code class="code">/src/routes/+page.svelte</code> - this page,
					you can replace the contents
				</li>
			</ul>
		</div>
	</div>
</div>

<style>
</style>
