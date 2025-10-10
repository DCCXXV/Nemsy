<script>
	let { text, query } = $props();

	const parts = $derived(
		(() => {
			if (!query || !query.trim()) {
				return [{ text, highlight: false }];
			}

			const lowerText = text.toLowerCase();
			const lowerQuery = query.toLowerCase();

			const index = lowerText.indexOf(lowerQuery);

			if (index === -1) {
				return [{ text: text, highlight: false }];
			}

			return [
				{ text: text.slice(0, index), highlight: false },
				{ text: text.slice(index, index + query.length), highlight: true },
				{ text: text.slice(index + query.length), highlight: false }
			];
		})()
	);
</script>

{#each parts as part, i (i)}
	{#if part.highlight}
		<mark class="bg-amber-200 text-amber-900">{part.text}</mark>
	{:else}
		{part.text}
	{/if}
{/each}
