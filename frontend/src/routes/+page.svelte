<script lang="ts">
	import type { PageData } from './$types';
	import { Tabs } from 'bits-ui';
	import type { Subject } from '$lib/types';
	import { page } from '$app/stores';
	import { goto } from '$app/navigation';
	import { onMount } from 'svelte';

	let { data }: { data: PageData } = $props();

	const subjectsByYear = $derived(
		data.subjects.reduce(
			(acc, subject) => {
				const year = subject.year || 'Otros';
				if (!acc[year]) acc[year] = [];
				acc[year].push(subject);
				return acc;
			},
			{} as Record<string, Subject[]>
		)
	);

	const tabIds = $derived(['Fijadas', ...Object.keys(subjectsByYear).sort()]);

	const selectedSubjectID = $derived($page.url.searchParams.get('subject'));
	const selectedSubject = $derived(data.subjects.find((s) => s.id == selectedSubjectID));

	let selectedTab = $state(tabIds[0]);

	function selectSubject(id: string) {
		localStorage.setItem('lastSubject', id);
	}

	function selectTab(id: string) {
		selectedTab = id;
		localStorage.setItem('lastTab', id);
	}

	onMount(() => {
		const savedTab = localStorage.getItem('lastTab');
		if (savedTab && tabIds.includes(savedTab)) {
			selectedTab = savedTab;
		}

		if (!selectedSubjectID) {
			const saved = localStorage.getItem('lastSubject');
			if (saved) {
				goto(`?subject=${saved}`, { replaceState: true });
			}
		}
	});
</script>

{#if data.me}
	<div class="bg-zinc-200 flex items-start justify-center pt-4 pb-6">
		<div class="bg-zinc-100 border-2 border-zinc-900 rounded w-1/4 ml-4 sticky top-4">
			<div class="p-2 flex gap-4 items-center border-b-2 border-zinc-900">
				<img
					src="https://www.google.com/s2/favicons?domain={data.me?.hd}&sz=64"
					alt="Logo de {data.me?.hd}"
					class="rounded border-2 border-zinc-900"
				/>
				<p class="text-xl">Universidad Complutense de Madrid</p>
			</div>
			<div class="p-2 flex gap-4 items-center border-b-2 border-zinc-900">
				<p class="text-lg">Grado en Ingeniería de Software</p>
			</div>
			<div class="p-2 flex gap-4 items-center">
				<Tabs.Root value={selectedTab} onValueChange={selectTab} class="w-full">
					<Tabs.List class="flex w-full gap-2">
						{#each tabIds as id (id)}
							<Tabs.Trigger
								value={id}
								class="flex-1 rounded px-2 py-1 cursor-pointer transition-colors bg-zinc-200 text-zinc-950 hover:bg-zinc-300 data-[state=active]:bg-zinc-700 data-[state=active]:text-white hover:data-[state=active]:bg-zinc-900 text-center"
							>
								{id}
							</Tabs.Trigger>
						{/each}
					</Tabs.List>

					{#each tabIds as id (id)}
						<Tabs.Content value={id}>
							<div>
								<ul class="pt-2">
									{#if id === 'Fijadas'}
										<li class="text-zinc-500 py-2 px-2">
											No has fijado ninguna asignatura todavía
										</li>
									{:else if subjectsByYear[id]?.length}
										{#each subjectsByYear[id] as subject (subject.id)}
											<a
												href="?subject={subject.id}"
												onclick={() => selectSubject(subject.id)}
												class="block rounded py-2 px-2 mb-2 cursor-pointer
												{selectedSubject?.name == subject?.name
													? 'bg-zinc-700 text-zinc-100 hover:bg-zinc-900'
													: 'bg-zinc-100 hover:bg-zinc-200'}"
											>
												{subject.name}
											</a>
										{/each}
									{:else}
										<li class="text-zinc-500 py-2 px-2">No hay asignaturas</li>
									{/if}
								</ul>
							</div>
						</Tabs.Content>
					{/each}
				</Tabs.Root>
			</div>
		</div>
		<div class="bg-zinc-100 border-2 border-zinc-900 rounded w-1/2 mx-4 h-screen">
			<div class="p-2 border-b-2 border-zinc-900">
				<h1 class="text-2xl">
					{selectedSubject ? selectedSubject.name : 'Fijadas'}
				</h1>
			</div>

			<div class="p-4">
				{#if selectedSubject}{:else}
					<p class="text-zinc-500">Selecciona una asignatura</p>
				{/if}
			</div>
		</div>
		<div class="bg-zinc-100 border-2 border-zinc-900 rounded w-1/4 mr-4 h-24 sticky top-4"></div>
	</div>
{:else}
	<div class="relative min-h-screen bg-zinc-100 flex justify-center items-center">
		<div class="absolute inset-x-0 top-0 h-[calc(3/7*100vh)] bg-zinc-200 z-0"></div>

		<div
			class="relative z-10 flex flex-col lg:flex-row w-full max-w-6xl p-4 lg:p-0 mt-24 lg:mt-0 mb-24"
		>
			<div
				class="flex-1 h-full flex justify-center items-center lg:items-start py-12 border-2 border-oatmeal-900 rounded lg:px-10 bg-zinc-100 shadow-[-8px_8px_0px_#000000] transition-all hover:shadow-none hover:translate-x-[-8px] hover:translate-y-[8px]"
			>
				<div
					class="max-w-md flex flex-col items-center lg:items-start text-left mx-auto lg:mx-0 w-full p-4"
				>
					<h1 class="text-5xl font-bold text-zinc-700">nemsy</h1>
					<p class="py-6">
						Comparte y accede a <mark>apuntes universitarios</mark> con facilidad. Todo lo que necesitas
						para estudiar mejor, en un solo lugar.
					</p>

					<div class="grid gap-3 mb-6 w-full">
						<div class="flex items-center gap-2 justify-start">
							<span><mark>Open Source</mark> = transparente y colaborativo.</span>
						</div>
						<div class="flex items-center gap-2 justify-start">
							<span><mark>Sin anuncios</mark> = tu atención en lo que importa.</span>
						</div>
						<div class="flex items-center gap-2 justify-start">
							<span><mark>Rápido y ligero</mark> = acceso instantáneo a tus apuntes.</span>
						</div>
					</div>
				</div>
			</div>

			<div class="flex-1 h-full flex justify-center items-center p-4 rounded-xl">
				<div class="w-full max-w-md text-base-100">
					<div style="height: 400px;"></div>
				</div>
			</div>
		</div>
	</div>
{/if}
