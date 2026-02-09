<script lang="ts">
	import { PUBLIC_API_BASE_URL } from '$env/static/public';

	import type { PageData } from './$types';
	import { Tabs } from 'bits-ui';
	import type { Subject } from '$lib/types';
	import { page } from '$app/state';
	import { goto } from '$app/navigation';
	import { onMount } from 'svelte';

	import { Toggle } from 'melt/components';
	import { Tooltip } from 'melt/components';

	import DownloadSimpleIcon from 'phosphor-svelte/lib/DownloadSimpleIcon';
	import ArrowFatUpIcon from 'phosphor-svelte/lib/ArrowFatUpIcon';
	import ArrowFatDownIcon from 'phosphor-svelte/lib/ArrowFatDownIcon';
	import ArchiveIcon from 'phosphor-svelte/lib/ArchiveIcon';
	import RowsIcon from 'phosphor-svelte/lib/RowsIcon';
	import SquareIcon from 'phosphor-svelte/lib/SquareIcon';
	import PushPinIcon from 'phosphor-svelte/lib/PushPinIcon';
	import ImagesIcon from 'phosphor-svelte/lib/ImagesIcon';
	import PencilRulerIcon from 'phosphor-svelte/lib/PencilRulerIcon';

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

	const selectedSubjectID = $derived(page.url.searchParams.get('subject'));
	const selectedSubject = $derived(
		data.subjects.find((s) => s.id.toString() === selectedSubjectID)
	);

	let selectedTab = $state(tabIds[0]);
	let compactMode = $state(false);

	function selectSubject(id: string) {
		localStorage.setItem('lastSubject', id);
	}

	function selectTab(id: string) {
		selectedTab = id;
		localStorage.setItem('lastTab', id);
	}

	function selectViewMode(compact: boolean) {
		compactMode = compact;
		localStorage.setItem('compactMode', compact.toString());
	}

	onMount(() => {
		const savedTab = localStorage.getItem('lastTab');
		if (savedTab && tabIds.includes(savedTab)) {
			selectedTab = savedTab;
		}

		const savedCompactMode = localStorage.getItem('compactMode');
		if (savedCompactMode === 'true') {
			compactMode = true;
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
	<div class="relative bg-zinc-100 min-h-screen overflow-hidden">
		<div class="relative z-10 flex items-start justify-center pt-4 pb-6">
			<div class="bg-zinc-50 border border-zinc-300 rounded-none w-1/4 ml-4 sticky top-4">
				<div class="p-2 flex gap-4 items-center border-b border-zinc-300">
					<img
						src="https://www.google.com/s2/favicons?domain={data.me?.hd}&sz=64"
						alt="Logo de {data.me?.hd}"
						class="rounded-none border border-zinc-300"
					/>
					<p class="text-xl">Universidad Complutense de Madrid</p>
				</div>
				<div class="p-2 flex gap-4 items-center border-b border-zinc-300">
					<p class="text-lg">Grado en Ingeniería de Software</p>
				</div>
				<div class="p-2 flex gap-4 items-center">
					<Tabs.Root value={selectedTab} onValueChange={selectTab} class="w-full">
						<Tabs.List class="flex w-full gap-2">
							{#each tabIds as id (id)}
								<Tabs.Trigger
									value={id}
									class="flex-1 rounded-none px-2 py-1 cursor-pointer border border-zinc-200 transition-colors bg-zinc-100 text-zinc-950 hover:bg-zinc-200 data-[state=active]:bg-violet-200 data-[state=active]:text-violet-900 hover:data-[state=active]:bg-violet-200 text-center"
								>
									{id}
								</Tabs.Trigger>
							{/each}
						</Tabs.List>

						{#each tabIds as id (id)}
							<Tabs.Content value={id}>
								<div class="max-h-[calc(50vh-2rem)] overflow-auto">
									<ul class="pt-2">
										{#if id === 'Fijadas'}
											<li class="text-zinc-500 py-2 px-2">
												No has fijado ninguna asignatura todavía
											</li>
										{:else if subjectsByYear[id]?.length}
											{#each subjectsByYear[id] as subject (subject.id)}
												<a
													href="?subject={subject.id}"
													onclick={() => selectSubject(subject.id.toString())}
													class="block rounded-none py-2 px-2 mb-2 border cursor-pointer
												{selectedSubject?.name == subject?.name
														? 'bg-lime-200 border-lime-200 text-lime-800'
														: 'text-zinc-700 bg-zinc-50 hover:bg-zinc-100 border-zinc-50 hover:border-zinc-200'}"
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
			<div
				class="bg-zinc-50 border border-zinc-300 rounded-none w-1/2 mx-4 {data.resources.length
					? ''
					: 'min-h-screen'}"
			>
				<div class="p-2 border-b border-zinc-300 text-zinc-700 flex items-center justify-between">
					<h1 class="text-2xl">
						{selectedSubject ? selectedSubject.name : 'Fijadas'}
					</h1>

					<Tooltip>
						{#snippet children(tooltip)}
							<span {...tooltip.trigger}>
								<Toggle>
									{#snippet children(toggle)}
										<button
											{...toggle.trigger}
											class="flex items-center justify-center cursor-pointer"
										>
											{#if toggle.value}
												<PushPinIcon weight="fill" class="size-6 text-red-400" />
											{:else}
												<PushPinIcon
													weight="regular"
													class="size-6 text-zinc-500 hover:text-zinc-900"
												/>
											{/if}
										</button>
									{/snippet}
								</Toggle>
							</span>
							<div {...tooltip.content}>
								<div {...tooltip.arrow}></div>
								<p class="border border-zinc-300 bg-zinc-50 p-2 text-zinc-500 rounded-none">
									Fijar Asignatura
								</p>
							</div>
						{/snippet}
					</Tooltip>
				</div>
				<div class="bg-zinc-100 border-b border-zinc-300 flex items-center py-1 justify-between">
					<div></div>
					<div class="text-zinc-500 bg-zinc-100 border-l hover:text-zinc-900 border-zinc-300">
						<button
							class="flex gap-1 items-center justify-center cursor-pointer px-2"
							onclick={() => selectViewMode(!compactMode)}
						>
							<ImagesIcon class="size-6" />
							<span class="w-20 text-left">{compactMode ? 'Compacto' : 'Comforte'}</span>
						</button>
					</div>
				</div>

				<div class="">
					{#if selectedSubject}
						{#each data.resources as resource (resource.id)}
							{#if compactMode}
								<div class="border-b last:border-b-0 p-2 border-zinc-200 flex gap-3">
									{#if resource.fileUrl.match(/\.(jpg|jpeg|png|gif|webp)$/i)}
										<img
											src="{PUBLIC_API_BASE_URL}{resource.fileUrl}"
											alt="{resource.title} thumbnail"
											class="rounded-none border border-zinc-300 w-20 object-cover self-stretch"
										/>
									{/if}
									<div class="flex flex-col flex-1 justify-between py-1">
										<div>
											<h2 class="text-base">{resource.title}</h2>
											<p class="text-sm text-zinc-500">
												@{resource.owner?.email?.split('@')[0]}
											</p>
										</div>
										<div class="flex justify-end gap-2">
											<!--
											<button
												class="bg-zinc-100 hover:bg-zinc-200 border border-zinc-300 text-zinc-600 px-2 py-0.5 flex items-center cursor-pointer text-sm"
												><ArchiveIcon class="size-4 mr-1" />Guardar</button
											>-->
											<a href="{PUBLIC_API_BASE_URL}/api/resources/{resource.id}/download">
												<button
													class="bg-blue-200 border border-blue-100 hover:bg-blue-100 text-blue-900 px-2 py-0.5 flex items-center cursor-pointer text-sm"
													><DownloadSimpleIcon class="size-4 mr-1" />Descargar</button
												></a
											>
										</div>
									</div>
								</div>
							{:else}
								<div class="border-b last:border-b-0 p-2 border-zinc-200 flex gap-2">
									<div class="flex flex-col gap-2">
										<div class="flex align-middle">
											<img
												class="size-12 border border-zinc-300 mr-2 rounded-none"
												src={resource.owner?.pfp}
												alt="{resource.owner?.fullName}'s profile picture"
												referrerpolicy="no-referrer"
											/>
											<div class="flex flex-col">
												<h2 class="text-xl -mb-1">{resource.title}</h2>
												<p class="text-md text-zinc-500">
													@{resource.owner?.email?.split('@')[0]}
												</p>
											</div>
										</div>
										<div>
											{#if resource.fileUrl.match(/\.(jpg|jpeg|png|gif|webp)$/i)}
												<img
													src="{PUBLIC_API_BASE_URL}{resource.fileUrl}"
													alt="{resource.title} thumbnail"
													class="rounded-none border border-zinc-300 w-full"
												/>
											{/if}
										</div>
										<p class="text-zinc-700">{resource.description}</p>
										<div class="flex justify-end mb-4 gap-2">
											<!--
											<button
												class="bg-zinc-100 hover:bg-zinc-200 border border-zinc-300 text-zinc-600 px-3 py-1 flex items-center cursor-pointer"
												><ArchiveIcon class="size-5 mr-2" />Guardar</button
											>-->
											<a href="{PUBLIC_API_BASE_URL}/api/resources/{resource.id}/download">
												<button
													class="bg-blue-200 border border-blue-100 hover:bg-blue-100 text-blue-900 px-3 py-1 flex items-center cursor-pointer"
													><DownloadSimpleIcon class="size-5 mr-2" />Descargar</button
												></a
											>
										</div>
									</div>
								</div>
							{/if}
						{/each}
					{:else}
						<p class="text-zinc-500">Selecciona una asignatura</p>
					{/if}
				</div>
			</div>
			<div
				class="text-yellow-600 bg-yellow-100 border border-yellow-400 rounded-none w-1/4 mr-4 min-h-136 sticky top-4 flex flex-col items-center"
			>
				<div class="m-auto aspect-square text-center p-2">
					<PencilRulerIcon weight="thin" class="size-14 mx-auto" />
					<h3>WIP</h3>
				</div>
			</div>
		</div>
	</div>
{:else}
	<div class="relative min-h-screen bg-zinc-200 flex justify-center items-center">
		<div class="absolute inset-x-0 top-0 h-[calc(3/7*100vh)] bg-zinc-100 z-0"></div>

		<div
			class="relative z-10 flex flex-col lg:flex-row w-full max-w-6xl p-4 lg:p-0 mt-24 lg:mt-0 mb-24"
		>
			<div
				class="flex-1 h-full flex justify-center items-center lg:items-start py-12 border border-zinc-300 rounded-none lg:px-10 bg-zinc-50 shadow-[-8px_8px_0px_#d4d4d8] transition-all hover:shadow-none hover:translate-x-[-8px] hover:translate-y-[8px]"
			>
				<div
					class="max-w-md flex flex-col items-center lg:items-start text-left mx-auto lg:mx-0 w-full p-4"
				>
					<h1 class="text-5xl text-zinc-700">nemsy</h1>
					<p class="py-6">
						Comparte y accede a <mark class="bg-red-200 text-red-900">apuntes universitarios</mark> con
						facilidad. Todo lo que necesitas para estudiar mejor, en un solo lugar.
					</p>

					<div class="grid gap-3 mb-6 w-full">
						<div class="flex items-center gap-2 justify-start">
							<span
								><mark class="bg-yellow-200 text-yellow-900">Open Source</mark> = transparente y colaborativo.</span
							>
						</div>
						<div class="flex items-center gap-2 justify-start">
							<span
								><mark class="bg-blue-200 text-blue-900">Sin anuncios</mark> = tu atención en lo que importa.</span
							>
						</div>
						<div class="flex items-center gap-2 justify-start">
							<span
								><mark class="bg-lime-200 text-lime-900">Rápido y ligero</mark> = acceso instantáneo a
								tus apuntes.</span
							>
						</div>
					</div>
				</div>
			</div>

			<div class="flex-1 h-full flex justify-center items-center p-4 rounded-none-xl">
				<div class="w-full max-w-md text-base-100">
					<div style="height: 400px;"></div>
				</div>
			</div>
		</div>
	</div>
{/if}
