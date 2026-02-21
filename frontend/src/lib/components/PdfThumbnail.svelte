<script lang="ts">
	import { onMount } from 'svelte';

	let { url }: { url: string } = $props();

	let canvas: HTMLCanvasElement;
	let loading = $state(true);
	let error = $state(false);

	onMount(async () => {
		const pdfjsLib = await import('pdfjs-dist');
		pdfjsLib.GlobalWorkerOptions.workerSrc = new URL(
			'pdfjs-dist/build/pdf.worker.min.mjs',
			import.meta.url
		).href;

		try {
			const res = await fetch(url, { credentials: 'include' });
			if (!res.ok) throw new Error(`HTTP ${res.status}`);
			const data = new Uint8Array(await res.arrayBuffer());

			const pdf = await pdfjsLib.getDocument({ data }).promise;
			const page = await pdf.getPage(1);

			const viewport = page.getViewport({ scale: 1.5 });
			canvas.width = viewport.width;
			canvas.height = viewport.height;

			const ctx = canvas.getContext('2d')!;
			await page.render({ canvasContext: ctx, viewport, canvas: canvas }).promise;

			loading = false;
		} catch (e) {
			console.error('Failed to render PDF thumbnail:', e);
			loading = false;
			error = true;
		}
	});
</script>

<div class="w-full max-h-100 overflow-hidden">
	{#if loading && !error}
		<div class="h-48 flex items-center justify-center text-zinc-400 animate-pulse">
			Cargando previsualización...
		</div>
	{/if}
	{#if error}
		<div class="h-24 flex items-center justify-center text-red-300">
			Error al cargar la previsualización
		</div>
	{/if}
	<canvas bind:this={canvas} class="w-full {loading ? 'hidden' : ''}"></canvas>
</div>
