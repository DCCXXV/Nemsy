<script lang="ts">
	let { children } = $props();
</script>

<div
	class="flex flex-col lg:flex-row min-h-screen bg-lime-200 p-2 items-center lg:items-stretch justify-center lg:justify-normal"
>
	<div class="invisible lg:visible lg:w-1/2 bg-lime-200"></div>

	<div
		class="w-full lg:w-1/2 bg-zinc-100 rounded flex items-center justify-center border-2 border-zinc-900 px-4 py-20"
	>
		{@render children()}
	</div>
</div>
