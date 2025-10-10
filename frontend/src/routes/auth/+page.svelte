<script lang="ts">
	import { env } from '$env/dynamic/public';
	import { Search } from '@lucide/svelte';
	import type { LayoutData } from './$types';
	import type { Study } from '$lib/types';
	import { onMount } from 'svelte';
	import HighlightText from '$lib/components/HighlightText.svelte';

	let { data }: { data: LayoutData } = $props();

	let query = $state('');
	let loading = $state(false);
	let studies = $state<Study[]>([]);
	let errorMsg = $state('');
	let selectedStudy = $state<Study | null>(null);

	onMount(() => {
		if (data.me?.StudyID === null) {
			loadStudies();
		}
	});

	async function loadStudies() {
		loading = true;
		errorMsg = '';
		try {
			const response = await fetch(`${env.PUBLIC_API_BASE_URL}/api/studies`, {
				credentials: 'include'
			});
			if (response.ok) {
				studies = await response.json();
			} else {
				errorMsg = 'No se pudieron cargar los grados';
			}
		} catch {
			errorMsg = 'Error de conexión';
		} finally {
			loading = false;
		}
	}

	let filteredStudies = $derived(
		studies.filter((study) => study.Name?.toLowerCase().includes(query.toLocaleLowerCase()))
	);

	function selectStudy(study: Study) {
		selectedStudy = study;
		console.log('Selected study:', study);
	}

	async function finish(study: Study | null) {
		if (study == null) {
			errorMsg = 'Debes seleccionar un grado.';
		} else {
			try {
				const response = await fetch(`${env.PUBLIC_API_BASE_URL}/api/me/study`, {
					method: 'PUT',
					headers: {
						'Content-Type': 'application/json'
					},
					credentials: 'include',
					body: JSON.stringify({ studyId: study.ID })
				});

				if (response.ok) {
					window.location.href = '/';
				} else {
					errorMsg = 'No se pudo seleccionar el grado';
				}
			} catch {
				errorMsg = 'Error al seleccionar el grado';
			}
		}
	}
</script>

<div class="lg:mx-20 w-full h-full flex flex-col relative">
	{#if data.me?.StudyID === null}
		{#if data.me?.Hd === null}
			<p class="text-xl">
				No hemos podido detectar tu universidad a partir de tu correo (<em>{data.me?.Email}</em>).
				Por el momento solo esta disponible la UCM.
			</p>
			<div class="rounded bg-slate-200 my-4 p-2 flex border-2 border-slate-900 gap-4 items-center">
				<img
					src="https://logo.clearbit.com/{data.me?.Hd}"
					alt="Logo de {data.me?.Hd}"
					class="rounded"
				/>
				<h1 class="text-xl">Universidad Complutense de Madrid</h1>
			</div>
		{:else}
			<p class="text-xl">Hemos detectado que tu universidad es:</p>
			<div class="rounded bg-slate-200 my-4 p-2 flex border-2 border-slate-500 gap-4 items-center">
				<img
					src="https://logo.clearbit.com/{data.me?.Hd}"
					alt="Logo de {data.me?.Hd}"
					class="rounded"
				/>
				<!-- TODO: JetBrains SWOT -->
				<h1 class="text-xl">Universidad Complutense de Madrid</h1>
			</div>
		{/if}
		<p class="text-xl">Para continuar, selecciona tu grado:</p>
		<div class="w-full flex rounded border-2 border-slate-900 mb-6 mt-2 items-center">
			<Search class="mx-2" />
			<input
				type="search"
				class="flex-grow border-0 focus:outline-none focus:ring-0 bg-zinc-100"
				placeholder="Buscar grado"
				spellcheck="false"
				bind:value={query}
			/>
		</div>
		{#if loading}
			<!-- skeleton later -->
			<p class="text-center text-slate-600">Cargando grados...</p>
		{:else if filteredStudies.length > 0}
			<div class="mb-6 bg-zinc-100 border-slate-900 border-2 rounded overflow-auto max-h-1/2">
				{#each filteredStudies as study (study.ID)}
					<button
						onclick={() => selectStudy(study)}
						class={`w-full text-left p-2 hover:cursor-pointer transition-colors ${
							selectedStudy?.Name == study.Name
								? 'bg-slate-200 hover:bg-slate-300 border-slate-500'
								: 'bg-zinc-100 hover:bg-zinc-200 border-zinc-100'
						} ${filteredStudies.length == 1 ? 'border-0' : 'border-y-2'}
						}`}
					>
						<h3 class="text-lg text-nowrap">
							<HighlightText text={study.Name} {query} />
						</h3>
					</button>
				{/each}
			</div>
			<p
				class={`bg-rose-200 border-l-4 border-rose-900 pl-4 py-2 mb-6 rounded ${errorMsg == '' ? 'invisible' : ''}`}
			>
				{errorMsg}
			</p>
		{/if}
		<div class="w-full text-right absolute bottom-0 ml-auto">
			<button
				onclick={() => finish(selectedStudy)}
				class="h-10 px-10 py-2 bg-rose-300 hover:brightness-90 border-zinc-900 border-2 rounded transition-colors inline-flex items-center cursor-pointer text-lg"
			>
				Terminar
			</button>
		</div>
	{:else}
		<div class="my-auto">
			<h1 class="text-4xl text-slate-900 select-none">Bienvenido a Nemsy</h1>
			<hr class="text-slate-900 mb-8 border-1 border-slate-900" />
			<p class="bg-zinc-200 border-l-4 border-slate-900 pl-4 py-2 rounded">
				Si es su primer acceso es recomendable usar el correo institucional (por ejemplo:
				usuario@ucm.es), en caso de disponer de él, para facilitar el proceso de <i>onboarding</i>.
			</p>
			<div class="text-right mt-4">
				<button
					onclick={() => (window.location.href = `${env.PUBLIC_API_BASE_URL}/auth/login`)}
					class="h-10 px-4 py-2 bg-rose-300 hover:brightness-90 border-zinc-900 border-2 rounded transition-colors inline-flex items-center cursor-pointer"
				>
					<svg
						class="mr-2 ml-1 w-4 h-4"
						aria-hidden="true"
						focusable="false"
						data-prefix="fab"
						data-icon="google"
						role="img"
						xmlns="http://www.w3.org/2000/svg"
						viewBox="0 0 488 512"
					>
						<path
							fill="currentColor"
							d="M488 261.8C488 403.3 391.1 504 248 504 110.8 504 0 393.2 0 256S110.8 8 248 8c66.8 0 123 24.5 166.3 64.9l-67.5 64.9C258.5 52.6 94.3 116.6 94.3 256c0 86.5 69.1 156.6 153.7 156.6 98.2 0 135-70.4 140.8-106.9H248v-85.3h236.1c2.3 12.7 3.9 24.9 3.9 41.4z"
						></path>
					</svg>
					Iniciar sesión con Google
				</button>
			</div>
		</div>
	{/if}
</div>

<style>
</style>
