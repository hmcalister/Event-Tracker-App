<script>
	import Header from "./header.svelte";
</script>

<div class="container h-full mx-auto p-12">
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
