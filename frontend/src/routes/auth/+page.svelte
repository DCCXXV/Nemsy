<script lang="ts">
	import { env } from '$env/dynamic/public';
	import { Search } from '@lucide/svelte';
	const loginUrl = `${env.PUBLIC_API_BASE_URL}/auth/login`;
	const listStudiesUrl = `${env.PUBLIC_API_BASE_URL}/api/studies`;

	import type { LayoutData } from './$types';

	let { data }: { data: LayoutData } = $props();
</script>

<div class="my-aut lg:mx-20">
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
			<div class="rounded bg-slate-200 my-4 p-2 flex border-2 border-slate-900 gap-4 items-center">
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
			/>
		</div>

		<div class="w-full text-right">
			<button
				onclick={() => (window.location.href = '/')}
				class="h-10 px-4 py-2 bg-rose-300 hover:brightness-90 border-zinc-900 border-2 rounded transition-colors inline-flex items-center cursor-pointer"
			>
				Terminar
			</button>
		</div>
	{:else}
		<h1 class="text-4xl text-slate-900 select-none">Bienvenido a Nemsy</h1>
		<hr class="text-slate-900 mb-8 border-1 border-slate-900" />
		<p class="bg-zinc-200 border-l-4 border-slate-900 pl-4 py-2 rounded">
			Para facilitar el proceso de <em>onboarding</em>, se recomienda utilizar el correo
			institucional (por ejemplo: usuario@ucm.es), en caso de disponer de él.
		</p>
		<div class="text-right mt-4">
			<button
				onclick={() => (window.location.href = loginUrl)}
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
	{/if}
</div>
